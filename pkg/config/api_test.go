package config

import (
	_ "embed"
	"errors"
	"io/fs"
	"testing"
)

func Test_validSchema(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "valid path to valid config",
			args: args{
				path: "testdata/valid-config.json",
			},
			wantErr: nil,
		},
		{
			name: "valid path to invalid config (missing input)",
			args: args{
				path: "testdata/invalid-config-missing-input.json",
			},
			wantErr: errSchemaValidation,
		},
		{
			name: "valid path to invalid config (missing output)",
			args: args{
				path: "testdata/invalid-config-missing-output.json",
			},
			wantErr: errSchemaValidation,
		},
		{
			name: "valid path to invalid config (no path in file input)",
			args: args{
				path: "testdata/invalid-config-file-input-no-path.json",
			},
			wantErr: errSchemaValidation,
		},
		{
			name: "valid path to invalid config (no path in file output)",
			args: args{
				path: "testdata/invalid-config-file-output-no-path.json",
			},
			wantErr: errSchemaValidation,
		},
		{
			name: "invalid path",
			args: args{
				path: "not a valid path",
			},
			wantErr: fs.ErrNotExist,
		},
	}
	t.Parallel()
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			err := validSchema(test.args.path)
			if err != nil && !errors.Is(err, test.wantErr) {
				tt.Errorf("validSchema() got error: %v, want error: %v", err, test.wantErr)
			}
		})
	}
}
