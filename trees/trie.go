package trees

type TrieNode struct {
	char           rune
	children       map[rune]*TrieNode
	terminatesWord bool
}

func NewTrieNode(c rune, terminal bool) *TrieNode {
	children := make(map[rune]*TrieNode)
	return &TrieNode{char: c, children: children, terminatesWord: terminal}
}

type Trie struct {
	base map[rune]*TrieNode
}

func NewTrie() *Trie {
	m := make(map[rune]*TrieNode)
	return &Trie{base: m}
}

func (t *Trie) insert(word string) {
	if len(word) == 0 {
		return
	}
	//convert string to array of characters
	chars := []rune(word)

	//check if first char is in the base layer
	c := chars[0]

	term := false
	if len(chars) == 1 {
		term = true
	}

	_, ok := t.base[c]
	if !ok {
		t.base[c] = NewTrieNode(c, term)
	} else {
		if term {
			t.base[c].terminatesWord = true
		}
	}
	node := t.base[c]
	//otherwise iterate through chars and add to trie
	for idx, char := range chars[1:] {
		//if it's the last character, terminate the word
		term := false
		if idx == len(chars)-2 {
			term = true
		}

		//check if the child node exists
		_, ok := node.children[char]
		if !ok {
			// if it doesn't create a new node
			node.children[char] = NewTrieNode(char, term)
		} else {
			// if it does check if we need to terminate a word
			if term {
				node.children[char].terminatesWord = true
			}
		}
		//go to the next child node
		node = node.children[char]
	}

}

func generateWords(node *TrieNode, prefix string, output chan string) {
	if node == nil {
		return
	}
	for c, n := range node.children {
		tmp := prefix + string(c)
		if n.terminatesWord {
			output <- tmp
		}
		go generateWords(n, tmp, output)
	}
}

func autocomplete(prefix string, T Trie, output chan string) {
	//basic autocomplete given a prefix
	inputChars := []rune(prefix)
	node, ok := T.base[inputChars[0]]
	if !ok {
		return
	}

	n := node
	//iterate to lowest node of prefix
	for _, c := range inputChars[1:] {
		node, ok := node.children[c]
		n = node
		if !ok {
			return
		}
	}
	go generateWords(n, prefix, output)
}
