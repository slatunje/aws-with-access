package env

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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
				"aws_profile":                    "PROFILE",
				"aws_iam_role_name":              "WITH.",
				"aws_shell_interactive":          false,
				"aws_access_key_id_previous":     "",
				"aws_secret_access_key_previous": "",
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
	if err := os.Setenv("AWS_ACCESS_KEY_ID", ""); err != nil {
		log.Println(err)
	}
	if err := os.Setenv("AWS_SECRET_ACCESS_KEY", ""); err != nil {
		log.Println(err)
	}
	if err := os.Setenv("AWS_SESSION_TOKEN", ""); err != nil {
		log.Println(err)
	}
	if err := os.Setenv("AWS_DEFAULT_REGION", ""); err != nil {
		log.Println(err)
	}
	if err := os.Setenv("AWS_DEFAULT_OUTPUT", ""); err != nil {
		log.Println(err)
	}
	if err := os.Setenv("AWS_CA_BUNDLE", ""); err != nil {
		log.Println(err)
	}
	if err := os.Setenv("AWS_PROFILE", ""); err != nil {
		log.Println(err)
	}
	if err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", ""); err != nil {
		log.Println(err)
	}
	if err := os.Setenv("AWS_CONFIG_FILE", ""); err != nil {
		log.Println(err)
	}
}

func initConfig() {
	DefaultEnv()
	DefaultProfile("PROFILE", false)
	DefaultConfigReady()
}
