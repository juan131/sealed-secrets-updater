package output

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"

	ssv1alpha1 "github.com/bitnami-labs/sealed-secrets/pkg/apis/sealedsecrets/v1alpha1"
	"github.com/bitnami-labs/sealed-secrets/pkg/kubeseal"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

// ApplyConfig represents the configuration for directly applying sealed secret
type ApplyConfig struct {
	CreateOnly bool `json:"createOnly"`
}

// UpdateSealedSecret updates a sealed secret by applying it directly to the cluster
func (c *ApplyConfig) UpdateSealedSecret(
	ctx context.Context,
	k8sConfig clientcmd.ClientConfig,
	pubKey *rsa.PublicKey,
	secret *v1.Secret,
) error {
	var inputSecret bytes.Buffer
	if err := json.NewEncoder(&inputSecret).Encode(secret); err != nil {
		return err
	}

	input := bytes.NewReader(inputSecret.Bytes())
	var buf bytes.Buffer
	output := bufio.NewWriter(&buf)
	if err := kubeseal.Seal(k8sConfig, "", input, output, scheme.Codecs, pubKey, ssv1alpha1.StrictScope, true, "", ""); err != nil {
		return err
	}

	output.Flush()
	restConf, err := k8sConfig.ClientConfig()
	if err != nil {
		return err
	}

	k8sClient, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return err
	}

	ssApiPath := "/apis/bitnami.com/v1alpha1/namespaces/" + secret.GetNamespace() + "/sealedsecrets/"
	err = k8sClient.RESTClient().
		Get().
		AbsPath(ssApiPath + secret.GetName()).
		Do(ctx).
		Error()
	if err == nil {
		if c.CreateOnly {
			klog.Info("=> Sealed secret already exists, skipping...")
			return nil
		}
		err = k8sClient.RESTClient().
			Patch(types.MergePatchType).
			AbsPath(ssApiPath + secret.GetName()).
			Body(buf.Bytes()).
			Do(ctx).
			Error()
		if err != nil {
			return fmt.Errorf("cannot update sealed secret: %v", err)
		}
	} else if k8serrors.IsNotFound(err) {
		err = k8sClient.RESTClient().
			Post().
			AbsPath(ssApiPath).
			Body(buf.Bytes()).
			Do(ctx).
			Error()
		if err != nil {
			return fmt.Errorf("cannot create sealed secret: %v", err)
		}
	} else {
		return fmt.Errorf("cannot get sealed secrets: %v", err)
	}

	return nil
}
