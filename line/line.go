package line

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"unicode/utf8"
)

type Line struct {
	Cs []termbox.Cell
	Bg termbox.Attribute
	Fg termbox.Attribute
}

func (p *Line) GetString() string {
	s := ""
	for _, v := range p.Cs {
		s += string(v.Ch)
	}
	return s
}

func (p *Line) PushCell(c termbox.Cell) {
	p.Cs = append(p.Cs, c)
}

func (p *Line) PopCell() int {
	if len(p.Cs) > 0 {
		w := runewidth.RuneWidth(p.Cs[len(p.Cs)-1].Ch)
		p.Cs = p.Cs[0 : len(p.Cs)-1]
		return w
	}
	return 0
}

func (p *Line) PushBytes(b []byte) {
	var utf8str []rune
	for i := 0; i < len(b); {
		r, l := utf8.DecodeRune(b[i:])
		if l == 0 || r == utf8.RuneError {
			break
		}
		i += l
		if r == '\t' {
			utf8str = append(utf8str, ' ')
			utf8str = append(utf8str, ' ')
		} else {
			utf8str = append(utf8str, r)
		}
	}
	for _, v := range utf8str {
		p.Cs = append(p.Cs, termbox.Cell{Ch: v, Bg: p.Bg, Fg: p.Fg})
	}
}

func (p *Line) GetHeight(width int) int {
	l := 0
	h := 1
	for _, v := range p.Cs {
		if l+runewidth.RuneWidth(v.Ch) > width {
			h++
			l = runewidth.RuneWidth(v.Ch)
		} else {
			l += runewidth.RuneWidth(v.Ch)
		}
	}
	return h
}
