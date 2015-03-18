package physics

func (frame *Frame) Largest() Values     { return (*Largest)(frame) }
func (frame *Frame) Interacting() Values { return (*Interacting)(frame) }
func (frame *Frame) ABC() Values         { return (*ABC)(frame) }
func (frame *Frame) XYZ() Values         { return (*XYZ)(frame) }

type Largest Frame

func (frame *Largest) Index(i int) int    { return i }
func (frame *Largest) Len() int           { return len(frame.Cubes) }
func (frame *Largest) Get(i int) int32    { return frame.Cubes[i].Largest }
func (frame *Largest) Set(i int, v int32) { frame.Cubes[i].Largest = v }

type Interacting Frame

func (frame *Interacting) Index(i int) int    { return i }
func (frame *Interacting) Len() int           { return len(frame.Cubes) }
func (frame *Interacting) Get(i int) int32    { return frame.Cubes[i].Interacting }
func (frame *Interacting) Set(i int, v int32) { frame.Cubes[i].Interacting = v }

type ABC Frame

func (frame *ABC) Index(i int) int { return i / 3 }
func (frame *ABC) Len() int        { return len(frame.Cubes) * 3 }

func (frame *ABC) Get(i int) int32 {
	switch i % 3 {
	case 0:
		return frame.Cubes[i/3].A
	case 1:
		return frame.Cubes[i/3].B
	case 2:
		return frame.Cubes[i/3].C
	}
	return 0
}

func (frame *ABC) Set(i int, v int32) {
	switch i % 3 {
	case 0:
		frame.Cubes[i/3].A = v
	case 1:
		frame.Cubes[i/3].B = v
	case 2:
		frame.Cubes[i/3].C = v
	}
}

type XYZ Frame

func (frame *XYZ) Index(i int) int { return i / 3 }
func (frame *XYZ) Len() int        { return len(frame.Cubes) * 3 }

func (frame *XYZ) Get(i int) int32 {
	switch i % 3 {
	case 0:
		return frame.Cubes[i/3].X
	case 1:
		return frame.Cubes[i/3].Y
	case 2:
		return frame.Cubes[i/3].Z
	}
	return 0
}

func (frame *XYZ) Set(i int, v int32) {
	switch i % 3 {
	case 0:
		frame.Cubes[i/3].X = v
	case 1:
		frame.Cubes[i/3].Y = v
	case 2:
		frame.Cubes[i/3].Z = v
	}
}
