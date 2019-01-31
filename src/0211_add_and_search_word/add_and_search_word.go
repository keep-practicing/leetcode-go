/*
211. Add and Search Word - Data structure design
https://leetcode.com/problems/add-and-search-word-data-structure-design/
*/
// time: 2019-01-31

package aasw

type node struct {
	isWord bool
	next   map[byte]*node
}

// WordDictionary supports two operations(addWord, search)
type WordDictionary struct {
	root *node
}

// Constructor initialize data structure here.
func Constructor() WordDictionary {
	return WordDictionary{&node{next: make(map[byte]*node)}}
}

// AddWord adds a word into the trie.
func (trie *WordDictionary) AddWord(word string) {
	cur := trie.root
	for i := 0; i < len(word); i++ {
		if _, ok := cur.next[word[i]]; !ok {
			cur.next[word[i]] = &node{next: make(map[byte]*node)}
		}
		cur, _ = cur.next[word[i]]
	}
	if !cur.isWord {
		cur.isWord = true
	}
}

// Search returns if the word is in the trie.
// a word could contain the dot character '.' to represent any ont letter.
func (trie *WordDictionary) Search(word string) bool {
	return match(trie.root, word, 0)
}

func match(n *node, word string, index int) bool {
	if index == len(word) {
		return n.isWord
	}
	if word[index] != '.' {
		if _, ok := n.next[word[index]]; !ok {
			return false
		}
		nextNode, _ := n.next[word[index]]
		return match(nextNode, word, index+1)
	} else {
		for _, nextNode := range n.next {
			if match(nextNode, word, index+1) {
				return true
			}
		}
		return false
	}
}
