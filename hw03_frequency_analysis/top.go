package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type freqDict map[string]int
type freqSlice struct {
	word  string
	count int
}

func newFreqSlice(w string) *freqSlice {
	return &freqSlice{
		word:  w,
		count: 1,
	}
}

type word struct {
	wcDict  freqDict
	wcSlice []freqSlice
}

func newWord() *word {
	return &word{
		wcDict:  make(freqDict),
		wcSlice: []freqSlice{},
	}
}

func (w *word) add(s string) {
	if _, ok := w.wcDict[s]; ok {
		w.wcDict[s]++
	} else {
		w.wcDict[s] = 1
		w.wcSlice = append(w.wcSlice, *newFreqSlice(s))
	}
}

func (w *word) populate() {
	for i, word := range w.wcSlice {
		w.wcSlice[i].count = w.wcDict[word.word]
	}
	// clear map
	w.wcDict = make(freqDict)
}

func (w *word) sort() []string {
	sort.Slice(w.wcSlice, func(i, j int) bool {
		if w.wcSlice[i].count == w.wcSlice[j].count {
			return w.wcSlice[i].word < w.wcSlice[j].word
		}
		return w.wcSlice[i].count > w.wcSlice[j].count
	})
	slc := w.wcSlice[:10]
	result := []string{}
	for _, val := range slc {
		result = append(result, val.word)
	}

	return result
}

func Top10(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	stat := newWord()
	words := strings.Fields(str)

	for _, word := range words {
		stat.add(word)
	}

	stat.populate()

	return stat.sort()
}
