#!/bin/bash
unalias cd
alias dcd='$HOME/.dacecd/dcd'
if [ $# == 0 ];then
    cd
else
    cd $1
fi
dcd `pwd`
unalias dcd
alias cd='source $HOME/.dacecd/dcd.sh'
