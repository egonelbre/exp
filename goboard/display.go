package main

func (board *Board) String() string {
	var content [1 + 2*Size + 3 +
		Size + 2*Count + 3*Size +
		1 + 2*Size + 3]rune

	k := 0
	p := func(r rune) {
		content[k] = r
		k++
	}

	p('┏')
	for x := 0; x < Size; x++ {
		p('━')
		p('━')
	}
	p('━')
	p('┓')
	p('\n')

	for y := 0; y < Size; y++ {
		p('┃')
		for x := 0; x < Size; x++ {
			p(' ')
			p(board.At(x, y).Rune())
		}
		p(' ')
		p('┃')
		p('\n')
	}

	p('┗')
	for x := 0; x < Size; x++ {
		p('━')
		p('━')
	}
	p('━')
	p('┛')
	p('\n')

	return string(content[:])
}
