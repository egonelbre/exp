package main

import (
	"os"
)

// typedef unsigned long long BlobHandle;
// typedef unsigned long long FileHandle;
import "C"

type Blob struct {
	file *os.File
}

type File struct {
	file *os.File
}

var counter uint64
var files = map[C.FileHandle]*File{}
var blobs = map[C.BlobHandle]*Blob{}

//export OpenFile
func OpenFile(name *C.char) C.FileHandle {
	underlying, err := os.Open(C.GoString(name))
	if err != nil {
		panic(err)
	}

	file := &File{file: underlying}
	counter++

	handle := C.FileHandle(counter)
	files[handle] = file
	return handle
}

//export CloseFile
func CloseFile(handle C.FileHandle) {
	file, ok := files[handle]
	if !ok {
		panic("invalid handle")
	}

	err := file.file.Close()
	if err != nil {
		panic(err)
	}
	file.file = nil
}

//export OpenBlob
func OpenBlob(name *C.char) C.BlobHandle {
	underlying, err := os.Open(C.GoString(name))
	if err != nil {
		panic(err)
	}

	blob := &Blob{file: underlying}
	counter++

	handle := C.BlobHandle(counter)
	blobs[handle] = blob
	return handle
}

//export CloseBlob
func CloseBlob(handle C.BlobHandle) {
	blob, ok := blobs[handle]
	if !ok {
		panic("invalid handle")
	}

	err := blob.file.Close()
	if err != nil {
		panic(err)
	}
	blob.file = nil
}

func main() {}
