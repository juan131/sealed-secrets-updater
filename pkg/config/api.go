package config

import (
	_ "embed"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/juan131/sealed-secrets-updater/pkg/input"
	"github.com/juan131/sealed-secrets-updater/pkg/output"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed config.schema.json
var schema string

const (
	defaultControllerName string = "sealed-secrets-controller"
	defaultControllerNs   string = "kube-system"
)

// Config represents the configuration for the sealed-secrets-updater
type Config struct {
	KubesealConfig *KubesealConfig `json:"kubeseal"`
	Secrets        []*Secret       `json:"secrets"`
}

// KubesealConfig represents the configuration for the kubeseal command
type KubesealConfig struct {
	ControllerName      string `json:"controllerName"`
	ControllerNamespace string `json:"controllerNamespace"`
	Certificate         string `json:"certificate,omitempty"`
}

// Secret represents the configuration for a secret
type Secret struct {
	Name      string         `json:"name"`
	Namespace string         `json:"namespace"`
	Input     *input.Input   `json:"input"`
	Output    *output.Output `json:"output"`
	Metadata  *Metadata      `json:"metadata"`
}

// Metadata represents a secret metadata
type Metadata struct {
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

var errSchemaValidation = errors.New("schema validation failed")

// validSchema ensure a config file is valid against the schema
func validSchema(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	schemaLoader := gojsonschema.NewStringLoader(schema)
	documentLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", absPath))
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		// Wrap every error in a single error
		wrappedErr := errSchemaValidation
		for _, err := range result.Errors() {
			wrappedErr = fmt.Errorf("%w\n- %s", wrappedErr, err)
		}

		return wrappedErr
	}

	return nil
}
