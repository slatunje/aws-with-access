package cue

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/slatunje/aws-with-access/pkg/utils"
	"github.com/spf13/viper"
	"github.com/slatunje/aws-with-access/pkg/env"
	"github.com/aws/aws-sdk-go-v2/aws/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var (
	DefaultRoleARN         string
	DefaultDuration        = 15 * time.Minute
	DefaultRoleSession     = viper.GetString(env.RoleSession)
	DefaultRoleProfile     = viper.GetString(env.RoleProfile)
	DefaultRoleSessionName = fmt.Sprintf("%s%d", DefaultRoleSession, time.Now().UTC().UnixNano())
)

var (
	args []string
	cmd  *exec.Cmd
)

// Credentials loads the required credentials
func Credentials() {

	var profile = viper.GetString(env.Profile)

	cfg, err := external.LoadDefaultAWSConfig(
		external.WithRegion(viper.GetString(env.Region)),
		external.WithSharedConfigFiles([]string{filepath.Join(utils.HomeDir(), ".aws", "credentials")}),
		external.WithSharedConfigProfile(DefaultRoleProfile),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] cannot load configuration due to, %v", err)
		os.Exit(utils.ExitCredentialsFailure)
	}

	cm := fmt.Sprintf("[info] using aws profile: '%v' ", profile)
	
	role := viper.GetString(env.Role)
	if profile != "default" {
		DefaultRoleARN = fmt.Sprintf("arn:aws:iam::%s:role/%s", viper.GetString(env.Account), role)
		credentials := stscreds.NewAssumeRoleProvider(sts.New(cfg), DefaultRoleARN)
		credentials.Duration = DefaultDuration
		credentials.RoleSessionName = DefaultRoleSessionName
		cfg.Credentials = credentials
		cm += fmt.Sprintf("but will override using '%v' for this current session.", role)
	}

	log.Println(cm)

	// spew.Dump(viper.AllSettings())

	cfgWriteEnvironment(cfg)

	// spew.Dump(viper.AllSettings())

	cfgExecuteCommand()

}

func cfgWriteEnvironment(cfg aws.Config) {
	c, err := cfg.Credentials.Retrieve()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] failed find credentials, %v", err)
		os.Exit(1)
	}
	os.Setenv("AWS_ACCESS_KEY_ID", c.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", c.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", c.SessionToken)
}

func cfgExecuteCommand() {
	args = os.Args[1:]
	if len(args) < 1 {
		return
	}
	cmd = exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // output errors to stderr
		os.Exit(utils.ExitCommandlineFailure)
	}
}
