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
				"aws_profile": "AKIA",
				"aws_iam_role_name": "A-I-R-N",
			},
		},
	}
	for i, tc := range c {
		t.Run(fmt.Sprintf("#%d %s", i, tc.name), func(t *testing.T) {
			 a.Equal(viper.AllSettings(), tc.expect)
		})
	}
}

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
	DefaultProfile("AKIA")
	DefaultConfigReady()
}
