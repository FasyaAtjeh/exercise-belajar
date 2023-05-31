package main

import (
	"errors"
	"fmt"
	"strings"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
		return
	}
	if !website.Valid {
		chErr <- errors.New("domain not valid")
		return
	}
	if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
		return
	}
	TLD, IDN_TLD := GetTLD(website.Domain)
	website.TLD, website.IDN_TLD = TLD, IDN_TLD
	ch <- website
	chErr <- nil
}

var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error, 1)
	var filterData []RowData
	var dataErr error
	for _, info := range data {
		if info.Domain == "" {
			return []RowData{}, errors.New("domain name is empty")
		}
		go ProcessGetTLD(info, ch, errCh)
		if strings.Contains(info.Domain, TLD) {
			filterData = append(filterData, <-ch)
			dataErr = <-errCh
			if dataErr != nil {
				return nil, dataErr
			}
		}
	}
	return filterData, nil //TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
