package main

import (
	"dcd/cd"
	"dcd/config"
	"dcd/panel"
	"github.com/nsf/termbox-go"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		cd.PushHis(os.Args[1])
		return
	}
	err := termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()

	home := config.GetConfig().Home

	os.Mkdir(home+"/.dacecd", os.ModeDir|os.ModePerm)

	var inputPanel panel.Panel
	inputPanel.Init(0, 0, w, 2, termbox.ColorWhite, termbox.ColorBlack, panel.InputType, 0, 0)

	var outputPanel panel.Panel
	w, h := termbox.Size()
	outputPanel.Init(0, 2, w, h-2, termbox.ColorWhite, termbox.ColorBlack, panel.OutputType, 0, 0)

	dirs := cd.GetDirs()
	outputPanel.InitBuffers(dirs)

	inputPanel.Draw()
	outputPanel.Draw()
	termbox.Flush()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				ioutil.WriteFile(home+"/.dacecd/command.sh", []byte("cd ."), os.ModePerm)
				break mainloop
			case termbox.KeyArrowDown, termbox.KeyCtrlJ, termbox.KeyCtrlN:
				outputPanel.Down()
			case termbox.KeyCtrlD:
				if config.GetConfig().M == config.HisMode {
					config.GetConfig().M = config.GlobalMode
				} else {
					config.GetConfig().M = config.HisMode
				}
				dirs = cd.GetDirs()
				outputPanel.InitBuffers(dirs)

				for _, v := range inputPanel.GetSelectLine().Cs {
					outputPanel.FilterPush(string(v.Ch))
				}
			case termbox.KeyArrowUp, termbox.KeyCtrlK, termbox.KeyCtrlP:
				outputPanel.Up()
			case termbox.KeyBackspace:
				inputPanel.Pop()
				outputPanel.InitFilter()
				for _, v := range inputPanel.GetSelectLine().Cs {
					outputPanel.FilterPush(string(v.Ch))
				}
			case termbox.KeyEnter:
				sl := outputPanel.GetSelectLine()
				godir := sl.GetString()
				ioutil.WriteFile(home+"/.dacecd/command.sh", []byte("cd "+godir), os.ModePerm)
				cd.PushHis(godir)
				break mainloop
			default:
				if ev.Ch != 0 {
					inputPanel.Push(ev.Ch)
					outputPanel.FilterPush(string(ev.Ch))
				}
			}
		}
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		outputPanel.Draw()
		inputPanel.Draw()
		termbox.Flush()
	}
}
