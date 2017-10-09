# cgolGo

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d25d558df2b1437c8893a9b23f207e7e)](https://www.codacy.com/app/SimonWaldherr/cgolGo?utm_source=github.com&utm_medium=referral&utm_content=SimonWaldherr/cgolGo&utm_campaign=badger)
[![Travis CI](https://img.shields.io/travis/SimonWaldherr/cgolGo.svg?style=flat)](https://travis-ci.org/SimonWaldherr/cgolGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/simonwaldherr/cgolGo)](https://goreportcard.com/report/github.com/simonwaldherr/cgolGo)
[![codebeat badge](https://codebeat.co/badges/a20ab70f-2baa-490b-8fcf-69ac1961e969)](https://codebeat.co/projects/github-com-simonwaldherr-cgolgo-master)
[![codecov](https://codecov.io/gh/SimonWaldherr/cgolGo/branch/master/graph/badge.svg)](https://codecov.io/gh/SimonWaldherr/cgolGo)
[![Coverage Status](https://coveralls.io/repos/github/SimonWaldherr/cgolGo/badge.svg?branch=master)](https://coveralls.io/github/SimonWaldherr/cgolGo?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/SimonWaldherr/cgolGo/life) 
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/SimonWaldherr/cgolGo/master/LICENSE)  

## conway's game of life in golang

[Conway's Game of Life](http://en.wikipedia.org/wiki/Conway's_Game_of_Life) 
in [golang](http://en.wikipedia.org/wiki/Go_(programming_language))  

Conway's Game of Life is a [zero-player game](https://en.wikipedia.org/wiki/Zero-player_game) - a [cellular automaton](https://en.wikipedia.org/wiki/Cellular_automaton) simulation invented by [John Horton Conway](https://en.wikipedia.org/wiki/John_Horton_Conway). 
There are many implementations in every important programming language [here on GitHub](https://github.com/search?q=topic%3Aconway-game&type=Repositories). 
The map of a Game of Life consists of an two-dimensional grid of square cells. 
Each cell can have one of to two possible states - dead or alive. 
The future of a cell is determined by its own current status and that of the eight direct neighbors - vertically, horizontally and diagonally. 
* a living cell with two or three living neighbours stays alive
* a dead cell with three living neighbours becomes a live cell
* every other cell will be a dead cell in the next round

## examples

01.gif

[![01.gif](http://simonwaldherr.github.io/cgolGo/output/01.gif)](https://github.com/SimonWaldherr/cgolGo/blob/master/structures/01.txt)  

02.gif

[![02.gif](http://simonwaldherr.github.io/cgolGo/output/02.gif)](https://github.com/SimonWaldherr/cgolGo/blob/master/structures/02.txt)  

03.gif

[![03.gif](http://simonwaldherr.github.io/cgolGo/output/03.gif)](https://github.com/SimonWaldherr/cgolGo/blob/master/structures/03.txt)  

15.gif

[![15.gif](http://simonwaldherr.github.io/cgolGo/output/15.gif)](https://github.com/SimonWaldherr/cgolGo/blob/master/structures/15.txt)  

## license

[MIT](https://github.com/SimonWaldherr/cgolGo/blob/master/LICENSE)  
