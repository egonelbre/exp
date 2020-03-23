package check

import "encoding/binary"

func Uint64Switch(name string) bool {
	var key [8]byte
	copy(key[:], name)
	h := binary.LittleEndian.Uint64(key[:])
	switch h {
	case 461195539042:
		return len(name) == 5
	case 1702060387:
		return len(name) == 4
	case 1851877475:
		return len(name) == 4
	case 500152823651:
		return len(name) == 5
	case 7310870969309884259:
		return len(name) == 8
	case 32770348699510116:
		return len(name) == 7
	case 491327481188:
		return len(name) == 5
	case 1702063205:
		return len(name) == 4
	case 8030595934799552870:
		return name[8:] == "ugh"
	case 7499622:
		return len(name) == 3
	case 1668183398:
		return len(name) == 4
	case 28519:
		return len(name) == 2
	case 1869901671:
		return len(name) == 4
	case 26217:
		return len(name) == 2
	case 128034844732777:
		return len(name) == 6
	case 7161117524010233449:
		return name[8:] == "e"
	case 7364973:
		return len(name) == 3
	case 28542640758940016:
		return len(name) == 7
	case 435526984050:
		return len(name) == 5
	case 121437875889522:
		return len(name) == 6
	case 127970252186995:
		return len(name) == 6
	case 127970521019507:
		return len(name) == 6
	case 114776364119923:
		return len(name) == 6
	case 1701869940:
		return len(name) == 4
	case 7496054:
		return len(name) == 3
	}
	return false
}
