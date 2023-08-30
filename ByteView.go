package qcache

type ByteView struct {
	bytes []byte
}

// Len return length
func (view ByteView) Len() int {
	return len(view.bytes)
}

func (view ByteView) String() string {
	return string(view.bytes)
}

func (view ByteView) ByteSlice() []byte {
	return cloneBytes(view.bytes)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
