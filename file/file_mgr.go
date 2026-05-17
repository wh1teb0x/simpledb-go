package file

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type FileMgr struct {
	mu          sync.Mutex
	dbDirectory string
	blockSize   int
	isNew       bool
	openFiles   map[string]*os.File
}

func NewFileMgr(dbDirectory string, blockSize int) (*FileMgr, error) {
	fm := &FileMgr{
		dbDirectory: dbDirectory,
		blockSize:   blockSize,
		openFiles:   make(map[string]*os.File),
	}

	// create the directory if the database is new
	if _, err := os.Stat(dbDirectory); errors.Is(err, fs.ErrNotExist) {
		if err := os.MkdirAll(dbDirectory, 0755); err != nil {
			return nil, err
		}
		fm.isNew = true
	}

	// remove any leftover temporary tables
	entries, err := os.ReadDir(dbDirectory)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "temp") {
			if err := os.Remove(filepath.Join(dbDirectory, entry.Name())); err != nil {
				return nil, err
			}
		}
	}

	return fm, nil
}

func (fm *FileMgr) Read(blk BlockId, p *Page) error {
	return nil
}

func (fm *FileMgr) Write(blk BlockId, p *Page) error {
	return nil
}

func (fm *FileMgr) Append(filename string) (BlockId, error) {
	return BlockId{}, nil
}

func (fm *FileMgr) Length(filename string) (int, error) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	f, err := fm.getFile(filename)
	if err != nil {
		return 0, err
	}

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	length := int(info.Size() / int64(fm.blockSize))

	return length, nil
}

func (fm *FileMgr) IsNew() bool {
	return fm.isNew
}

func (fm *FileMgr) BlockSize() int {
	return fm.blockSize
}

func (fm *FileMgr) getFile(filename string) (*os.File, error) {
	f := fm.openFiles[filename]
	if f != nil {
		return f, nil
	}
	pathName := filepath.Join(fm.dbDirectory, filename)
	f, err := os.OpenFile(pathName, os.O_RDWR|os.O_CREATE|os.O_SYNC, 0644)
	if err != nil {
		return nil, err
	}
	fm.openFiles[filename] = f
	return f, nil
}
