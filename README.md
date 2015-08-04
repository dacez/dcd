# Dacecd
## 中文介绍
    代替cd的一个工具，平时经常遇到在多个文件夹之前切换的问题，每次cd都要打很长的命令，费时费力。
    该工具采用fuzzyfind的方法解决该问题，让你闪电般打开想要的文件夹。
    两种模式：历史模式和搜索模式
    历史模式记录你打开过的文件夹，运行cd命令就会被记录，然后fuzzyfind。
    搜索模式打开所有文件夹，然后fuzzyfind，如果文件夹太多，加载速度较慢，fuzzyfind还是很快。
    默认模式为历史模式。
    
## Introduction
    like cd command
    but it can save the cd history and fuzzy find dirs
    cdl to play
    Ctrl-x to hash all directories
    Ctrl-d to change fuzzy find mode

## How To Use
    Type cdl to launch
    Ctrl - x to switch println all files or history files
    Ctrl - d to switch fuzzy find mode path or name
    Esc      to exit
    Enter    to go to the select directory
    Ctrl - j down the select line
    Ctrl - n down the select line
    Down     down the select line
    Ctrl - k up the select line
    Ctrl - p up the select line
    Up       up the select line
<img src="./dacecd.gif" width="800">

    
## How To Install （Linux 64bit Only）
    Download the execute file dcd (I compile the 64bit for linux only)
    Download the dcd.sh 
    mkdir .dacecd in your homepath then move dcd.sh and dcd in it
    vim $Home/.dacecd/.dcd.sh
    
        #!/bin/bash
        unalias cd
        alias dcd='$Home/.dacecd/dcd'
        if [ $# == 0 ];then
           cd
        else
            cd $1
        fi
        dcd `pwd`
        unalias dcd
        alias cd='source $Home/.dacecd/dcd.sh'
    
## Source Install （32bit or 64bit）
    go get -u github.com/dacez/dcd
    
## How To Config

###First:
    vim $Home/.dacecd/.dacecdrc

        {
          "ContainDirs": [
            "~/QQMail/micromsg",
            "~/QQMail/mmcomm"
            ],
            "HisCount":100
        }

    Modify the ContainDirs that you need to fuzzy. The .dacecdrc must be a json.
###Second:
    vim .profile or .bashrc
    Add two lines below:
    
        alias cd='source $Home/.dacecd/dcd.sh'
        alias cdl='$Home/.dacecd/dcd;source $Home/.dacecd/command.sh'
