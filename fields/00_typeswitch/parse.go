package main

import (
	"encoding/json"
	"errors"
	"io"
)

//gistsnip:start:parse
func ParseFields(r io.Reader) (map[string]Field, error) {
	var config struct {
		Fields []struct {
			Name       string
			Type       string
			Val        interface{}
			Multiplier interface{}
		}
	}

	err := json.NewDecoder(r).Decode(&config)
	if err != nil {
		return nil, err
	}

	fields := map[string]Field{}

	for _, jsonField := range config.Fields {
		val, valok := jsonField.Val.(float64)
		if !valok {
			return nil, errors.New("unsupported type " + jsonField.Type)
		}

		var field Field
		switch jsonField.Type {
		case "uint":
			field = &Uint{
				ID:    jsonField.Name,
				Value: uint64(val),
			}
		case "float":
			field = &Float{
				ID:    jsonField.Name,
				Value: float64(val),
			}
		default:
			return nil, errors.New("unsupported type " + jsonField.Type)
		}

		fields[field.Name()] = field
	}

	return fields, nil
}

//gistsnip:end:parse
