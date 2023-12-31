package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// MapStringInterfaceToStruct transforms a map[string]interface{} into a struct of the given type
func MapStringInterfaceToStruct(m map[string]interface{}, s interface{}) error {
	if reflect.TypeOf(s).Kind() != reflect.Ptr ||
		reflect.TypeOf(reflect.ValueOf(s).Elem()).Kind() != reflect.Struct {
		return fmt.Errorf("s must be a pointer to a struct")
	}

	buf, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(buf, s); err != nil {
		return err
	}

	return nil
}

// StructToMapStringInterface transforms a struct of the given type into a map[string]interface{}
func StructToMapStringInterface(s interface{}) (map[string]interface{}, error) {
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return nil, fmt.Errorf("s must be a struct")
	}

	buf, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var mapVal map[string]interface{}
	if err = json.Unmarshal(buf, &mapVal); err != nil {
		return nil, err
	}

	return mapVal, nil
}

// StringSliceContains is a helper function to detect whether a string slice contains a string or not
func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// StringSliceContainsAll is a helper function to detect whether a string slice contains all elements of a second string slice
func StringSliceContainsAll(s []string, ref []string) bool {
	allContained := true
	for _, a := range s {
		if !StringSliceContains(ref, a) {
			allContained = false
			break
		}
	}
	return allContained
}
