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

	w, h := termbox.Size()

	var inputPanel panel.Panel
	inputPanel.Init(0, 0, w, 2, termbox.ColorWhite, termbox.ColorBlack, panel.InputType, 0, 0)

	var outputPanel panel.Panel
	outputPanel.Init(0, 2, w, h-4, termbox.ColorWhite, termbox.ColorBlack, panel.OutputType, 0, 0)
	dirs := cd.GetDirs()
	outputPanel.InitBuffers(dirs)

	var statePanel panel.Panel
	statePanel.Init(0, h-2, w, 2, termbox.ColorWhite, termbox.ColorBlack, panel.OutputType, 0, 0)
	statePanel.InitBuffer(config.GetStateLine())

	inputPanel.Draw()
	outputPanel.Draw()
	statePanel.Draw()
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
			case termbox.KeyCtrlX:
				if config.GetConfig().DirectoryMode == config.HisMode {
					config.GetConfig().DirectoryMode = config.AllMode
				} else if config.GetConfig().DirectoryMode == config.AllMode {
					config.GetConfig().DirectoryMode = config.HisMode
				}
				dirs = cd.GetDirs()
				outputPanel.InitBuffers(dirs)
				for _, v := range inputPanel.GetSelectLine().Cs {
					outputPanel.FilterPush(string(v.Ch))
				}
			case termbox.KeyCtrlD:
				if config.GetConfig().FuzzyFindMode == config.NameMode {
					config.GetConfig().FuzzyFindMode = config.PathMode
				} else if config.GetConfig().FuzzyFindMode == config.PathMode {
					config.GetConfig().FuzzyFindMode = config.NameMode
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
		inputPanel.Draw()
		outputPanel.Draw()
		statePanel.InitBuffer(config.GetStateLine())
		statePanel.Draw()
		termbox.Flush()
	}
}
