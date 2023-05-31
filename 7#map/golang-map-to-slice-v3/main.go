package main
import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	result := make([][]string, 0, len(mapData))
    for k, v := range mapData {
        result = append(result, []string{k, v})
    }
    return result
}

func main() {
	data1 := map[string]string{"hello": "world", "John": "Doe", "age": "14"}
	fmt.Println(MapToSlice(data1))

	data2 := map[string]string{"foo": "33", "bar": "44", "baz": "55"}
	fmt.Println(MapToSlice(data2))

	data3 := map[string]string{}
	fmt.Println(MapToSlice(data3))
}