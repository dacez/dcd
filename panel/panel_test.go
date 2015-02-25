package panel

import (
	"dcd/cd"
	"github.com/nsf/termbox-go"
	"testing"
)

func TestInit(t *testing.T) {
	err := termbox.Init()
	if err != nil {
		t.Fatal("termboxInit", err)
	}
	defer termbox.Close()

	var outputPanel Panel
	outputPanel.Init(0, 0, 80, 20, termbox.ColorWhite, termbox.ColorBlack, OutputType, 0, 0)
	var dirs []string
	cd.GetAllDir("E:\\GoProject\\src", &dirs)
	for _, v := range dirs {
		outputPanel.PushLine([]byte(v))
	}

	var inputPanel Panel
	inputPanel.Init(0, 20, 80, 2, termbox.ColorWhite, termbox.ColorBlack, InputType, 0, 0)
	termbox.Flush()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowDown:
				outputPanel.Down()
			case termbox.KeyArrowUp:
				outputPanel.Up()
			case termbox.KeyBackspace:
				inputPanel.Pop()
			default:
				if ev.Ch != 0 {
					inputPanel.Push(ev.Ch)
				}
			}
		}
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		outputPanel.Draw()
		inputPanel.Draw()
		termbox.Flush()
	}
}
