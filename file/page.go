package file

import "encoding/binary"

const (
	// BytesPerInt is the number of bytes used to store an int32 value.
	BytesPerInt = 4
	// maxBytesPerChar is the max bytes per UTF-8 char (internal use only)
	maxBytesPerChar = 4
)

type Page struct {
	b []byte
}

func NewPage(blockSize int) *Page {
	return &Page{b: make([]byte, blockSize)}
}

func NewPageFromBytes(b []byte) *Page {
	return &Page{b: b}
}
func (p *Page) GetInt(offset int) int32 {
	return int32(binary.BigEndian.Uint32(p.b[offset:]))
}

func (p *Page) SetInt(offset int, n int32) {
	binary.BigEndian.PutUint32(p.b[offset:], uint32(n))
}

func (p *Page) GetBytes(offset int) []byte {
	length := p.GetInt(offset)
	return p.b[offset+BytesPerInt : offset+BytesPerInt+int(length)]
}

func (p *Page) SetBytes(offset int, b []byte) {
	p.SetInt(offset, int32(len(b)))
	copy(p.b[offset+BytesPerInt:], b)
}

func (p *Page) GetString(offset int) string {
	return string(p.GetBytes(offset))
}

func (p *Page) SetString(offset int, s string) {
	p.SetBytes(offset, []byte(s))
}

func MaxLength(strlen int) int {
	return BytesPerInt + strlen*maxBytesPerChar
}

func (p *Page) contents() []byte {
	return p.b
}
