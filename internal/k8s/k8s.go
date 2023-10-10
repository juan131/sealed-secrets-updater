package k8s

import (
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// NewClientConfig returns a new k8s client config
func NewClientConfig() clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	return clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{}, os.Stdout)
}

// NewMetadata returns a new k8s object metadata
func NewMetadata(name, namespace string, annotations, labels map[string]string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:        name,
		Namespace:   namespace,
		Annotations: annotations,
		Labels:      labels,
	}
}

// NewSecret returns a new k8s secret
func NewSecret(metadata metav1.ObjectMeta, data map[string]string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metadata,
		Data:       encodeSecretData(data),
	}
}

// encodeSecretData encodes the secret data using base64
func encodeSecretData(data map[string]string) map[string][]byte {
	encodedData := map[string][]byte{}
	for k, v := range data {
		encodedData[k] = []byte(v)
	}

	return encodedData
}
