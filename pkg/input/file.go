package input

import (
	"encoding/json"
	"os"
	"path"
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
	if err := json.NewDecoder(inputFile).Decode(&secretsData); err != nil {
		return nil, err
	}

	return secretsData, nil
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
