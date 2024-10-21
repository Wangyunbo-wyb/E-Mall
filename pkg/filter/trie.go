package filter

import (
	"unicode/utf8"
)

type trie struct {
	next   map[rune]*trie
	fail   *trie
	length uint8 // the length of all the words in the trie
	end    bool
}

// TrieWriter includes the operation of the trie tree
type TrieWriter struct {
	size     int   //the number of words in the trie
	skip     *Skip //the char in the sortedSkipList will be skipped
	trieRoot *trie
}

func (t *TrieWriter) trie() *trie {
	return t.trieRoot
}

func (t *TrieWriter) Skip() *Skip {
	return t.skip
}

func NewTrieWriter() *TrieWriter {
	return &TrieWriter{
		trieRoot: &trie{next: map[rune]*trie{}},
	}
}

func (t *TrieWriter) Insert(word string) *TrieWriter {
	move := t.trieRoot
	wLen := len(word)
	for _, w := range word {
		//if the char is in the sortedSkipList, skip it
		if t.skip.ShouldSkip(w) {
			continue
		}
		//if the next node is nil, create a new node
		if move.next[w] == nil {
			move.next[w] = &trie{next: map[rune]*trie{}}
		}
		wLen += utf8.RuneLen(w)
		move = move.next[w]
		move.length = uint8(wLen)
	}
	if wLen > 0 && !move.end {
		move.end = true
		t.size++
	}
	return t
}

func (t *TrieWriter) InsertWords(words []string) *TrieWriter {
	for _, word := range words {
		t.Insert(word)
	}
	return t
}

func (t *TrieWriter) BuildFail() int {
	//create a queue to store the nodes
	queue := make([]*trie, 0, len(t.trieRoot.next))
	for _, node := range t.trieRoot.next {
		node.fail = t.trieRoot // set the fail of the node to the root node
		queue = append(queue, node)
	}

	count := 0
	for level := 1; len(queue) > 0; level++ {
		temp := make([]*trie, len(queue))
		copy(temp, queue)
		// clear the queue and store the next level nodes
		queue = queue[:0]
		for _, prev := range temp {
			for c, curr := range prev.next {
				count++
				queue = append(queue, curr)
				fail := prev.fail
				for fail != nil && fail.next[c] == nil {
					fail = fail.fail
				}
				if fail == nil {
					curr.fail = t.trieRoot
				} else {
					curr.fail = fail.next[c]
				}
			}
		}
	}
	return count
}

func (t *TrieWriter) setSkip(skip *Skip) *TrieWriter {
	t.skip = skip
	return t
}
