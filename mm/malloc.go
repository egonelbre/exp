package mm

import "unsafe"

// #include <stdlib.h>
import "C"

const PlatformAlignment = 4 // todo: add build tags

type Malloc struct{}

func (m *Malloc) Alignment() int { return PlatformAlignment }

func (m *Malloc) Alloc(size int) unsafe.Pointer {
	if size == 0 {
		return nil
	}
	return unsafe.Pointer(C.malloc(C.size_t(size)))
}

func (m *Malloc) Dealloc(p unsafe.Pointer) bool {
	C.free(p)
	return true
}

func (m *Malloc) Realloc(p *unsafe.Pointer, size int) bool {
	if size == 0 {
		m.Dealloc(*p)
		*p = nil
		return true
	}

	*p = unsafe.Pointer(C.realloc(*p, C.size_t(size)))
	return *p != nil
}

type Aligned struct{}

func (m *Aligned) Alignment() int { return PlatformAlignment }

func (m *Aligned) Alloc(size int) unsafe.Pointer {
	if size == 0 {
		return nil
	}

	return unsafe.Pointer(C._aligned_malloc(C.size_t(size), C.size_t(PlatformAlignment)))
}

func (m *Aligned) Dealloc(p unsafe.Pointer) bool {
	C._aligned_free(p)
	return true
}

func (m *Aligned) Realloc(p *unsafe.Pointer, size int) bool {
	if size == 0 {
		m.Dealloc(*p)
		*p = nil
		return true
	}

	*p = unsafe.Pointer(C._aligned_realloc(*p, C.size_t(size), C.size_t(PlatformAlignment)))
	return *p != nil
}
