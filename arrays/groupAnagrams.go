package arrays

import (
	"sort"
)

type runesSlice []rune

func (s runesSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s runesSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s runesSlice) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(runesSlice(r))
	return string(r)
}

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		sorted_str := SortString(str)
		groups[sorted_str] = append(groups[sorted_str], str)
	}
	result := make([][]string, len(groups))
	for _, arr := range groups {
		if len(arr) > 0 {
			result = append(result, arr)
		}
	}
	return result
}
