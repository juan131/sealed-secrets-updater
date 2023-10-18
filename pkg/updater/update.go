package updater

import (
	"context"
	"crypto/rsa"
	"errors"

	"github.com/bitnami-labs/sealed-secrets/pkg/kubeseal"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	"github.com/juan131/sealed-secrets-updater/internal/k8s"
	"github.com/juan131/sealed-secrets-updater/internal/utils"
	"github.com/juan131/sealed-secrets-updater/pkg/config"
	"github.com/juan131/sealed-secrets-updater/pkg/input"
	"github.com/juan131/sealed-secrets-updater/pkg/output"
)

// UpdateSealedSecrets iterates over all the secrets in the secrets manager and
// updates the sealed secrets manifests
func UpdateSealedSecrets(ctx context.Context, config *config.Config, skipSecrets []string) error {
	k8sConfig := k8s.NewClientConfig()
	klog.Info("Obtaining public key...")
	pubKey, err := getPublicKey(ctx, k8sConfig, config.KubesealConfig)
	if err != nil {
		return err
	}

	klog.Info("Updating sealed secrets...")
	for _, secret := range config.Secrets {
		if utils.StringSliceContains(skipSecrets, secret.Name) {
			klog.Infof("=> Skipping sealed secret \"%s\"", secret.Name)
			continue
		}

		klog.Infof("=> Updating sealed secret \"%s\"", secret.Name)
		var secretsData map[string]string
		switch secret.Input.Type {
		case input.TypeFile:
			var fileConfig input.FileConfig
			err := utils.MapStringInterfaceToStruct(secret.Input.Config, &fileConfig)
			if err != nil {
				return err
			}

			secretsData, err = fileConfig.GetSecretsData()
			if err != nil {
				return err
			}
		}

		metadata := k8s.NewMetadata(secret.Name, secret.Namespace, secret.Metadata.Annotations, secret.Metadata.Labels)
		switch secret.Output.Type {
		case output.TypeApply:
			var applyConfig output.ApplyConfig
			if err := utils.MapStringInterfaceToStruct(secret.Output.Config, &applyConfig); err != nil {
				return err
			}

			if err := applyConfig.UpdateSealedSecret(ctx, k8sConfig, pubKey, k8s.NewSecret(metadata, secretsData)); err != nil {
				return err
			}
		case output.TypeFile:
			var fileConfig output.FileConfig
			if err := utils.MapStringInterfaceToStruct(secret.Output.Config, &fileConfig); err != nil {
				return err
			}

			if err := fileConfig.UpdateSealedSecret(ctx, k8sConfig, pubKey, k8s.NewSecret(metadata, secretsData)); err != nil {
				return err
			}
		}
	}

	klog.Info("Sealed secrets updated successfully!")
	return nil
}

// getPublicKey returns the public key of the sealed secrets controller
func getPublicKey(
	ctx context.Context,
	k8sConfig clientcmd.ClientConfig,
	kubesealConfig *config.KubesealConfig,
) (*rsa.PublicKey, error) {
	certFile, err := kubeseal.OpenCert(ctx, k8sConfig, kubesealConfig.ControllerNamespace, kubesealConfig.ControllerName, kubesealConfig.Certificate)
	if err != nil {
		// Ignore kubeseal error message, it's too verbose
		return nil, errors.New("unable to obtain public key")
	}
	defer certFile.Close()

	pubKey, err := kubeseal.ParseKey(certFile)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}
