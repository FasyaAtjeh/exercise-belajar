package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	result := make([]map[string]interface{}, len(data))
	for i, d := range data {
		parts := strings.Split(d, ";")
		if len(parts) < 3 {
			result[i] = nil
			continue
		}
		age, err := strconv.Atoi(parts[1])
		if err != nil {
			result[i] = nil
			continue
		}
		item := make(map[string]interface{})
		item["name"] = parts[0]
		item["age"] = age
		item["address"] = parts[2]
		if len(parts) > 3 && parts[3] != "" {
			height, err := strconv.ParseFloat(parts[3], 64)
			if err == nil {
				item["height"] = height
			}
		}
		if len(parts) > 4 && parts[4] != "" {
			isMarried, err := strconv.ParseBool(parts[4])
			if err == nil {
				item["isMarried"] = isMarried
			}
		}
		result[i] = item
	}
	return result
}
func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	result := PopulationData(data)
	fmt.Println(result)
}
