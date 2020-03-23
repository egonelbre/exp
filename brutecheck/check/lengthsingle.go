package check

func Len1Ch(name string) bool {
	switch len(name) {
	case 2: // go, if
		switch name[0] {
		case 'g':
			return name == "go"
		case 'i':
			return name == "if"
		}
		return false
	case 3: // for, map, var
		switch name[0] {
		case 'f':
			return name == "for"
		case 'm':
			return name == "map"
		case 'v':
			return name == "var"
		}
		return false
	case 4: // case, chan, else, func, goto, type
		switch name[1] {
		case 'a':
			return name == "case"
		case 'h':
			return name == "chan"
		case 'l':
			return name == "else"
		case 'u':
			return name == "func"
		case 'o':
			return name == "goto"
		case 'y':
			return name == "type"
		}
		return false
	case 5: // break, const, defer, range
		switch name[0] {
		case 'b':
			return name == "break"
		case 'c':
			return name == "const"
		case 'd':
			return name == "defer"
		case 'r':
			return name == "range"
		}
		return false
	case 6: // import, return, select, struct, switch
		switch name[2] {
		case 'p':
			return name == "import"
		case 't':
			return name == "return"
		case 'l':
			return name == "select"
		case 'r':
			return name == "struct"
		case 'i':
			return name == "switch"
		}
		return false
	case 7: // default, package
		switch name[0] {
		case 'd':
			return name == "default"
		case 'p':
			return name == "package"
		}
		return false
	case 8: // continue
		return name == "continue"
	case 9: // interface
		return name == "interface"
	}
	return false
}
