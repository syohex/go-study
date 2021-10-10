package trie

import "unicode/utf8"

type Trie struct {
	nodes map[rune]*Trie
	end   bool
}

func NewTrie() *Trie {
	nodes := make(map[rune]*Trie)
	return &Trie{
		nodes: nodes,
		end:   false,
	}
}

func (t *Trie) Insert(word string) {
	r, size := utf8.DecodeRuneInString(word)
	if _, ok := t.nodes[r]; !ok {
		t.nodes[r] = NewTrie()
	}

	word = word[size:]
	if len(word) == 0 {
		t.nodes[r].end = true
	} else {
		t.nodes[r].Insert(word)
	}
}

func (t *Trie) Search(word string) bool {
	r, size := utf8.DecodeRuneInString(word)
	if _, ok := t.nodes[r]; !ok {
		return false
	}

	word = word[size:]
	if len(word) == 0 {
		return t.nodes[r].end
	}

	return t.nodes[r].Search(word)
}
