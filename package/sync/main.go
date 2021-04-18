package main

import (
	"bytes"
	"sync"
)

var buffers = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// GetBuffer _
func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

// PutBuffer _
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()

	buffers.Put(buf)
}

func main() {

}
