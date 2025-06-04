package qdxl

import (
	"reflect"
	"sort"
	"testing"
)

func compareStringArrays(got, want [][]string) bool {
	if len(got) != len(want) {
		return false
	}

	// Sort each inner array
	sortArrays := func(arr [][]string) [][]string {
		sorted := make([][]string, len(arr))
		for i, v := range arr {
			sorted[i] = make([]string, len(v))
			copy(sorted[i], v)
			sort.Strings(sorted[i])
		}
		sort.Slice(sorted, func(i, j int) bool {
			if len(sorted[i]) != len(sorted[j]) {
				return len(sorted[i]) < len(sorted[j])
			}
			return sorted[i][0] < sorted[j][0]
		})
		return sorted
	}

	return reflect.DeepEqual(sortArrays(got), sortArrays(want))
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		words []string
		want  [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}
	for _, test := range tests {
		if got := groupAnagrams(test.words); !compareStringArrays(got, test.want) {
			t.Errorf("groupAnagrams(%v) = %v, want %v", test.words, got, test.want)
		}
	}
}

func TestGroupAnagrams2(t *testing.T) {
	tests := []struct {
		words []string
		want  [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}
	for _, test := range tests {
		if got := groupAnagrams2(test.words); !compareStringArrays(got, test.want) {
			t.Errorf("groupAnagrams2(%v) = %v, want %v", test.words, got, test.want)
		}
	}
}

func TestGroupAnagrams1(t *testing.T) {
	tests := []struct {
		words []string
		want  [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}
	for _, test := range tests {
		if got := groupAnagrams1(test.words); !compareStringArrays(got, test.want) {
			t.Errorf("groupAnagrams1(%v) = %v, want %v", test.words, got, test.want)
		}
	}
}
