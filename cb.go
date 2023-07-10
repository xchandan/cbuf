package main

type ContextBuffer struct {
	data  []string
	index int
	size  int
}

func NewContextBuffer(size int) *ContextBuffer {
	return &ContextBuffer{
		data:  make([]string, size),
		index: 0,
		size:  size,
	}
}

func (c *ContextBuffer) Add(line string) {
	c.data[c.index] = line
	if c.index+2 > c.size {
		c.index = 0
	} else {
		c.index += 1
	}
}

func (c *ContextBuffer) Get() []string {
	return append(c.data[c.index:], c.data[0:c.index]...)
}
