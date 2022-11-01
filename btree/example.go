package main

import "os"

func main() {
	tree := &BinaryTree{}
	tree.Insert(100).
		Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)

	Print(os.Stdout, tree.root, 0, 'M')
}
