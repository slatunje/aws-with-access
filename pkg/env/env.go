// Copyright © 2018 Sylvester La-Tunje. All rights reserved.

package env

import (
	"log"

	"github.com/spf13/viper"
	"github.com/slatunje/aws-with-access/pkg/utils"
	"os"
)

// https://docs.aws.amazon.com/cli/latest/userguide/cli-environment.html
const (
	AccessKeyID             = "aws_access_key_id"
	AccessSecretKey         = "aws_secret_access_key"
	SessionToken            = "aws_session_token"
	SessionDuration         = "aws_session_duration"
	Region                  = "aws_default_region"
	Output                  = "aws_default_output"
	CaBundle                = "aws_ca_bundle"
	SharedCredentialsFile   = "aws_shared_credentials_file"
	ConfigFile              = "aws_config_file"
	Profile                 = "aws_profile"
	RoleSession             = "aws_iam_role_name"
	Interactive             = "aws_shell_interactive"
	PreviousAccessKeyID     = "aws_access_key_id_previous"
	PreviousAccessSecretKey = "aws_secret_access_key_previous"
)

// requiredKeys defines required keys
var requiredKeys = []string{Profile}

// DefaultEnv
func DefaultEnv() {
	viper.AutomaticEnv()
	//viper.SetDefault(AccessKeyID, "")
	//viper.SetDefault(AccessSecretKey, "")
	//viper.SetDefault(SessionToken, nil)
	//viper.SetDefault(SessionDuration, 15*time.Minute)
	//viper.SetDefault(Region, "eu-west-1")
	//viper.SetDefault(Output, "json")
	//viper.SetDefault(CaBundle, nil)
	//viper.SetDefault(SharedCredentialsFile, "~/.aws/credentials")
	//viper.SetDefault(ConfigFile, "~/.aws/config")
	viper.SetDefault(RoleSession, "WITH.")
}

// DefaultConfigFile
func DefaultProfile(profile string, interactive bool) {
	viper.Set(Profile, profile)
	viper.Set(Interactive, interactive)
}

// DefaultConfigReady
func DefaultConfigReady() {
	storePreviousKeys()
	if missingRequiredKeys() {
		os.Exit(utils.ExitRequireKeys)
	}
}

// missingRequiredKeys
func missingRequiredKeys() bool {
	var missing []string
	for _, k := range requiredKeys {
		if v := viper.GetString(k); len(v) == 0 {
			missing = append(missing, k)
		}
	}
	if len(missing) != 0 {
		log.Printf("missing required keys %s.", utils.ToUpper(missing))
		return true
	}
	return false
}

// storePreviousKeys
func storePreviousKeys() {
	viper.Set(PreviousAccessKeyID, viper.GetString(AccessKeyID))
	viper.Set(PreviousAccessSecretKey, viper.GetString(AccessSecretKey))
}