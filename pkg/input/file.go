package input

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

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

	secretsData, err := getSecretsDataFromFile(inputFile, c.getFormat())
	if err != nil {
		return nil, err
	}

	return secretsData, nil
}

// getFormat returns the sealed secret format based on the file extension
func (c *FileConfig) getFormat() string {
	switch strings.ToLower(path.Ext(c.Path)) {
	case ".yaml", ".yml":
		return "yaml"
	case ".csv":
		return "csv"
	default:
		return "json"
	}
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

// getSecretsDataFromFile returns the secrets data from the given file path
func getSecretsDataFromFile(r io.Reader, format string) (map[string]string, error) {
	var secretsData map[string]string
	switch format {
	case "yaml":
		if err := yaml.NewDecoder(r).Decode(&secretsData); err != nil {
			return nil, err
		}
	case "csv":
		reader := csv.NewReader(r)
		reader.FieldsPerRecord = 2
		records, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		secretsData = make(map[string]string, len(records))
		for _, record := range records {
			secretsData[record[0]] = record[1]
		}
	case "json":
		if err := json.NewDecoder(r).Decode(&secretsData); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported input file format: %s", format)
	}

	return secretsData, nil
}
