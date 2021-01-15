package main

import (
	"flag"
)

var (
	path = flag.String("file", "./Quiz/problems.csv", "Path to csv file containing tasks")
	//testTime = flag.Int("testTime", 20, "Set time in seconds allowed for the game")
)

func main() {
	var game quiz
	game.loadTasks(*path)
	game.playWithTimer()
	game.gameResults()
}
