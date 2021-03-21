package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		tasks    []task
		answered int
		score    int

		path     = flag.String("file", "./Quiz/problems.csv", "Path to csv file containing tasks")
		testTime = flag.Duration("testTime", 25, "Set time in seconds allowed for the game")
	)

	err := loadTasks(*path, &tasks)
	if err != nil {
		fmt.Println("Sorry, we have problems with loading tasks")
		return
	}
	playWithTimer(testTime, &tasks, &score, &answered)
	fmt.Printf("\nYou answered %v question(s) and got %v correct from %v questions", answered, score, len(tasks))
}
