package filter

import (
	"bufio"
	"io"
	"os"
	"sort"
)

type options struct {
	writer *TrieWriter
	skip   *Skip
}

type Option func(options *options)

func NewSearch(opts ...Option) *Search {
	opt := &options{
		skip:   &Skip{list: []rune(sortedSkipList)},
		writer: NewTrieWriter(),
	}
	for _, o := range opts {
		o(opt)
	}
	opt.writer.setSkip(opt.skip)
	return &Search{opt.writer}
}

func LoadWordDict(path string, skip ...string) (*Search, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	buf := bufio.NewReader(f)
	var words []string
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		words = append(words, string(line))
	}
	search := NewSearch(SetSortedRunesSkip(skipStr(skip...)))
	search.TrieWriter().InsertWords(words).BuildFail()
	return search, err
}

func SetSortedRunesSkip(s []rune) Option {
	return func(options *options) {
		skip := &Skip{list: s}
		options.skip = skip
	}
}

func skipStr(skip ...string) []rune {
	if len(skip) > 0 {
		runes := []rune(skip[0])
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		return runes
	}
	return []rune(SortedSkipList())
}
