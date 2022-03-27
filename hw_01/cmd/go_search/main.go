package main

import (
	"flag"
	"fmt"
	"go_hw_01/pkg/crawler"
	"go_hw_01/pkg/crawler/spider"
	"strings"
)

func main() {
	sFlag := flag.String("s", "", "help message for flag s")
	flag.Parse()

	urls := []string{"https://go.dev/", "https://golang.org/"}
	listAll := []crawler.Document{}

	s := spider.New()

	for _, v := range urls {
		list, err := s.Scan(v, 1)
		if err != nil {
			fmt.Println(v, " по ссылке ничего не найдено")
		}
		listAll = append(listAll, list...)
	}

	for _, v := range listAll {
		containTitle := strings.Contains(v.URL, *sFlag)
		if containTitle && *sFlag != "" {
			fmt.Println("Вхождения найдены в: ", v.Title, v.URL)
		}
	}
}
