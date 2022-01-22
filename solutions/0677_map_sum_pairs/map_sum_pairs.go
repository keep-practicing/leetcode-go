/*
677. Map Sum Pairs
https://leetcode.com/problems/map-sum-pairs/

Implement a MapSum class with insert, and sum methods.

For the method insert, you'll be given a pair of (string, integer).
The string represents the key and the integer represents the value.
If the key already existed, then the original key-value pair will be overridden to the new one.

For the method sum, you'll be given a string representing the prefix,
and you need to return the sum of all the pairs' value whose key starts with the prefix.
*/
// time: 2019-02-01

package mapsumpairs

type node struct {
	val  int
	next map[rune]*node
}

// MapSum data structure for solution.
type MapSum struct {
	root *node
}

// Constructor initialize data structure here.
func Constructor() MapSum {
	return MapSum{&node{next: make(map[rune]*node)}}
}

// Insert inserts a word into the trie.
func (ms *MapSum) Insert(key string, val int) {
	cur := ms.root
	for _, c := range key {
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = &node{next: make(map[rune]*node)}
		}
		cur = cur.next[c]
	}
	cur.val = val
}

// Sum sum of all the pairs' value whose key starts with the prefix.
func (ms *MapSum) Sum(prefix string) int {
	cur := ms.root
	for _, c := range prefix {
		if _, ok := cur.next[c]; !ok {
			return 0
		}
		cur = cur.next[c]
	}
	return sum(cur)
}

func sum(n *node) int {
	res := n.val
	for _, nextNode := range n.next {
		res += sum(nextNode)
	}
	return res
}
