package ui

import (
	"bytes"
	"fmt"
)

type writer struct {
	buf    bytes.Buffer
	stack  []string
	inattr bool
	invoid bool
}

func NewWriter() Writer { return &writer{} }

func (w *writer) String() string {
	if len(w.stack) > 0 {
		panic("writing is not finished")
	}
	return w.buf.String()
}

func (w *writer) Bytes() []byte {
	if len(w.stack) > 0 {
		panic("writing is not finished")
	}
	return w.buf.Bytes()
}

func (w *writer) Open(tag string) {
	w.CloseAttributes()

	w.stack = append(w.stack, tag)
	w.invoid = voidElements[tag]
	w.inattr = true

	w.buf.WriteByte('<')
	w.buf.WriteString(tag)
}

func (w *writer) CloseAttributes() {
	if !w.inattr {
		return
	}

	w.inattr = false
	w.buf.WriteByte('>')
}

func (w *writer) Attr(name, value string) {
	if !w.inattr {
		panic("not in attributes section")
	}

	w.buf.WriteByte(' ')
	w.buf.WriteString(name)
	w.buf.WriteByte('=')
	w.buf.WriteByte('"')
	w.buf.WriteString(EscapeAttribute(value))
	w.buf.WriteByte('"')
}

func (w *writer) Text(text string) {
	w.CloseAttributes()
	w.buf.WriteString(EscapeCharData(text))
}

func (w *writer) UnsafeWrite(text string) {
	w.buf.WriteString(text)
}

func (w *writer) UnsafeContent(text string) {
	w.CloseAttributes()
	w.buf.WriteString(text)
}

func (w *writer) Close(tag string) {
	w.CloseAttributes()

	if len(w.stack) == 0 {
		panic("no unclosed tags")
	}

	var current string
	n := len(w.stack) - 1
	current, w.stack = w.stack[n], w.stack[:n]
	if current != tag {
		panic(fmt.Sprintf("closing tag %v expected %v", tag, current))
	}

	w.invoid = (len(w.stack) > 0) && voidElements[w.stack[len(w.stack)-1]]
	// void elements have only a single tag
	if voidElements[tag] {
		return
	}

	w.buf.WriteString("</")
	w.buf.WriteString(tag)
	w.buf.WriteByte('>')
}

// Section 12.1.2, "Elements", gives this list of void elements. Void elements
// are those that can't have any contents.
var voidElements = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}
