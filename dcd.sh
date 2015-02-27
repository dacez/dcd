#!/bin/bash
#alias cd='source ~/.dacecd/dcd.sh'
unalias cd
alias dcd='~/.dacecd/dcd'
if [ $# == 0 ];then
    dcd;source ~/.dacecd/command.sh
else
    cd $1
    dcd `pwd`
fi
alias cd='source ~/.dacecd/dcd.sh'