package env

import (
	"log"

	"github.com/spf13/viper"
	"fmt"
	"github.com/slatunje/aws-with-access/pkg/utils"
	"time"
	"os"
)

// export AWS_ACCESS_KEY_ID=___ACCESSKEY___
// export AWS_SECRET_ACCESS_KEY=___SECRET___
// export AWS_SESSION_TOKEN=""
// export AWS_DEFAULT_REGION=eu-west-1
// export AWS_DEFAULT_OUTPUT=json
// export AWS_CA_BUNDLE=""
// export AWS_PROFILE=default
// export AWS_SHARED_CREDENTIALS_FILE=~/.aws/credentials
// export AWS_CONFIG_FILE=~/.aws/config
const (
	AccessKeyID           = "aws_access_key_id"
	AccessSecretKey       = "aws_secret_access_key"
	SessionToken          = "aws_session_token"
	SessionDuration       = "aws_session_duration"
	Region                = "aws_default_region"
	Output                = "aws_default_output"
	CaBundle              = "aws_ca_bundle"
	Profile               = "aws_profile"
	SharedCredentialsFile = "aws_shared_credentials_file"
	ConfigFile            = "aws_config_file"
	Account               = "aws_iam_account"
	Role                  = "aws_iam_role"
	RoleSession           = "aws_iam_role_name"
	RoleProfile           = "aws_iam_role_profile"
)

// default values
const (
	DefaultProfile = "default"
)

// default values
const (
	DefaultConfigFilename   = "config"   // default name of the configuration file
	DefaultConfigFileType   = "toml"     // default configuration file extension
	DefaultProjectConfigDir = ".private" // default configuration directory name
	DefaultProjectDir       = "../.."    // default path to project directory
)

var requiredKeys []string
//var requiredKeys = []string{AccessKeyID, AccessSecretKey}

func DefaultEnv() {
	viper.AutomaticEnv()
	viper.SetDefault(AccessKeyID, "")
	viper.SetDefault(AccessSecretKey, "")
	viper.SetDefault(SessionToken, nil)
	viper.SetDefault(SessionDuration, 15*time.Minute)
	viper.SetDefault(Region, "eu-west-1")
	viper.SetDefault(Output, "json")
	viper.SetDefault(CaBundle, nil)
	viper.SetDefault(Profile, DefaultProfile)
	viper.SetDefault(SharedCredentialsFile, "~/.aws/credentials")
	viper.SetDefault(ConfigFile, "~/.aws/config")
}

// DefaultConfigFile
func DefaultConfigFile(configFile string) {
	viper.SetConfigName(DefaultConfigFilename)
	viper.SetConfigType(DefaultConfigFileType)
	var configPaths = [...]string{
		utils.ProjectDefaultConfigDir(fmt.Sprintf(".%s/", configFile)),
		utils.ProjectConfigDir(fmt.Sprintf("%s/%s/", DefaultProjectDir, DefaultProjectConfigDir)),
		utils.ProjectDir(DefaultProjectDir),
	}
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	if err := viper.ReadInConfig(); err != nil {
		logMsg := "[warning] unable to load configuration file '%s' from any of the following paths: '%s' due to %v"
		log.Printf(logMsg, DefaultConfigFilename, configFile, err)
	}
	log.Printf("[info] using configuration from path: %s", viper.ConfigFileUsed())
}

// DefaultConfigReady
func DefaultConfigReady() {
	ensureDefaultSettings()
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

// ensureDefaultSettings
func ensureDefaultSettings() {}
