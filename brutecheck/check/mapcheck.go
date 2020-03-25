package check

var __mapcheck = map[string]struct{}{
	"break":       struct{}{},
	"case":        struct{}{},
	"chan":        struct{}{},
	"const":       struct{}{},
	"continue":    struct{}{},
	"default":     struct{}{},
	"defer":       struct{}{},
	"else":        struct{}{},
	"fallthrough": struct{}{},
	"for":         struct{}{},
	"func":        struct{}{},
	"go":          struct{}{},
	"goto":        struct{}{},
	"if":          struct{}{},
	"import":      struct{}{},
	"interface":   struct{}{},
	"map":         struct{}{},
	"package":     struct{}{},
	"range":       struct{}{},
	"return":      struct{}{},
	"select":      struct{}{},
	"struct":      struct{}{},
	"switch":      struct{}{},
	"type":        struct{}{},
	"var":         struct{}{},
}

func MapCheck(name string) bool {
	_, ok := __mapcheck[name]
	return ok
}
