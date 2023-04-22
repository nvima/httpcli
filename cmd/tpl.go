package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/nvima/httpcli/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type FunctionConfig struct {
	Header []string               `mapstructure:"header"`
	Data   map[string]interface{} `mapstructure:"data"`
	Env    []string               `mapstructure:"env"`
	Url    string                 `mapstructure:"url"`
	Output string                 `mapstructure:"output"`
}

type AppConfig struct {
	Functions map[string]FunctionConfig
}

func tplCommand(cmd *cobra.Command, args []string) {
	config := viper.AllSettings()
	if len(config) == 0 {
		panic("No config found")
	}

	if len(args) == 0 {
		panic("No function name provided")
	}

	var appConfig AppConfig
	err := mapstructure.Decode(config, &appConfig.Functions)
	if err != nil {
		panic("Failed to decode config: " + err.Error())
	}

	fc, ok := appConfig.Functions[args[0]]
	if !ok {
		panic("No config found for function: " + args[0])
	}
	
    fmt.Fprintf(cmd.OutOrStdout(), fc.handleFunc())
}

func (fc *FunctionConfig) handleFunc() string {
	jsonData := fc.getJSONData()

	req, err := http.NewRequest("POST", fc.replaceEnvVariables(fc.Url), bytes.NewBuffer(jsonData))
	if err != nil {
        panic("Failed to create request: " + err.Error())
	}

	for _, header := range fc.Header {
		header = fc.replaceEnvVariables(header)
		headerParts := strings.SplitN(header, ":", 2)
		if len(headerParts) == 2 {
			req.Header.Set(strings.TrimSpace(headerParts[0]), strings.TrimSpace(headerParts[1]))
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        panic("Failed to send request: " + err.Error())
	}
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        panic("Failed to read response body: " + err.Error())
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
        panic("Request failed, Status: " + resp.Status + ", Body: " + string(body))
	}

	// Parse the JSON response
	responseData, err := util.ParseJSONResponse(body)
	if err != nil {
		panic(err)
	}

	// Extract the desired output from the JSON response
	return util.GetOutputField(responseData, fc.Output)
}

func (fc *FunctionConfig) getJSONData() []byte {
	jsonData, err := json.Marshal(fc.Data)
	if err != nil {
		panic(err)
	}

	jsonData, err = util.ReplaceStdIn(jsonData)
	if err != nil {
		panic(err)
	}

	return jsonData
}

func (fc *FunctionConfig) replaceEnvVariables(value string) string {
	for _, envVar := range fc.Env {
		envValue := os.Getenv(envVar)
		placeholder := fmt.Sprintf("${%s}", envVar)
		value = strings.Replace(value, placeholder, envValue, -1)
	}
	return value
}
