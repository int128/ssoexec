package awsconfig

import "fmt"

func Find(name string) (*Profile, error) {
	config, err := Load()
	if err != nil {
		return nil, fmt.Errorf("could not load config: %w", err)
	}
	for _, profile := range config.Profiles {
		if profile.Name == name {
			return profile, nil
		}
	}
	return nil, nil
}
