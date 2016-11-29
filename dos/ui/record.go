package ui

import "reflect"

func ReflectLoadWidgets(record Record) []Widget {
	if record, ok := record.(interface {
		Widgets() []Widget
	}); ok {
		return record.Widgets()
	}

	widgets := []Widget{}

	v := reflect.ValueOf(record)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// get address if not one
		if field.Kind() != reflect.Ptr {
			field = field.Addr()
		}

		if widget, ok := field.Interface().(Widget); ok {
			widgets = append(widgets, widget)
		}
	}

	return widgets
}
