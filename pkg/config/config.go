package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/bitnami-labs/sealed-secrets/pkg/kubeseal"
	"k8s.io/klog/v2"
)

// New returns a new Config given a path to a config file
func New(path string) (*Config, error) {
	// Ensure config follows the schema before loading it
	if err := validSchema(path); err != nil {
		return nil, err
	}

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

	for _, secret := range config.Secrets {
		if secret.Metadata == nil {
			secret.Metadata = &Metadata{}
		}
	}

	return &config, nil
}

// Validate validates the config with extra validations that are not covered by the JSON schema
func (c *Config) Validate() error {
	if c == nil || c.KubesealConfig == nil {
		return errors.New("no config defined")
	}

	if c.KubesealConfig.Certificate != "" {
		if c.KubesealConfig.ControllerName != "" || c.KubesealConfig.ControllerNamespace != "" {
			klog.Warning("controller name and namespace will be ignored since a certificate was provided")
		}

		if err := isValidCertificate(c.KubesealConfig.Certificate); err != nil {
			return fmt.Errorf("invalid certificate: %w", err)
		}
	}

	if len(c.Secrets) == 0 {
		return errors.New("no secrets defined")
	}

	return nil
}

// isValidCertificate checks if a certificate is valid
func isValidCertificate(filenameOrURI string) error {
	var certFile io.ReadCloser
	if _, err := os.Stat(filenameOrURI); err != nil {
		if _, err := url.ParseRequestURI(filenameOrURI); err != nil {
			return err
		}

		// TODO: download certificate from URI
		return nil
	} else {
		certFile, err = os.Open(filenameOrURI)
		if err != nil {
			return err
		}
	}

	defer certFile.Close()
	if _, err := kubeseal.ParseKey(certFile); err != nil {
		return err
	}

	return nil
}
