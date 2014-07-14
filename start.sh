#!/bin/sh

echo "INSERT A FRAMERATE AND CONFIRM WITH \"ENTER\" [15]"
read -s -t 30 FPS

echo "PRESS \"R\" FOR A RANDOM MAP OR \"F\" TO LOAD FROM A FILE [R]"
read -n 1 -s -t 30 INPUT

if [ "$FPS" == "" ]
  then
  FPS=15
fi

if [ "$INPUT" == "" ]
  then
  INPUT="R"
fi

if [ "$INPUT" == "R" ] || [ "$INPUT" == "r" ]
  then
  go run cgol.go $(tput cols) $(tput lines) 0 $FPS
elif [ "$INPUT" == "F" ] || [ "$INPUT" == "f" ]
  then
  echo "ENTER THE NAME OF A FILE"
  ls "structures"
  read -n 2 -s INPUT
  go run cgol.go $(tput cols) $(tput lines) 0 $FPS "structures/$INPUT.txt"
fi
