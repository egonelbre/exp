package sync2

import (
	"sync"
)

type BiasedMutex struct {
	readerThreshold int32
	writerThreshold int32

	mu      sync.Mutex
	readers sync.Cond
	writers sync.Cond

	reading        int32
	writing        int32
	writersWaiting int32

	readCount  int32
	writeCount int32
}

func NewBiasedMutex() *BiasedMutex {
	m := &BiasedMutex{}
	m.readers.L = &m.mu
	m.writers.L = &m.mu

	m.readerThreshold = 4
	m.writerThreshold = 4

	return m
}

func (m *BiasedMutex) SetReaderThreshold(x int) { m.readerThreshold = int32(x) }
func (m *BiasedMutex) SetWriterThreshold(x int) { m.writerThreshold = int32(x) }

func (m *BiasedMutex) RLock() {
	m.mu.Lock()
	m.readCount++
	for m.writing+m.writersWaiting > 0 {
		if m.writing == 0 && m.readCount < m.readerThreshold {
			break
		}
		m.readers.Wait()
	}
	m.reading++
	m.mu.Unlock()
}

func (m *BiasedMutex) RUnlock() {
	m.mu.Lock()
	m.reading--
	if m.reading == 0 && m.writersWaiting > 0 {
		m.writers.Signal()
	}
	m.mu.Unlock()
}

func (m *BiasedMutex) Lock() {
	m.mu.Lock()
	m.writersWaiting++
	// wait for other readers or writers to finish
	for m.reading > 0 || m.writing > 0 {
		m.writers.Wait()
	}
	// mark us as writing
	m.writersWaiting--
	m.writing++
	m.mu.Unlock()
}

func (m *BiasedMutex) Unlock() {
	m.mu.Lock()
	m.writing--
	m.writeCount++
	if m.writeCount >= m.writerThreshold {
		m.readers.Broadcast()
		m.readCount = 0
		m.writeCount = 0
	}
	if m.writersWaiting > 0 {
		m.writers.Signal()
	} else {
		m.readCount = 0
		m.readers.Broadcast()
	}
	m.mu.Unlock()
}
