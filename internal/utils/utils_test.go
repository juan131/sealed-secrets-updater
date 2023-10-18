package utils

import (
	"reflect"
	"testing"
)

type testStruct struct {
	One   string
	Two   int
	Three bool
	Four  []testSubStruct
	Five  string
}

type testSubStruct struct {
	Six string
}

type jsonStruct struct {
	One   string        `json:"uno"`
	Two   int           `json:"dos"`
	Three bool          `json:"tres"`
	Four  jsonSubStruct `json:"cuatro"`
	Five  string        `json:"cinco"`
}

type jsonSubStruct struct {
	Six string `json:"seis"`
}

func Test_MapStringInterfaceToStruct(t *testing.T) {
	type args struct {
		m map[string]interface{}
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Successful transformation",
			args: args{
				m: map[string]interface{}{
					"One":   "one",
					"Two":   2,
					"Three": true,
					"Four": []map[string]interface{}{{
						"Six": "six",
					}},
					"Five": "five",
				},
				s: &testStruct{},
			},
			want: &testStruct{
				One:   "one",
				Two:   2,
				Three: true,
				Four: []testSubStruct{{
					Six: "six",
				}},
				Five: "five",
			},
			wantErr: false,
		},
		{
			name: "Successful transformation (JSON)",
			args: args{
				m: map[string]interface{}{
					"uno":  "one",
					"dos":  2,
					"tres": true,
					"cuatro": map[string]interface{}{
						"seis": "six",
					},
					"cinco": "five",
				},
				s: &jsonStruct{},
			},
			want: &jsonStruct{
				One:   "one",
				Two:   2,
				Three: true,
				Four: jsonSubStruct{
					Six: "six",
				},
				Five: "five",
			},
			wantErr: false,
		},
	}
	t.Parallel()
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			if err := MapStringInterfaceToStruct(test.args.m, test.args.s); (err != nil) != test.wantErr {
				tt.Errorf("MapStringInterfaceToStruct() error = %v, wantErr %v", err, test.wantErr)
			}
			if !reflect.DeepEqual(test.args.s, test.want) {
				tt.Errorf("MapStringInterfaceToStruct() = %v, want %v", test.args.s, test.want)
			}
		})
	}
}

func TestStructToMapStringInterface(t *testing.T) {
	type args struct {
		s interface{}
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Successful transformation",
			args: args{
				s: testStruct{
					One:   "one",
					Two:   2,
					Three: true,
					Four:  nil,
					Five:  "five",
				},
			},
			want: map[string]interface{}{
				"One":   "one",
				"Two":   float64(2),
				"Three": true,
				"Four":  nil,
				"Five":  "five",
			},
			wantErr: false,
		},
		{
			name: "Successful transformation (JSON)",
			args: args{
				s: jsonStruct{
					One:   "one",
					Two:   2,
					Three: true,
					Four: jsonSubStruct{
						Six: "six",
					},
					Five: "five",
				},
			},
			want: map[string]interface{}{
				"uno":  "one",
				"dos":  float64(2),
				"tres": true,
				"cuatro": map[string]interface{}{
					"seis": "six",
				},
				"cinco": "five",
			},
			wantErr: false,
		},
	}
	t.Parallel()
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			got, err := StructToMapStringInterface(test.args.s)
			if (err != nil) != test.wantErr {
				tt.Errorf("StructToMapStringInterface() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				tt.Errorf("StructToMapStringInterface() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestStringSliceContains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Successful search",
			args: args{
				s: []string{"one", "two", "three"},
				e: "two",
			},
			want: true,
		},
		{
			name: "Unsuccessful search",
			args: args{
				s: []string{"one", "two", "three"},
				e: "four",
			},
			want: false,
		},
	}
	t.Parallel()
	for _, testToRun := range tests {
		test := testToRun
		t.Run(test.name, func(tt *testing.T) {
			tt.Parallel()
			if got := StringSliceContains(test.args.s, test.args.e); got != test.want {
				tt.Errorf("StringSliceContains() = %v, want %v", got, test.want)
			}
		})
	}
}
