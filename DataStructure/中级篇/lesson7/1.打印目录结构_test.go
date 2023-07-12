package lesson7

import "testing"

func TestDirStructure(t *testing.T) {
	root := NewPrefixTree()

	root.Insert("b\\\\cst")
	root.Insert("d\\\\")
	root.Insert("a\\\\d\\\\e")
	root.Insert("a\\\\b\\\\c")

	root.DFSPrintFilePath()
}
