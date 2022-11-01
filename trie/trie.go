// Effecient Trie data structure implementation using a map.
//
// It can store all UTF-8 characters in runes as supported in golang.package trie
package trie

import "sort"

// Trie holds the data in a prefix tree using an unordered map
type Trie struct {
	hash      map[rune]*Trie
	isWordEnd bool
}

// initializes a new Trie
func NewTrie() *Trie {
	return &Trie{
		hash:      map[rune]*Trie{},
		isWordEnd: false,
	}
}

// Insert one or more words into the Tri.
//
// Insertion is O(m) where m is the number of letters in the word.
func (t *Trie) Insert(words ...string) {
	for _, word := range words {
		var temp *Trie
		temp = t

		for i := 0; i < len(word); i++ {
			x := rune(word[i])

			// make the node if there is no path
			if _, ok := temp.hash[x]; !ok {
				temp.hash[x] = NewTrie()
			}
			temp = temp.hash[x]
		}
		temp.isWordEnd = true
	}
}

// Returns true if word exists in the trie.
//
// O(m) time complexity
func (t *Trie) Exists(word string) bool {
	var temp *Trie
	temp = t

	for i := 0; i < len(word); i++ {
		x := rune(word[i])
		node, ok := temp.hash[x]
		if !ok {
			return false
		}
		temp = node

	}
	return temp.isWordEnd
}

// Recursive function to find all words on the node root.
// Successive runes are concatednated until end of word. At end of word
// a new word is appended to words slice.
func (t *Trie) suggestionRec(root *Trie, prefix string, words *[]string) {
	// If we are at the end of the word, append to words list
	if root.isWordEnd {
		*words = append(*words, prefix)
	}

	// recursively explore all the nodes in the map
	for char, node := range root.hash {
		// base case as there is no recursion if node is nil
		if node != nil {
			t.suggestionRec(node, prefix+string(char), words)
		}
	}
}

// Returns a slice of words matching query from the trie.
// Worst case: O(m+n) where m is the number of characters in query
// and n is the number of nodes in the trie.
func (t *Trie) Suggestions(query string) []string {
	words := []string{}
	wordLen := len(query)

	currentNode := t

	for i := 0; i < wordLen; i++ {
		x := rune(query[i])
		node, ok := currentNode.hash[x]
		if !ok {
			return []string{}
		}
		currentNode = node
	}

	t.suggestionRec(currentNode, query, &words)
	return words
}

// Returns all words in the trie in no particular order.
func (t *Trie) Words() []string {
	words := []string{}
	t.suggestionRec(t, "", &words)
	return words
}

// Returns a slice of words sorted alphabetically
func (t *Trie) SortedWords() []string {
	words := t.Words()
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
	return words
}

// Recursive function to find wordcount along the node root in trie.
func (t *Trie) sizeRec(root *Trie, wordCount *int, prefix string) {
	if root.isWordEnd {
		(*wordCount)++
	}

	for char, node := range root.hash {
		// base case as there is no recursion if node is nil
		if node != nil {
			t.sizeRec(node, wordCount, prefix+string(char))
		}
	}
}

// Returns the number of words in the trie.
func (t *Trie) Size() int {
	n := 0
	t.sizeRec(t, &n, "")
	return n
}

// Recursively counts the number of child nodes in root node
// adding to the count.
func countNodesRec(root *Trie, count *int) {
	for _, v := range root.hash {
		if v != nil {
			(*count)++
			countNodesRec(v, count)
		}
	}
}

// Returns the number of nodes in the trie
func (t *Trie) Nodes() int {
	count := 0
	countNodesRec(t, &count)
	return count
}
