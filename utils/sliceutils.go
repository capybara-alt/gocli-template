package utils

import "sort"

func Contains(src []string, target string) bool {
	sort.Strings(src)
	return sort.SearchStrings(src, target) != len(src)
}
