package main

import "github.com/mikuh/gotext/trie"
import "fmt"

func main() {
	t := trie.New()

	t.Add("清华大学", []string{"清华大学","北大","浙大","中科大"})
	t.Add("春节", []string{"春节","中秋","国庆","清明"})

	a := t.MultiReplace("我考上了清华,所以在清华大学打游戏春节不回家")
	fmt.Println(a)
	
}