package main

import "fmt"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children  [26]*TrieNode
	endOfWord bool
}

func newTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (trie *Trie) Insert(word string) {
	curr := trie.root

	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')

		if curr.children[idx] == nil {
			curr.children[idx] = &TrieNode{}
		}
		if i == len(word)-1 {
			curr.children[idx].endOfWord = true
		}
		curr = curr.children[idx]
	}
}

func (trie *Trie) Search(key string) bool {
	curr := trie.root

	for i := 0; i < len(key); i++ {
		idx := int(key[i] - 'a')

		if curr.children[idx] == nil {
			return false
		}
		if i == len(key)-1 && curr.children[idx].endOfWord == false {
			return false
		}
		curr = curr.children[idx]
	}

	return true
}

func (trie *Trie) wordBreak(s string) bool {
	if s == "" {
		return true
	}
	for i := 1; i <= len(s); i++ {
		firstPart := s[:i]
		secondPart := s[i:]
		if trie.Search(firstPart) && trie.wordBreak(secondPart) {
			return true
		}
	}
	return false
}

func (trie *Trie) countUniqueSubstring(root *TrieNode) int {
	if root == nil {
		return 0
	}
	count := 0
	for i := 0; i < 26; i++ {
		if root.children[i] != nil {
			count += trie.countUniqueSubstring(root.children[i])
		}
	}
	return count + 1

}

func (t *Trie) StartWith(prefix string) bool {
	curr := t.root
	for i := 0; i < len(prefix); i++ {
		idx := int(prefix[i] - 'a')
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}

var ans string

func (trie *Trie) LongestWord(root *TrieNode, temp string) {
	if root == nil {
		return
	}
	for i := 0; i < 26; i++ {
		if root.children[i] != nil && root.children[i].endOfWord {
			temp += string('a' + i)
			if len(temp) > len(ans) {
				ans = temp
			}
			trie.LongestWord(root.children[i], temp)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println("----------Trie----------------")
	t := newTrie()

	// words := []string{"the", "a", "there", "their", "any", "thee"}
	// for _, word := range words {
	// 	t.Insert(word)
	// }
	// fmt.Println("there ->", t.Search("there"))
	// fmt.Println("thor ->", t.Search("thor"))
	// fmt.Println("thee ->", t.Search("thee"))
	// fmt.Println("a ->", t.Search("a"))
	// fmt.Println("three ->", t.Search("three"))

	var words = []string{"i", "like", "sam", "samsung", "mobile"}
	key := "ilikesamsung"
	for _, word := range words {
		t.Insert(word)
	}
	fmt.Println("i ->", t.Search("i"))
	fmt.Println("like ->", t.Search("like"))
	fmt.Println("samsung ->", t.Search("samsung"))
	fmt.Println("Word break-->", t.wordBreak(key))
	fmt.Println("startwith li-->", t.StartWith("li"))
	fmt.Println("startwith mobile-->", t.StartWith("mobile"))
	fmt.Println("startwith likx-->", t.StartWith("likx"))

	s := "ababa"
	for i := 0; i < len(s); i++ {
		suffix := s[i:]
		t.Insert(suffix)
	}

	fmt.Printf("Substring count is :: %v\n", t.countUniqueSubstring(t.root))

	words = []string{"a", "banana", "app", "appl", "ap", "apple", "apply"}
	for _, word := range words {
		t.Insert(word)
	}

	ans = ""
	t.LongestWord(t.root, "")
	fmt.Println(ans)
}
