package input

import "testing"

func TestInput_Validate(t *testing.T) {
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
			name: "valid input",
			fields: fields{
				Type: "file",
				Config: map[string]interface{}{
					"path": "/path/to/secret",
				},
			},
			wantErr: false,
		},
		{
			name: "no secret input type defined",
			fields: fields{
				Config: map[string]interface{}{
					"path": "/path/to/secret",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid secret input type defined",
			fields: fields{
				Type: "invalid-type",
				Config: map[string]interface{}{
					"path": "/path/to/secret",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid file input config",
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
			i := &Input{
				Type:   test.fields.Type,
				Config: test.fields.Config,
			}
			if err := i.Validate(); (err != nil) != test.wantErr {
				tt.Errorf("Input.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
