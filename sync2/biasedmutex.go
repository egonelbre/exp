package sync2

import (
	"sync"
)

// BiasedMutex is a RWMutex that allows tuning progress making on both
// writing and reading side. Tries to let through multiple readers and then
// multiple writers.
//
// BiasedMutex controls the flow by keeping track on how many reads / writes
// there have been. Once there have been RLock()-s called over `ReaderThreshold`,
// then RLock will wait in cased there are any writes pending.
// BiasedMutex will continue letting through at most `WriterThreshold` writers.
// Once writers have been exhausted the variables are reset.
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
	for m.writing+m.writersWaiting > 0 {
		if m.writing == 0 && m.readCount < m.readerThreshold {
			break
		}
		m.readers.Wait()
	}
	m.readCount++
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
