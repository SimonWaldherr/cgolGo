#!/bin/bash

echo "INSERT A FRAMERATE AND CONFIRM WITH \"ENTER\" [15]"
read -s -t 30 FPS

echo "INSERT AN AMOUNT OF FRAMES AND CONFIRM WITH \"ENTER\" [200]"
read -s -t 30 FRAMES

echo "INSERT A WIDTH AND CONFIRM WITH \"ENTER\" [250]"
read -s -t 30 WIDTH

echo "INSERT A HEIGHT AND CONFIRM WITH \"ENTER\" [250]"
read -s -t 30 HEIGHT

echo "PRESS \"R\" FOR A RANDOM MAP OR \"F\" TO LOAD FROM A FILE [R]"
read -n 1 -s -t 30 INPUT



if [ "$FPS" == "" ]
  then
  FPS=15
fi

if [ "$FRAMES" == "" ]
  then
  FRAMES=200
fi

if [ "$WIDTH" == "" ]
  then
  WIDTH=250
fi

if [ "$HEIGHT" == "" ]
  then
  HEIGHT=250
fi

if [ "$INPUT" == "" ]
  then
  INPUT="R"
fi

if [ "$INPUT" == "R" ] || [ "$INPUT" == "r" ]
  then
  go run cgol.go -w $((WIDTH/4)) -h $((HEIGHT/4)) -d $FRAMES -g cgol
elif [ "$INPUT" == "F" ] || [ "$INPUT" == "f" ]
  then
  echo "ENTER THE NAME OF A FILE"
  ls "structures"
  read -n 2 -s INPUT
  go run cgol.go -w $((WIDTH/4)) -h $((HEIGHT/4)) -d $FRAMES -g cgol -o "structures/$INPUT.txt"
fi

gifsicle --resize-width $WIDTH -O --careful -d $((100/FPS)) -o cgol.gif cgol.gif
