package address

import (
	"sort"
	"strings"
)

type sorter struct {
	records    *[][]string
	idFieldIdx int
}

func (s sorter) Sort() {
	records := *s.records
	sort.Slice(records, func(i, j int) bool {
		return strings.Compare(records[i][s.idFieldIdx], records[j][s.idFieldIdx]) == -1
	})
}
