package config

import (
	"fmt"
	"reflect"
)

func (cfg Config) String() string {
	t := reflect.TypeOf(cfg)
	v := reflect.ValueOf(cfg)

	if t.Kind() != reflect.Struct {
		return "Config is invalid struct"
	}

	returnString := ""

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		returnString += fmt.Sprintf("%s: %s\n", field.Name, value)
	}

	return returnString
}
