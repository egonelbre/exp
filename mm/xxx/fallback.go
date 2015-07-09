// DO NOT EDIT
// GENERATED CODE
package xxx

import "github.com/egonelbre/exp/mm"

type Fallback struct {
	Primary  mm.Malloc
	Fallback *mm.Region
}

func (m *Fallback) Alignment() int {
	return 4
}

func (m *Fallback) Alloc(size int) unsafe.Pointer {
	if size == 0 {
		return nil
	}
	p := m.Primary.Allocate(size)
	if p == nil {
		p = m.Fallback.Allocate(size)
	}
	return p
}
