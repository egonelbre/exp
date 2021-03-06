package check

func TwoHashTableAlt_XorXor_Shift015(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 1) ^ (name[1] << 5))
	if h > 249 {
		return false
	}
	return [...]string{
		0:   "switch",
		16:  "if",
		27:  "interface",
		35:  "const",
		42:  "goto",
		44:  "go",
		46:  "continue",
		47:  "for",
		64:  "select",
		66:  "return",
		78:  "else",
		96:  "struct",
		104: "func",
		109: "defer",
		111: "default",
		116: "import",
		129: "break",
		193: "range",
		194: "chan",
		199: "package",
		204: "type",
		207: "var",
		226: "case",
		231: "fallthrough",
		249: "map",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift016(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 1) ^ (name[1] << 6))
	if h > 224 {
		return false
	}
	return [...]string{
		3:   "const",
		10:  "goto",
		12:  "go",
		14:  "continue",
		15:  "for",
		32:  "switch",
		65:  "break",
		80:  "if",
		91:  "interface",
		130: "case",
		135: "fallthrough",
		136: "func",
		141: "defer",
		143: "default",
		148: "import",
		153: "map",
		160: "select",
		161: "range",
		162: "return",
		167: "package",
		172: "type",
		175: "var",
		194: "chan",
		206: "else",
		224: "struct",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift021(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 2) ^ (name[1] << 1))
	if h > 120 {
		return false
	}
	return [...]string{
		0:   "select",
		4:   "return",
		5:   "package",
		15:  "range",
		25:  "var",
		34:  "struct",
		36:  "switch",
		38:  "type",
		64:  "go",
		69:  "for",
		70:  "goto",
		72:  "else",
		74:  "case",
		81:  "fallthrough",
		87:  "const",
		88:  "chan",
		90:  "continue",
		93:  "default",
		95:  "defer",
		105: "break",
		106: "if",
		113: "interface",
		117: "map",
		118: "func",
		120: "import",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift024(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 2) ^ (name[1] << 4))
	if h > 221 {
		return false
	}
	return [...]string{
		8:   "chan",
		68:  "type",
		77:  "interface",
		80:  "else",
		104: "goto",
		107: "for",
		110: "go",
		114: "import",
		116: "continue",
		121: "const",
		131: "fallthrough",
		138: "struct",
		152: "case",
		154: "select",
		158: "return",
		167: "map",
		173: "break",
		186: "switch",
		197: "defer",
		198: "if",
		199: "default",
		203: "var",
		204: "func",
		215: "package",
		221: "range",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift025(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 2) ^ (name[1] << 5))
	if h > 251 {
		return false
	}
	return [...]string{
		2:   "import",
		16:  "else",
		42:  "switch",
		53:  "defer",
		55:  "default",
		60:  "func",
		74:  "struct",
		100: "continue",
		102: "if",
		105: "const",
		106: "select",
		109: "interface",
		110: "return",
		120: "goto",
		123: "for",
		126: "go",
		136: "chan",
		151: "map",
		168: "case",
		179: "fallthrough",
		205: "break",
		231: "package",
		237: "range",
		244: "type",
		251: "var",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift026(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 2) ^ (name[1] << 6))
	if h > 247 {
		return false
	}
	return [...]string{
		10:  "switch",
		13:  "break",
		38:  "if",
		45:  "interface",
		68:  "continue",
		73:  "const",
		88:  "goto",
		91:  "for",
		94:  "go",
		135: "package",
		136: "chan",
		138: "select",
		141: "range",
		142: "return",
		144: "else",
		148: "type",
		155: "var",
		200: "case",
		202: "struct",
		211: "fallthrough",
		213: "defer",
		215: "default",
		220: "func",
		226: "import",
		247: "map",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift032(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) ^ (name[1] << 2))
	if h > 250 {
		return false
	}
	return [...]string{
		2:   "return",
		3:   "package",
		10:  "select",
		17:  "range",
		55:  "var",
		64:  "type",
		66:  "switch",
		78:  "struct",
		128: "goto",
		134: "go",
		143: "for",
		152: "case",
		156: "else",
		161: "const",
		172: "continue",
		177: "defer",
		179: "default",
		188: "chan",
		191: "fallthrough",
		210: "if",
		221: "break",
		224: "func",
		239: "map",
		249: "interface",
		250: "import",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift034(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) ^ (name[1] << 4))
	if h > 238 {
		return false
	}
	return [...]string{
		12:  "case",
		42:  "if",
		43:  "fallthrough",
		52:  "type",
		53:  "break",
		100: "func",
		117: "defer",
		119: "default",
		123: "map",
		133: "range",
		151: "package",
		156: "chan",
		158: "import",
		161: "interface",
		163: "var",
		195: "for",
		198: "return",
		202: "go",
		204: "goto",
		206: "select",
		222: "struct",
		224: "continue",
		236: "else",
		237: "const",
		238: "switch",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift035(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) ^ (name[1] << 5))
	if h > 253 {
		return false
	}
	return [...]string{
		27:  "fallthrough",
		28:  "chan",
		30:  "struct",
		54:  "return",
		60:  "case",
		62:  "select",
		75:  "map",
		85:  "break",
		126: "switch",
		129: "interface",
		132: "type",
		133: "defer",
		135: "default",
		138: "if",
		147: "var",
		148: "func",
		167: "package",
		172: "else",
		181: "range",
		211: "for",
		218: "go",
		220: "goto",
		238: "import",
		240: "continue",
		253: "const",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) ^ (name[1] << 1))
	if h > 252 {
		return false
	}
	return [...]string{
		17:  "map",
		69:  "interface",
		76:  "import",
		94:  "if",
		140: "else",
		141: "default",
		142: "func",
		143: "defer",
		161: "var",
		169: "fallthrough",
		170: "goto",
		172: "go",
		182: "type",
		189: "for",
		193: "break",
		197: "package",
		216: "switch",
		222: "struct",
		228: "chan",
		230: "continue",
		231: "range",
		235: "const",
		236: "return",
		246: "case",
		252: "select",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift043(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) ^ (name[1] << 3))
	if h > 254 {
		return false
	}
	return [...]string{
		10:  "go",
		12:  "goto",
		14:  "return",
		15:  "package",
		27:  "for",
		30:  "select",
		45:  "range",
		52:  "else",
		60:  "case",
		64:  "continue",
		77:  "const",
		99:  "fallthrough",
		107: "var",
		109: "defer",
		111: "default",
		116: "chan",
		140: "type",
		142: "switch",
		150: "struct",
		162: "if",
		181: "break",
		204: "func",
		219: "map",
		233: "interface",
		254: "import",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift045(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) ^ (name[1] << 5))
	if h > 243 {
		return false
	}
	return [...]string{
		5:   "range",
		20:  "case",
		39:  "package",
		52:  "chan",
		54:  "import",
		67:  "var",
		75:  "fallthrough",
		82:  "if",
		89:  "interface",
		100: "type",
		101: "break",
		131: "for",
		134: "return",
		146: "go",
		148: "goto",
		150: "select",
		182: "struct",
		196: "func",
		212: "else",
		213: "const",
		214: "switch",
		216: "continue",
		229: "defer",
		231: "default",
		243: "map",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift046(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) ^ (name[1] << 6))
	if h > 248 {
		return false
	}
	return [...]string{
		4:   "type",
		5:   "defer",
		7:   "default",
		18:  "if",
		25:  "interface",
		35:  "var",
		36:  "func",
		43:  "fallthrough",
		52:  "chan",
		54:  "struct",
		71:  "package",
		84:  "else",
		101: "range",
		102: "return",
		116: "case",
		118: "select",
		147: "map",
		163: "for",
		165: "break",
		178: "go",
		180: "goto",
		214: "import",
		245: "const",
		246: "switch",
		248: "continue",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 5) ^ (name[1] << 0))
	if h > 253 {
		return false
	}
	return [...]string{
		3:   "select",
		5:   "case",
		7:   "continue",
		10:  "const",
		12:  "chan",
		17:  "switch",
		18:  "struct",
		35:  "return",
		36:  "range",
		55:  "break",
		68:  "if",
		71:  "interface",
		75:  "import",
		102: "package",
		139: "goto",
		141: "go",
		162: "var",
		170: "fallthrough",
		172: "for",
		177: "func",
		194: "map",
		200: "else",
		224: "defer",
		226: "default",
		253: "type",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 5) ^ (name[1] << 1))
	if h > 252 {
		return false
	}
	return [...]string{
		1:   "var",
		9:   "fallthrough",
		29:  "for",
		46:  "func",
		58:  "goto",
		60:  "go",
		77:  "default",
		79:  "defer",
		97:  "map",
		118: "type",
		124: "else",
		135: "range",
		136: "switch",
		140: "return",
		142: "struct",
		161: "break",
		166: "case",
		172: "select",
		180: "chan",
		182: "continue",
		187: "const",
		197: "package",
		238: "if",
		245: "interface",
		252: "import",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 6) ^ (name[1] << 0))
	if h > 247 {
		return false
	}
	return [...]string{
		34:  "map",
		36:  "if",
		39:  "interface",
		40:  "else",
		43:  "import",
		96:  "defer",
		98:  "default",
		102: "package",
		125: "type",
		163: "select",
		165: "case",
		167: "continue",
		170: "const",
		171: "goto",
		172: "chan",
		173: "go",
		177: "switch",
		178: "struct",
		226: "var",
		227: "return",
		228: "range",
		234: "fallthrough",
		236: "for",
		241: "func",
		247: "break",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 6) ^ (name[1] << 2))
	if h > 244 {
		return false
	}
	return [...]string{
		1:   "range",
		7:   "var",
		15:  "fallthrough",
		18:  "return",
		22:  "struct",
		26:  "switch",
		63:  "for",
		64:  "case",
		77:  "break",
		80:  "func",
		82:  "select",
		100: "chan",
		116: "continue",
		120: "goto",
		121: "const",
		126: "go",
		131: "package",
		145: "defer",
		147: "default",
		199: "map",
		218: "if",
		224: "type",
		241: "interface",
		242: "import",
		244: "else",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 7) ^ (name[1] << 3))
	if h > 253 {
		return false
	}
	return [...]string{
		3:   "fallthrough",
		11:  "var",
		13:  "range",
		15:  "package",
		38:  "struct",
		45:  "defer",
		46:  "return",
		47:  "default",
		62:  "switch",
		123: "for",
		139: "map",
		140: "case",
		149: "break",
		172: "func",
		174: "select",
		178: "if",
		196: "chan",
		204: "type",
		228: "else",
		238: "import",
		240: "continue",
		249: "interface",
		250: "go",
		252: "goto",
		253: "const",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift130(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 3) ^ (name[1] << 0))
	if h > 251 {
		return false
	}
	return [...]string{
		15:  "map",
		41:  "import",
		42:  "if",
		52:  "interface",
		71:  "fallthrough",
		75:  "default",
		76:  "else",
		77:  "func",
		79:  "defer",
		83:  "go",
		89:  "for",
		95:  "goto",
		103: "continue",
		104: "break",
		113: "case",
		120: "chan",
		125: "const",
		209: "type",
		215: "var",
		224: "struct",
		227: "switch",
		239: "package",
		241: "select",
		249: "return",
		251: "range",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift132(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 3) ^ (name[1] << 2))
	if h > 240 {
		return false
	}
	return [...]string{
		0:   "select",
		8:   "return",
		10:  "package",
		30:  "range",
		50:  "var",
		68:  "struct",
		72:  "switch",
		76:  "type",
		128: "go",
		138: "for",
		140: "goto",
		144: "else",
		148: "case",
		162: "fallthrough",
		174: "const",
		176: "chan",
		180: "continue",
		186: "default",
		190: "defer",
		210: "break",
		212: "if",
		226: "interface",
		234: "map",
		236: "func",
		240: "import",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift135(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 3) ^ (name[1] << 5))
	if h > 242 {
		return false
	}
	return [...]string{
		6:   "fallthrough",
		16:  "chan",
		20:  "struct",
		48:  "case",
		52:  "select",
		60:  "return",
		78:  "map",
		90:  "break",
		116: "switch",
		136: "type",
		138: "defer",
		140: "if",
		142: "default",
		150: "var",
		152: "func",
		154: "interface",
		160: "else",
		174: "package",
		186: "range",
		208: "goto",
		214: "for",
		220: "go",
		228: "import",
		232: "continue",
		242: "const",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift143(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) ^ (name[1] << 3))
	if h > 244 {
		return false
	}
	return [...]string{
		0:   "goto",
		4:   "return",
		6:   "package",
		12:  "go",
		20:  "select",
		30:  "for",
		34:  "range",
		48:  "case",
		56:  "else",
		66:  "const",
		88:  "continue",
		98:  "defer",
		102: "default",
		110: "var",
		120: "chan",
		126: "fallthrough",
		128: "type",
		132: "switch",
		156: "struct",
		164: "if",
		186: "break",
		192: "func",
		222: "map",
		242: "interface",
		244: "import",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift145(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) ^ (name[1] << 5))
	if h > 246 {
		return false
	}
	return [...]string{
		10:  "range",
		24:  "case",
		46:  "package",
		56:  "chan",
		60:  "import",
		66:  "interface",
		70:  "var",
		84:  "if",
		86:  "fallthrough",
		104: "type",
		106: "break",
		134: "for",
		140: "return",
		148: "go",
		152: "goto",
		156: "select",
		188: "struct",
		192: "continue",
		200: "func",
		216: "else",
		218: "const",
		220: "switch",
		234: "defer",
		238: "default",
		246: "map",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift146(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) ^ (name[1] << 6))
	if h > 252 {
		return false
	}
	return [...]string{
		2:   "interface",
		8:   "type",
		10:  "defer",
		14:  "default",
		20:  "if",
		38:  "var",
		40:  "func",
		54:  "fallthrough",
		56:  "chan",
		60:  "struct",
		78:  "package",
		88:  "else",
		106: "range",
		108: "return",
		120: "case",
		124: "select",
		150: "map",
		166: "for",
		170: "break",
		180: "go",
		184: "goto",
		220: "import",
		224: "continue",
		250: "const",
		252: "switch",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 6) ^ (name[1] << 1))
	if h > 250 {
		return false
	}
	return [...]string{
		6:   "select",
		10:  "case",
		14:  "continue",
		20:  "const",
		22:  "goto",
		24:  "chan",
		26:  "go",
		34:  "switch",
		36:  "struct",
		68:  "var",
		70:  "return",
		72:  "range",
		84:  "fallthrough",
		88:  "for",
		98:  "func",
		110: "break",
		132: "map",
		136: "if",
		142: "interface",
		144: "else",
		150: "import",
		192: "defer",
		196: "default",
		204: "package",
		250: "type",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift240(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) ^ ((name[0] << 4) ^ (name[1] << 0))
	if h > 254 {
		return false
	}
	return [...]string{
		3:   "for",
		5:   "func",
		13:  "var",
		15:  "goto",
		23:  "go",
		41:  "type",
		44:  "else",
		45:  "fallthrough",
		49:  "defer",
		57:  "default",
		65:  "case",
		70:  "break",
		72:  "chan",
		75:  "const",
		77:  "select",
		85:  "range",
		92:  "struct",
		93:  "return",
		95:  "switch",
		125: "package",
		127: "continue",
		189: "map",
		218: "interface",
		229: "import",
		254: "if",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift246(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) ^ ((name[0] << 4) ^ (name[1] << 6))
	if h > 232 {
		return false
	}
	return [...]string{
		12:  "fallthrough",
		16:  "type",
		20:  "defer",
		24:  "if",
		28:  "default",
		32:  "chan",
		40:  "struct",
		44:  "var",
		48:  "func",
		52:  "interface",
		64:  "else",
		92:  "package",
		96:  "case",
		104: "select",
		116: "range",
		120: "return",
		156: "map",
		160: "goto",
		172: "for",
		180: "break",
		184: "go",
		200: "import",
		208: "continue",
		228: "const",
		232: "switch",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift250(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) ^ ((name[0] << 5) ^ (name[1] << 0))
	if h > 249 {
		return false
	}
	return [...]string{
		12:  "struct",
		15:  "switch",
		17:  "case",
		24:  "chan",
		27:  "const",
		29:  "select",
		38:  "break",
		47:  "continue",
		53:  "range",
		61:  "return",
		78:  "if",
		85:  "import",
		106: "interface",
		125: "package",
		135: "go",
		141: "fallthrough",
		159: "goto",
		163: "for",
		165: "func",
		173: "var",
		205: "map",
		220: "else",
		233: "type",
		241: "defer",
		249: "default",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift260(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) ^ ((name[0] << 6) ^ (name[1] << 0))
	if h > 253 {
		return false
	}
	return [...]string{
		10:  "interface",
		45:  "map",
		46:  "if",
		53:  "import",
		60:  "else",
		105: "type",
		113: "defer",
		121: "default",
		125: "package",
		143: "continue",
		167: "go",
		172: "struct",
		175: "switch",
		177: "case",
		184: "chan",
		187: "const",
		189: "select",
		191: "goto",
		205: "fallthrough",
		227: "for",
		229: "func",
		230: "break",
		237: "var",
		245: "range",
		253: "return",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift303(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) ^ ((name[0] << 0) ^ (name[1] << 3))
	if h > 251 {
		return false
	}
	return [...]string{
		3:   "chan",
		6:   "for",
		15:  "go",
		37:  "else",
		49:  "import",
		51:  "const",
		54:  "fallthrough",
		63:  "goto",
		64:  "package",
		73:  "if",
		75:  "case",
		81:  "interface",
		82:  "range",
		91:  "continue",
		100: "defer",
		102: "var",
		106: "return",
		107: "select",
		116: "default",
		125: "map",
		156: "type",
		218: "break",
		227: "struct",
		238: "func",
		251: "switch",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift304(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) ^ ((name[0] << 0) ^ (name[1] << 4))
	if h > 211 {
		return false
	}
	return [...]string{
		3:   "struct",
		12:  "default",
		18:  "return",
		19:  "select",
		22:  "func",
		25:  "if",
		28:  "defer",
		46:  "fallthrough",
		51:  "switch",
		74:  "range",
		83:  "case",
		88:  "package",
		101: "map",
		106: "break",
		126: "var",
		133: "else",
		135: "go",
		137: "import",
		142: "for",
		183: "goto",
		187: "const",
		193: "interface",
		195: "chan",
		196: "type",
		211: "continue",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift351(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) ^ ((name[0] << 5) ^ (name[1] << 1))
	if h > 254 {
		return false
	}
	return [...]string{
		6:   "for",
		10:  "func",
		26:  "var",
		30:  "goto",
		46:  "go",
		82:  "type",
		88:  "else",
		90:  "fallthrough",
		98:  "defer",
		114: "default",
		122: "map",
		130: "case",
		140: "break",
		144: "chan",
		150: "const",
		154: "select",
		170: "range",
		180: "interface",
		184: "struct",
		186: "return",
		190: "switch",
		202: "import",
		250: "package",
		252: "if",
		254: "continue",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift361(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) ^ ((name[0] << 6) ^ (name[1] << 1))
	if h > 250 {
		return false
	}
	return [...]string{
		14:  "go",
		24:  "struct",
		26:  "fallthrough",
		30:  "switch",
		34:  "case",
		48:  "chan",
		54:  "const",
		58:  "select",
		62:  "goto",
		70:  "for",
		74:  "func",
		76:  "break",
		90:  "var",
		94:  "continue",
		106: "range",
		122: "return",
		154: "map",
		156: "if",
		170: "import",
		184: "else",
		210: "type",
		212: "interface",
		226: "defer",
		242: "default",
		250: "package",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift402(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) ^ ((name[0] << 0) ^ (name[1] << 2))
	if h > 251 {
		return false
	}
	return [...]string{
		65:  "interface",
		82:  "fallthrough",
		95:  "continue",
		128: "default",
		131: "chan",
		132: "package",
		134: "return",
		135: "select",
		143: "const",
		149: "else",
		155: "goto",
		160: "defer",
		166: "range",
		167: "case",
		189: "import",
		194: "var",
		195: "struct",
		207: "switch",
		208: "type",
		209: "if",
		217: "map",
		234: "for",
		242: "func",
		250: "break",
		251: "go",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift403(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) ^ ((name[0] << 0) ^ (name[1] << 3))
	if h > 252 {
		return false
	}
	return [...]string{
		8:   "package",
		28:  "defer",
		42:  "range",
		43:  "case",
		46:  "for",
		58:  "return",
		59:  "select",
		60:  "default",
		63:  "go",
		69:  "else",
		75:  "const",
		78:  "var",
		85:  "map",
		95:  "goto",
		97:  "import",
		99:  "chan",
		121: "if",
		137: "interface",
		142: "func",
		155: "continue",
		162: "break",
		171: "switch",
		179: "struct",
		222: "fallthrough",
		252: "type",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift404(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) ^ ((name[0] << 0) ^ (name[1] << 4))
	if h > 229 {
		return false
	}
	return [...]string{
		16:  "package",
		18:  "break",
		19:  "continue",
		25:  "interface",
		41:  "if",
		50:  "range",
		51:  "case",
		66:  "return",
		67:  "select",
		68:  "default",
		77:  "map",
		83:  "struct",
		86:  "var",
		99:  "switch",
		100: "defer",
		118: "func",
		163: "chan",
		164: "type",
		166: "for",
		183: "go",
		195: "const",
		198: "fallthrough",
		215: "goto",
		217: "import",
		229: "else",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift414(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) ^ ((name[0] << 1) ^ (name[1] << 4))
	if h > 250 {
		return false
	}
	return [...]string{
		6:   "chan",
		12:  "for",
		30:  "go",
		56:  "type",
		74:  "else",
		98:  "import",
		102: "const",
		108: "fallthrough",
		126: "goto",
		128: "package",
		146: "if",
		150: "case",
		162: "interface",
		164: "range",
		180: "break",
		182: "continue",
		198: "struct",
		200: "defer",
		204: "var",
		212: "return",
		214: "select",
		220: "func",
		232: "default",
		246: "switch",
		250: "map",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift501(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 0) ^ (name[1] << 1))
	if h > 249 {
		return false
	}
	return [...]string{
		6:   "type",
		12:  "func",
		14:  "defer",
		16:  "range",
		29:  "const",
		33:  "case",
		38:  "break",
		51:  "chan",
		57:  "goto",
		61:  "else",
		78:  "default",
		82:  "package",
		91:  "struct",
		93:  "switch",
		115: "import",
		120: "return",
		121: "select",
		149: "interface",
		189: "continue",
		196: "fallthrough",
		207: "map",
		212: "var",
		216: "for",
		229: "if",
		249: "go",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift503(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 0) ^ (name[1] << 3))
	if h > 236 {
		return false
	}
	return [...]string{
		5:   "map",
		11:  "switch",
		14:  "fallthrough",
		19:  "struct",
		25:  "if",
		27:  "continue",
		30:  "var",
		57:  "interface",
		60:  "type",
		78:  "func",
		82:  "break",
		95:  "go",
		126: "for",
		133: "else",
		152: "package",
		154: "return",
		155: "select",
		159: "goto",
		163: "chan",
		172: "default",
		187: "const",
		193: "import",
		218: "range",
		235: "case",
		236: "defer",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift505(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 0) ^ (name[1] << 5))
	if h > 242 {
		return false
	}
	return [...]string{
		7:   "goto",
		9:   "import",
		18:  "return",
		19:  "select",
		35:  "const",
		36:  "default",
		38:  "fallthrough",
		45:  "map",
		51:  "struct",
		54:  "var",
		70:  "func",
		83:  "switch",
		100: "defer",
		101: "else",
		130: "break",
		131: "continue",
		137: "interface",
		176: "package",
		195: "case",
		199: "go",
		212: "type",
		227: "chan",
		230: "for",
		233: "if",
		242: "range",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift510(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 1) ^ (name[1] << 0))
	if h > 244 {
		return false
	}
	return [...]string{
		9:   "const",
		13:  "defer",
		17:  "type",
		22:  "break",
		33:  "goto",
		37:  "range",
		38:  "else",
		39:  "case",
		46:  "chan",
		57:  "func",
		65:  "return",
		67:  "select",
		77:  "default",
		81:  "switch",
		82:  "struct",
		97:  "package",
		127: "import",
		156: "interface",
		169: "continue",
		195: "for",
		205: "fallthrough",
		219: "map",
		225: "go",
		237: "var",
		244: "if",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift513(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 1) ^ (name[1] << 3))
	if h > 246 {
		return false
	}
	return [...]string{
		0:   "default",
		6:   "chan",
		8:   "package",
		12:  "return",
		14:  "select",
		30:  "const",
		42:  "else",
		54:  "goto",
		64:  "defer",
		76:  "range",
		78:  "case",
		122: "import",
		130: "interface",
		132: "var",
		134: "struct",
		158: "switch",
		160: "type",
		162: "if",
		164: "fallthrough",
		178: "map",
		190: "continue",
		212: "for",
		228: "func",
		244: "break",
		246: "go",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift525(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 2) ^ (name[1] << 5))
	if h > 252 {
		return false
	}
	return [...]string{
		0:   "package",
		12:  "chan",
		24:  "for",
		36:  "if",
		44:  "case",
		60:  "go",
		68:  "interface",
		72:  "range",
		104: "break",
		108: "continue",
		112: "type",
		140: "struct",
		144: "defer",
		148: "else",
		152: "var",
		168: "return",
		172: "select",
		184: "func",
		196: "import",
		204: "const",
		208: "default",
		216: "fallthrough",
		236: "switch",
		244: "map",
		252: "goto",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift602(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 0) ^ (name[1] << 2))
	if h > 234 {
		return false
	}
	return [...]string{
		26:  "for",
		34:  "fallthrough",
		35:  "struct",
		41:  "map",
		47:  "switch",
		48:  "default",
		50:  "var",
		52:  "package",
		91:  "go",
		93:  "import",
		102: "return",
		103: "select",
		113: "if",
		144: "type",
		145: "interface",
		159: "const",
		176: "defer",
		178: "func",
		182: "range",
		195: "chan",
		213: "else",
		219: "goto",
		223: "continue",
		231: "case",
		234: "break",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift603(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 0) ^ (name[1] << 3))
	if h > 222 {
		return false
	}
	return [...]string{
		5:   "else",
		12:  "defer",
		27:  "continue",
		31:  "goto",
		35:  "chan",
		58:  "range",
		75:  "switch",
		83:  "struct",
		89:  "interface",
		91:  "const",
		107: "case",
		129: "import",
		140: "default",
		159: "go",
		165: "map",
		174: "fallthrough",
		178: "break",
		184: "package",
		188: "type",
		190: "var",
		206: "func",
		217: "if",
		218: "return",
		219: "select",
		222: "for",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift604(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 0) ^ (name[1] << 4))
	if h > 244 {
		return false
	}
	return [...]string{
		2:   "break",
		23:  "go",
		34:  "range",
		54:  "func",
		57:  "import",
		86:  "for",
		115: "case",
		116: "defer",
		131: "switch",
		137: "if",
		147: "continue",
		151: "goto",
		160: "package",
		162: "return",
		163: "select",
		165: "else",
		166: "var",
		179: "struct",
		182: "fallthrough",
		189: "map",
		201: "interface",
		211: "const",
		227: "chan",
		228: "type",
		244: "default",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift605(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 0) ^ (name[1] << 5))
	if h > 233 {
		return false
	}
	return [...]string{
		4:   "default",
		7:   "go",
		18:  "range",
		19:  "switch",
		41:  "if",
		67:  "case",
		70:  "for",
		73:  "import",
		82:  "return",
		83:  "select",
		84:  "type",
		98:  "break",
		99:  "chan",
		115: "struct",
		131: "continue",
		132: "defer",
		134: "fallthrough",
		135: "goto",
		141: "map",
		144: "package",
		150: "var",
		195: "const",
		198: "func",
		229: "else",
		233: "interface",
	}[h] == name
}

func TwoHashTableAlt_XorXor_Shift630(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 3) ^ (name[1] << 0))
	if h > 217 {
		return false
	}
	return [...]string{
		5:   "defer",
		17:  "var",
		33:  "package",
		34:  "break",
		55:  "const",
		68:  "else",
		69:  "func",
		87:  "goto",
		102: "interface",
		108: "struct",
		111: "switch",
		112: "chan",
		117: "return",
		119: "continue",
		121: "case",
		125: "select",
		133: "default",
		145: "fallthrough",
		159: "for",
		165: "import",
		174: "if",
		177: "range",
		201: "map",
		215: "go",
		217: "type",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift015(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 1) + (name[1] << 5))
	if h > 253 {
		return false
	}
	return [...]string{
		12:  "switch",
		20:  "if",
		27:  "interface",
		43:  "const",
		46:  "continue",
		47:  "for",
		48:  "go",
		50:  "goto",
		74:  "return",
		76:  "select",
		78:  "else",
		108: "struct",
		109: "defer",
		111: "default",
		112: "func",
		120: "import",
		137: "break",
		199: "package",
		201: "range",
		202: "chan",
		204: "type",
		207: "var",
		234: "case",
		247: "fallthrough",
		253: "map",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift021(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 2) + (name[1] << 1))
	if h > 132 {
		return false
	}
	return [...]string{
		8:   "return",
		9:   "package",
		12:  "select",
		15:  "range",
		29:  "var",
		38:  "type",
		40:  "switch",
		42:  "struct",
		68:  "go",
		70:  "goto",
		73:  "for",
		80:  "else",
		82:  "case",
		87:  "const",
		90:  "continue",
		95:  "defer",
		96:  "chan",
		97:  "default",
		101: "fallthrough",
		106: "if",
		113: "break",
		118: "func",
		121: "map",
		129: "interface",
		132: "import",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift024(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 2) + (name[1] << 4))
	if h > 221 {
		return false
	}
	return [...]string{
		16:  "chan",
		68:  "type",
		77:  "interface",
		88:  "else",
		107: "for",
		110: "go",
		112: "goto",
		122: "import",
		129: "const",
		132: "continue",
		146: "struct",
		147: "fallthrough",
		158: "return",
		160: "case",
		162: "select",
		167: "map",
		173: "break",
		194: "switch",
		197: "defer",
		198: "if",
		199: "default",
		203: "var",
		204: "func",
		215: "package",
		221: "range",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift025(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 2) + (name[1] << 5))
	if h > 251 {
		return false
	}
	return [...]string{
		10:  "import",
		24:  "else",
		50:  "switch",
		53:  "defer",
		55:  "default",
		60:  "func",
		82:  "struct",
		102: "if",
		109: "interface",
		110: "return",
		113: "const",
		114: "select",
		116: "continue",
		123: "for",
		126: "go",
		128: "goto",
		144: "chan",
		151: "map",
		176: "case",
		195: "fallthrough",
		205: "break",
		231: "package",
		237: "range",
		244: "type",
		251: "var",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift026(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 2) + (name[1] << 6))
	if h > 247 {
		return false
	}
	return [...]string{
		13:  "break",
		18:  "switch",
		38:  "if",
		45:  "interface",
		81:  "const",
		84:  "continue",
		91:  "for",
		94:  "go",
		96:  "goto",
		135: "package",
		141: "range",
		142: "return",
		144: "chan",
		146: "select",
		148: "type",
		152: "else",
		155: "var",
		208: "case",
		210: "struct",
		213: "defer",
		215: "default",
		220: "func",
		227: "fallthrough",
		234: "import",
		247: "map",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift030(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) + (name[1] << 0))
	if h > 251 {
		return false
	}
	return [...]string{
		3:   "select",
		12:  "map",
		43:  "import",
		47:  "interface",
		48:  "if",
		72:  "else",
		73:  "func",
		74:  "defer",
		76:  "default",
		89:  "go",
		91:  "goto",
		92:  "fallthrough",
		98:  "for",
		103: "break",
		116: "chan",
		124: "const",
		125: "case",
		127: "continue",
		212: "var",
		221: "type",
		232: "package",
		242: "struct",
		245: "switch",
		246: "range",
		251: "return",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift032(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) + (name[1] << 2))
	if h > 249 {
		return false
	}
	return [...]string{
		2:   "import",
		10:  "return",
		11:  "package",
		18:  "select",
		25:  "range",
		55:  "var",
		72:  "type",
		74:  "switch",
		78:  "struct",
		134: "go",
		136: "goto",
		143: "for",
		156: "else",
		160: "case",
		169: "const",
		172: "continue",
		185: "defer",
		187: "default",
		188: "chan",
		191: "fallthrough",
		210: "if",
		221: "break",
		232: "func",
		239: "map",
		249: "interface",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift034(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) + (name[1] << 4))
	if h > 240 {
		return false
	}
	return [...]string{
		12:  "case",
		42:  "if",
		43:  "fallthrough",
		52:  "type",
		53:  "break",
		100: "func",
		117: "defer",
		119: "default",
		123: "map",
		133: "range",
		151: "package",
		156: "chan",
		158: "import",
		163: "var",
		177: "interface",
		195: "for",
		198: "return",
		202: "go",
		204: "goto",
		206: "select",
		222: "struct",
		236: "else",
		237: "const",
		238: "switch",
		240: "continue",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift035(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) + (name[1] << 5))
	if h > 253 {
		return false
	}
	return [...]string{
		0:   "continue",
		27:  "fallthrough",
		28:  "chan",
		30:  "struct",
		54:  "return",
		60:  "case",
		62:  "select",
		75:  "map",
		85:  "break",
		126: "switch",
		132: "type",
		133: "defer",
		135: "default",
		138: "if",
		145: "interface",
		147: "var",
		148: "func",
		167: "package",
		172: "else",
		181: "range",
		211: "for",
		218: "go",
		220: "goto",
		238: "import",
		253: "const",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift040(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) + (name[1] << 0))
	if h > 248 {
		return false
	}
	return [...]string{
		3:   "import",
		4:   "var",
		7:   "interface",
		12:  "fallthrough",
		18:  "for",
		25:  "func",
		33:  "go",
		35:  "goto",
		42:  "defer",
		44:  "default",
		61:  "type",
		64:  "else",
		70:  "range",
		74:  "struct",
		75:  "return",
		77:  "switch",
		85:  "case",
		87:  "break",
		91:  "select",
		92:  "chan",
		100: "const",
		103: "continue",
		104: "package",
		180: "map",
		248: "if",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift043(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) + (name[1] << 3))
	if h > 254 {
		return false
	}
	return [...]string{
		10:  "go",
		12:  "goto",
		14:  "return",
		15:  "package",
		27:  "for",
		30:  "select",
		45:  "range",
		52:  "else",
		60:  "case",
		77:  "const",
		80:  "continue",
		107: "var",
		109: "defer",
		111: "default",
		115: "fallthrough",
		116: "chan",
		140: "type",
		142: "switch",
		150: "struct",
		162: "if",
		181: "break",
		204: "func",
		219: "map",
		233: "interface",
		254: "import",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift045(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) + (name[1] << 5))
	if h > 243 {
		return false
	}
	return [...]string{
		5:   "range",
		20:  "case",
		39:  "package",
		52:  "chan",
		54:  "import",
		67:  "var",
		75:  "fallthrough",
		82:  "if",
		89:  "interface",
		100: "type",
		101: "break",
		131: "for",
		134: "return",
		146: "go",
		148: "goto",
		150: "select",
		182: "struct",
		196: "func",
		212: "else",
		213: "const",
		214: "switch",
		216: "continue",
		229: "defer",
		231: "default",
		243: "map",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift046(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) + (name[1] << 6))
	if h > 248 {
		return false
	}
	return [...]string{
		4:   "type",
		5:   "defer",
		7:   "default",
		18:  "if",
		25:  "interface",
		35:  "var",
		36:  "func",
		43:  "fallthrough",
		52:  "chan",
		54:  "struct",
		71:  "package",
		84:  "else",
		101: "range",
		102: "return",
		116: "case",
		118: "select",
		147: "map",
		163: "for",
		165: "break",
		178: "go",
		180: "goto",
		214: "import",
		245: "const",
		246: "switch",
		248: "continue",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 5) + (name[1] << 0))
	if h > 253 {
		return false
	}
	return [...]string{
		5:   "case",
		11:  "select",
		12:  "chan",
		20:  "const",
		23:  "continue",
		26:  "struct",
		29:  "switch",
		38:  "range",
		43:  "return",
		55:  "break",
		72:  "if",
		83:  "import",
		87:  "interface",
		104: "package",
		145: "go",
		147: "goto",
		164: "var",
		172: "fallthrough",
		178: "for",
		185: "func",
		196: "map",
		208: "else",
		234: "defer",
		236: "default",
		253: "type",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift052(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 5) + (name[1] << 2))
	if h > 250 {
		return false
	}
	return [...]string{
		20:  "else",
		24:  "func",
		25:  "defer",
		27:  "default",
		39:  "map",
		71:  "var",
		79:  "fallthrough",
		94:  "go",
		96:  "goto",
		104: "type",
		127: "for",
		139: "package",
		141: "break",
		154: "import",
		161: "interface",
		182: "struct",
		186: "if",
		194: "switch",
		196: "chan",
		201: "range",
		218: "return",
		225: "const",
		228: "continue",
		232: "case",
		250: "select",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 6) + (name[1] << 0))
	if h > 249 {
		return false
	}
	return [...]string{
		36:  "map",
		40:  "if",
		48:  "else",
		51:  "import",
		55:  "interface",
		104: "package",
		106: "defer",
		108: "default",
		125: "type",
		165: "case",
		171: "select",
		172: "chan",
		177: "go",
		179: "goto",
		180: "const",
		183: "continue",
		186: "struct",
		189: "switch",
		228: "var",
		230: "range",
		235: "return",
		236: "fallthrough",
		242: "for",
		247: "break",
		249: "func",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift061(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 6) + (name[1] << 1))
	if h > 246 {
		return false
	}
	return [...]string{
		6:   "case",
		16:  "select",
		20:  "chan",
		32:  "go",
		34:  "goto",
		35:  "const",
		38:  "continue",
		46:  "struct",
		52:  "switch",
		69:  "var",
		71:  "range",
		77:  "fallthrough",
		80:  "return",
		97:  "for",
		105: "break",
		110: "func",
		133: "map",
		142: "if",
		156: "else",
		160: "import",
		165: "interface",
		201: "package",
		207: "defer",
		209: "default",
		246: "type",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 6) + (name[1] << 2))
	if h > 250 {
		return false
	}
	return [...]string{
		1:   "interface",
		7:   "var",
		9:   "range",
		15:  "fallthrough",
		22:  "struct",
		26:  "return",
		34:  "switch",
		63:  "for",
		72:  "case",
		77:  "break",
		88:  "func",
		90:  "select",
		100: "chan",
		126: "go",
		128: "goto",
		129: "const",
		132: "continue",
		139: "package",
		153: "defer",
		155: "default",
		199: "map",
		218: "if",
		232: "type",
		244: "else",
		250: "import",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 7) + (name[1] << 3))
	if h > 253 {
		return false
	}
	return [...]string{
		0:   "continue",
		11:  "var",
		13:  "range",
		15:  "package",
		19:  "fallthrough",
		38:  "struct",
		45:  "defer",
		46:  "return",
		47:  "default",
		62:  "switch",
		123: "for",
		139: "map",
		140: "case",
		149: "break",
		172: "func",
		174: "select",
		178: "if",
		196: "chan",
		204: "type",
		228: "else",
		238: "import",
		249: "interface",
		250: "go",
		252: "goto",
		253: "const",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift132(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 3) + (name[1] << 2))
	if h > 242 {
		return false
	}
	return [...]string{
		2:   "interface",
		8:   "import",
		16:  "return",
		18:  "package",
		24:  "select",
		30:  "range",
		58:  "var",
		76:  "type",
		80:  "switch",
		84:  "struct",
		136: "go",
		140: "goto",
		146: "for",
		160: "else",
		164: "case",
		174: "const",
		180: "continue",
		190: "defer",
		192: "chan",
		194: "default",
		202: "fallthrough",
		212: "if",
		226: "break",
		236: "func",
		242: "map",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift135(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 3) + (name[1] << 5))
	if h > 244 {
		return false
	}
	return [...]string{
		2:   "const",
		8:   "continue",
		32:  "chan",
		36:  "struct",
		38:  "fallthrough",
		60:  "return",
		64:  "case",
		68:  "select",
		78:  "map",
		90:  "break",
		132: "switch",
		136: "type",
		138: "defer",
		140: "if",
		142: "default",
		150: "var",
		152: "func",
		154: "interface",
		174: "package",
		176: "else",
		186: "range",
		214: "for",
		220: "go",
		224: "goto",
		244: "import",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift141(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) + (name[1] << 1))
	if h > 254 {
		return false
	}
	return [...]string{
		6:   "select",
		24:  "map",
		86:  "import",
		94:  "interface",
		96:  "if",
		144: "else",
		146: "func",
		148: "defer",
		152: "default",
		168: "var",
		178: "go",
		182: "goto",
		184: "fallthrough",
		186: "type",
		196: "for",
		206: "break",
		208: "package",
		228: "struct",
		232: "chan",
		234: "switch",
		236: "range",
		246: "return",
		248: "const",
		250: "case",
		254: "continue",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift143(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) + (name[1] << 3))
	if h > 242 {
		return false
	}
	return [...]string{
		4:   "import",
		12:  "go",
		16:  "goto",
		20:  "return",
		22:  "package",
		30:  "for",
		36:  "select",
		50:  "range",
		56:  "else",
		64:  "case",
		82:  "const",
		88:  "continue",
		110: "var",
		114: "defer",
		118: "default",
		120: "chan",
		126: "fallthrough",
		144: "type",
		148: "switch",
		156: "struct",
		164: "if",
		186: "break",
		208: "func",
		222: "map",
		242: "interface",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift145(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) + (name[1] << 5))
	if h > 246 {
		return false
	}
	return [...]string{
		10:  "range",
		24:  "case",
		46:  "package",
		56:  "chan",
		60:  "import",
		70:  "var",
		84:  "if",
		86:  "fallthrough",
		98:  "interface",
		104: "type",
		106: "break",
		134: "for",
		140: "return",
		148: "go",
		152: "goto",
		156: "select",
		188: "struct",
		200: "func",
		216: "else",
		218: "const",
		220: "switch",
		224: "continue",
		234: "defer",
		238: "default",
		246: "map",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift146(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) + (name[1] << 6))
	if h > 252 {
		return false
	}
	return [...]string{
		0:   "continue",
		8:   "type",
		10:  "defer",
		14:  "default",
		20:  "if",
		34:  "interface",
		38:  "var",
		40:  "func",
		54:  "fallthrough",
		56:  "chan",
		60:  "struct",
		78:  "package",
		88:  "else",
		106: "range",
		108: "return",
		120: "case",
		124: "select",
		150: "map",
		166: "for",
		170: "break",
		180: "go",
		184: "goto",
		220: "import",
		250: "const",
		252: "switch",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift150(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 5) + (name[1] << 0))
	if h > 243 {
		return false
	}
	return [...]string{
		1:   "type",
		9:   "case",
		16:  "chan",
		17:  "select",
		25:  "const",
		31:  "continue",
		32:  "struct",
		35:  "switch",
		43:  "range",
		49:  "return",
		60:  "break",
		74:  "if",
		89:  "import",
		96:  "interface",
		111: "package",
		147: "go",
		151: "goto",
		167: "var",
		181: "for",
		183: "fallthrough",
		189: "func",
		199: "map",
		212: "else",
		239: "defer",
		243: "default",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift151(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 5) + (name[1] << 1))
	if h > 240 {
		return false
	}
	return [...]string{
		6:   "import",
		8:   "var",
		14:  "interface",
		24:  "fallthrough",
		36:  "for",
		50:  "func",
		66:  "go",
		70:  "goto",
		84:  "defer",
		88:  "default",
		104: "map",
		122: "type",
		128: "else",
		140: "range",
		148: "struct",
		150: "return",
		154: "switch",
		170: "case",
		174: "break",
		182: "select",
		184: "chan",
		200: "const",
		206: "continue",
		208: "package",
		240: "if",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 6) + (name[1] << 1))
	if h > 250 {
		return false
	}
	return [...]string{
		10:  "case",
		22:  "select",
		24:  "chan",
		34:  "go",
		38:  "goto",
		40:  "const",
		46:  "continue",
		52:  "struct",
		58:  "switch",
		72:  "var",
		76:  "range",
		86:  "return",
		88:  "fallthrough",
		100: "for",
		110: "break",
		114: "func",
		136: "map",
		144: "if",
		160: "else",
		166: "import",
		174: "interface",
		208: "package",
		212: "defer",
		216: "default",
		250: "type",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift240(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) ^ ((name[0] << 4) + (name[1] << 0))
	if h > 254 {
		return false
	}
	return [...]string{
		13:  "var",
		21:  "import",
		27:  "for",
		34:  "interface",
		37:  "func",
		39:  "go",
		45:  "fallthrough",
		47:  "goto",
		57:  "defer",
		65:  "default",
		73:  "type",
		76:  "else",
		85:  "range",
		92:  "struct",
		93:  "return",
		95:  "switch",
		97:  "case",
		102: "break",
		104: "chan",
		109: "select",
		115: "const",
		125: "package",
		127: "continue",
		189: "map",
		254: "if",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift262(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) ^ ((name[0] << 6) + (name[1] << 2))
	if h > 244 {
		return false
	}
	return [...]string{
		0:   "else",
		12:  "import",
		16:  "var",
		24:  "range",
		28:  "interface",
		40:  "struct",
		44:  "return",
		48:  "fallthrough",
		52:  "switch",
		72:  "for",
		84:  "case",
		92:  "break",
		100: "func",
		108: "select",
		112: "chan",
		132: "go",
		140: "goto",
		144: "const",
		156: "continue",
		160: "package",
		168: "defer",
		176: "default",
		208: "map",
		224: "if",
		244: "type",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift304(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) ^ ((name[0] << 0) + (name[1] << 4))
	if h > 233 {
		return false
	}
	return [...]string{
		3:   "chan",
		4:   "type",
		25:  "if",
		51:  "switch",
		82:  "return",
		83:  "select",
		86:  "func",
		92:  "defer",
		99:  "struct",
		106: "break",
		108: "default",
		126: "var",
		138: "range",
		147: "case",
		149: "map",
		152: "package",
		167: "go",
		174: "for",
		183: "goto",
		187: "const",
		197: "else",
		206: "fallthrough",
		209: "interface",
		211: "continue",
		233: "import",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift350(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) ^ ((name[0] << 5) + (name[1] << 0))
	if h > 249 {
		return false
	}
	return [...]string{
		13:  "defer",
		25:  "type",
		29:  "default",
		33:  "case",
		40:  "chan",
		53:  "select",
		55:  "const",
		68:  "struct",
		71:  "switch",
		73:  "range",
		79:  "continue",
		85:  "return",
		86:  "if",
		90:  "break",
		125: "import",
		150: "interface",
		153: "package",
		159: "go",
		175: "goto",
		185: "var",
		199: "for",
		213: "func",
		217: "map",
		236: "else",
		249: "fallthrough",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift351(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) ^ ((name[0] << 5) + (name[1] << 1))
	if h > 254 {
		return false
	}
	return [...]string{
		26:  "var",
		42:  "import",
		54:  "for",
		68:  "interface",
		74:  "func",
		78:  "go",
		90:  "fallthrough",
		94:  "goto",
		114: "defer",
		122: "map",
		130: "default",
		146: "type",
		152: "else",
		170: "range",
		184: "struct",
		186: "return",
		190: "switch",
		194: "case",
		204: "break",
		208: "chan",
		218: "select",
		230: "const",
		250: "package",
		252: "if",
		254: "continue",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift403(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) ^ ((name[0] << 0) + (name[1] << 3))
	if h > 252 {
		return false
	}
	return [...]string{
		14:  "func",
		30:  "fallthrough",
		43:  "switch",
		51:  "struct",
		63:  "go",
		66:  "break",
		69:  "else",
		78:  "for",
		95:  "goto",
		97:  "import",
		99:  "chan",
		107: "const",
		121: "if",
		149: "map",
		155: "continue",
		156: "defer",
		169: "interface",
		171: "case",
		174: "var",
		186: "return",
		187: "select",
		188: "default",
		202: "range",
		232: "package",
		252: "type",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift405(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) ^ ((name[0] << 0) + (name[1] << 5))
	if h > 246 {
		return false
	}
	return [...]string{
		3:   "continue",
		6:   "func",
		20:  "defer",
		37:  "else",
		41:  "import",
		50:  "return",
		51:  "select",
		52:  "default",
		57:  "interface",
		83:  "struct",
		114: "break",
		125: "map",
		131: "case",
		134: "var",
		148: "type",
		162: "range",
		163: "chan",
		167: "go",
		182: "for",
		192: "package",
		199: "goto",
		201: "if",
		211: "const",
		243: "switch",
		246: "fallthrough",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift460(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) ^ ((name[0] << 6) + (name[1] << 0))
	if h > 255 {
		return false
	}
	return [...]string{
		5:   "select",
		17:  "var",
		20:  "struct",
		23:  "switch",
		31:  "for",
		47:  "continue",
		49:  "range",
		53:  "func",
		66:  "break",
		69:  "return",
		70:  "if",
		81:  "map",
		108: "else",
		141: "import",
		145: "fallthrough",
		181: "defer",
		185: "type",
		190: "interface",
		207: "go",
		209: "package",
		213: "default",
		225: "case",
		232: "chan",
		239: "goto",
		255: "const",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift502(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 0) + (name[1] << 2))
	if h > 241 {
		return false
	}
	return [...]string{
		16:  "type",
		27:  "go",
		49:  "if",
		50:  "func",
		58:  "for",
		66:  "fallthrough",
		67:  "chan",
		73:  "map",
		74:  "break",
		82:  "var",
		85:  "else",
		91:  "goto",
		99:  "struct",
		103: "case",
		111: "switch",
		127: "const",
		144: "defer",
		150: "range",
		157: "import",
		166: "return",
		167: "select",
		208: "default",
		212: "package",
		223: "continue",
		241: "interface",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift505(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 0) + (name[1] << 5))
	if h > 242 {
		return false
	}
	return [...]string{
		7:   "goto",
		35:  "const",
		48:  "package",
		70:  "func",
		83:  "switch",
		100: "defer",
		101: "else",
		131: "continue",
		137: "import",
		146: "return",
		147: "select",
		164: "default",
		166: "fallthrough",
		173: "map",
		179: "struct",
		182: "var",
		194: "break",
		195: "case",
		199: "go",
		201: "interface",
		212: "type",
		227: "chan",
		230: "for",
		233: "if",
		242: "range",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift510(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) ^ ((name[0] << 1) + (name[1] << 0))
	if h > 244 {
		return false
	}
	return [...]string{
		3:   "for",
		13:  "fallthrough",
		17:  "type",
		27:  "map",
		33:  "goto",
		37:  "range",
		38:  "else",
		39:  "case",
		46:  "chan",
		57:  "func",
		65:  "return",
		67:  "select",
		73:  "const",
		77:  "defer",
		81:  "switch",
		82:  "struct",
		86:  "break",
		97:  "package",
		127: "import",
		141: "default",
		169: "continue",
		220: "interface",
		225: "go",
		237: "var",
		244: "if",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift603(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 0) + (name[1] << 3))
	if h > 222 {
		return false
	}
	return [...]string{
		5:   "else",
		12:  "default",
		27:  "continue",
		31:  "goto",
		35:  "chan",
		37:  "map",
		46:  "fallthrough",
		50:  "break",
		56:  "package",
		62:  "var",
		75:  "switch",
		83:  "struct",
		89:  "interface",
		91:  "const",
		107: "case",
		129: "import",
		140: "defer",
		159: "go",
		186: "range",
		188: "type",
		206: "func",
		217: "if",
		218: "return",
		219: "select",
		222: "for",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift605(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 0) + (name[1] << 5))
	if h > 233 {
		return false
	}
	return [...]string{
		4:   "defer",
		6:   "fallthrough",
		7:   "go",
		13:  "map",
		16:  "package",
		19:  "switch",
		22:  "var",
		41:  "if",
		67:  "case",
		70:  "for",
		73:  "import",
		82:  "return",
		83:  "select",
		84:  "type",
		98:  "break",
		99:  "chan",
		115: "struct",
		131: "continue",
		132: "default",
		135: "goto",
		146: "range",
		195: "const",
		198: "func",
		229: "else",
		233: "interface",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift606(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 0) + (name[1] << 6))
	if h > 246 {
		return false
	}
	return [...]string{
		34:  "break",
		35:  "case",
		38:  "func",
		39:  "go",
		41:  "interface",
		51:  "switch",
		52:  "type",
		99:  "chan",
		100: "defer",
		101: "else",
		102: "for",
		105: "if",
		114: "range",
		163: "continue",
		167: "goto",
		169: "import",
		178: "return",
		179: "select",
		227: "const",
		228: "default",
		230: "fallthrough",
		237: "map",
		240: "package",
		243: "struct",
		246: "var",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift616(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 1) + (name[1] << 6))
	if h > 228 {
		return false
	}
	return [...]string{
		6:   "continue",
		14:  "goto",
		18:  "import",
		36:  "return",
		38:  "select",
		70:  "const",
		72:  "default",
		76:  "fallthrough",
		90:  "map",
		96:  "package",
		102: "struct",
		108: "var",
		132: "break",
		134: "case",
		140: "func",
		142: "go",
		146: "interface",
		166: "switch",
		168: "type",
		198: "chan",
		200: "defer",
		202: "else",
		204: "for",
		210: "if",
		228: "range",
	}[h] == name
}

func TwoHashTableAlt_XorAdd_Shift630(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 3) + (name[1] << 0))
	if h > 217 {
		return false
	}
	return [...]string{
		5:   "default",
		17:  "fallthrough",
		31:  "for",
		49:  "range",
		68:  "else",
		69:  "func",
		87:  "goto",
		102: "interface",
		108: "struct",
		111: "switch",
		112: "chan",
		117: "return",
		119: "continue",
		121: "case",
		125: "select",
		133: "defer",
		145: "var",
		161: "package",
		162: "break",
		165: "import",
		174: "if",
		183: "const",
		201: "map",
		215: "go",
		217: "type",
	}[h] == name
}

func TwoHashTableAlt_XorOr_Shift034(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) | (name[1] << 4))
	if h > 238 {
		return false
	}
	return [...]string{
		12:  "case",
		42:  "if",
		43:  "fallthrough",
		52:  "type",
		53:  "break",
		100: "func",
		117: "defer",
		119: "default",
		123: "map",
		133: "range",
		151: "package",
		156: "chan",
		158: "import",
		163: "var",
		169: "interface",
		195: "for",
		198: "return",
		202: "go",
		204: "goto",
		206: "select",
		222: "struct",
		232: "continue",
		236: "else",
		237: "const",
		238: "switch",
	}[h] == name
}

func TwoHashTableAlt_XorOr_Shift035(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 3) | (name[1] << 5))
	if h > 253 {
		return false
	}
	return [...]string{
		27:  "fallthrough",
		28:  "chan",
		30:  "struct",
		54:  "return",
		60:  "case",
		62:  "select",
		75:  "map",
		85:  "break",
		126: "switch",
		132: "type",
		133: "defer",
		135: "default",
		137: "interface",
		138: "if",
		147: "var",
		148: "func",
		167: "package",
		172: "else",
		181: "range",
		211: "for",
		218: "go",
		220: "goto",
		238: "import",
		248: "continue",
		253: "const",
	}[h] == name
}

func TwoHashTableAlt_XorOr_Shift045(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) | (name[1] << 5))
	if h > 243 {
		return false
	}
	return [...]string{
		5:   "range",
		20:  "case",
		39:  "package",
		52:  "chan",
		54:  "import",
		67:  "var",
		75:  "fallthrough",
		82:  "if",
		89:  "interface",
		100: "type",
		101: "break",
		131: "for",
		134: "return",
		146: "go",
		148: "goto",
		150: "select",
		182: "struct",
		196: "func",
		212: "else",
		213: "const",
		214: "switch",
		216: "continue",
		229: "defer",
		231: "default",
		243: "map",
	}[h] == name
}

func TwoHashTableAlt_XorOr_Shift046(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) ^ ((name[0] << 4) | (name[1] << 6))
	if h > 248 {
		return false
	}
	return [...]string{
		4:   "type",
		5:   "defer",
		7:   "default",
		18:  "if",
		25:  "interface",
		35:  "var",
		36:  "func",
		43:  "fallthrough",
		52:  "chan",
		54:  "struct",
		71:  "package",
		84:  "else",
		101: "range",
		102: "return",
		116: "case",
		118: "select",
		147: "map",
		163: "for",
		165: "break",
		178: "go",
		180: "goto",
		214: "import",
		245: "const",
		246: "switch",
		248: "continue",
	}[h] == name
}

func TwoHashTableAlt_XorOr_Shift145(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) | (name[1] << 5))
	if h > 246 {
		return false
	}
	return [...]string{
		10:  "range",
		24:  "case",
		46:  "package",
		56:  "chan",
		60:  "import",
		70:  "var",
		82:  "interface",
		84:  "if",
		86:  "fallthrough",
		104: "type",
		106: "break",
		134: "for",
		140: "return",
		148: "go",
		152: "goto",
		156: "select",
		188: "struct",
		200: "func",
		208: "continue",
		216: "else",
		218: "const",
		220: "switch",
		234: "defer",
		238: "default",
		246: "map",
	}[h] == name
}

func TwoHashTableAlt_XorOr_Shift146(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) ^ ((name[0] << 4) | (name[1] << 6))
	if h > 252 {
		return false
	}
	return [...]string{
		8:   "type",
		10:  "defer",
		14:  "default",
		18:  "interface",
		20:  "if",
		38:  "var",
		40:  "func",
		54:  "fallthrough",
		56:  "chan",
		60:  "struct",
		78:  "package",
		88:  "else",
		106: "range",
		108: "return",
		120: "case",
		124: "select",
		150: "map",
		166: "for",
		170: "break",
		180: "go",
		184: "goto",
		220: "import",
		240: "continue",
		250: "const",
		252: "switch",
	}[h] == name
}

func TwoHashTableAlt_XorOr_Shift306(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) ^ ((name[0] << 0) | (name[1] << 6))
	if h > 249 {
		return false
	}
	return [...]string{
		35:  "case",
		38:  "func",
		44:  "defer",
		50:  "return",
		51:  "select",
		52:  "type",
		56:  "package",
		57:  "import",
		58:  "range",
		60:  "default",
		61:  "map",
		62:  "var",
		99:  "chan",
		101: "else",
		115: "struct",
		126: "fallthrough",
		167: "goto",
		171: "const",
		179: "switch",
		183: "go",
		190: "for",
		227: "continue",
		233: "interface",
		234: "break",
		249: "if",
	}[h] == name
}

