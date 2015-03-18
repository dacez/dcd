# Dacecd
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

## Notice
    Can't use ~ instead of $Home
    
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

    Replace $Home to your real homepath.
    
## How To Config

###First:
    vim $Home/.dacecd/.dacecdrc

        {
          "ContainDirs": [
            "/home/qspace_system/QQMail/micromsg",
            "/home/qspace_system/QQMail/mmcomm"
            ],
            "HisCount":100
        }

    Modify the ContainDirs that you need to fuzzy. The .dacecdrc must be a json.
###Second:
    vim .profile or .bashrc
    Add two lines below:
    
        alias cd='source $Home/.dacecd/dcd.sh'
        alias cdl='$Home/.dacecd/dcd;source $Home/.dacecd/command.sh'
        
    Replace $Home to your real homepath and reconnnect your terminal.
