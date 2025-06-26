package main

import (
	"fmt"
	"quiz/m/src/question"
)

func main() {
	quiz := question.NewQuiz("questions.json")
	for {
		fmt.Print("\033[H\033[2J")
		quiz.AskNextQuestion()
	}
}
