#!/bin/bash
unalias cd
alias dcd='YOURHOMEPATH/.dacecd/dcd'
if [ $# == 0 ];then
    cd
else
    cd $1
fi
dcd `pwd`
unalias dcd
alias cd='source YOURHOMEPATH/.dacecd/dcd.sh'
