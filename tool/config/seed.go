package tconfig

import (
	"encoding/json"
	"io"
	"os"
)

type SeedConfig struct {
}

var SeedConfigPath string = "config/seed.json"
var SeedConfigTestPath string = "config/seed_test.json"
var SeedConfigLocalPath string = "config/seed_local.json"

var SeedConfigInstance *SeedConfig

func GetSeedConfigInstance() *SeedConfig {
	if SeedConfigInstance == nil {
		SeedConfigInstance = &SeedConfig{}
		SeedConfigInstance.ReadConfig()
	}
	return SeedConfigInstance
}

func (c *SeedConfig) ReadConfig() {
	configPath := SeedConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = SeedConfigLocalPath
		}
	} else {
		if os.Getenv("TEST") != "" {
			if os.Getenv("TEST") == "true" {
				configPath = SeedConfigTestPath
			}
		}
	}

	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	byteValue, _ := io.ReadAll(f)

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		panic(err)
	}
}
