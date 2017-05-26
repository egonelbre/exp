package element

type Element byte

const (
	Invalid = Element(0)
	Ac      = Element(1)
	Ag      = Element(2)
	Al      = Element(3)
	Am      = Element(4)
	Ar      = Element(5)
	As      = Element(6)
	At      = Element(7)
	Au      = Element(8)
	B       = Element(9)
	Ba      = Element(10)
	Be      = Element(11)
	Bh      = Element(12)
	Bi      = Element(13)
	Bk      = Element(14)
	Br      = Element(15)
	C       = Element(16)
	Ca      = Element(17)
	Cd      = Element(18)
	Ce      = Element(19)
	Cf      = Element(20)
	Cl      = Element(21)
	Cm      = Element(22)
	Co      = Element(23)
	Cr      = Element(24)
	Cs      = Element(25)
	Cu      = Element(26)
	Db      = Element(27)
	Ds      = Element(28)
	Dy      = Element(29)
	Er      = Element(30)
	Es      = Element(31)
	Eu      = Element(32)
	F       = Element(33)
	Fe      = Element(34)
	Fm      = Element(35)
	Fr      = Element(36)
	Ga      = Element(37)
	Gd      = Element(38)
	Ge      = Element(39)
	H       = Element(40)
	He      = Element(41)
	Hf      = Element(42)
	Hg      = Element(43)
	Ho      = Element(44)
	Hs      = Element(45)
	I       = Element(46)
	In      = Element(47)
	Ir      = Element(48)
	K       = Element(49)
	Kr      = Element(50)
	La      = Element(51)
	Li      = Element(52)
	Lr      = Element(53)
	Lu      = Element(54)
	Md      = Element(55)
	Mg      = Element(56)
	Mn      = Element(57)
	Mo      = Element(58)
	Mt      = Element(59)
	N       = Element(60)
	Na      = Element(61)
	Nb      = Element(62)
	Nd      = Element(63)
	Ne      = Element(64)
	Ni      = Element(65)
	No      = Element(66)
	Np      = Element(67)
	O       = Element(68)
	Os      = Element(69)
	P       = Element(70)
	Pa      = Element(71)
	Pb      = Element(72)
	Pd      = Element(73)
	Pm      = Element(74)
	Po      = Element(75)
	Pr      = Element(76)
	Pt      = Element(77)
	Pu      = Element(78)
	Ra      = Element(79)
	Rb      = Element(80)
	Re      = Element(81)
	Rf      = Element(82)
	Rg      = Element(83)
	Rh      = Element(84)
	Rn      = Element(85)
	Ru      = Element(86)
	S       = Element(87)
	Sb      = Element(88)
	Sc      = Element(89)
	Se      = Element(90)
	Sg      = Element(91)
	Si      = Element(92)
	Sm      = Element(93)
	Sn      = Element(94)
	Sr      = Element(95)
	Ta      = Element(96)
	Tb      = Element(97)
	Tc      = Element(98)
	Te      = Element(99)
	Th      = Element(100)
	Ti      = Element(101)
	Tl      = Element(102)
	Tm      = Element(103)
	U       = Element(104)
	Uub     = Element(105)
	Uuh     = Element(106)
	Uuo     = Element(107)
	Uup     = Element(108)
	Uuq     = Element(109)
	Uus     = Element(110)
	Uut     = Element(111)
	V       = Element(112)
	W       = Element(113)
	Xe      = Element(114)
	Y       = Element(115)
	Yb      = Element(116)
	Zn      = Element(117)
	Zr      = Element(118)
)

var AtomSymbol = [...]string{"<INVALID>",
	"Ac", "Ag", "Al", "Am", "Ar", "As", "At", "Au", "B", "Ba",
	"Be", "Bh", "Bi", "Bk", "Br", "C", "Ca", "Cd", "Ce", "Cf",
	"Cl", "Cm", "Co", "Cr", "Cs", "Cu", "Db", "Ds", "Dy", "Er",
	"Es", "Eu", "F", "Fe", "Fm", "Fr", "Ga", "Gd", "Ge", "H",
	"He", "Hf", "Hg", "Ho", "Hs", "I", "In", "Ir", "K", "Kr",
	"La", "Li", "Lr", "Lu", "Md", "Mg", "Mn", "Mo", "Mt", "N",
	"Na", "Nb", "Nd", "Ne", "Ni", "No", "Np", "O", "Os", "P",
	"Pa", "Pb", "Pd", "Pm", "Po", "Pr", "Pt", "Pu", "Ra", "Rb",
	"Re", "Rf", "Rg", "Rh", "Rn", "Ru", "S", "Sb", "Sc", "Se",
	"Sg", "Si", "Sm", "Sn", "Sr", "Ta", "Tb", "Tc", "Te", "Th",
	"Ti", "Tl", "Tm", "U", "Uub", "Uuh", "Uuo", "Uup", "Uuq", "Uus",
	"Uut", "V", "W", "Xe", "Y", "Yb", "Zn", "Zr",
}

