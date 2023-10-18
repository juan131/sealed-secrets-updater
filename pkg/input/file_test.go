package input

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_getSecretsDataFromFile(t *testing.T) {
	type args struct {
		r      io.Reader
		format string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "valid yaml input",
			args: args{
				r: strings.NewReader(`key1: value1
key2: value2`),
				format: "yaml",
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name: "invalid yaml input",
			args: args{
				r: strings.NewReader(`key1: value1
key2 value2`),
				format: "yaml",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid csv input",
			args: args{
				r: strings.NewReader(`key1,value1
key2,value2`),
				format: "csv",
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name: "invalid csv input",
			args: args{
				r: strings.NewReader(`key1,value1
key2,value2,value3`),
				format: "csv",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid json input",
			args: args{
				r:      strings.NewReader(`{"key1":"value1","key2":"value2"}`),
				format: "json",
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name: "invalid json input",
			args: args{
				r:      strings.NewReader(`{"key1":"value1","key2":"value2"`),
				format: "json",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsupported properties format",
			args: args{
				r: strings.NewReader(`key1 = value1
key2 = value2`),
				format: "properties",
			},
			want:    nil,
			wantErr: true,
		},
	}
	t.Parallel()
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			got, err := getSecretsDataFromFile(test.args.r, test.args.format)
			if (err != nil) != test.wantErr {
				tt.Errorf("getSecretsDataFromFile() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				tt.Errorf("getSecretsDataFromFile() = %v, want %v", got, test.want)
			}
		})
	}
}
