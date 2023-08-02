package output

import "testing"

func TestOutput_Validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		Type   string
		Config map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid output",
			fields: fields{
				Type: "file",
				Config: map[string]interface{}{
					"path": "/path/to/secret",
				},
			},
			wantErr: false,
		},
		{
			name: "no secret output type defined",
			fields: fields{
				Config: map[string]interface{}{
					"path": "/path/to/manifest",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid secret output type defined",
			fields: fields{
				Type: "invalid-type",
				Config: map[string]interface{}{
					"path": "/path/to/manifest",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid file output config",
			fields: fields{
				Type: "file",
				Config: map[string]interface{}{
					"foo": "bar",
				},
			},
			wantErr: true,
		},
	}
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			o := &Output{
				Type:   test.fields.Type,
				Config: test.fields.Config,
			}
			if err := o.Validate(); (err != nil) != test.wantErr {
				tt.Errorf("Output.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
