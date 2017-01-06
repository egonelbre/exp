package audio

type Node interface {
	Node32
	Node64
}

type Node32 interface {
	Process32(*Buffer32) (int, error)
}

type Node64 interface {
	Process64(*Buffer64) (int, error)
}
