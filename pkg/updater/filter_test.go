package updater

import (
	"testing"

	"github.com/juan131/sealed-secrets-updater/pkg/config"
)

func TestFilter_Validate(t *testing.T) {
	type fields struct {
		OnlySecrets []string
		SkipSecrets []string
	}
	type args struct {
		secrets []*config.Secret
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Valid filters",
			fields{
				OnlySecrets: []string{"secret1"},
				SkipSecrets: []string{"secret2"},
			},
			args{
				secrets: []*config.Secret{{
					Name: "secret1",
				}, {
					Name: "secret2",
				}},
			},
			false,
		},
		{
			"OnlySecrets and SkipSecrets are empty",
			fields{},
			args{
				secrets: []*config.Secret{{
					Name: "secret1",
				}},
			},
			false,
		},
		{
			"OnlySecrets is set and some secrets do not exist",
			fields{
				OnlySecrets: []string{"secret1"},
			},
			args{
				secrets: []*config.Secret{{}},
			},
			true,
		},
		{
			"OnlySecrets & SkipSecrets are set and some secrets are in both lists",
			fields{
				OnlySecrets: []string{"secret1"},
				SkipSecrets: []string{"secret1"},
			},
			args{
				secrets: []*config.Secret{{
					Name: "secret1",
				}},
			},
			true,
		},
	}
	t.Parallel()
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			f := Filter{
				OnlySecrets: test.fields.OnlySecrets,
				SkipSecrets: test.fields.SkipSecrets,
			}
			if err := f.Validate(test.args.secrets); (err != nil) != test.wantErr {
				t.Errorf("Filter.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
