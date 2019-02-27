// Copyright © 2018 Sylvester La-Tunje. All rights reserved.

package cue

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/aws/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/slatunje/aws-with-access/pkg/env"
	"github.com/slatunje/aws-with-access/pkg/term"
	"github.com/slatunje/aws-with-access/pkg/utils"
	"github.com/spf13/viper"
)

const (
	escapedFlag  = "\\-"
	escapeSymbol = "\\"
)

// Credentials loads the required credentials
func Credentials(args []string) {

	if !viper.GetBool(env.Interactive) {
		WriteEnvironment(cfgByProfile())
		ExecuteCommand(args)
		return
	}

	var t = term.NewTerminal()

	WriteEnvironment(cfgByProfile())
	ExecuteCommand(args)

	var p = term.NewProcess(t)

	p.Start()
	p.Wait()

}

// WriteEnvironment
func WriteEnvironment(cfg aws.Config) {
	c, err := cfg.Credentials.Retrieve()
	if err != nil {
		if _, err = fmt.Fprintf(os.Stderr, "[error] failed find credentials, %v", err); err != nil {
			log.Println(err)
		}
		os.Exit(utils.ExitCredentialsFailure)
	}
	if err := os.Setenv("AWS_ACCESS_KEY_ID", c.AccessKeyID); err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("AWS_SECRET_ACCESS_KEY", c.SecretAccessKey); err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("AWS_SESSION_TOKEN", c.SessionToken); err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("AWS_SECURITY_TOKEN", c.SessionToken); err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("AWS_DEFAULT_PROFILE", viper.GetString(env.Profile)); err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("AWS_CONFIG_FILE", file("config")); err != nil {
		log.Fatalln(err)
	}
}

// ExecuteCommand
func ExecuteCommand(args []string) {
	var cmd *exec.Cmd
	if len(args) < 1 {
		return
	}
	cmd = exec.Command(args[0], flags(args[1:])...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		if _, err = fmt.Fprintln(os.Stderr, err); err != nil { // output errors to stderr
			log.Println(err)
		}
		os.Exit(utils.ExitCommandlineFailure)
	}
}

// cfgByProfile is responsible for returning the correct `aws.Config`
// first it establishes which credentials to load, then it loads the
// actual credentials. finally, if profile is assumed role scenario,
// then switch role by setting a session.
func cfgByProfile() (cfg aws.Config) {
	var role = viper.GetString(env.Profile)
	share, err := external.NewSharedConfig(role, files())
	if err != nil {
		log.Printf("[error] cannot load configuration due to, %v", err.Error())
		os.Exit(utils.ExitShareConfigFailure)
	}
	cfg, err = external.LoadDefaultAWSConfig(share)
	if err != nil {
		if _, err = fmt.Fprintf(os.Stderr, "[error] cannot load configuration due to, %v", err); err != nil {
			log.Println(err)
		}
		os.Exit(utils.ExitCredentialsFailure)
	}
	if !reflect.DeepEqual(share.AssumeRole, external.AssumeRoleConfig{}) {
		cfg.Credentials = credentials(cfg, share)
	}
	if !viper.GetBool(env.QuietMode) {
		log.Printf("[info] using aws profile: '%v' ", share.Profile)
	}
	return
}

// credentials returns a `aws.CredentialsProvider`
func credentials(cfg aws.Config, share external.SharedConfig) aws.CredentialsProvider {
	c := stscreds.NewAssumeRoleProvider(sts.New(cfg), share.AssumeRole.RoleARN)
	c.Duration = 15 * time.Minute
	c.RoleSessionName = fmt.Sprintf("%s%d", viper.GetString(env.RoleSession), time.Now().UTC().UnixNano())
	return c
}

// files return `[]string` of paths to search for files
func files() []string {
	return []string{file("config"), file("credentials"),}
}

// file sets the full path to a file
func file(filename string) string {
	return filepath.Join(utils.HomeDir(), ".aws", filename)
}

func flags(args []string) (data []string) {
	data = make([]string, 0)
	for _, a := range args {
		if strings.HasPrefix(a, escapedFlag) {
			a = strings.TrimPrefix(a, escapeSymbol)
		}
		data = append(data, a)
	}
	return
}
