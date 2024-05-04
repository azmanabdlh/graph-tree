package main

import "fmt"

type Node struct {
	Base  string
	Key   int
	Value string // as path
	Left  *Node
	Right *Node
}

func isExist(node *Node, path string) bool {
	ok, tmp := search(node, path)
	return ok && tmp != ""
}

func search(node *Node, path string) (bool, string) {
	
	//if node == nil {
	//	return false
	//}
	//
	//if node.Key == key {
	//	return true
	//}
	//
	//if key < node.Key {
	//	return search(node.Left, key)
	//}
	//
	//if key > node.Key {
	//	return search(node.Right, key)
	//}
	
	// 1. Binary tree DFS
	
	if node == nil {
		return false, ""
	}
	
	if node.Value == path {
		return true, node.Value
	}
	
	if node.Left != nil {
		ok, temp := search(node.Left, path)
		if ok {
			return ok, node.Value + "/" + temp
		}
	}
	
	if node.Right != nil {
		ok, temp := search(node.Right, path)
		if ok {
			return ok, node.Value + "/" + temp
		}
	}
	
	return false, ""
}

func main() {
	
	node := &Node{
		Key:   4,
		Value: "src",
		Left: &Node{
			Key:   2,
			Value: "docs",
			Left: &Node{
				Key:   1,
				Value: "README.md",
			},
			Right: &Node{
				Key:   3,
				Value: "CONTRIBUTE.md",
			},
		},
		Right: &Node{
			Key:   6,
			Value: "config",
			Left: &Node{
				Key:   5,
				Value: "app.yaml",
			},
			Right: &Node{
				Key:   7,
				Value: "auth.yaml",
			},
		},
	}
	
	ok, path := search(node, "autha.yaml")
	fmt.Println("ok ==> ", ok)
	fmt.Println("path ==> ", path)
	
	exist := isExist(node, "auth.yaml")
	fmt.Println("exist ==> ", exist)
}
