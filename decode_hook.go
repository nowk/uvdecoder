package uvdecoder

import "reflect"
import "github.com/mitchellh/mapstructure"

// Hook is a decode hook for mapstructure to help properly map url.Values to a
// struct.
// Because url.Values is a map[string][]string, mapstructure skips any non slice
// string conversions.
// This requires WeaklyTypedInput = true
func Hook(f, t reflect.Kind, data interface{}) (interface{}, error) {
	if reflect.Slice != t {
		switch f {
		case reflect.Slice:
			dataVal := reflect.ValueOf(data)
			return dataVal.Interface().([]string)[0], nil
		}
	}

	return data, nil
}

// Config returns a preconfigured DecoderConfig for Hook
func Config(d interface{}) *mapstructure.DecoderConfig {
	return &mapstructure.DecoderConfig{
		DecodeHook:       Hook,
		Result:           d,
		WeaklyTypedInput: true,
	}
}
