package awsconfig

import (
	"fmt"
	"io"
	"strings"

	"github.com/int128/ssoexec/awsconfig/parser"
)

func Parse(r io.Reader) (*Config, error) {
	sections, err := parser.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("could not parse config: %w", err)
	}
	var cfg Config
	for _, section := range sections {
		var p Profile
		if err := parseSectionHeader(&p, section.Header); err != nil {
			return nil, fmt.Errorf("invalid config: %s", err)
		}
		mapSectionValues(&p, section.Values)
		cfg.Profiles = append(cfg.Profiles, &p)
	}
	return &cfg, nil
}

func parseSectionHeader(p *Profile, header string) error {
	if header == "default" {
		p.Name = "default"
		return nil
	}
	if strings.HasPrefix(header, "profile ") {
		name := strings.TrimPrefix(header, "profile ")
		name = strings.TrimSpace(name)
		p.Name = name
		return nil
	}
	return fmt.Errorf("could not parse section header `%s`", header)
}

func mapSectionValues(p *Profile, values map[string]string) {
	for k, v := range values {
		if k == "region" {
			p.Region = v
			continue
		}
		if k == "sso_start_url" {
			p.SSOStartURL = v
			continue
		}
		if k == "sso_region" {
			p.SSORegion = v
			continue
		}
		if k == "sso_account_id" {
			p.SSOAccountID = v
			continue
		}
		if k == "sso_role_name" {
			p.SSORoleName = v
			continue
		}
	}
}
