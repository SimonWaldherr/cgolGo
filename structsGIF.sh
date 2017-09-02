#!/bin/bash

for i in `seq -w 1 20` ; do 
  go run cgol.go -w 50 -h 50 -d 250 -o structures/$i.txt -g $i -l 200;
  gifsicle --resize-width 250 -O --careful -d 5 -o $i.gif $i.gif;
done