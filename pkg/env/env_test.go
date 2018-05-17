package env

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/spf13/viper"
	"fmt"
)

func TestMain(m *testing.M) {
	setupMain()
	os.Exit(m.Run())
}

func TestOnInitiation(t *testing.T) {

	initConfig()

	a := assert.New(t)

	c := []struct {
		name   string
		expect map[string]interface{}
	}{
		{
			"Default: No Environment set",
			map[string]interface{}{
				"aws_access_key_id": "AKIA",
			},
		},
	}
	for i, tc := range c {
		t.Run(fmt.Sprintf("#%d %s", i, tc.name), func(t *testing.T) {
			 a.Equal(viper.AllSettings(), tc.expect)
		})
	}

}

// export AWS_ACCESS_KEY_ID=___ACCESSKEY___
// export AWS_SECRET_ACCESS_KEY=___SECRETACCESSKEY___
// export AWS_SESSION_TOKEN=""
// export AWS_DEFAULT_REGION=eu-west-1
// export AWS_DEFAULT_OUTPUT=json
// export AWS_CA_BUNDLE=""
// export AWS_PROFILE=default
// export AWS_SHARED_CREDENTIALS_FILE=~/.aws/credentials
// export AWS_CONFIG_FILE=~/.aws/config

//func setupMain(key, secret, token, region, output, ca, profile, credFile, confFile string) {
//	os.Setenv("AWS_ACCESS_KEY_ID", key)
//	os.Setenv("AWS_SECRET_ACCESS_KEY", secret)
//	os.Setenv("AWS_SESSION_TOKEN", token)
//	os.Setenv("AWS_DEFAULT_REGION", region)
//	os.Setenv("AWS_DEFAULT_OUTPUT", output)
//	os.Setenv("AWS_CA_BUNDLE", ca)
//	os.Setenv("AWS_PROFILE", profile)
//	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", credFile)
//	os.Setenv("AWS_CONFIG_FILE", confFile)
//}

func setupMain() {
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")
	os.Setenv("AWS_SESSION_TOKEN", "")
	os.Setenv("AWS_DEFAULT_REGION", "")
	os.Setenv("AWS_DEFAULT_OUTPUT", "")
	os.Setenv("AWS_CA_BUNDLE", "")
	os.Setenv("AWS_PROFILE", "")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "")
	os.Setenv("AWS_CONFIG_FILE", "")
}

func initConfig() {
	DefaultEnv()
	DefaultConfigFile("with")
	DefaultConfigReady()
}
