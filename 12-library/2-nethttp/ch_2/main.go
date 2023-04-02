package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Package struct {
	ModulePath    string  `json:"module_path"`
	DownloadCount float64 `json:"download_count"`
}

type ByDownloadCount []Package

func (a ByDownloadCount) Len() int           { return len(a) }
func (a ByDownloadCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDownloadCount) Less(i, j int) bool { return a[i].DownloadCount > a[j].DownloadCount }

func main() {
	url := "https://goproxy.cn/stats/trends/latest"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var packages []Package
	err = json.Unmarshal(body, &packages)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	sort.Sort(ByDownloadCount(packages))

	for _, p := range packages {
		fmt.Printf("module_path: %s, download_count: %f\n", p.ModulePath, p.DownloadCount)
	}
}
