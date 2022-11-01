package trie_test

import (
	"testing"

	"github.com/abiiranathan/algo/trie"
)

func TestTrie(t *testing.T) {
	// Initialize a new trie
	tr := trie.NewTrie()

	// must be empty
	if tr.Size() != 0 {
		t.Errorf("empty trie should have a size of zero, got: %d", tr.Size())
	}

	// Insert items into the trie
	tr.Insert("cat")
	tr.Insert("cattle", "cow")
	tr.Insert("cowflower")

	// multiple insertions
	tr.Insert("mundane", "mundane fact")
	tr.Insert("123 word")

	// must be a valid size
	if tr.Size() != 7 {
		t.Errorf("trie.Size() after insert should be 7, got: %d", tr.Size())
	}

	// test the number of words after insertion
	words := tr.Words()

	if len(words) != 7 {
		t.Errorf("trie.Words() returned from trie do not match the expected length: %d", tr.Size())
	}

	// Check existence of word in trie
	if !tr.Exists("cat") {
		t.Errorf("trie.Exists() returned false unexpectedly")
	}

	// words that don't exist in trie
	if tr.Exists("not found") {
		t.Errorf("trie.Exists() returned true on non-existing word")
	}

	// Sorted words
	sorted := tr.SortedWords()

	if len(sorted) != len(words) {
		t.Errorf("SortedWords() and Words() should return equal elements")
	}

	// auto complete suggestions
	matches := tr.Suggestions("ca")
	expected := []string{"cat", "cattle"}

	wordExists := func(w string) (found bool) {
		for _, word := range expected {
			if w == word {
				found = true
				return
			}
		}
		return
	}

	for _, match := range matches {
		if !wordExists(match) {
			t.Errorf("expected a match %s, not found in trie.", match)
		}
	}

	// word not in trie
	m := tr.Suggestions("not found")
	if len(m) != 0 {
		t.Errorf("expected no suggestions for a word not in trie.")
	}

	// Test number of nodes
	treeSize := tr.Nodes()
	if treeSize != 34 {
		t.Errorf("expected number of nodes to be %d, got %d", 43, treeSize)
	}

	tr2 := trie.NewTrie()
	tr2.Insert("git", "github", "github.com")

	if tr2.Nodes() != 10 {
		t.Errorf("expected number of nodes to be %d, got %d", 10, tr2.Nodes())
	}
}
