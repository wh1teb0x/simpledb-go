package file

import "fmt"

type BlockId struct {
	fileName string
	blkNum   int
}

func NewBlockId(fileName string, blkNum int) BlockId {
	return BlockId{fileName, blkNum}
}

func (b BlockId) FileName() string {
	return b.fileName
}

func (b BlockId) Number() int {
	return b.blkNum
}

func (b BlockId) String() string {
	return fmt.Sprintf("[file %s, block %d]", b.fileName, b.blkNum)
}
