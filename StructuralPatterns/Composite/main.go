package main

import (
	"github.com/w3gop2p/go_design_patterns/StructuralPatterns/Composite/cmd"
)

func main() {
	file1 := &cmd.File{Name: "File1"}
	file2 := &cmd.File{Name: "File2"}
	file3 := &cmd.File{Name: "File3"}

	folder1 := &cmd.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &cmd.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
