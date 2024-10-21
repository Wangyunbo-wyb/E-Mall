package filter

import (
	"unicode/utf8"
)

type Search struct {
	trieWriter *TrieWriter
}

func (s *Search) TrieWriter() *TrieWriter {
	return s.trieWriter
}

func (s *Search) Replace(word []byte, new byte) []byte {
	data := make([]byte, len(word))
	start := 0
	for _, r := range s.Find(word) {
		copy(data[start:], word[start:r.Start])
		start = r.Start
		copy(data[start:], repeatByte(new, r.End-r.Start+1))
		start = r.End + 1
	}
	copy(data[start:], word[start:])
	return data
}

func (s *Search) Find(word []byte) []*Result {
	return s.findByAC(word, false)
}

func repeatByte(b byte, n int) []byte {
	res := make([]byte, n)
	for i := range res {
		res[i] = b
	}
	return res
}

func findSub(s []byte, end int, sub string) (idx int) {
	idx, j := end, len(sub)-1
	for ; j >= 0; idx-- {
		if s[idx] == sub[j] {
			j--
		}
	}
	idx++
	return
}

func decodeBytes(s []byte) (r rune, size int) {
	return utf8.DecodeRune(s)
}

func (s *Search) findByAC(word []byte, single bool) (list []*Result) {
	n := len(word)
	trieRoot := s.trieWriter.trie()
	skipper := s.trieWriter.Skip()

	for i := 0; i < n; {
		v, l := decodeBytes(word[i:]) // Decode the byte array starting at s[i:] and return the first rune and its byte length.
		node, ok := trieRoot.next[v]  // Find the next node in the trie tree
		if !ok {                      // If the next node is not found, skip the current character
			i += l
			continue
		}

		j := i
		var (
			words []byte  // the word found
			res   *Result // record the matched sensitive words and related information.
			skip  int     // record the number of characters skipped
		)
		for {
			words = append(words, word[j:j+l]...) // put the found word into the words slice
			if node.end {                         // if the node is the end of a sensitive word, record the sensitive word and related information
				sub := string(words[len(words)-int(node.length):])          // record the matched substring
				end := j + l - 1                                            // record the end position of the sensitive word in s
				start := findSub(word, end, sub)                            // record the start position of the sensitive word in s
				res = &Result{sub, string(word[start : end+1]), start, end} // record the sensitive word and related information
			}

			j += l
			// skip the meaningless characters
			for {
				v, l = decodeBytes(word[j:])
				if j < n && skipper.ShouldSkip(v) {
					j += l
					skip += l
				} else {
					break
				}
			}

			// if the next node is found, try to use the fail pointer to find the next node
			if next := node.next[v]; next != nil {
				node = next
			} else {
				if res == nil && node.fail.end {
					sub := string(words[len(words)-int(node.fail.length):])
					end := j - 1
					start := findSub(words, end, sub)
					res = &Result{sub, string(word[start : end+1]), start, end}
				}
				if res != nil {
					list = append(list, res)
					words = words[:0]
					if single {
						return
					}
				}
				res = nil
				if node = node.fail.next[v]; node == nil {
					i = j
					break
				} else {
					i = j + l - int(node.length) - skip
				}
			}
		}
	}
	return
}
