package memba

import (
	"os"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func LoadConfig(c *Config) error {
	configFile := os.Getenv("MEMBA_CONF")
	if configFile == "" {
		configFile = "/usr/src/app/config.yml"
	}
	config.AddDriver(yaml.Driver)
	err := config.LoadFiles(configFile)
	if err != nil {
		err = config.LoadFiles("./config.yml")
		if err != nil {
			return err
		}
	}

	err = config.BindStruct("", c)
	if err != nil {
		return err
	}
	return nil
}
