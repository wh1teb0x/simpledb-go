package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewFileCreatesDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "newDB")
	fm, err := NewFileMgr(dir, 400)
	if err != nil {
		t.Fatalf("NreFileMgr: %v", err)
	}
	if !fm.IsNew() {
		t.Errorf("IsNew() = false, want true")
	}
	if _, err := os.Stat(dir); err != nil {
		t.Errorf("directory not created: %v", err)
	}
}

func TestNewFileMgrExistingDir(t *testing.T) {
	dir := t.TempDir()
	fm, err := NewFileMgr(dir, 400)
	if err != nil {
		t.Fatalf("NewFileMgr: %v", err)
	}
	if fm.IsNew() {
		t.Errorf("IsNew() = true, want false")
	}
}

func TestNewFileMgrRemovesTempFiles(t *testing.T) {
	dir := t.TempDir()
	tempPath := filepath.Join(dir, "tempfoo")
	if err := os.WriteFile(tempPath, []byte("garbage"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := NewFileMgr(dir, 400); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(tempPath); !os.IsNotExist(err) {
		t.Errorf("temp file should be removed")
	}
}

func TestFileMgrReadWriteRoundTrip(t *testing.T) {
	fm, err := NewFileMgr(t.TempDir(), 400)
	if err != nil {
		t.Fatal(err)
	}

	blk, err := fm.Append("testfile")
	if err != nil {
		t.Fatal(err)
	}

	p1 := NewPage(400)
	p1.SetInt(0, 12345)
	p1.SetString(100, "hello")
	if err := fm.Write(blk, p1); err != nil {
		t.Fatal(err)
	}

	p2 := NewPage(400)
	if err := fm.Read(blk, p2); err != nil {
		t.Fatal(err)
	}

	if got := p2.GetInt(0); got != 12345 {
		t.Errorf("GetInt(0) = %d, want 12345", got)
	}
	if got := p2.GetString(100); got != "hello" {
		t.Errorf("GetString(100) = %q, want %q", got, "hello")
	}
}

func TestFileMgrLengthGrowsWithAppend(t *testing.T) {
	fm, err := NewFileMgr(t.TempDir(), 400)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		if _, err := fm.Append("growfile"); err != nil {
			t.Fatal(err)
		}
	}

	n, err := fm.Length("growfile")
	if err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Errorf("Length = %d, want 3", n)
	}
}
