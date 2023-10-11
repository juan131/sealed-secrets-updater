package input

import (
	"encoding/json"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// FileConfig represents the configuration for using a file as secret input
type FileConfig struct {
	Path     string `json:"path"`
	Relative bool   `json:"relative"`
}

// GetSecretsData returns the secrets data from the input file
func (c *FileConfig) GetSecretsData() (map[string]string, error) {
	inputFile, err := os.Open(c.getFilePath())
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	var secretsData map[string]string
	switch c.getFormat() {
	case "yaml":
		if err := yaml.NewDecoder(inputFile).Decode(&secretsData); err != nil {
			return nil, err
		}
	default:
		if err := json.NewDecoder(inputFile).Decode(&secretsData); err != nil {
			return nil, err
		}
	}

	return secretsData, nil
}

// getFormat returns the sealed secret format based on the file extension
func (c *FileConfig) getFormat() string {
	if ext := path.Ext(c.Path); ext == ".yaml" || ext == ".yml" {
		return "yaml"
	}

	return "json"
}

// getFilePath returns the input file path
func (c *FileConfig) getFilePath() string {
	if !c.Relative {
		wd, err := os.Getwd()
		if err != nil {
			return c.Path
		}

		return path.Join(wd, c.Path)
	}

	return c.Path
}
