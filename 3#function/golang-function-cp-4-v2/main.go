package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := make([]string, 0)
	for _, v := range data {
		apakahMengandungInput := strings.Contains(v, input)

		if apakahMengandungInput {
                result = append(result, v)
            }
		}
	return strings.Join(result, ",") // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
