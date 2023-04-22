package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "httpcli",
	Short: "A flexible CLI tool for making API requests using YAML configuration",
	Long: `httpcli is a versatile command-line tool that allows users to define and manage API requests
using YAML configuration files. The tool supports controlling API requests through environment variables
and standard input (stdin). Users can create functions in the YAML configuration, such as the "translate" function,
which sends a request to an Translation API and outputs the response to stdout.

For example, the following command:

	cat german.txt | httpcli translate > english.txt

reads from the YAML configuration file, processes the "translate" function, and sends an API request to an API.
The response is then printed to stdout. For more Information about YAML Configuration, visit https://github.com/nvima/httpcli.`,
	Run: tplCommand,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.httpcli.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".httpcli")
	}
	viper.ReadInConfig()
}
