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
	"github.com/davecgh/go-spew/spew"
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

	//spew.Dump(viper.AllSettings())
	//os.Exit(utils.ExitOnDebug)



	// spew.Dump(viper.AllSettings())

	// store

	cfgWriteEnvironment(cfg())

	// spew.Dump(viper.AllSettings())

	// execute

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

// cfg is responsible for returning the correct `aws.Config`
func cfg() (cfg aws.Config) {

	// TODO: must have profile here

	DefaultRoleProfile = "vit-prod"
	fmt.Println("PROFILE", DefaultRoleProfile,)


	sharecfg, err := external.NewSharedConfig(DefaultRoleProfile, []string{
		file("config"), file("credentials"),
	})
	if err != nil {
		log.Println("\n\n\nSHARE CONFIG LOAD ERROR ===========", err.Error())
		os.Exit(-200)
	}





	//os.Exit(utils.ExitOnDebug)

	// TODO: use this instead ... LoadSharedConfig
	
	cfg, err = external.LoadDefaultAWSConfig(
		external.WithRegion(viper.GetString(env.Region)),
		external.WithSharedConfigFiles([]string{file("credentials")}),
		external.WithSharedConfigProfile(sharecfg.AssumeRole.Source.Profile),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] cannot load configuration due to, %v", err)
		os.Exit(utils.ExitCredentialsFailure)
	}
	cm := fmt.Sprintf("[info] using aws profile: '%v' ", profile())
	if profile() != env.DefaultProfile {

		spew.Dump(cfg.Credentials)
		
		//sharecfg, err := external.NewSharedConfig(DefaultRoleProfile, []string{
		//		file("config"), file("credentials"),
		//})
		//if err != nil {
		//	log.Println("\n\n\nSHARE CONFIG LOAD ERROR ===========", err.Error())
		//	os.Exit(-200)
		//}

		spew.Dump(sharecfg.AssumeRole)
		spew.Dump(sharecfg.Credentials)
		spew.Dump(sharecfg.Region)
		spew.Dump(sharecfg.Profile)

		fmt.Println("===DONE===")

		//os.Exit(utils.ExitOnDebug)


		cfg.Credentials = credentialWithShare(cfg, sharecfg)
		//cfg.Credentials = credentials(cfg)
		cm += fmt.Sprintf("but will override using '%v' for this current session.", role())
	}
	log.Println(cm)
	return
}

// profile will resolve and return the correct profile
func profile() string {
	return DefaultRoleProfile
  	//return viper.GetString(env.Profile)
}

//// credentials returns a `aws.CredentialsProvider`
//func credentials(cfg aws.Config) aws.CredentialsProvider {
//	DefaultRoleARN = fmt.Sprintf("arn:aws:iam::%s:role/%s", account(), role())
//	c := stscreds.NewAssumeRoleProvider(sts.New(cfg), DefaultRoleARN)
//	c.Duration = DefaultDuration
//	c.RoleSessionName = DefaultRoleSessionName
//	return c
//}

// credentials returns a `aws.CredentialsProvider`
func credentialWithShare(cfg aws.Config, sharecfg external.SharedConfig) aws.CredentialsProvider {
	c := stscreds.NewAssumeRoleProvider(sts.New(cfg), sharecfg.AssumeRole.RoleARN)
	c.Duration = DefaultDuration
	c.RoleSessionName = DefaultRoleSessionName
	return c
}

func role() string {
	return "TODO"
	//return "js_roles_js-aap-vit-prod_admin" // return viper.GetString(env.Role) // TODO: resolve
}

//func account() string {
//	return "883300774050" // return viper.GetString(env.Account) // TODO: resolve
//}
func file(filename string) string {
	return filepath.Join(utils.HomeDir(), ".aws", filename)
}