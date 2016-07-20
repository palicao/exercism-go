// Package circular implements a circular buffer
package circular

import "errors"

const testVersion = 3

// Buffer is the data structure for our circular buffer
type Buffer struct {
	data  []byte
	read  int
	write int
	size  int
}

// NewBuffer creates a new empty buffer of a given size
func NewBuffer(size int) *Buffer {
	return &Buffer{make([]byte, size), 0, 0, 0}
}

// ReadByte reads a byte in the circular buffer, if the slot is not empty
func (b *Buffer) ReadByte() (byte, error) {
	if b.size <= 0 {
		return 0, errors.New("Buffer empty")
	}
	data := b.data[b.read]
	l := len(b.data)
	b.read = (b.read + 1) % l
	b.size--
	return data, nil
}

// WriteByte writes a byte in the buffer, if the slot is empty
func (b *Buffer) WriteByte(c byte) error {
	l := len(b.data)
	if b.size >= l {
		return errors.New("Buffer full")
	}
	b.data[b.write] = c
	b.write = (b.write + 1) % l
	b.size++
	return nil
}

// Overwrite writes a byte in the buffer, eventually overwriting the existing content
func (b *Buffer) Overwrite(c byte) {
	b.data[b.write] = c
	l := len(b.data)
	b.write = (b.write + 1) % l
	if b.size >= l {
		b.read = (b.read + 1) % l
	} else {
		b.size++
	}
}

// Reset resets the buffer
func (b *Buffer) Reset() {
	b.read = 0
	b.write = 0
	b.size = 0
}
