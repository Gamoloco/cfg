// Package cfg is an utility library to help with the parsing of json configuration files
package cfg

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strings"
)

var uncheckableTypes = []reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
	reflect.Bool,
}

func isNumericOrBool(typ reflect.Kind) bool {
	for _, tp := range uncheckableTypes {
		if tp == typ {
			return true
		}
	}
	return false
}

func check(typ reflect.Type, val reflect.Value, paths []string) error {
	// Loop through type fields
	for i := 0; i < typ.NumField(); i++ {
		kd := typ.Field(i).Type.Kind()
		// If type of field is struct recursively check it
		if kd == reflect.Struct {
			if err := check(val.Field(i).Type(), val.Field(i), append(paths, typ.Field(i).Name)); err != nil {
				return err
			}
		} else if isNumericOrBool(typ.Field(i).Type.Kind()) {
			// As Golang has default values for numeric and boolean types we can't check if there are missing
			continue
		} else {
			// If field is missing and not tagged as optional
			if val.Field(i).Len() == 0 && typ.Field(i).Tag.Get("cfg") != "optional" {
				if len(paths) > 0 {
					return errors.New("Missing required field: " + strings.Join(paths, ".") + "." + typ.Field(i).Name)
				}
				return errors.New("Missing required field: " + typ.Field(i).Name)
			}
		}
	}
	return nil
}

// Load load configuration file at path in container and check if there is missing fields
func Load(path string, container interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&container); err != nil {
		return err
	}

	if err := check(reflect.TypeOf(container).Elem(), reflect.ValueOf(container).Elem(), []string{}); err != nil {
		return err
	}

	return nil
}
