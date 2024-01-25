package main

import (
	"fmt"
	cmd "github.com/w3gop2p/go_design_patterns/CreationalPatterns/prototype/cmd"
)

func main() {
	file1 := &cmd.File{Name: "File1"}
	file2 := &cmd.File{Name: "File2"}
	file3 := &cmd.File{Name: "File3"}

	folder1 := &cmd.Folder{
		Children: []cmd.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &cmd.Folder{
		Children: []cmd.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
