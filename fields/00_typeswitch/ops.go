package main

func Add(a, b Field) Field {
	if aerr, ok := a.(*Error); ok {
		return aerr
	}
	if berr, ok := b.(*Error); ok {
		return berr
	}

	switch x := a.(type) {
	case *Float:
		if y, ok := b.(*Float); ok {
			return x.Add(y)
		} else {
			return &Error{"add type-mismatch"}
		}
	case *Uint:
		if y, ok := b.(*Uint); ok {
			return x.Add(y)
		} else {
			return &Error{"add type-mismatch"}
		}
	default:
		return &Error{"unhandled types"}
	}
}

func Sub(a, b Field) Field {
	if aerr, ok := a.(*Error); ok {
		return aerr
	}
	if berr, ok := b.(*Error); ok {
		return berr
	}

	switch x := a.(type) {
	case *Float:
		if y, ok := b.(*Float); ok {
			return x.Sub(y)
		} else {
			return &Error{"add type-mismatch"}
		}
	case *Uint:
		if y, ok := b.(*Uint); ok {
			return x.Sub(y)
		} else {
			return &Error{"add type-mismatch"}
		}
	default:
		return &Error{"unhandled types"}
	}
}
