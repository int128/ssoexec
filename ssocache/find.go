package ssocache

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Find(startURL, region string) (*Cache, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not determine home directory: %w", err)
	}
	ssoCachePath := filepath.Join(homeDir, ".aws", "sso", "cache", "*.json")
	g, err := filepath.Glob(ssoCachePath)
	if err != nil {
		return nil, fmt.Errorf("could not find sso cache directory: %w", err)
	}
	for _, name := range g {
		c, err := Load(name)
		if err != nil {
			log.Printf("invalid sso cache: %s", err)
			continue
		}
		if c.StartURL == startURL && c.Region == region {
			return c, nil
		}
	}
	return nil, nil
}
