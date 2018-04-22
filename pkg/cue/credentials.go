package cue

import (
	"time"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"path/filepath"
	"os"
	"github.com/slatunje/aws-role/pkg/utils"
	"github.com/aws/aws-sdk-go-v2/aws/stscreds"
	"os/exec"
)

const (
	DefaultProfile  = "jsmgmt"
	DefaultRole     = "arn:aws:iam::883300774050:role/js_roles_js-aap-vit-prod_admin"
	DefaultRegion   = "eu-west-1"
	DefaultRoleName = "AssumeRoleSession"
)

var (
	DefaultBucket          = "vit-prod-data"
	DefaultDuration        = 15 * time.Minute
	DefaultRoleSessionName = fmt.Sprintf("%s%d", DefaultRoleName, time.Now().UTC().UnixNano())
)

var (
	args []string // os.Args[0] is _this_ program's name
	cmd  *exec.Cmd
)

// Credentials loads the required credentials
func Credentials() {

	cfgInitEnvironment()

	cfg, err := external.LoadDefaultAWSConfig(
		external.WithRegion(DefaultRegion),
		external.WithSharedConfigFiles([]string{filepath.Join(utils.UserHomeDir(), ".aws", "credentials")}),
		external.WithSharedConfigProfile(DefaultProfile),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load configuration, %v", err)
		os.Exit(1)
	}

	credentials := stscreds.NewAssumeRoleProvider(sts.New(cfg), DefaultRole)
	credentials.Duration = DefaultDuration
	credentials.RoleSessionName = DefaultRoleSessionName

	cfg.Credentials = credentials
	cfgWriteEnvironment(cfg)

	args = os.Args[1:]
	if len(args) < 1 {
		return
	}

	cmd = exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // output errors to stderr
		os.Exit(1) // exit with non-zero status to indicate command failure
	}

}

func cfgInitEnvironment() {
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")
	os.Unsetenv("AWS_SESSION_TOKEN")
}

func cfgWriteEnvironment(cfg aws.Config) {
	c, err := cfg.Credentials.Retrieve()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed find credentials, %v", err)
		os.Exit(1)
	}
	os.Setenv("AWS_ACCESS_KEY_ID", c.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", c.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", c.SessionToken)
}

