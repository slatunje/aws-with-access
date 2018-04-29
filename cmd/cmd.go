// Copyright © 2018 Sylvester La-Tunje

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
	"github.com/slatunje/aws-role/pkg/cue"
)

var app = "with"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   app,
	Short: fmt.Sprintf("=> %s makes it easy to obtain temporary AWS credentials", app),
	Long: fmt.Sprintf(`
Description:
  %s makes it easy to obtain temporary AWS credentials whenever you are required to access an AWS 
  account through 'AssumeRole'.
`, app),
	Run: func(cmd *cobra.Command, args []string) {
		cue.Credentials()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	os.Setenv("TZ", "")
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.%s.yaml)", app))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	switch true {
	case cfgFile != "":
		viper.SetConfigFile(cfgFile)
		break
	default:
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(fmt.Sprintf(".%s", app)) // $HOME is default ".awsrole" (without extension).
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
