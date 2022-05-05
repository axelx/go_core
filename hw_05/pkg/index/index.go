package index

import (
	"go_hw_03/pkg/crawler"
	"strings"
)

func Index(docs []crawler.Document) map[string][]int {
	id := 1
	var in = make(map[string][]int)

	for i, v := range docs {
		docs[i].ID = id
		v.ID = id
		tmp := strings.Fields(v.Title)
		for i := 0; i < len(tmp); i++ {
			k := strings.ToLower(tmp[i])
			in[k] = append(in[k], v.ID)
		}
		id++
	}

	return in
}
