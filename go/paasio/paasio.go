package paasio

import (
	"io"
	"sync"
)

func NewReadCounter(reader io.Reader) ReadCounter {
	return &ReadState{reader: reader}
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &WriteState{writer: writer}
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	return &ReadWriteState{ReadState{reader: readWriter}, WriteState{writer: readWriter}}
}

type ReadState struct {
	reader io.Reader
	bytesRead int64
	readOps int
	mux sync.Mutex
}

func (s *ReadState) Read(p []byte) (n int, err error) {
	bytes, err := s.reader.Read(p)
	s.mux.Lock()
	s.bytesRead += int64(bytes)
	s.readOps++
	s.mux.Unlock()
	return bytes, err
}

func (s *ReadState) ReadCount() (n int64, nops int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.bytesRead, s.readOps
}

type WriteState struct {
	writer io.Writer
	bytesWritten int64
	writeOps int
	mux sync.Mutex
}

func (s *WriteState) Write(p []byte) (n int, err error) {
	bytes, err := s.writer.Write(p)
	s.mux.Lock()
	s.bytesWritten += int64(bytes)
	s.writeOps++
	s.mux.Unlock()
	return bytes, err
}

func (s *WriteState) WriteCount() (n int64, nops int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.bytesWritten, s.writeOps
}

type ReadWriteState struct {
	ReadState
	WriteState
}
