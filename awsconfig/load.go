package awsconfig

import (
	"fmt"
	"os"
	"path/filepath"
)

func Load() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not determine home directory: %w", err)
	}
	awsConfigPath := filepath.Join(homeDir, ".aws", "config")
	f, err := os.Open(awsConfigPath)
	if err != nil {
		return nil, fmt.Errorf("could not open aws config: %w", err)
	}
	defer f.Close()
	return Parse(f)
}
