package question

import (
	"fmt"
	"math/rand"
)

type Statistics struct {
	wrongQuestionsHistory []*Question
	answerHistory         [30]bool
	index                 int
	justRepeatedQuestion  bool
}

func (s *Statistics) addWrongAnswer(q *Question) {
	s.answerHistory[s.index] = false
	s.index++
	s.index %= len(s.answerHistory)

	q.RetryCount += 3
	if !s.contains(q) {
		fmt.Println("Dont worry, I will ask you again.")
		s.wrongQuestionsHistory = append(s.wrongQuestionsHistory, q)
	}
}

func (s *Statistics) addCorrectAnswer(q *Question) {
	s.answerHistory[s.index] = true
	s.index++
	s.index %= len(s.answerHistory)
	if q.RetryCount > 0 && s.contains(q) {
		q.RetryCount -= 1
		if q.RetryCount <= 0 {
			s.removeQuestion(q)
		}
	}
}

func (s *Statistics) Debug() {
	fmt.Println(len(s.wrongQuestionsHistory))
}

func (s *Statistics) ShouldIUseSameQuestionAgain() bool {
	if s.justRepeatedQuestion {
		s.justRepeatedQuestion = false
		return false
	}
	s.justRepeatedQuestion = true
	return rand.Intn(5) < len(s.wrongQuestionsHistory)
}

func (s *Statistics) GetRandomWrongQuestion() *Question {
	idx := rand.Intn(len(s.wrongQuestionsHistory))
	return s.wrongQuestionsHistory[idx]
}

func (s *Statistics) contains(q *Question) bool {
	for _, item := range s.wrongQuestionsHistory {
		if item == q { // Pointer-Vergleich: gleiche Adresse?
			return true
		}
	}
	return false
}

func (s *Statistics) removeQuestion(q *Question) {
	for i, question := range s.wrongQuestionsHistory {
		if question == q {
			s.wrongQuestionsHistory = append(s.wrongQuestionsHistory[:i], s.wrongQuestionsHistory[i+1:]...)
			return
		}
	}

}

func (s *Statistics) PrintRatio() {
	length := len(s.answerHistory)
	count := 0
	for i := 0; i < length; i++ {
		if s.answerHistory[i] {
			count++
		}
	}

	fmt.Printf("Ratio: %.0f%%\n", float64(count)/float64(length)*100)

}
