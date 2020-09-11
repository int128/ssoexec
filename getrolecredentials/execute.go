package getrolecredentials

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/int128/ssoexec/awsconfig"
	"github.com/int128/ssoexec/ssocache"
)

type output struct {
	RoleCredentials RoleCredentials `json:"roleCredentials"`
}

type RoleCredentials struct {
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	SessionToken    string `json:"sessionToken"`
	Expiration      int    `json:"expiration"`
}

func Execute(profile awsconfig.Profile, cache ssocache.Cache) (*RoleCredentials, error) {
	var b bytes.Buffer
	c := exec.Command("aws", "sso", "get-role-credentials",
		"--role-name", profile.SSORoleName,
		"--region", profile.SSORegion,
		"--account-id", profile.SSOAccountID,
		"--access-token", cache.AccessToken,
	)
	c.Stdin = os.Stdin
	c.Stdout = &b
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return nil, fmt.Errorf("could not run aws sso get-role-credentials command: %w", err)
	}
	var o output
	if err := json.Unmarshal(b.Bytes(), &o); err != nil {
		return nil, fmt.Errorf("could not parse output of aws sso get-role-credentials command: %w", err)
	}
	return &o.RoleCredentials, nil
}