func TwoHashTableAlt_XorOr_Shift603(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) ^ ((name[0] << 0) | (name[1] << 3))
	if h > 254 {
		return false
	}
	return [...]string{
		5:   "else",
		27:  "continue",
		31:  "goto",
		35:  "chan",
		76:  "defer",
		89:  "interface",
		91:  "const",
		107: "case",
		122: "range",
		129: "import",
		159: "go",
		188: "type",
		203: "switch",
		204: "default",
		206: "func",
		211: "struct",
		217: "if",
		218: "return",
		219: "select",
		222: "for",
		229: "map",
		238: "fallthrough",
		242: "break",
		248: "package",
		254: "var",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift016(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 1) ^ (name[1] << 6))
	if h > 224 {
		return false
	}
	return [...]string{
		2:   "case",
		7:   "fallthrough",
		8:   "func",
		13:  "defer",
		15:  "default",
		20:  "import",
		25:  "map",
		32:  "select",
		33:  "range",
		34:  "return",
		39:  "package",
		44:  "type",
		47:  "var",
		65:  "break",
		80:  "if",
		91:  "interface",
		131: "const",
		138: "goto",
		140: "go",
		142: "continue",
		143: "for",
		160: "switch",
		194: "chan",
		206: "else",
		224: "struct",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift024(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 2) ^ (name[1] << 4))
	if h > 236 {
		return false
	}
	return [...]string{
		6:   "if",
		8:   "chan",
		10:  "struct",
		26:  "select",
		30:  "return",
		58:  "switch",
		80:  "else",
		100: "type",
		114: "import",
		116: "continue",
		121: "const",
		136: "goto",
		139: "for",
		141: "interface",
		142: "go",
		152: "case",
		163: "fallthrough",
		173: "break",
		199: "map",
		215: "package",
		221: "range",
		229: "defer",
		231: "default",
		235: "var",
		236: "func",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift025(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 2) ^ (name[1] << 5))
	if h > 251 {
		return false
	}
	return [...]string{
		16:  "else",
		53:  "defer",
		55:  "default",
		60:  "func",
		66:  "import",
		74:  "struct",
		100: "continue",
		102: "if",
		105: "const",
		106: "select",
		109: "interface",
		110: "return",
		120: "goto",
		123: "for",
		126: "go",
		136: "chan",
		168: "case",
		170: "switch",
		179: "fallthrough",
		205: "break",
		215: "map",
		231: "package",
		237: "range",
		244: "type",
		251: "var",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift031(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 3) ^ (name[1] << 1))
	if h > 254 {
		return false
	}
	return [...]string{
		4:   "else",
		13:  "for",
		18:  "goto",
		20:  "go",
		22:  "if",
		30:  "func",
		36:  "import",
		41:  "map",
		45:  "interface",
		69:  "package",
		87:  "range",
		92:  "return",
		100: "select",
		113: "var",
		128: "switch",
		134: "struct",
		150: "type",
		222: "case",
		236: "chan",
		237: "default",
		239: "defer",
		241: "break",
		243: "const",
		249: "fallthrough",
		254: "continue",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift033(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 3) ^ (name[1] << 3))
	if h > 220 {
		return false
	}
	return [...]string{
		36:  "case",
		51:  "fallthrough",
		62:  "struct",
		77:  "defer",
		79:  "default",
		86:  "switch",
		92:  "chan",
		108: "type",
		115: "map",
		122: "if",
		140: "else",
		143: "package",
		149: "const",
		152: "continue",
		157: "range",
		165: "break",
		171: "for",
		177: "interface",
		178: "go",
		180: "goto",
		182: "import",
		187: "var",
		190: "return",
		198: "select",
		220: "func",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 4) ^ (name[1] << 1))
	if h > 252 {
		return false
	}
	return [...]string{
		1:   "break",
		4:   "chan",
		6:   "continue",
		11:  "const",
		13:  "default",
		15:  "defer",
		24:  "switch",
		30:  "struct",
		33:  "var",
		41:  "fallthrough",
		44:  "else",
		54:  "type",
		61:  "for",
		74:  "goto",
		76:  "go",
		78:  "func",
		94:  "if",
		101: "interface",
		108: "import",
		145: "map",
		197: "package",
		231: "range",
		236: "return",
		246: "case",
		252: "select",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift042(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 4) ^ (name[1] << 2))
	if h > 239 {
		return false
	}
	return [...]string{
		4:   "else",
		6:   "struct",
		10:  "switch",
		31:  "for",
		32:  "type",
		40:  "goto",
		42:  "if",
		46:  "go",
		48:  "func",
		65:  "interface",
		66:  "import",
		87:  "map",
		131: "package",
		161: "range",
		176: "case",
		178: "return",
		194: "select",
		209: "defer",
		211: "default",
		212: "chan",
		228: "continue",
		231: "var",
		233: "const",
		237: "break",
		239: "fallthrough",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 5) ^ (name[1] << 0))
	if h > 253 {
		return false
	}
	return [...]string{
		2:   "map",
		8:   "else",
		34:  "var",
		42:  "fallthrough",
		44:  "for",
		49:  "func",
		75:  "goto",
		77:  "go",
		102: "package",
		132: "if",
		135: "interface",
		139: "import",
		163: "return",
		164: "range",
		183: "break",
		195: "select",
		197: "case",
		199: "continue",
		202: "const",
		204: "chan",
		209: "switch",
		210: "struct",
		224: "defer",
		226: "default",
		253: "type",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 5) ^ (name[1] << 1))
	if h > 252 {
		return false
	}
	return [...]string{
		7:   "range",
		12:  "return",
		33:  "break",
		38:  "case",
		44:  "select",
		52:  "chan",
		54:  "continue",
		59:  "const",
		72:  "switch",
		77:  "default",
		78:  "struct",
		79:  "defer",
		97:  "map",
		118: "type",
		124: "else",
		129: "var",
		137: "fallthrough",
		157: "for",
		174: "func",
		186: "goto",
		188: "go",
		197: "package",
		238: "if",
		245: "interface",
		252: "import",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 6) ^ (name[1] << 0))
	if h > 247 {
		return false
	}
	return [...]string{
		35:  "select",
		37:  "case",
		39:  "continue",
		42:  "const",
		43:  "goto",
		44:  "chan",
		45:  "go",
		49:  "switch",
		50:  "struct",
		96:  "defer",
		98:  "default",
		102: "package",
		125: "type",
		162: "map",
		164: "if",
		167: "interface",
		168: "else",
		171: "import",
		226: "var",
		227: "return",
		228: "range",
		234: "fallthrough",
		236: "for",
		241: "func",
		247: "break",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 6) ^ (name[1] << 2))
	if h > 244 {
		return false
	}
	return [...]string{
		1:   "range",
		7:   "var",
		15:  "fallthrough",
		18:  "return",
		63:  "for",
		64:  "case",
		77:  "break",
		80:  "func",
		82:  "select",
		100: "chan",
		116: "continue",
		120: "goto",
		121: "const",
		126: "go",
		131: "package",
		145: "defer",
		147: "default",
		150: "struct",
		154: "switch",
		199: "map",
		218: "if",
		224: "type",
		241: "interface",
		242: "import",
		244: "else",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 7) ^ (name[1] << 3))
	if h > 253 {
		return false
	}
	return [...]string{
		3:   "fallthrough",
		11:  "var",
		13:  "range",
		15:  "package",
		38:  "struct",
		45:  "defer",
		46:  "return",
		47:  "default",
		62:  "switch",
		123: "for",
		139: "map",
		140: "case",
		149: "break",
		172: "func",
		174: "select",
		178: "if",
		196: "chan",
		204: "type",
		228: "else",
		238: "import",
		240: "continue",
		249: "interface",
		250: "go",
		252: "goto",
		253: "const",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift120(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) + ((name[0] << 2) ^ (name[1] << 0))
	if h > 255 {
		return false
	}
	return [...]string{
		0:   "interface",
		1:   "for",
		3:   "goto",
		5:   "func",
		8:   "else",
		14:  "if",
		15:  "go",
		19:  "map",
		29:  "import",
		33:  "return",
		35:  "range",
		47:  "package",
		61:  "select",
		63:  "var",
		65:  "type",
		76:  "struct",
		79:  "switch",
		229: "case",
		235: "continue",
		239: "fallthrough",
		240: "break",
		241: "const",
		251: "default",
		252: "chan",
		255: "defer",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) + ((name[0] << 6) ^ (name[1] << 1))
	if h > 250 {
		return false
	}
	return [...]string{
		4:   "map",
		8:   "if",
		14:  "interface",
		16:  "else",
		22:  "import",
		68:  "var",
		70:  "return",
		72:  "range",
		84:  "fallthrough",
		88:  "for",
		98:  "func",
		110: "break",
		134: "select",
		138: "case",
		142: "continue",
		148: "const",
		150: "goto",
		152: "chan",
		154: "go",
		162: "switch",
		164: "struct",
		192: "defer",
		196: "default",
		204: "package",
		250: "type",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift231(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) + ((name[0] << 3) ^ (name[1] << 1))
	if h > 254 {
		return false
	}
	return [...]string{
		0:   "interface",
		2:   "for",
		6:   "goto",
		10:  "func",
		16:  "else",
		28:  "if",
		30:  "go",
		38:  "map",
		58:  "import",
		66:  "return",
		70:  "range",
		94:  "package",
		122: "select",
		126: "var",
		130: "type",
		152: "struct",
		158: "switch",
		202: "case",
		214: "continue",
		222: "fallthrough",
		224: "break",
		226: "const",
		246: "default",
		248: "chan",
		254: "defer",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift260(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) + ((name[0] << 6) ^ (name[1] << 0))
	if h > 253 {
		return false
	}
	return [...]string{
		15:  "continue",
		39:  "go",
		44:  "struct",
		47:  "switch",
		49:  "case",
		56:  "chan",
		59:  "const",
		61:  "select",
		63:  "goto",
		105: "type",
		113: "defer",
		121: "default",
		125: "package",
		138: "interface",
		173: "map",
		174: "if",
		181: "import",
		188: "else",
		205: "fallthrough",
		227: "for",
		229: "func",
		230: "break",
		237: "var",
		245: "range",
		253: "return",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift303(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) + ((name[0] << 0) ^ (name[1] << 3))
	if h > 255 {
		return false
	}
	return [...]string{
		27:  "switch",
		28:  "type",
		35:  "struct",
		46:  "func",
		54:  "fallthrough",
		64:  "package",
		75:  "case",
		82:  "range",
		102: "var",
		109: "map",
		131: "chan",
		137: "if",
		145: "interface",
		155: "continue",
		164: "defer",
		170: "return",
		171: "select",
		180: "default",
		198: "for",
		207: "go",
		218: "break",
		225: "import",
		229: "else",
		243: "const",
		255: "goto",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift460(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) + ((name[0] << 6) ^ (name[1] << 0))
	if h > 236 {
		return false
	}
	return [...]string{
		15:  "go",
		17:  "package",
		21:  "default",
		53:  "defer",
		57:  "type",
		62:  "interface",
		69:  "select",
		81:  "fallthrough",
		84:  "struct",
		87:  "switch",
		97:  "case",
		104: "chan",
		111: "goto",
		127: "const",
		133: "return",
		134: "if",
		145: "map",
		162: "break",
		175: "continue",
		177: "range",
		181: "func",
		205: "import",
		209: "var",
		223: "for",
		236: "else",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift500(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) + ((name[0] << 0) ^ (name[1] << 0))
	if h > 247 {
		return false
	}
	return [...]string{
		22:  "import",
		23:  "return",
		24:  "select",
		39:  "struct",
		41:  "default",
		42:  "switch",
		49:  "package",
		68:  "case",
		75:  "chan",
		81:  "else",
		86:  "goto",
		91:  "func",
		105: "defer",
		109: "type",
		114: "const",
		115: "range",
		116: "break",
		143: "if",
		150: "go",
		167: "fallthrough",
		174: "map",
		181: "for",
		183: "var",
		210: "continue",
		247: "interface",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift601(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 0) ^ (name[1] << 1))
	if h > 248 {
		return false
	}
	return [...]string{
		1:   "const",
		5:   "interface",
		6:   "break",
		37:  "case",
		51:  "chan",
		61:  "else",
		65:  "continue",
		69:  "goto",
		80:  "func",
		102: "type",
		110: "defer",
		116: "range",
		132: "for",
		181: "if",
		188: "return",
		189: "select",
		195: "import",
		197: "go",
		219: "struct",
		225: "switch",
		232: "fallthrough",
		238: "default",
		239: "map",
		242: "package",
		248: "var",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift603(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 0) ^ (name[1] << 3))
	if h > 223 {
		return false
	}
	return [...]string{
		14:  "func",
		25:  "if",
		26:  "return",
		27:  "select",
		30:  "for",
		58:  "range",
		60:  "type",
		76:  "default",
		81:  "import",
		95:  "go",
		107: "case",
		147: "struct",
		153: "interface",
		155: "const",
		163: "chan",
		171: "switch",
		174: "fallthrough",
		178: "break",
		181: "map",
		184: "package",
		190: "var",
		197: "else",
		204: "defer",
		219: "continue",
		223: "goto",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift605(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 0) ^ (name[1] << 5))
	if h > 229 {
		return false
	}
	return [...]string{
		3:   "const",
		6:   "func",
		67:  "continue",
		68:  "defer",
		70:  "fallthrough",
		71:  "goto",
		77:  "map",
		80:  "package",
		86:  "var",
		99:  "chan",
		105: "interface",
		115: "struct",
		131: "case",
		134: "for",
		137: "import",
		146: "return",
		147: "select",
		148: "type",
		169: "if",
		196: "default",
		199: "go",
		210: "range",
		211: "switch",
		226: "break",
		229: "else",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift620(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 2) ^ (name[1] << 0))
	if h > 251 {
		return false
	}
	return [...]string{
		0:   "else",
		11:  "goto",
		13:  "func",
		53:  "default",
		57:  "fallthrough",
		73:  "type",
		82:  "interface",
		105: "range",
		138: "if",
		139: "go",
		145: "import",
		173: "return",
		177: "select",
		181: "defer",
		186: "break",
		187: "const",
		192: "struct",
		195: "switch",
		199: "for",
		213: "map",
		225: "package",
		237: "case",
		244: "chan",
		249: "var",
		251: "continue",
	}[h] == name
}

func TwoHashTableAlt_AddXor_Shift630(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 3) ^ (name[1] << 0))
	if h > 246 {
		return false
	}
	return [...]string{
		9:   "map",
		25:  "type",
		33:  "package",
		39:  "go",
		46:  "if",
		53:  "import",
		69:  "default",
		81:  "fallthrough",
		95:  "for",
		117: "return",
		121: "case",
		125: "select",
		128: "chan",
		135: "continue",
		140: "struct",
		143: "switch",
		148: "else",
		165: "func",
		167: "goto",
		177: "range",
		194: "break",
		197: "defer",
		199: "const",
		209: "var",
		246: "interface",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift022(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 2) + (name[1] << 2))
	if h > 184 {
		return false
	}
	return [...]string{
		20:  "case",
		39:  "fallthrough",
		41:  "defer",
		43:  "default",
		48:  "chan",
		59:  "map",
		62:  "if",
		72:  "else",
		75:  "package",
		77:  "const",
		80:  "continue",
		81:  "range",
		85:  "break",
		87:  "for",
		90:  "go",
		92:  "goto",
		94:  "import",
		95:  "var",
		98:  "return",
		101: "interface",
		102: "select",
		112: "func",
		162: "struct",
		174: "switch",
		184: "type",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift024(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 2) + (name[1] << 4))
	if h > 236 {
		return false
	}
	return [...]string{
		6:   "if",
		16:  "chan",
		18:  "struct",
		30:  "return",
		34:  "select",
		66:  "switch",
		88:  "else",
		100: "type",
		122: "import",
		129: "const",
		132: "continue",
		139: "for",
		141: "interface",
		142: "go",
		144: "goto",
		160: "case",
		173: "break",
		179: "fallthrough",
		199: "map",
		215: "package",
		221: "range",
		229: "defer",
		231: "default",
		235: "var",
		236: "func",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift025(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 2) + (name[1] << 5))
	if h > 251 {
		return false
	}
	return [...]string{
		24:  "else",
		53:  "defer",
		55:  "default",
		60:  "func",
		74:  "import",
		82:  "struct",
		102: "if",
		109: "interface",
		110: "return",
		113: "const",
		114: "select",
		116: "continue",
		123: "for",
		126: "go",
		128: "goto",
		144: "chan",
		176: "case",
		178: "switch",
		195: "fallthrough",
		205: "break",
		215: "map",
		231: "package",
		237: "range",
		244: "type",
		251: "var",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift033(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 3) + (name[1] << 3))
	if h > 220 {
		return false
	}
	return [...]string{
		36:  "case",
		62:  "struct",
		67:  "fallthrough",
		77:  "defer",
		79:  "default",
		86:  "switch",
		92:  "chan",
		108: "type",
		115: "map",
		122: "if",
		140: "else",
		143: "package",
		149: "const",
		152: "continue",
		157: "range",
		165: "break",
		171: "for",
		178: "go",
		180: "goto",
		182: "import",
		187: "var",
		190: "return",
		193: "interface",
		198: "select",
		220: "func",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 4) + (name[1] << 1))
	if h > 246 {
		return false
	}
	return [...]string{
		0:   "select",
		4:   "chan",
		9:   "break",
		15:  "defer",
		17:  "default",
		19:  "const",
		22:  "continue",
		30:  "struct",
		36:  "switch",
		37:  "var",
		44:  "else",
		45:  "fallthrough",
		54:  "type",
		65:  "for",
		78:  "func",
		80:  "go",
		82:  "goto",
		94:  "if",
		112: "import",
		117: "interface",
		149: "map",
		201: "package",
		231: "range",
		240: "return",
		246: "case",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift042(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 4) + (name[1] << 2))
	if h > 244 {
		return false
	}
	return [...]string{
		4:   "else",
		6:   "struct",
		18:  "switch",
		31:  "for",
		40:  "type",
		42:  "if",
		46:  "go",
		48:  "goto",
		56:  "func",
		74:  "import",
		81:  "interface",
		87:  "map",
		139: "package",
		169: "range",
		184: "case",
		186: "return",
		202: "select",
		212: "chan",
		217: "defer",
		219: "default",
		231: "var",
		237: "break",
		239: "fallthrough",
		241: "const",
		244: "continue",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 5) + (name[1] << 0))
	if h > 253 {
		return false
	}
	return [...]string{
		4:   "map",
		16:  "else",
		36:  "var",
		44:  "fallthrough",
		50:  "for",
		57:  "func",
		81:  "go",
		83:  "goto",
		104: "package",
		136: "if",
		147: "import",
		151: "interface",
		166: "range",
		171: "return",
		183: "break",
		197: "case",
		203: "select",
		204: "chan",
		212: "const",
		215: "continue",
		218: "struct",
		221: "switch",
		234: "defer",
		236: "default",
		253: "type",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 5) + (name[1] << 1))
	if h > 238 {
		return false
	}
	return [...]string{
		0:   "import",
		5:   "interface",
		7:   "range",
		16:  "return",
		38:  "case",
		41:  "break",
		48:  "select",
		52:  "chan",
		67:  "const",
		70:  "continue",
		78:  "struct",
		79:  "defer",
		81:  "default",
		84:  "switch",
		101: "map",
		118: "type",
		124: "else",
		133: "var",
		141: "fallthrough",
		161: "for",
		174: "func",
		192: "go",
		194: "goto",
		201: "package",
		238: "if",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 6) + (name[1] << 0))
	if h > 249 {
		return false
	}
	return [...]string{
		37:  "case",
		43:  "select",
		44:  "chan",
		49:  "go",
		51:  "goto",
		52:  "const",
		55:  "continue",
		58:  "struct",
		61:  "switch",
		104: "package",
		106: "defer",
		108: "default",
		125: "type",
		164: "map",
		168: "if",
		176: "else",
		179: "import",
		183: "interface",
		228: "var",
		230: "range",
		235: "return",
		236: "fallthrough",
		242: "for",
		247: "break",
		249: "func",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift061(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 6) + (name[1] << 1))
	if h > 246 {
		return false
	}
	return [...]string{
		5:   "map",
		14:  "if",
		28:  "else",
		32:  "import",
		37:  "interface",
		69:  "var",
		71:  "range",
		77:  "fallthrough",
		80:  "return",
		97:  "for",
		105: "break",
		110: "func",
		134: "case",
		144: "select",
		148: "chan",
		160: "go",
		162: "goto",
		163: "const",
		166: "continue",
		174: "struct",
		180: "switch",
		201: "package",
		207: "defer",
		209: "default",
		246: "type",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 6) + (name[1] << 2))
	if h > 250 {
		return false
	}
	return [...]string{
		1:   "interface",
		7:   "var",
		9:   "range",
		15:  "fallthrough",
		26:  "return",
		63:  "for",
		72:  "case",
		77:  "break",
		88:  "func",
		90:  "select",
		100: "chan",
		126: "go",
		128: "goto",
		129: "const",
		132: "continue",
		139: "package",
		150: "struct",
		153: "defer",
		155: "default",
		162: "switch",
		199: "map",
		218: "if",
		232: "type",
		244: "else",
		250: "import",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 7) + (name[1] << 3))
	if h > 253 {
		return false
	}
	return [...]string{
		0:   "continue",
		11:  "var",
		13:  "range",
		15:  "package",
		19:  "fallthrough",
		38:  "struct",
		45:  "defer",
		46:  "return",
		47:  "default",
		62:  "switch",
		123: "for",
		139: "map",
		140: "case",
		149: "break",
		172: "func",
		174: "select",
		178: "if",
		196: "chan",
		204: "type",
		228: "else",
		238: "import",
		249: "interface",
		250: "go",
		252: "goto",
		253: "const",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift130(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) + ((name[0] << 3) + (name[1] << 0))
	if h > 251 {
		return false
	}
	return [...]string{
		1:   "return",
		9:   "select",
		23:  "var",
		24:  "struct",
		27:  "switch",
		33:  "type",
		129: "case",
		136: "chan",
		140: "break",
		143: "defer",
		145: "const",
		147: "default",
		151: "continue",
		156: "else",
		165: "for",
		167: "fallthrough",
		171: "go",
		173: "func",
		175: "goto",
		178: "if",
		193: "import",
		200: "interface",
		207: "map",
		239: "package",
		251: "range",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift133(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) + ((name[0] << 3) + (name[1] << 3))
	if h > 224 {
		return false
	}
	return [...]string{
		40:  "case",
		68:  "struct",
		78:  "fallthrough",
		82:  "defer",
		86:  "default",
		92:  "switch",
		96:  "chan",
		112: "type",
		118: "map",
		124: "if",
		144: "else",
		150: "package",
		154: "const",
		160: "continue",
		162: "range",
		170: "break",
		174: "for",
		180: "go",
		184: "goto",
		188: "import",
		190: "var",
		196: "return",
		202: "interface",
		204: "select",
		224: "func",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift150(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) + ((name[0] << 5) + (name[1] << 0))
	if h > 243 {
		return false
	}
	return [...]string{
		1:   "type",
		7:   "map",
		20:  "else",
		39:  "var",
		53:  "for",
		55:  "fallthrough",
		61:  "func",
		83:  "go",
		87:  "goto",
		111: "package",
		138: "if",
		153: "import",
		160: "interface",
		171: "range",
		177: "return",
		188: "break",
		201: "case",
		208: "chan",
		209: "select",
		217: "const",
		223: "continue",
		224: "struct",
		227: "switch",
		239: "defer",
		243: "default",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 1) + ((name[0] << 6) + (name[1] << 1))
	if h > 250 {
		return false
	}
	return [...]string{
		8:   "map",
		16:  "if",
		32:  "else",
		38:  "import",
		46:  "interface",
		72:  "var",
		76:  "range",
		86:  "return",
		88:  "fallthrough",
		100: "for",
		110: "break",
		114: "func",
		138: "case",
		150: "select",
		152: "chan",
		162: "go",
		166: "goto",
		168: "const",
		174: "continue",
		180: "struct",
		186: "switch",
		208: "package",
		212: "defer",
		216: "default",
		250: "type",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift230(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 2) + ((name[0] << 3) + (name[1] << 0))
	if h > 253 {
		return false
	}
	return [...]string{
		5:   "range",
		13:  "return",
		21:  "select",
		29:  "var",
		36:  "struct",
		39:  "switch",
		41:  "type",
		137: "case",
		144: "chan",
		150: "break",
		153: "defer",
		155: "const",
		161: "default",
		164: "else",
		167: "continue",
		171: "for",
		175: "go",
		181: "func",
		182: "if",
		183: "goto",
		189: "fallthrough",
		205: "import",
		213: "map",
		218: "interface",
		253: "package",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift302(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) + ((name[0] << 0) + (name[1] << 2))
	if h > 127 {
		return false
	}
	return [...]string{
		7:   "case",
		9:   "map",
		17:  "if",
		18:  "var",
		30:  "range",
		32:  "defer",
		35:  "chan",
		44:  "package",
		48:  "default",
		51:  "go",
		53:  "else",
		54:  "return",
		55:  "select",
		58:  "for",
		66:  "fallthrough",
		67:  "goto",
		71:  "const",
		77:  "import",
		82:  "break",
		90:  "func",
		95:  "continue",
		105: "interface",
		115: "struct",
		120: "type",
		127: "switch",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift303(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 3) + ((name[0] << 0) + (name[1] << 3))
	if h > 255 {
		return false
	}
	return [...]string{
		1:   "import",
		3:   "const",
		26:  "break",
		27:  "continue",
		33:  "interface",
		46:  "func",
		67:  "struct",
		91:  "switch",
		92:  "type",
		139: "case",
		141: "map",
		150: "var",
		162: "range",
		169: "if",
		176: "package",
		180: "defer",
		195: "chan",
		196: "default",
		198: "fallthrough",
		202: "return",
		203: "select",
		229: "else",
		239: "go",
		246: "for",
		255: "goto",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift400(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) + ((name[0] << 0) + (name[1] << 0))
	if h > 254 {
		return false
	}
	return [...]string{
		4:   "case",
		5:   "for",
		7:   "var",
		11:  "chan",
		17:  "else",
		22:  "goto",
		25:  "defer",
		27:  "func",
		34:  "const",
		35:  "range",
		36:  "break",
		45:  "type",
		54:  "import",
		55:  "return",
		56:  "select",
		57:  "default",
		65:  "package",
		71:  "struct",
		74:  "switch",
		82:  "continue",
		103: "interface",
		119: "fallthrough",
		239: "if",
		246: "go",
		254: "map",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift403(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) + ((name[0] << 0) + (name[1] << 3))
	if h > 255 {
		return false
	}
	return [...]string{
		5:   "else",
		14:  "for",
		30:  "fallthrough",
		31:  "goto",
		43:  "const",
		49:  "import",
		66:  "break",
		78:  "func",
		91:  "continue",
		105: "interface",
		115: "struct",
		124: "type",
		139: "switch",
		165: "map",
		171: "case",
		174: "var",
		185: "if",
		202: "range",
		220: "defer",
		227: "chan",
		232: "package",
		250: "return",
		251: "select",
		252: "default",
		255: "go",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift413(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 4) + ((name[0] << 1) + (name[1] << 3))
	if h > 254 {
		return false
	}
	return [...]string{
		14:  "case",
		18:  "map",
		34:  "if",
		36:  "var",
		60:  "range",
		64:  "defer",
		70:  "chan",
		88:  "package",
		96:  "default",
		102: "go",
		106: "else",
		108: "return",
		110: "select",
		116: "for",
		132: "fallthrough",
		134: "goto",
		142: "const",
		154: "import",
		164: "break",
		180: "func",
		190: "continue",
		210: "interface",
		230: "struct",
		240: "type",
		254: "switch",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift500(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) + ((name[0] << 0) + (name[1] << 0))
	if h > 247 {
		return false
	}
	return [...]string{
		15:  "if",
		22:  "go",
		39:  "fallthrough",
		46:  "map",
		53:  "for",
		55:  "var",
		68:  "case",
		75:  "chan",
		81:  "else",
		86:  "goto",
		91:  "func",
		105: "defer",
		109: "type",
		114: "const",
		115: "range",
		116: "break",
		150: "import",
		151: "return",
		152: "select",
		167: "struct",
		169: "default",
		170: "switch",
		177: "package",
		210: "continue",
		247: "interface",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift511(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 5) + ((name[0] << 1) + (name[1] << 1))
	if h > 252 {
		return false
	}
	return [...]string{
		8:   "case",
		10:  "for",
		14:  "var",
		22:  "chan",
		34:  "else",
		44:  "goto",
		50:  "defer",
		54:  "func",
		68:  "const",
		70:  "range",
		72:  "break",
		90:  "type",
		108: "import",
		110: "return",
		112: "select",
		114: "default",
		130: "package",
		142: "struct",
		148: "switch",
		164: "continue",
		206: "interface",
		222: "if",
		236: "go",
		238: "fallthrough",
		252: "map",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift601(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 0) + (name[1] << 1))
	if h > 248 {
		return false
	}
	return [...]string{
		4:   "for",
		37:  "case",
		51:  "chan",
		61:  "else",
		65:  "continue",
		69:  "goto",
		80:  "func",
		102: "type",
		110: "defer",
		116: "range",
		129: "const",
		133: "interface",
		134: "break",
		181: "if",
		188: "return",
		189: "select",
		195: "import",
		197: "go",
		219: "struct",
		225: "switch",
		232: "fallthrough",
		238: "default",
		239: "map",
		242: "package",
		248: "var",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift602(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 0) + (name[1] << 2))
	if h > 231 {
		return false
	}
	return [...]string{
		3:   "chan",
		21:  "else",
		31:  "continue",
		35:  "goto",
		54:  "range",
		56:  "defer",
		58:  "func",
		88:  "type",
		95:  "const",
		97:  "interface",
		106: "break",
		129: "if",
		134: "return",
		135: "select",
		157: "import",
		163: "go",
		170: "fallthrough",
		177: "map",
		180: "package",
		184: "default",
		186: "var",
		195: "struct",
		207: "switch",
		226: "for",
		231: "case",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift620(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 2) + (name[1] << 0))
	if h > 251 {
		return false
	}
	return [...]string{
		0:   "else",
		11:  "goto",
		13:  "func",
		53:  "defer",
		58:  "break",
		59:  "const",
		73:  "type",
		82:  "interface",
		105: "range",
		138: "if",
		139: "go",
		145: "import",
		173: "return",
		177: "select",
		181: "default",
		185: "fallthrough",
		192: "struct",
		195: "switch",
		199: "for",
		213: "map",
		225: "package",
		237: "case",
		244: "chan",
		249: "var",
		251: "continue",
	}[h] == name
}

func TwoHashTableAlt_AddAdd_Shift630(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 6) + ((name[0] << 3) + (name[1] << 0))
	if h > 246 {
		return false
	}
	return [...]string{
		25:  "type",
		39:  "go",
		46:  "if",
		49:  "range",
		53:  "import",
		69:  "default",
		81:  "fallthrough",
		95:  "for",
		117: "return",
		121: "case",
		125: "select",
		128: "chan",
		135: "continue",
		137: "map",
		140: "struct",
		143: "switch",
		148: "else",
		161: "package",
		165: "func",
		167: "goto",
		194: "break",
		197: "defer",
		199: "const",
		209: "var",
		246: "interface",
	}[h] == name
}

func TwoHashTableAlt_AddOr_Shift033(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) + ((name[0] << 3) | (name[1] << 3))
	if h > 220 {
		return false
	}
	return [...]string{
		36:  "case",
		59:  "fallthrough",
		62:  "struct",
		77:  "defer",
		79:  "default",
		86:  "switch",
		92:  "chan",
		108: "type",
		115: "map",
		122: "if",
		140: "else",
		143: "package",
		149: "const",
		152: "continue",
		157: "range",
		165: "break",
		171: "for",
		178: "go",
		180: "goto",
		182: "import",
		185: "interface",
		187: "var",
		190: "return",
		198: "select",
		220: "func",
	}[h] == name
}

func TwoHashTableAlt_OrXor_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) | ((name[0] << 7) ^ (name[1] << 3))
	if h > 253 {
		return false
	}
	return [...]string{
		3:   "fallthrough",
		11:  "var",
		13:  "range",
		15:  "package",
		45:  "defer",
		46:  "return",
		47:  "default",
		123: "for",
		139: "map",
		140: "case",
		149: "break",
		166: "struct",
		172: "func",
		174: "select",
		178: "if",
		190: "switch",
		196: "chan",
		204: "type",
		228: "else",
		238: "import",
		240: "continue",
		249: "interface",
		250: "go",
		252: "goto",
		253: "const",
	}[h] == name
}

func TwoHashTableAlt_OrAdd_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	h := (byte(len(name)) << 0) | ((name[0] << 7) + (name[1] << 3))
	if h > 253 {
		return false
	}
	return [...]string{
		0:   "continue",
		11:  "var",
		13:  "range",
		15:  "package",
		19:  "fallthrough",
		45:  "defer",
		46:  "return",
		47:  "default",
		123: "for",
		139: "map",
		140: "case",
		149: "break",
		166: "struct",
		172: "func",
		174: "select",
		178: "if",
		190: "switch",
		196: "chan",
		204: "type",
		228: "else",
		238: "import",
		249: "interface",
		250: "go",
		252: "goto",
		253: "const",
	}[h] == name
}
