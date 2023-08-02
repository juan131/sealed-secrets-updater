package output

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"os"
	"path"

	ssv1alpha1 "github.com/bitnami-labs/sealed-secrets/pkg/apis/sealedsecrets/v1alpha1"
	"github.com/bitnami-labs/sealed-secrets/pkg/kubeseal"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

// FileConfig represents the configuration for using a file as sealed secret output
type FileConfig struct {
	Path       string `json:"path"`
	Relative   bool   `json:"relative"`
	CreateOnly bool   `json:"createOnly"`
}

// UpdateSealedSecret updates a sealed secret and writes it to a file
func (c *FileConfig) UpdateSealedSecret(
	ctx context.Context,
	k8sConfig clientcmd.ClientConfig,
	pubKey *rsa.PublicKey,
	secret *v1.Secret,
) error {
	if c.CreateOnly {
		if _, err := os.Stat(c.getFilePath()); err == nil {
			klog.Info("=> Sealed secret already exists, skipping...")
			return nil
		}
	}
	output, err := os.OpenFile(c.getFilePath(), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer output.Close()

	var inputSecret bytes.Buffer
	err = json.NewEncoder(&inputSecret).Encode(secret)
	if err != nil {
		return err
	}

	input := bytes.NewReader(inputSecret.Bytes())
	err = kubeseal.Seal(k8sConfig, c.getFormat(), input, output, scheme.Codecs, pubKey, ssv1alpha1.StrictScope, true, "", "")
	if err != nil {
		return err
	}

	return nil
}

// getFormat returns the sealed secret format based on the file extension
func (c *FileConfig) getFormat() string {
	if ext := path.Ext(c.Path); ext == ".yaml" || ext == ".yml" {
		return "yaml"
	}

	return "json"
}

// getFilePath returns the output file path
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
