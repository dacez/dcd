package filter

import (
	"fmt"
	"github.com/dacez/dcd/config"
	"github.com/dacez/dcd/line"
	"runtime"
	"sort"
	"strings"
)

type FilterType int

type ResultItem struct {
	Index int
	Line  line.Line
	Pos   int
}

type ResultItemSlice []ResultItem

func (p ResultItemSlice) Len() int {
	return len(p)
}

func (p ResultItemSlice) Less(i, j int) bool {
	return len(p[i].Line.Cs)-p[i].Index < len(p[j].Line.Cs)-p[j].Index
}

func (p ResultItemSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Filter struct {
	Results ResultItemSlice
	Type    FilterType
}

func (p *Filter) Init(ls []line.Line) {
	p.Results = make(ResultItemSlice, 0)
	for i, v := range ls {
		p.Results = append(p.Results, ResultItem{Pos: i, Index: 0, Line: v})
	}
}

func (p *Filter) Push(k string) {
	partResult := make(ResultItemSlice, 0)
	s := "/"
	if runtime.GOOS == "windows" {
		s = "\\"
	}
	for _, v := range p.Results {
		if v.Index >= len(v.Line.Cs) {
			continue
		}
		l := v.Line.GetString()
		if config.GetConfig().FuzzyFindMode == config.NameMode && v.Index == 0 {
			li := strings.LastIndex(l, s)
			if li != -1 && li <= len(l)-1 {
				if li == len(l)-1 {
					v.Index = li
				} else {
					v.Index = li + 1
				}
			}
		}
		k = strings.ToLower(k)
		l = strings.ToLower(l[v.Index:])
		if strings.Contains(l, k) {
			v.Index += strings.Index(l, k) + len(k)
			partResult = append(partResult, v)
		}
	}
	sort.Sort(partResult)
	p.Results = partResult
}

func (p *Filter) Print() {
	for _, v := range p.Results {
		fmt.Println(v.Line.GetString())
	}
}
