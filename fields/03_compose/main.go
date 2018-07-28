package main

import (
	"encoding/json"
	"fmt"
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

// using spec types
type Spec interface {
	FieldName() string
	Assign(field *Field) error
}

type Float struct {
	Name  string
	Value *float64
}

func (s Float) FieldName() string { return s.Name }
func (s Float) Assign(field *Field) error {
	uv, ok := field.Val.(float64)
	if !ok || field.Type != "float" {
		return fmt.Errorf("expected float, got %T and %v", field.Val, field.Type)
	}
	*s.Value = uv
	return nil
}

type Uint struct {
	Name  string
	Value *uint
}

func (s Uint) FieldName() string { return s.Name }
func (s Uint) Assign(field *Field) error {
	uv, ok := field.Val.(float64)
	if !ok || field.Type != "uint" {
		return fmt.Errorf("expected uint, got %T and %v", field.Val, field.Type)
	}
	*s.Value = uint(uv)
	return nil
}

func (config *Config) Scan(specs ...Spec) error {
	for _, r := range specs {
		name := r.FieldName()
		field, err := config.findField(name)
		if err != nil {
			return err
		}

		err = r.Assign(field)
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

func main() {
	var err error

	config := &Config{}
	err = json.Unmarshal([]byte(exampleInput), config)
	dieIf(err)

	var (
		alpha float64
		gamma float64
		beta  uint
	)

	err = config.Scan(
		Float{"Alpha", &alpha},
		Float{"Gamma", &gamma},
		Uint{"Beta", &beta},
	)
	dieIf(err)

	// use example as usual
	fmt.Println(alpha + gamma)
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
