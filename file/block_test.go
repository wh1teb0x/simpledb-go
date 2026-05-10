package file

import "testing"

func TestBlockId(t *testing.T) {
	block1 := NewBlockId("file1", 1)
	block2 := NewBlockId("file1", 1)

	if block1 != block2 {
		t.Errorf("Expected block1 and block2 to be equal")
	}
}

func TestBlockIdString(t *testing.T) {
	block := NewBlockId("file1", 1)
	expected := "[file file1, block 1]"
	if block.String() != expected {
		t.Errorf("wanted %s, got %s", expected, block.String())
	}
}

func TestBlockIdAsMapKey(t *testing.T) {
	block1 := NewBlockId("file1", 1)
	block2 := NewBlockId("file1", 1)
	blockMap := make(map[BlockId]string)
	blockMap[block1] = "Block 1"
	if blockMap[block2] != "Block 1" {
		t.Errorf("Expected block2 to retrieve the same value as block1 from the map")
	}
}
