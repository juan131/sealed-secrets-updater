package updater

import (
	"errors"
	"fmt"

	"github.com/juan131/sealed-secrets-updater/internal/utils"
	"github.com/juan131/sealed-secrets-updater/pkg/config"
)

// Filter is used to filter the secrets to update
type Filter struct {
	OnlySecrets []string
	SkipSecrets []string
}

// Validate validates the filter
func (f Filter) Validate(secrets []*config.Secret) error {
	for _, secret := range f.OnlySecrets {
		if utils.StringSliceContains(f.SkipSecrets, secret) {
			return fmt.Errorf("secret \"%s\" cannot be in both --only-secrets and --skip-secrets lists", secret)
		}
	}

	existingSecrets := make([]string, 0, len(secrets))
	for _, secret := range secrets {
		existingSecrets = append(existingSecrets, secret.Name)
	}

	if !utils.StringSliceContainsAll(f.OnlySecrets, existingSecrets) {
		return errors.New("some secrets in --only-secrets list do not exist")
	}

	return nil
}
