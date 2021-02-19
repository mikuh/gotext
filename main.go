package main

import "github.com/mikuh/gotext/trie"

func main() {
	t := trie.New()

	t.Add("foobar", 1)

	node, ok := t.Find("foobar")
	meta := node.Meta()
}