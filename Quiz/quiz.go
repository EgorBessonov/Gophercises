package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

type task struct {
	question string
	answer   string
}

type quiz struct {
	tasks    []task
	answered int
	score    int
}

func (quizGame *quiz) loadTasks(path string) {
	csvFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {

			break
		} else if err != nil {
			panic(err)
		}
		question := task{line[0], line[1]}
		quizGame.tasks = append(quizGame.tasks, question)
	}
	return
}

func (quizGame *quiz) play() {
	var answer string
	for _, t := range quizGame.tasks {
		fmt.Printf("What %v, sir? ", t.question)
		fmt.Fscan(os.Stdin, &answer)
		if answer == t.answer {
			quizGame.score++
		}
		quizGame.answered++
	}
	return
}

func (quizGame *quiz) playWithTimer() {
	timer := time.NewTimer(time.Second * 10)
	answer := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)


	func(){
		for _, t := range quizGame.tasks {
			go func() {
				fmt.Printf("What %v, sir?", t.question)
				scanner.Scan()
				ans := scanner.Text()
				answer <- ans
			}()
			select {
			case <-timer.C:
				return
			case ans := <-answer:
				if ans == t.answer {
					quizGame.score++
				}
				quizGame.answered++
			}
		}
	}()

	return
}

func (quizGame *quiz) gameResults() {
	fmt.Printf("You answered %v questions and got %v correct",
		quizGame.answered,
		quizGame.score)
}
