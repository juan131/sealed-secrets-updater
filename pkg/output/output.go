package output

import (
	"errors"

	"github.com/juan131/sealed-secrets-updater/internal/utils"
)

const (
	TypeApply = "apply"
	TypeFile  = "file"
)

// Output represents a sealed secret output
type Output struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}

// Validate validates a secret output
func (o *Output) Validate() error {
	if o == nil {
		return errors.New("no output defined")
	}

	switch o.Type {
	case "":
		return errors.New("no output type defined")
	case TypeApply:
		return nil
	case TypeFile:
		var fileConfig FileConfig
		if err := utils.MapStringInterfaceToStruct(o.Config, &fileConfig); err != nil {
			return err
		}
		if fileConfig.Path == "" {
			return errors.New("no output file path defined")
		}
	default:
		return errors.New("unsupported output type")
	}

	return nil
}
