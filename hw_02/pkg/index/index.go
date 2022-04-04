package index

import (
	"fmt"
	"go_hw_02/pkg/crawler"
	"go_hw_02/pkg/crawler/spider"
	"sort"
	"strings"
)

type ByID []crawler.Document

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func Index(l []string) (map[string][]int, []crawler.Document) {
	id := 1
	s := parseDocument(l)
	in := make(map[string][]int)

	for i, v := range s {
		s[i].ID = id
		v.ID = id
		tmp := strings.Fields(v.Title)
		for i := 0; i < len(tmp); i++ {
			k := strings.ToLower(tmp[i])
			in[k] = append(in[k], v.ID)
		}
		id++
	}
	sort.Sort(ByID(s))

	return in, s
}

func parseDocument(u []string) []crawler.Document {
	ad := []crawler.Document{}

	s := spider.New()

	for _, v := range u {
		list, err := s.Scan(v, 1)
		if err != nil {
			fmt.Println(v, " по ссылке ничего не найдено")
		}
		ad = append(ad, list...)
	}

	return ad
}
