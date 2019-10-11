/*
Quiz game on terminal in Go.
Task 1 : Read quiz questions from a csv file and display them one by one. Take the answer and print how many in the end were correct /incorrect. Questions can themselves have commas in them and that should be handled.
Task 2 : Add a universal timer which starts ticking once the user presses say enter. Then quiz questions should only be displayed until the timer goes to that amount. Still, number of right and wrong answes should be printed. 
Program should exit after timeout even if scanf is waiting on user input for a certain question. 
*/

package main 

import (
		"fmt"
		"os"
		"log"
		"io"
		"bufio"
		"encoding/csv"
		"strconv"
		"time"
)

// A set of one question and its answer.  
type QuizItem struct {
	question string 
	answer string
}

func main() {
	fmt.Println("Welcome to akshata's quiz game!")
	correct := 0
	incorrect := 0
	useranswer := 0
	
	//STEP 1 : Read in the file and store contents in a list of type QuizItem struct 
	
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("failed to open csv\n", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var quiz []QuizItem
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal("failed to read csv\n", error)
		}		
		quiz = append(quiz, QuizItem{
			question: line[0],
			answer: line[1],
		})
	}
	
	//STEP 2 : Keep a timer (time.After) which will essentially create a channel that will receive input after specified number of seconds. Iterate quiz questions in a range loop. To read user input in a non-blocking way, spawn a routine for it. Then add a select in the main routine waiting for input from 2 channels - either the user input channel or the timeup channel. The moment timeup input comes, code exists because the go routine accepting input from scanf is separate. Amazing go routines :)
	
	timeup := time.After(10 * time.Second)
	ch := make(chan int)
	
	for _, quizitem := range quiz {
		fmt.Println(quizitem.question)
		go func(ch chan int) {
			fmt.Scanf("%d\n", &useranswer)
			ch <- useranswer
		}(ch)
		
		select {
			case <-timeup :
				//number of incorrects is everything that is incorrect/unanswered. Basically everthing that is not "correct".
				
				fmt.Printf("Oops! time up :( \nYou got %d correct ,and %d incorrect\n", correct, len(quiz) - correct)
				return
			case uanswer := <- ch :
				qanswer, err := strconv.Atoi(quizitem.answer)
				if err != nil {
					log.Fatal("failed to Atoi\n", err)
				} 
				if uanswer == qanswer {
					correct++
				} else {
					incorrect++
				}
		}
	}
	fmt.Printf("You got %d correct ,and %d incorrect out of %d", correct, incorrect, len(quiz))
}
