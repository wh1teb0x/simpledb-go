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

func TestPageStringRoundTrip(t *testing.T) {
	p := NewPage(100)
	p.SetString(50, "hello")
	if got := p.GetString(50); got != "hello" {
		t.Errorf("expected %s, got %s", "hello", got)
	}
}

func TestNewPageFromBytes(t *testing.T) {
	buf := make([]byte, 100)
	p := NewPage(100)
	p.SetInt(0, 42)
	copy(buf, p.contents()) // Copy the contents of p into buf

	p2 := NewPageFromBytes(buf)
	if got := p2.GetInt(0); got != 42 {
		t.Errorf("expected 42, got %d", got)
	}
}

func TestMaxLength(t *testing.T) {
	if got := MaxLength(10); got != 44 {
		t.Errorf("expected 44, got %d", got)
	}
}
