package main

import (
	"fmt"
	"strings"
)

const (
	JKT = "JKT"
	BDG = "BDG"
	BKS = "BKS"
	DPK = "DPK"
)

var (
	hariBuka = map[string][]string{
		JKT: {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu"},
		BDG: {"rabu", "kamis", "sabtu"},
		BKS: {"selasa", "kamis", "jumat"},
		DPK: {"senin", "selasa"},
	}

	biayaAdmin = map[string]float32{
		"senin":   0.1,
		"selasa":  0.05,
		"rabu":    0.1,
		"kamis":   0.05,
		"jumat":   0.1,
		"sabtu":   0.05,
		"minggu":  0,
		"default": 0.1,
	}
)

func DeliveryOrder(data []string, day string) map[string]float32 {
	filteredData := make(map[string]float32)

	for _, d := range data {
		items := strings.Split(d, ":")
		name := items[0] + "-" + items[1]
		price := parseFloat(items[2])
		location := items[3]

		if isValidLocationDay(location, day) {
			total := calculateTotalPrice(price, day)
			filteredData[name] += total
		}
	}

	return filteredData
} // TODO: replace this
func parseFloat(s string) float32 {
	var f float32
	fmt.Sscanf(s, "%f", &f)
	return f
}

func isValidLocationDay(location, day string) bool {
	days, found := hariBuka[location]
	if !found {
		return false
	}

	for _, d := range days {
		if d == day {
			return true
		}
	}

	return false
}
func calculateTotalPrice(price float32, day string) float32 {
	adminFee, found := biayaAdmin[day]
	if !found {
		adminFee = biayaAdmin["default"]
	}

	return price + (price * adminFee)
}
func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
