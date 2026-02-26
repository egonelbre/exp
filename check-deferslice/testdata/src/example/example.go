package example

import "io"

func use(s []int)    {}
func cleanup(s []int) {}

type myReader struct{}

func (myReader) Read(p []byte) (int, error)            { return 0, nil }
func (myReader) ReadAt(p []byte, off int64) (int, error) { return 0, nil }

// Modifications after defer should be flagged.

func indexAssignAfterDefer() {
	s := []int{1, 2, 3}
	defer use(s)
	s[0] = 99 // want `slice s is modified after being passed to defer`
}

func appendAfterDefer() {
	s := []int{1, 2, 3}
	defer use(s)
	s = append(s, 4) // want `slice s is modified after being passed to defer`
}

func reassignAfterDefer() {
	s := []int{1, 2, 3}
	defer use(s)
	s = make([]int, 10) // want `slice s is modified after being passed to defer`
}

func modifyInNestedBlock() {
	s := []int{1, 2, 3}
	defer use(s)
	if true {
		s[0] = 5 // want `slice s is modified after being passed to defer`
	}
}

func multipleDefers() {
	a := []int{1}
	b := []int{2}
	defer use(a)
	defer use(b)
	a[0] = 10 // want `slice a is modified after being passed to defer`
	b[0] = 20 // want `slice b is modified after being passed to defer`
}

// Modifications before defer are OK.

func modifyBeforeDefer() {
	s := []int{1, 2, 3}
	s[0] = 99
	defer use(s)
}

func appendBeforeDefer() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	defer use(s)
}

// No modification at all is OK.

func noModification() {
	s := []int{1, 2, 3}
	defer use(s)
}

// Non-slice arguments should not be flagged.

func deferWithNonSlice() {
	x := 42
	defer func(v int) {}(x)
	x = 100
}

// Modification inside a closure is not flagged (different scope).

func modifyInClosure() {
	s := []int{1, 2, 3}
	defer use(s)
	go func() {
		s[0] = 5 // want `slice s is modified after being passed to defer`
	}()
}

// Only the deferred slice is flagged, not others.

func differentSlice() {
	a := []int{1}
	b := []int{2}
	defer use(a)
	b[0] = 10 // ok: b is not deferred
}

// Passing deferred slice to known mutating functions.

func readAfterDefer() {
	buf := make([]byte, 1024)
	defer use(nil) // unrelated
	var r myReader
	_ = r
	buf2 := make([]byte, 512)
	defer func(b []byte) {}(buf2)
	r.Read(buf2) // want `slice buf2 is passed to Read after being passed to defer`
	_ = buf
}

func readAtAfterDefer() {
	buf := make([]byte, 1024)
	defer func(b []byte) {}(buf)
	var r myReader
	r.ReadAt(buf, 0) // want `slice buf is passed to ReadAt after being passed to defer`
}

func ioReadFullAfterDefer() {
	buf := make([]byte, 1024)
	defer func(b []byte) {}(buf)
	var r myReader
	io.ReadFull(r, buf) // want `slice buf is passed to ReadFull after being passed to defer`
}

func copyAfterDefer() {
	src := []int{1, 2, 3}
	dst := make([]int, 3)
	defer use(dst)
	copy(dst, src) // want `slice dst is passed to copy after being passed to defer`
}

func readBeforeDefer() {
	buf := make([]byte, 1024)
	var r myReader
	r.Read(buf)
	defer func(b []byte) {}(buf) // ok: Read is before defer
}

func readNonDeferredSlice() {
	buf := make([]byte, 1024)
	other := make([]byte, 512)
	defer func(b []byte) {}(buf)
	var r myReader
	r.Read(other) // ok: other is not deferred
}

// Save/restore pattern should not be flagged.

var global []string

func restorePattern() {
	defer func(old []string) { global = old }(global)
	global = []string{"a", "b"} // ok: defer restores the original value
}

func restorePatternAppend() {
	defer func(old []string) { global = old }(global)
	global = append(global, "c") // ok: defer restores the original value
}

// But a func literal that does NOT restore should still flag.

func nonRestoreFuncLit() {
	s := []int{1, 2, 3}
	defer func(b []int) {}(s)
	s[0] = 99 // want `slice s is modified after being passed to defer`
}

// Passing &slice to a function implies mutation.

func pointerArg(p *[]int) {}

func slicePointerAfterDefer() {
	s := []int{1, 2, 3}
	defer use(s)
	pointerArg(&s) // want `pointer to slice s is passed to pointerArg after being passed to defer`
}

func slicePointerAfterDeferInClosure() {
	s := []int{1, 2, 3}
	defer use(s)

	func(){
		pointerArg(&s) // want `pointer to slice s is passed to pointerArg after being passed to defer`
	}()
}

func slicePointerAfterDeferInClosure2() {
	s := []int{1, 2, 3}
	defer use(s)

	defer func(){
		func(){
			pointerArg(&s) // want `pointer to slice s is passed to pointerArg after being passed to defer`
		}()
	}()
}

func slicePointerBeforeDefer() {
	s := []int{1, 2, 3}
	pointerArg(&s)
	defer use(s) // ok: pointer use is before defer
}

func slicePointerDifferentVar() {
	s := []int{1, 2, 3}
	other := []int{4, 5}
	defer use(s)
	pointerArg(&other) // ok: other is not deferred
}
