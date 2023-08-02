package config

import (
	"encoding/json"
	"errors"
	"os"
)

// New returns a new Config
func New(path string) (*Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, err
	}

	if config.KubesealConfig == nil {
		config.KubesealConfig = &KubesealConfig{
			ControllerName:      defaultControllerName,
			ControllerNamespace: defaultControllerNs,
		}
	} else if config.KubesealConfig.Certificate == "" {
		if config.KubesealConfig.ControllerName == "" {
			config.KubesealConfig.ControllerName = defaultControllerName
		}
		if config.KubesealConfig.ControllerNamespace == "" {
			config.KubesealConfig.ControllerNamespace = defaultControllerNs
		}
	}

	return &config, nil
}

// Validate validates the config
func (c *Config) Validate() error {
	if c == nil {
		return errors.New("no config defined")
	}

	if len(c.Secrets) == 0 {
		return errors.New("no secrets defined")
	}

	for _, secret := range c.Secrets {
		if secret.Name == "" {
			return errors.New("no secret name defined")
		}
		if err := secret.Input.Validate(); err != nil {
			return err
		}

		if err := secret.Output.Validate(); err != nil {
			return err
		}
	}

	return nil
}
