// Copyright © 2018 Sylvester La-Tunje. All rights reserved.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/slatunje/aws-with-access/pkg/cue"
	"github.com/slatunje/aws-with-access/pkg/env"
	"github.com/slatunje/aws-with-access/pkg/utils"
	"github.com/spf13/cobra"
)

const (
	app = "with"
)

var (
	profile     string
	interactive bool
	quiet       bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: fmt.Sprintf(app),
	Short: fmt.Sprintf("=> %s makes it easy to obtain temporary AWS credentials", app),
	Long: fmt.Sprintf(`
Description:
  %s makes it easier to obtain temporary AWS credentials through 'AssumeRole'.
`, app),
	Example: fmt.Sprintf(`  %s --interactive --profile vit 
  %s --quiet --profile vit aws s3 ls
  %s --quiet --profile vit aws s3 ls "\--recursive"`, app, app, app),
	Run: func(cmd *cobra.Command, args []string) {
		cue.Credentials(args)
	},
}

// init is called in alphabetic order within this package
func init() {
	if err := os.Setenv("TZ", ""); err != nil {
		log.Fatalln(err)
	}
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().
		StringVarP(&profile, "profile", "p", "default", "set profile name.")
	rootCmd.PersistentFlags().
		BoolVarP(&interactive, "interactive", "i", false, "enter into interactive mode.")
	rootCmd.PersistentFlags().
		BoolVarP(&quiet, "quiet", "q", false, "suppress normal output quiet mode.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	env.DefaultEnv()
	env.DefaultProfile(env.ConfigOptions{Profile: profile, Interactive: interactive, QuietMode: quiet})
	env.DefaultConfigReady()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(utils.ExitExecute)
	}
}
