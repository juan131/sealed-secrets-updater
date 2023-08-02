package input

import (
	"errors"

	"github.com/juan131/sealed-secrets-updater/internal/utils"
)

const (
	TypeFile = "file"
)

// Input represents a secret input
type Input struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}

// Validate validates a secret input
func (i *Input) Validate() error {
	if i == nil {
		return errors.New("no input defined")
	}

	switch i.Type {
	case "":
		return errors.New("no input type defined")
	case TypeFile:
		var fileConfig FileConfig
		if err := utils.MapStringInterfaceToStruct(i.Config, &fileConfig); err != nil {
			return err
		}
		if fileConfig.Path == "" {
			return errors.New("no input path defined")
		}
	default:
		return errors.New("unsupported input type")
	}

	return nil
}
