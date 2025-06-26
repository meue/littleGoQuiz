package question

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

type Quiz struct {
	data            *QuizData
	currentQuestion *Question
	statistics      *Statistics
}

func NewQuiz(filename string) *Quiz {
	rand.Seed(time.Now().UnixNano())
	quiz := &Quiz{}
	quiz.data = loadQuiz(filename)
	quiz.statistics = &Statistics{}
	return quiz
}

func loadQuiz(filename string) *QuizData {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not find %s: %v", filename, err)

	}

	var quizdata *QuizData
	err = json.Unmarshal(data, &quizdata)
	if err != nil {
		log.Fatalf("Could not read %s: %v", filename, err)

	}

	return quizdata
}

func (q *Quiz) AskNextQuestion() {
	q.statistics.PrintRatio()
	q.currentQuestion = q.getRandomQuestion()
	fmt.Println(q.currentQuestion.Ask)
	/*for i, answer := range q.currentQuestion.Answers {
		fmt.Println(fmt.Sprintf("%s: %s", i, answer))
	}*/

	keys := make([]string, 0, len(q.currentQuestion.Answers))
	for k := range q.currentQuestion.Answers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %s\n", k, q.currentQuestion.Answers[k])
	}

	answer := q.readUserInput("Answer: ")
	if answer == q.currentQuestion.Correct {
		fmt.Println("")
		fmt.Println("Correct!")
		fmt.Println("")
		q.statistics.addCorrectAnswer(q.currentQuestion)
	} else {
		fmt.Println("")
		fmt.Println("Wrong Answer!")
		fmt.Println(fmt.Sprintf("The Correct answer is %s, %s.", q.currentQuestion.Correct, q.currentQuestion.Answers[q.currentQuestion.Correct]))
		fmt.Println("")
		q.statistics.addWrongAnswer(q.currentQuestion)

	}

	_ = q.readUserInput("Press Enter for the next question")
}

func (q *Quiz) getRandomQuestion() *Question {
	if q.statistics.ShouldIUseSameQuestionAgain() {
		return q.statistics.GetRandomWrongQuestion()
	}

	// take random question
	idx := rand.Intn(len(q.data.Questions))
	return &q.data.Questions[idx]
}

func (q *Quiz) readUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(input)
}
