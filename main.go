package main

import "github.com/mikuh/gotext/trie"
import "fmt"
import "time"

func main() {
	// t := trie.New()

	// t.Add("清华大学", []string{"清华大学","北大","浙大","中科大"})
	// t.Add("春节", []string{"春节","中秋","国庆","清明"})
	// t.Add("清华", []string{"清华", "复旦","交大"})
	// t.Add("我", []string{"我", "你", "他"})


	// a := t.MultiReplace("我考上了清华,所以在清华大学打游戏春节不回家")
	// fmt.Println(len(a))
	// for _,x := range a{
	// 	fmt.Println(x)
	// }

	t := trie.New()

	t.Add("分裂祖国", "Politics")
	t.Add("shabi", "Abuse")
	t.Add("a片", "Porn")
	var a []trie.MatchWord
	var sentence string = `那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了
	那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了那些分裂祖国又看a片的shabi们真是醉了`

	start_at := time.Now().UnixNano()
	for i:=0; i< 10000; i++{
		a =t.ExtractKeyWords(sentence)
	}
	
	fmt.Println(a)
	fmt.Println("Cost", float32(time.Now().UnixNano() - start_at)/1e9)
	
}