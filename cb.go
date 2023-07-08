package main

type CircularBuffer struct {
	data  []string
	index int
	size  int
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		data:  make([]string, size),
		index: 0,
		size:  size,
	}
}

func (c *CircularBuffer) Add(line string) {
	c.data[c.index] = line
	if c.index+2 > c.size {
		c.index = 0
	} else {
		c.index += 1
	}
}

func (c *CircularBuffer) Get() []string {
	return append(c.data[c.index:], c.data[0:c.index]...)
}
