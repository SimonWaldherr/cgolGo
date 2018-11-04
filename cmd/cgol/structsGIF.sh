#!/bin/bash

i=0

for filename in ../../structures/*; do
  ((i++))
  go run main.go -w 50 -h 50 -d 250 -o "$filename" -g $i -l 200;
  gifsicle --change-color "#000000" "#FFFFFF" --change-color "#FFFFFF" "#000000" --resize-width 250 -O --careful -d 5 -o $i.gif $i.gif;
done