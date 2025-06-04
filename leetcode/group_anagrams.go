package qdxl

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	seenMap := make(map[string]int)
	res := make([][]string, 0)
	for _, str := range strs {
		s := []rune(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		if i, ok := seenMap[sortedStr]; ok {
			res[i] = append(res[i], str)
		} else {
			res = append(res, []string{str})
			seenMap[sortedStr] = len(res) - 1
		}
	}
	return res
}

func groupAnagrams1(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	res := make([][]string, 0, len(mp))
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
