package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

type TailWriter struct {
	Head         int // writing head
	TotalWritten int
	Buffer       []byte
}

func NewTail(size int) *TailWriter {
	return &TailWriter{0, 0, make([]byte, size, size)}
}

func (w *TailWriter) Write(data []byte) (int, error) {
	w.TotalWritten += len(data)

	if len(data) > len(w.Buffer) {
		copy(w.Buffer, data[len(data)-len(w.Buffer):])
		w.Head = 0
	} else {
		n := copy(w.Buffer[w.Head:], data)
		if w.Head >= len(w.Buffer) {
			w.Head = 0
		}

		n = copy(w.Buffer[:], data[n:])
		w.Head += n
	}

	return len(data), nil
}

func (tw *TailWriter) WriteTo(w io.Writer) {
	if tw.TotalWritten <= len(tw.Buffer) {
		w.Write(tw.Buffer[0:tw.TotalWritten])
		return
	}
	w.Write(tw.Buffer[tw.Head:])
	w.Write(tw.Buffer[:tw.Head])
}

func monitor() {
	if os.Getenv("__NOMONITOR__") != "" {
		return
	}

	cmd := exec.Command(os.Args[0], os.Args...)
	cmd.Env = []string{"__NOMONITOR__=true"}
	cmd.Env = append(cmd.Env, os.Environ()...)

	tailout := NewTail(2 << 10)
	tailerr := NewTail(2 << 10)

	cmd.Stdin = os.Stdin
	cmd.Stdout = io.MultiWriter(os.Stdout, tailout)
	cmd.Stderr = io.MultiWriter(os.Stderr, tailerr)

	// TODO: also propagate signals

	err := cmd.Run()
	if err != nil {
		// send email here
		var data bytes.Buffer
		data.WriteString("=== STDOUT ===\n")
		tailout.WriteTo(&data)
		data.WriteString("\n\n=== STDERR ===\n")
		tailerr.WriteTo(&data)
		data.WriteString("\n\n")

		ioutil.WriteFile("error.txt", data.Bytes(), 0755)
	}

	// propagate exit-code
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
	}

	os.Exit(0)
}

func main() {
	monitor()

	fmt.Println("hello")
	panic("we broke stuff")
}
