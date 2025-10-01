type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func trieData() *Trie {
	t := new(Trie)
	t.root = new(TrieNode)

	return t
}

func (t *Trie) Insert(word string) {
	current := t.root

	for _, wr := range word {
		index := wr - 'a'

		if current.children[index] == nil {
			current.children[index] = new(TrieNode)
		}

		current = current.children[index]
	}

	current.isEnd = true
}

func (t *Trie) Search(word string) int {
	current := t.root

	for _, wr := range word {
		index := wr - 'a'

		if current.children[index] == nil {
			return 0
		}

		current = current.children[index]
	}

	if current.isEnd {
		return 1
	}

	return 0
}

func getSingleChild(node *TrieNode) (int, bool) {
	count := 0
	idx := -1

	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			count++
			idx = i
		}
	}

	if count == 1 {
		return idx, true
	}

	return -1, false
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	trie := trieData()

	for _, str := range strs {
		trie.Insert(str)
	}

	node := trie.root
	prefix := ""

	for {
		idx, ok := getSingleChild(node)

		if !ok || node.isEnd {
			break
		}

		prefix += string(rune('a' + idx))

		node = node.children[idx]
	}

	return prefix
}
