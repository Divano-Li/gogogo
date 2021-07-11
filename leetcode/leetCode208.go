package main

import "strings"

type Trie struct {
	isEnd bool
	children [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	a := []byte(word)
	for i:=0 ;i < len(a); i++ {
		index := a[i] - 97
		if this.children[index] == nil {
			this.children[index] = &Trie{}
		}
		this = this.children[index]
	}
	this.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	a := []byte(word)
	for i:=0; i <len(a); i++ {
		if this.children[a[i]-97] != nil {
			this = this.children[a[i] - 97]
		} else {
			return false
		}
	}
	if this.isEnd {
		return true
	} else {
		return false
	}
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	a := []byte(prefix)
	for i:=0; i <len(a); i++ {
		if this.children[a[i]-97] != nil {
			this = this.children[a[i] - 97]
		} else {
			return false
		}
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
