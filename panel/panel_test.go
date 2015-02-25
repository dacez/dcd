package panel

import (
	"github.com/nsf/termbox-go"
	"testing"
)

func TestInit(t *testing.T) {
	err := termbox.Init()
	if err != nil {
		t.Fatal("termboxInit", err)
	}
	defer termbox.Close()

	var p Panel
	p.Init(0, 0, 10, 10, termbox.ColorWhite, termbox.ColorBlack, OutputType, 0, 0)
	p.PushLine([]byte("abcdefghijklmn"))
	p.PushLine([]byte("1234"))
	p.PushLine([]byte("1234"))
	p.PushLine([]byte("1234"))
	p.PushLine([]byte("1234"))
	p.PushLine([]byte("1234"))
	p.PushLine([]byte("1234"))
	p.PushLine([]byte("1234"))
	p.PushLine([]byte("5678"))
	p.PushLine([]byte("10jkkkkkkkk"))
	p.PushLine([]byte("20j"))
	p.PushLine([]byte("30jkkkkkkkk"))
	p.PushLine([]byte("5678"))
	p.Draw()
	var p1 Panel
	p1.Init(10, 0, 10, 10, termbox.ColorWhite, termbox.ColorBlack, InputType, 0, 0)
	p1.Draw()

	termbox.Flush()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowDown:
				p.Down()
			case termbox.KeyArrowUp:
				p.Up()
			case termbox.KeyBackspace:
				p1.Pop()
			default:
				if ev.Ch != 0 {
					p1.Push(ev.Ch)
				}
			}
		}
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		p.Draw()
		p1.Draw()
		termbox.Flush()
	}
}
