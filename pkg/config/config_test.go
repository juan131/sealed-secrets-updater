package config

import (
	"testing"

	"github.com/juan131/sealed-secrets-updater/pkg/input"
	"github.com/juan131/sealed-secrets-updater/pkg/output"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		Secrets []*Secret
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid config",
			fields: fields{
				Secrets: []*Secret{{
					Name:      "secret-name",
					Namespace: "default",
					Input: &input.Input{
						Type: "file",
						Config: map[string]interface{}{
							"path": "/path/to/secret",
						},
					},
					Output: &output.Output{
						Type: "file",
						Config: map[string]interface{}{
							"path": "/path/to/manifest",
						},
					},
				}},
			},
			wantErr: false,
		},
		{
			name:    "no secrets defined",
			fields:  fields{},
			wantErr: true,
		},
		{
			name: "no secret name defined",
			fields: fields{
				Secrets: []*Secret{{
					Namespace: "default",
					Input: &input.Input{
						Type: "file",
						Config: map[string]interface{}{
							"path": "/path/to/secret",
						},
					},
					Output: &output.Output{
						Type: "file",
						Config: map[string]interface{}{
							"path": "/path/to/manifest",
						},
					},
				}},
			},
			wantErr: true,
		},
		{
			name: "no secret input defined",
			fields: fields{
				Secrets: []*Secret{{
					Name:      "secret-name",
					Namespace: "default",
					Output: &output.Output{
						Type: "file",
						Config: map[string]interface{}{
							"path": "/path/to/manifest",
						},
					},
				}},
			},
			wantErr: true,
		},
		{
			name: "no secret output defined",
			fields: fields{
				Secrets: []*Secret{{
					Name:      "secret-name",
					Namespace: "default",
					Input: &input.Input{
						Type: "file",
						Config: map[string]interface{}{
							"path": "/path/to/secret",
						},
					},
				}},
			},
			wantErr: true,
		},
	}
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			c := &Config{
				Secrets: test.fields.Secrets,
			}
			if err := c.Validate(); (err != nil) != test.wantErr {
				tt.Errorf("Config.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
