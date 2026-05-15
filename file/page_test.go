package file

import (
	"bytes"
	"testing"
)

func TestPageIntRountTrip(t *testing.T) {
	p := NewPage(100)
	p.SetInt(0, 42)
	if got := p.GetInt(0); got != 42 {
		t.Errorf("expected 42, got %d", got)
	}
}

func TestPageBytesRoundTrip(t *testing.T) {
	p := NewPage(100)
	data := []byte{1, 2, 3, 4, 5}
	p.SetBytes(0, data)
	got := p.GetBytes(0)
	if !bytes.Equal(got, data) {
		t.Errorf("expected %v, got %v", data, got)
	}
}
