package utils

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func SnakeUpperToCamelCase(str string) string {
	words := strings.Split(str, "_")
	results := make([]string, len(words))
	for index, word := range words {
		results[index] = strings.Title(strings.ToLower(word))
	} 

	return strings.Join(results, "")
}


func CamelToSnakeUpperCase(str string) string {
  snakeCaseStr := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
  snakeCaseStr  = matchAllCap.ReplaceAllString(snakeCaseStr, "${1}_${2}")

	return strings.ToUpper(snakeCaseStr)
}
