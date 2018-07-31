package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func Unmarshal(rd io.Reader, data interface{}) error {
	var config jsonConfig
	err := json.NewDecoder(rd).Decode(&config)
	if err != nil {
		return err
	}

	return config.Scan(data)
}

type jsonConfig struct {
	Fields []jsonField
}

type jsonField struct {
	Name       string
	Type       string
	Val        interface{}
	Multiplier interface{}
}

func (config *jsonConfig) Scan(r interface{}) error {
	// check that r is a pointer to some struct
	rv := reflect.ValueOf(r)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected pointer to a struct, got %T", r)
	}

	s := rv.Elem()
	t := s.Type()

	// iterate over all struct fields
	for i, n := 0, s.NumField(); i < n; i++ {
		resultField := s.Field(i)

		// find the corresponding field from config
		field, err := config.findField(t.Field(i).Name)
		if err != nil {
			return err
		}

		// assign field value to the struct field
		err = config.assignField(field, resultField.Addr().Interface())
		if err != nil {
			return err
		}
	}

	return nil
}

func (config *jsonConfig) findField(name string) (*jsonField, error) {
	for i := 0; i < len(config.Fields); i++ {
		field := &config.Fields[i]
		if strings.EqualFold(field.Name, name) {
			return field, nil
		}
	}
	return nil, fmt.Errorf("unable to find field " + name)
}

func (config *jsonConfig) assignField(field *jsonField, p interface{}) error {
	// p is a pointer to struct field
	switch p := p.(type) {
	case *uint:
		uv, ok := field.Val.(float64)
		if !ok || field.Type != "uint" {
			return fmt.Errorf("expected uint, got %T and %v", field.Val, field.Type)
		}
		*p = uint(uv)
	case *float64:
		uv, ok := field.Val.(float64)
		if !ok || field.Type != "float" {
			return fmt.Errorf("expected float, got %T and %v", field.Val, field.Type)
		}
		*p = uv
	default:
		return fmt.Errorf("unhandled field type %T", p)
	}
	return nil
}
