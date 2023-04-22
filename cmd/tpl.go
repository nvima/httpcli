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
	Header     []string               `mapstructure:"header"`
	Data       map[string]interface{} `mapstructure:"data"`
	Env        []string               `mapstructure:"env"`
	Url        string                 `mapstructure:"url"`
	Output     string                 `mapstructure:"output"`
	StatusCode int                    `mapstructure:"statuscode"`
}

type AppConfig struct {
	Functions map[string]FunctionConfig
}

var fc FunctionConfig

func tplCommand(cmd *cobra.Command, args []string) error {
	err := initFunctionConfig(cmd, args)
	if err != nil {
		return err
	}

	output, err := fc.handleFunc(cmd)
	if err != nil {
		return err
	}

	fmt.Fprintf(cmd.OutOrStdout(), output)
	return nil
}

func initFunctionConfig(cmd *cobra.Command, args []string) error {
	config := viper.AllSettings()

	if len(config) == 0 {
		return util.HandleError(cmd, util.NO_FUNC_NAME_ERR.Err(), util.NO_CONFIG_FILE_ERR)
	}

	if len(args) == 0 {
		return util.HandleError(cmd, util.NO_FUNC_NAME_ERR.Err(), util.NO_FUNC_NAME_ERR)
	}

	var appConfig AppConfig
	err := mapstructure.Decode(config, &appConfig.Functions)
	if err != nil {
		return util.HandleError(cmd, err, util.INVALID_CONFIG_ERR)
	}

	funcConfig, ok := appConfig.Functions[args[0]]
	if !ok {
		return util.HandleError(cmd, util.NO_FUNC_FOUND_ERR.Err(), util.NO_FUNC_FOUND_ERR)
	}

	if funcConfig.Url == "" {
		return util.HandleError(cmd, util.NO_URL_ERR.Err(), util.NO_URL_ERR)
	}

	fc = funcConfig
	return nil
}

func (fc *FunctionConfig) makeHttpCall(jsonData []byte, cmd *cobra.Command) ([]byte, error) {
	url := fc.replaceEnvVariables(fc.Url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the request was successful when a status code is provided
	if fc.StatusCode != 0 && resp.StatusCode != fc.StatusCode {
		err := fmt.Errorf("Request failed with status code %d, Body: %s", resp.StatusCode, string(body))
		util.HandleError(cmd, err, util.INVALID_RESP_CODE)
	}
	return body, nil
}

func (fc *FunctionConfig) handleFunc(cmd *cobra.Command) (string, error) {
	jsonData, err := fc.getJSONData()
	if err != nil {
		return "", util.HandleError(cmd, err, util.FAILED_TO_GET_DATA)
	}

	body, err := fc.makeHttpCall(jsonData, cmd)
	if err != nil {
		return "", util.HandleError(cmd, err, util.FAILED_TO_MAKE_HTTP_CALL)
	}

	responseData, err := util.ParseJSONResponse(body)
	if err != nil {
		return "", util.HandleError(cmd, err, util.FAILED_TO_PARSE_JSON)
	}

	output, err := util.GetOutputField(responseData, fc.Output)
	if err != nil {
		return "", util.HandleError(cmd, err, util.FAILED_TO_PARSE_OUTPUT_FIELD)
	}

	return output, nil
}

func (fc *FunctionConfig) getJSONData() ([]byte, error) {
	jsonData, err := json.Marshal(fc.Data)
	if err != nil {
		return nil, err
	}

	jsonData, err = util.ReplaceStdIn(jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (fc *FunctionConfig) replaceEnvVariables(value string) string {
	for _, envVar := range fc.Env {
		envValue := os.Getenv(envVar)
		placeholder := fmt.Sprintf("${%s}", envVar)
		value = strings.Replace(value, placeholder, envValue, -1)
	}

	return value
}
