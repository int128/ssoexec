package ssocache

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(name string) (*Cache, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("could not open: %w", err)
	}
	defer f.Close()
	d := json.NewDecoder(f)
	var c Cache
	if err := d.Decode(&c); err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}
	return &c, nil
}
