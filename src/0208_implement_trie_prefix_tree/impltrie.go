/*
208. Implement Trie (Prefix Tree)
https://leetcode.com/problems/implement-trie-prefix-tree/

Implement a trie with insert, search, and startsWith methods.

Note:
You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.
*/
// time: 2019-01-31

package impltree

type node struct {
	isWord bool
	next   map[byte]*node
}

// Trie prefix tree
type Trie struct {
	root *node
	size int
}

// Constructor initialize data structure here.
func Constructor() Trie {
	return Trie{root: &node{next: make(map[byte]*node)}}
}

// Insert inserts a word into the trie.
func (trie *Trie) Insert(word string) {
	cur := trie.root

	for i := 0; i < len(word); i++ {
		if _, ok := cur.next[word[i]]; !ok {
			cur.next[word[i]] = &node{next: make(map[byte]*node)}
		}
		cur = cur.next[word[i]]
	}
	if !cur.isWord {
		cur.isWord = true
		trie.size++
	}
}

// Search returns if the word is the trie.
func (trie *Trie) Search(word string) bool {
	cur := trie.root

	for i := 0; i < len(word); i++ {
		if _, ok := cur.next[word[i]]; !ok {
			return false
		}
		cur = cur.next[word[i]]
	}
	return cur.isWord
}

// StartsWith returns if there is any word in the trie that starts with the given prefix.
func (trie *Trie) StartsWith(prefix string) bool {
	cur := trie.root

	for i := 0; i < len(prefix); i++ {
		if _, ok := cur.next[prefix[i]]; !ok {
			return false
		}
		cur = cur.next[prefix[i]]
	}
	return true
}
