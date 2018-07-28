package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

var exampleInput = `{
	"Fields": [
		{
			"Name": "Alpha",
			"Type": "float",
			"Val": 15,
			"Multiplier": 0.5
		},
		{
			"Name": "Gamma",
			"Type": "float",
			"Val": 10,
			"Multiplier": 0.5
		},
		{
			"Name": "Beta",
			"Type": "uint",
			"Val": 10,
			"Multiplier": 10
		}
	]
}`

// json handling part

type Config struct {
	Fields []Field
}

type Field struct {
	Name       string
	Type       string
	Val        interface{}
	Multiplier interface{}
}

// assigning values to the struct

func (config *Config) Scan(r interface{}) error {
	// check that r is a pointer to some struct
	rv := reflect.ValueOf(r)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected pointer to a struct, got %+T", r)
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

func (config *Config) findField(name string) (*Field, error) {
	for i := 0; i < len(config.Fields); i++ {
		field := &config.Fields[i]
		if strings.EqualFold(field.Name, name) {
			return field, nil
		}
	}
	return nil, fmt.Errorf("unable to find field " + name)
}

func (config *Config) assignField(field *Field, p interface{}) error {
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
		return fmt.Errorf("unhandled field type %T", r)
	}
	return nil
}

func main() {
	var err error

	config := &Config{}
	err = json.Unmarshal([]byte(exampleInput), config)
	dieIf(err)

	var example struct {
		Alpha float64
		Gamma float64
		Beta  uint
	}

	err = config.Scan(&example)
	dieIf(err)

	// use example as usual
	fmt.Println(example.Alpha + example.Gamma)
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
