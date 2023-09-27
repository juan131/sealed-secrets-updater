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
				KubesealConfig: &KubesealConfig{
					defaultControllerName,
					defaultControllerNs,
					"",
				},
				Secrets: test.fields.Secrets,
			}
			if err := c.Validate(); (err != nil) != test.wantErr {
				tt.Errorf("Config.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func Test_isValidCertificate(t *testing.T) {
	t.Parallel()
	type args struct {
		filenameOrURI string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid path to valid certificate",
			args: args{
				filenameOrURI: "testdata/valid.pem",
			},
			wantErr: false,
		},
		{
			name: "invalid value",
			args: args{
				filenameOrURI: "not a valid path nor URI",
			},
			wantErr: true,
		},
		{
			name: "valid path to invalid certificate",
			args: args{
				filenameOrURI: "testdata/invalid.pem",
			},
			wantErr: true,
		},
	}
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			if err := isValidCertificate(test.args.filenameOrURI); (err != nil) != test.wantErr {
				tt.Errorf("isValidCertificate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
