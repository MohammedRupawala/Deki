package trie

import (
	// "fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Node struct {
	Children map[rune]*Node
	isEnd    bool
}

func NewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
	}
}

type Trie struct {
	RootNode *Node
}

func NewTrie() *Trie {
	return &Trie{
		RootNode: NewNode(),
	}
}

var commands = map[string]struct{}{
	"echo": {},
	"exit": {},
	"pwd":  {},
	"ls":   {},
}

func init() {

	paths := filepath.SplitList(os.Getenv("PATH"))
	for _, path := range paths {
		files, _ := os.ReadDir(path)
		for _, f := range files {
			info, _ := f.Info()
			if !info.IsDir() && info.Mode().Perm()&0111 != 0 {
				commands[info.Name()] = struct{}{}
			}

		}
	}

	// for k := range commands {
	// 	fmt.Println("Command " + k)
	// }
}

func (t *Trie) Insert(word string) error {
	current := t.RootNode
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	for _, char := range strippedWord {
		if current.Children == nil {
			current.Children = make(map[rune]*Node)
		}

		if current.Children[char] == nil {
			current.Children[char] = NewNode()
		}
		current = current.Children[char]
	}
	current.isEnd = true
	return nil
}

func (t *Trie) SearchWord(word string) bool {
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	current := t.RootNode

	for _, char := range strippedWord {
		if current == nil || current.Children[char] == nil {
			return false
		}
		current = current.Children[char]
	}

	return current != nil && current.isEnd
}

func (t *Trie) findOne(prefix string) *Node {
	node := t.RootNode
	for _, char := range prefix {
		if node == nil || node.Children[char] == nil {
			return nil
		}
		node = node.Children[char]
	}
	return node
}

func collect(node *Node, prefix string, results *[]string) {
	if node.isEnd {
		*results = append(*results, prefix)
	}

	for ch, child := range node.Children {
		if child != nil {
			collect(child, prefix+string(ch), results)
		}
	}
}

func (t *Trie) Autocomplete(prefix string) []string {
	node := t.findOne(prefix)
	if node == nil {
		return nil
	}

	var results []string
	collect(node, prefix, &results)
	sort.Strings(results)
	return results
}

func CreateTrie() *Trie {
	trie := NewTrie()
	for key := range commands {
		trie.Insert(key)
	}
	

	return trie
}
