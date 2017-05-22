package main

func main() {
	//n1 := Node{
	//	id:1,
	//	data:nil,
	//}
	//n2:=Node{
	//	id:2,
	//	data:nil,
	//	next:&n1,
	//}
	//println(n2)

	f := File{
		name: "text.txt",
		size: 213,
	}
	f.attr.owner = 1
	f.attr.perm = 075
}

type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

type File struct {
	name string
	size int
	attr struct {
		perm  int
		owner int
	}
}
