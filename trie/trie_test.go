package trie

import (
	"testing"
)

func TestTrie_Search(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("orange")
	trie.Insert("melon")

	tests := []struct {
		word string
		want bool
	}{
		{"apple", true},
		{"app", true},
		{"appl", false},
		{"orange", true},
		{"or", false},
		{"melon", true},
	}
	for _, tt := range tests {
		t.Run("search", func(t *testing.T) {
			if got := trie.Search(tt.word); got != tt.want {
				t.Errorf("Trie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
