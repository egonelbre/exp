package main

import "fmt"

// using spec types
type Field interface {
	Name() string
	Assign(typ string, value interface{}) error
}

type Float struct {
	FieldName string
	Value     *float64
}

func (s Float) Name() string { return s.FieldName }
func (s Float) Assign(typ string, val interface{}) error {
	uv, ok := val.(float64)
	if !ok || typ != "float" {
		return fmt.Errorf("expected float, got %T and %v", val, typ)
	}
	*s.Value = uv
	return nil
}

type Uint struct {
	FieldName string
	Value     *uint
}

func (s Uint) Name() string { return s.FieldName }
func (s Uint) Assign(typ string, val interface{}) error {
	uv, ok := val.(float64)
	if !ok || typ != "uint" {
		return fmt.Errorf("expected uint, got %T and %v", val, typ)
	}
	*s.Value = uint(uv)
	return nil
}
