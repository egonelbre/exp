package aq

import (
	"errors"
	"os"
	"sync"

	"github.com/edsrzf/mmap-go"
)

type DB struct {
	filename string
	size     int32
	file     *os.File

	data mmap.MMap
	mu   sync.Mutex
	head int32
}

func New(filename string, size int) (*DB, error) {
	db := &DB{filename: filename, size: int32(size)}
	return db, db.open()
}

func (db *DB) open() error {
	var err error
	db.file, err = os.OpenFile(db.filename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	if err := db.file.Truncate(int64(db.size)); err != nil {
		return err
	}

	db.data, err = mmap.Map(db.file, os.O_RDWR, 0)
	return err
}

func (db *DB) Flush() error { return db.data.Flush() }

func (db *DB) Close() error {
	err := db.data.Unmap()
	db.data = nil

	err2 := db.file.Close()
	db.file = nil

	if err == nil {
		return err2
	}
	return err
}

var (
	InvalidHeader   = Header{Off: -1, Len: -1}
	ErrFull         = errors.New("database is full")
	ErrInvalidValue = errors.New("empty value is not allowed")
)

type Header struct {
	Off int32
	Len int32
}

func (db *DB) Write(data []byte) (Header, error) {
	if len(data) == 0 {
		return InvalidHeader, ErrInvalidValue
	}
	db.mu.Lock()
	if db.head+4+int32(len(data)) > db.size {
		db.mu.Unlock()
		return InvalidHeader, ErrFull
	}
	hdr := Header{
		Off: int32(db.head + 4),
		Len: int32(len(data)),
	}

	db.data[db.head] = byte(hdr.Len >> 24)
	db.data[db.head+1] = byte(hdr.Len >> 16)
	db.data[db.head+2] = byte(hdr.Len >> 8)
	db.data[db.head+3] = byte(hdr.Len >> 0)

	db.head += 4 + hdr.Len
	db.mu.Unlock()

	copy(db.data[hdr.Off:hdr.Off+hdr.Len], data)
	return hdr, nil
}

func (db *DB) Read(hdr Header) ([]byte, error) { return db.data[hdr.Off : hdr.Off+hdr.Len], nil }

type Iterator struct {
	head int32
	clen int32
	data []byte
}

func (db *DB) Iterate() Iterator { return Iterator{0, 0, db.data[:db.head]} }

func (it *Iterator) Next() bool {
	// jump to next header
	it.head += it.clen
	it.clen = 0

	if it.head+4 > int32(len(it.data)) {
		return false
	}

	it.clen = int32(it.data[it.head+0])<<24 |
		int32(it.data[it.head+1])<<16 |
		int32(it.data[it.head+2])<<8 |
		int32(it.data[it.head+3])<<0
	it.head += 4
	return it.head+it.clen <= int32(len(it.data))
}

func (it *Iterator) Bytes() []byte { return it.data[it.head : it.head+it.clen] }
