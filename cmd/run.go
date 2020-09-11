package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/int128/ssoexec/awsconfig"
	"github.com/int128/ssoexec/getrolecredentials"
	"github.com/int128/ssoexec/ssocache"
)

func Run(osArgs []string) error {
	awsProfileName := os.Getenv("AWS_PROFILE")
	if awsProfileName == "" {
		awsProfileName = "default"
	}
	awsProfile, err := awsconfig.Find(awsProfileName)
	if err != nil {
		return fmt.Errorf("could not find the aws profile: %w", err)
	}
	if awsProfile == nil {
		return fmt.Errorf("could not find aws profile %s", awsProfileName)
	}
	ssoCache, err := ssocache.Find(awsProfile.SSOStartURL, awsProfile.SSORegion)
	if err != nil {
		return fmt.Errorf("could not parse aws config: %w", err)
	}
	if ssoCache == nil {
		return fmt.Errorf("could not find sso cache for profile %+v", awsProfile)
	}
	roleCredentials, err := getrolecredentials.Execute(*awsProfile, *ssoCache)
	if err != nil {
		return fmt.Errorf("could not run aws sso: %w", err)
	}

	if len(osArgs) == 1 {
		fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", roleCredentials.AccessKeyId)
		fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", roleCredentials.SecretAccessKey)
		fmt.Printf("export AWS_SESSION_TOKEN=%s\n", roleCredentials.SessionToken)
		return nil
	}

	c := exec.Command(osArgs[1], osArgs[2:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Env = os.Environ()
	c.Env = append(c.Env,
		fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", roleCredentials.AccessKeyId),
		fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", roleCredentials.SecretAccessKey),
		fmt.Sprintf("AWS_SESSION_TOKEN=%s", roleCredentials.SessionToken),
	)
	if err := c.Run(); err != nil {
		return fmt.Errorf("could not run the command: %w", err)
	}
	return nil
}
