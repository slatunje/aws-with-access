// Copyright © 2018 Sylvester La-Tunje

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/slatunje/aws-with-access/pkg/cue"
	"github.com/slatunje/aws-with-access/pkg/env"
	"github.com/slatunje/aws-with-access/pkg/utils"
)

const app = "with"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   app,
	Short: fmt.Sprintf("=> %s makes it easy to obtain temporary AWS credentials", app),
	Long: fmt.Sprintf(`
Description:
  %s makes it easier to obtain temporary AWS credentials through 'AssumeRole'.
`, app),
	Run: func(cmd *cobra.Command, args []string) {
		cue.Credentials()
	},
}

// init is called in alphabetic order within this package
func init() {
	os.Setenv("TZ", "")
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.with/config.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	env.DefaultEnv()
	env.DefaultConfigFile(app)
	env.DefaultConfigReady()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(utils.ExitExecute)
	}
}