func (e Element) String() string {
	return AtomSymbol[int(e)]
}

func (m *Matcher) acceptElement(p, t byte) {
	p0, p1, p2 := m.src[p], m.src[p+1], m.src[p+2]
	switch p0 {
	case 'a':
		switch p1 {
		case 'c':
			m.accepted(Ac, p+2, t)
		case 'g':
			m.accepted(Ag, p+2, t)
		case 'l':
			m.accepted(Al, p+2, t)
		case 'm':
			m.accepted(Am, p+2, t)
		case 'r':
			m.accepted(Ar, p+2, t)
		case 's':
			m.accepted(As, p+2, t)
		case 't':
			m.accepted(At, p+2, t)
		case 'u':
			m.accepted(Au, p+2, t)
		}
	case 'b':
		switch p1 {
		case 'a':
			m.accepted(Ba, p+2, t)
		case 'e':
			m.accepted(Be, p+2, t)
		case 'h':
			m.accepted(Bh, p+2, t)
		case 'i':
			m.accepted(Bi, p+2, t)
		case 'k':
			m.accepted(Bk, p+2, t)
		case 'r':
			m.accepted(Br, p+2, t)
		}
		m.accepted(B, p+1, t)
	case 'c':
		switch p1 {
		case 'a':
			m.accepted(Ca, p+2, t)
		case 'd':
			m.accepted(Cd, p+2, t)
		case 'e':
			m.accepted(Ce, p+2, t)
		case 'f':
			m.accepted(Cf, p+2, t)
		case 'l':
			m.accepted(Cl, p+2, t)
		case 'm':
			m.accepted(Cm, p+2, t)
		case 'o':
			m.accepted(Co, p+2, t)
		case 'r':
			m.accepted(Cr, p+2, t)
		case 's':
			m.accepted(Cs, p+2, t)
		case 'u':
			m.accepted(Cu, p+2, t)
		}
		m.accepted(C, p+1, t)
	case 'd':
		switch p1 {
		case 'b':
			m.accepted(Db, p+2, t)
		case 's':
			m.accepted(Ds, p+2, t)
		case 'y':
			m.accepted(Dy, p+2, t)
		}
	case 'e':
		switch p1 {
		case 'r':
			m.accepted(Er, p+2, t)
		case 's':
			m.accepted(Es, p+2, t)
		case 'u':
			m.accepted(Eu, p+2, t)
		}
	case 'f':
		switch p1 {
		case 'e':
			m.accepted(Fe, p+2, t)
		case 'm':
			m.accepted(Fm, p+2, t)
		case 'r':
			m.accepted(Fr, p+2, t)
		}
		m.accepted(F, p+1, t)
	case 'g':
		switch p1 {
		case 'a':
			m.accepted(Ga, p+2, t)
		case 'd':
			m.accepted(Gd, p+2, t)
		case 'e':
			m.accepted(Ge, p+2, t)
		}
	case 'h':
		switch p1 {
		case 'e':
			m.accepted(He, p+2, t)
		case 'f':
			m.accepted(Hf, p+2, t)
		case 'g':
			m.accepted(Hg, p+2, t)
		case 'o':
			m.accepted(Ho, p+2, t)
		case 's':
			m.accepted(Hs, p+2, t)
		}
		m.accepted(H, p+1, t)
	case 'i':
		switch p1 {
		case 'n':
			m.accepted(In, p+2, t)
		case 'r':
			m.accepted(Ir, p+2, t)
		}
		m.accepted(I, p+1, t)
	case 'k':
		if p1 == 'r' {
			m.accepted(Kr, p+2, t)
		}
		m.accepted(K, p+1, t)
	case 'l':
		switch p1 {
		case 'a':
			m.accepted(La, p+2, t)
		case 'i':
			m.accepted(Li, p+2, t)
		case 'r':
			m.accepted(Lr, p+2, t)
		case 'u':
			m.accepted(Lu, p+2, t)
		}
	case 'm':
		switch p1 {
		case 'd':
			m.accepted(Md, p+2, t)
		case 'g':
			m.accepted(Mg, p+2, t)
		case 'n':
			m.accepted(Mn, p+2, t)
		case 'o':
			m.accepted(Mo, p+2, t)
		case 't':
			m.accepted(Mt, p+2, t)
		}
	case 'n':
		switch p1 {
		case 'a':
			m.accepted(Na, p+2, t)
		case 'b':
			m.accepted(Nb, p+2, t)
		case 'd':
			m.accepted(Nd, p+2, t)
		case 'e':
			m.accepted(Ne, p+2, t)
		case 'i':
			m.accepted(Ni, p+2, t)
		case 'o':
			m.accepted(No, p+2, t)
		case 'p':
			m.accepted(Np, p+2, t)
		}
		m.accepted(N, p+1, t)
	case 'o':
		if p1 == 's' {
			m.accepted(Os, p+2, t)
		}
		m.accepted(O, p+1, t)
	case 'p':
		switch p1 {
		case 'a':
			m.accepted(Pa, p+2, t)
		case 'b':
			m.accepted(Pb, p+2, t)
		case 'd':
			m.accepted(Pd, p+2, t)
		case 'm':
			m.accepted(Pm, p+2, t)
		case 'o':
			m.accepted(Po, p+2, t)
		case 'r':
			m.accepted(Pr, p+2, t)
		case 't':
			m.accepted(Pt, p+2, t)
		case 'u':
			m.accepted(Pu, p+2, t)
		}
		m.accepted(P, p+1, t)
	case 'r':
		switch p1 {
		case 'a':
			m.accepted(Ra, p+2, t)
		case 'b':
			m.accepted(Rb, p+2, t)
		case 'e':
			m.accepted(Re, p+2, t)
		case 'f':
			m.accepted(Rf, p+2, t)
		case 'g':
			m.accepted(Rg, p+2, t)
		case 'h':
			m.accepted(Rh, p+2, t)
		case 'n':
			m.accepted(Rn, p+2, t)
		case 'u':
			m.accepted(Ru, p+2, t)
		}
	case 's':
		switch p1 {
		case 'b':
			m.accepted(Sb, p+2, t)
		case 'c':
			m.accepted(Sc, p+2, t)
		case 'e':
			m.accepted(Se, p+2, t)
		case 'g':
			m.accepted(Sg, p+2, t)
		case 'i':
			m.accepted(Si, p+2, t)
		case 'm':
			m.accepted(Sm, p+2, t)
		case 'n':
			m.accepted(Sn, p+2, t)
		case 'r':
			m.accepted(Sr, p+2, t)
		}
		m.accepted(S, p+1, t)
	case 't':
		switch p1 {
		case 'a':
			m.accepted(Ta, p+2, t)
		case 'b':
			m.accepted(Tb, p+2, t)
		case 'c':
			m.accepted(Tc, p+2, t)
		case 'e':
			m.accepted(Te, p+2, t)
		case 'h':
			m.accepted(Th, p+2, t)
		case 'i':
			m.accepted(Ti, p+2, t)
		case 'l':
			m.accepted(Tl, p+2, t)
		case 'm':
			m.accepted(Tm, p+2, t)
		}
	case 'u':
		if p1 == 'u' {
			switch p2 {
			case 'b':
				m.accepted(Uub, p+3, t)
			case 'h':
				m.accepted(Uuh, p+3, t)
			case 'o':
				m.accepted(Uuo, p+3, t)
			case 'p':
				m.accepted(Uup, p+3, t)
			case 'q':
				m.accepted(Uuq, p+3, t)
			case 's':
				m.accepted(Uus, p+3, t)
			case 't':
				m.accepted(Uut, p+3, t)
			}
		}
		m.accepted(U, p+1, t)
	case 'v':
		m.accepted(V, p+1, t)
	case 'w':
		m.accepted(W, p+1, t)
	case 'x':
		if p1 == 'e' {
			m.accepted(Xe, p+2, t)
		}
	case 'y':
		if p1 == 'b' {
			m.accepted(Yb, p+2, t)
		}
		m.accepted(Y, p+1, t)
	case 'z':
		switch p1 {
		case 'n':
			m.accepted(Zn, p+2, t)
		case 'r':
			m.accepted(Zr, p+2, t)
		}
	}
}
