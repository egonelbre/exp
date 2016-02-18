package arena

import (
	"syscall"
)

const (
	MEM_COMMIT  = 0x1000
	MEM_RESERVE = 0x2000
	MEM_RELEASE = 0x8000
	MEM_RESET   = 0x80000

	PAGE_NOACCESS          = 0x0001
	PAGE_EXECUTE_READWRITE = 0x40

	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

const ChunkSize = 8 * KB

var (
	dll          = syscall.MustLoadDLL("kernel32.dll")
	VirtualAlloc = dll.MustFindProc("VirtualAlloc")
	VirtualFree  = dll.MustFindProc("VirtualFree")
)

type chunk struct {
	offset uintptr
	data   uintptr
}

type Arena struct {
	chunks []chunk
}

func New(chunksize uint) *Arena {
	arena := &Arena{}
	arena.newblock()
	return arena
}

func alloc(size uint) (uintptr, error) {
	addr, _, err := VirtualAlloc.Call(0, uintptr(size), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		return 0, err
	}
	return addr, nil
}
func free(p uintptr) error {
	r, _, err := VirtualFree.Call(p, 0, MEM_RELEASE)
	if r == 0 {
		return err
	}
	return nil
}

func (a *Arena) last() *chunk {
	if len(a.chunks) == 0 {
		return a.newblock()
	}
	return &a.chunks[len(a.chunks)-1]
}

func (a *Arena) newblock() *chunk {
	data, err := alloc(ChunkSize)
	if err != nil {
		panic(err)
	}
	a.chunks = append(a.chunks, chunk{offset: 0, data: data})
	return &a.chunks[len(a.chunks)-1]
}

func (a *Arena) Alloc(size uintptr) uintptr {
	cursor := a.last()
	if cursor.offset+size > ChunkSize {
		cursor = a.newblock()
	}
	p := cursor.data + cursor.offset
	cursor.offset += size
	return uintptr(p)
}

func (a *Arena) Release() {
	for _, c := range a.chunks {
		err := free(c.data)
		if err != nil {
			panic(err)
		}
	}
}
