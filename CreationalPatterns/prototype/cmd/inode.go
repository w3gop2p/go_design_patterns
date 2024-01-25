package cmd

type Inode interface {
	Print(string)
	Clone() Inode
}
