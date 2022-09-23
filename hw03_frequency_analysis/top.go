package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const maxResultSliceLength = 10

func Top10(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	words := strings.Fields(str)

	wcDict := make(map[string]int, len(words))
	wcSlice := []string{}

	for _, word := range words {
		wcDict[word]++
	}

	for k := range wcDict {
		wcSlice = append(wcSlice, k)
	}

	sort.Slice(wcSlice, func(i, j int) bool {
		if wcDict[wcSlice[i]] == wcDict[wcSlice[j]] {
			return wcSlice[i] < wcSlice[j]
		}

		return wcDict[wcSlice[i]] > wcDict[wcSlice[j]]
	})

	if len(wcSlice) < maxResultSliceLength {
		return wcSlice
	}

	return wcSlice[:maxResultSliceLength]
}
