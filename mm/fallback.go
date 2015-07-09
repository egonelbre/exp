package mm

const FallbackDef = `
package xxx

import "github.com/egonelbre/exp/mm"

type Fallback struct {
	Primary  {{ .Primary.Name  }}
	Fallback {{ .Fallback.Name }}
}

func (m *Fallback) Alignment() int {
	return {{ min .Primary.Alignment .Fallback.Alignment }}
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

{{ if and .Primary.Owns (or .Primary.Dealloc .Fallback.Dealloc) }}
func (m *Fallback) Dealloc(p unsafe.Pointer) {
	if m.Primary.Owns(p) {
		{{ if .Primary.Dealloc }}
			return m.Primary.Dealloc(p)
		{{ else }}
			return false
		{{ end }}
	} else {
		{{ if .Fallback.Dealloc }}
			return m.Fallback.Dealloc(p)
		{{ else }}
			return false
		{{ end }}
	}
}
{{ end }}

{{ if and .Primary.Empty .Fallback.Empty }}
func (m *Fallback) Empty() bool {
	return m.Primary.Empty() && m.Fallback.Empty()
}
{{ end }}
`
