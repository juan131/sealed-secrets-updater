package config

import (
	"github.com/juan131/sealed-secrets-updater/pkg/input"
	"github.com/juan131/sealed-secrets-updater/pkg/output"
)

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
