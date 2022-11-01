package basic_trie

// The number of possible characters
const AlphabetSize = 26

type node struct {
	children [AlphabetSize]*node
	terminal bool
}

type trie struct {
	root *node
}

// creates a new trie with initialized root node.
func New() *trie {
	return &trie{root: &node{}}
}

// returns true if root is the leaf node.
func isLastNode(root *node) bool {
	for i := 0; i < AlphabetSize; i++ {
		if root.children[i] != nil {
			return false
		}
	}
	return true
}

// Recursively counts the number of child nodes in root node
// adding to the count.
func countNodesRec(root *node, count *int) {
	for i := 0; i < AlphabetSize; i++ {
		if root.children[i] != nil {
			(*count)++
			countNodesRec(root.children[i], count)
		}
	}
}

// Returns the number of nodes in the trie
func (t *trie) Nodes() int {
	count := 0
	countNodesRec(t.root, &count)
	return count
}

// inserts word into the trie
func (t *trie) Insert(word string) {
	wordLen := len(word)
	currentNode := t.root

	for i := 0; i < wordLen; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.terminal = true
}

// returns true if word is in the trie
func (t *trie) Exists(word string) bool {
	wordLen := len(word)
	currentNode := t.root

	for i := 0; i < wordLen; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.terminal
}

// recursively searches for words in child nodes of triNode
// and appends a found word to words slice.
func suggestionsRec(triNode *node, prefix string, words *[]string) {
	if triNode.terminal {
		*words = append(*words, prefix)
	}

	for i := 0; i < AlphabetSize; i++ {
		if triNode.children[i] != nil {
			// child node character value
			child := rune('a' + i)
			suggestionsRec(triNode.children[i], prefix+string(child), words)
		}
	}
}

// returns a slice of words with a prefix query from node.
func (t *trie) suggestions(node *node, query string) []string {
	words := []string{}
	wordLen := len(query)
	currentNode := t.root

	for i := 0; i < wordLen; i++ {
		charIndex := query[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return words
		}
		currentNode = currentNode.children[charIndex]
	}

	// If the prefix is present as a word but there is no subtree below the matching node.
	if isLastNode(currentNode) {
		return words
	}

	suggestionsRec(currentNode, query, &words)
	return words
}

// Returns a slice of words starting with prefix query.
func (t *trie) GetAutoSuggestions(query string) []string {
	if t.root == nil {
		return []string{}
	}

	return t.suggestions(t.root, query)
}

func (t *trie) sizeRec(root *node, wordCount *int, prefix string) {
	if root.terminal {
		(*wordCount)++
	}

	for i := 0; i < AlphabetSize; i++ {
		if root.children[i] != nil {
			child := rune('a' + i)
			t.sizeRec(root.children[i], wordCount, prefix+string(child))
		}
	}
}

// Returns the number of words in the trie.
func (t *trie) Size() int {
	n := 0
	t.sizeRec(t.root, &n, "")
	return n
}

func (t *trie) Memory() int {
	padding := 7 // 7 bytes of pading after the bool
	nodeSize := (8 * AlphabetSize) + 1 + padding
	return t.Nodes() * nodeSize
}
