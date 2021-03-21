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

func loadTasks(path string, tasks *[]task) (err error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return (err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {

			break
		} else if err != nil {
			return (err)
		}
		question := task{line[0], line[1]}
		*tasks = append(*tasks, question)
	}
	return
}

func play(tasks *[]task, score, answered *int) {
	var answer string
	for _, t := range *tasks {
		fmt.Printf("What %s, sir? ", t.question)
		fmt.Fscan(os.Stdin, &answer)
		if answer == t.answer {
			(*score)++
		}
		(*answered)++
	}
	return
}

/*func playWithTimer(testTime *time.Duration, tasks *[]task, score, answered *int) {
	timer := time.After(*testTime)
	answers := make(chan bool)
	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		defer close(answers)

		for _, t := range *tasks {
			select {
			case <-timer:
				return
			default:
				fmt.Printf("What %s, sir?", t.question)
				scanner.Scan()
				ans := scanner.Text()
				answers <- ans == t.answer
			}
		}
	}()

	for a := range answers {
		(*answered)++
		if a {
			(*score)++
		}
	}
	fmt.Println("end")
	return
}*/

func playWithTimer(testTime *time.Duration,tasks *[]task, score, answered *int) {
	timer := time.NewTimer(time.Second * (*testTime))
	answer := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)

	func() {
		for _, t := range *tasks {
			go func() {
				fmt.Printf("What %s, sir?", t.question)
				scanner.Scan()
				ans := scanner.Text()
				answer <- ans
			}()
			select {
			case <-timer.C:
				return
			case ans := <-answer:
				if ans == t.answer {
					(*score)++
				}
				(*answered)++
			}
		}
	}()

	return
}

