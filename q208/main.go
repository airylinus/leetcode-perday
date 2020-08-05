package main

import "fmt"

func main() {
	// t1 := Trie{Val: []rune("b")[0]}
	// t2 := Trie{Val: []rune("e")[0]}
	// t3 := Trie{Val: []rune("s")[0]}
	// t4 := Trie{Val: []rune("t")[0]}
	// t5 := Trie{Val: "e"}

	// t3.Children = append(t3.Children, &t4)
	// t2.Children = append(t2.Children, &t3)
	// t1.Children = append(t1.Children, &t2)
	// t1.Children = append(t1.Children, &t5)
	// t1.Display()
}

// Trie ...
type Trie struct {
	Val map[rune]*Trie
}

// Constructor ** Initialize your data structure here. */
func Constructor() Trie {
	t := new(Trie)
	return *t
}

//Insert Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	for _, v := range word {
		if v == rune(t.Val) {
			continue
		} else {

		}
	}
}

//Search * * Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	return true
}

//StartsWith ** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return true
}

// Display show up tier
func (t *Trie) Display() {
	// tmp := t
	fmt.Println(t.Val, " : ")
	if t.Children != nil {
		for _, x := range t.Children {
			fmt.Println("\t", x.Val, ",")
			x.Display()
		}
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
