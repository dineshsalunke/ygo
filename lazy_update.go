package ygo

type lazyBlockReader struct {
	UpdateDecoder
	filterSkips bool
	current     Block
}

func (self *lazyBlockReader) Current() Block {
	return self.current
}

func (self *lazyBlockReader) Next() (bool, error) {
	panic("not implemented")
}

type lazyBlockWriter struct {
	UpdateEncoder
}

func newLazyBlockReader(decoder UpdateDecoder) *lazyBlockReader {
	return &lazyBlockReader{
		UpdateDecoder: decoder,
	}
}

func newLazyBlockWriter(encoder UpdateEncoder) *lazyBlockWriter {
	return &lazyBlockWriter{
		UpdateEncoder: encoder,
	}
}

func (writer *lazyBlockWriter) WriteBlock(block Block, offset uint64) error {
	panic("not implemented")
}

func (writer *lazyBlockWriter) Flush() error {
	panic("not implemented")
}
