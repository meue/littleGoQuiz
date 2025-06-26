package question

import (
	"testing"
)

func TestAddWrongAnswerAddsQuestion(t *testing.T) {
	stats := &Statistics{}
	q := &Question{Ask: "Test"}

	stats.addWrongAnswer(q)

	if len(stats.wrongQuestionsHistory) != 1 {
		t.Errorf("Expected 1 wrong question, got %d", len(stats.wrongQuestionsHistory))
	}

	if q.RetryCount != 3 {
		t.Errorf("Expected RetryCount 3, got %d", q.RetryCount)
	}
}

func TestAddCorrectAnswerRemovesQuestionWhenRetryCountZero(t *testing.T) {
	q := &Question{Ask: "Test", RetryCount: 1}
	stats := &Statistics{
		wrongQuestionsHistory: []*Question{q},
	}

	stats.addCorrectAnswer(q)

	if len(stats.wrongQuestionsHistory) != 0 {
		t.Errorf("Expected question to be removed from history")
	}
}

func TestAddCorrectAnswerDecrementsRetryCount(t *testing.T) {
	q := &Question{Ask: "Test", RetryCount: 2}
	stats := &Statistics{
		wrongQuestionsHistory: []*Question{q},
	}

	stats.addCorrectAnswer(q)

	if q.RetryCount != 1 {
		t.Errorf("Expected RetryCount 1, got %d", q.RetryCount)
	}
}

func TestContainsReturnsTrue(t *testing.T) {
	q := &Question{Ask: "Test"}
	stats := &Statistics{
		wrongQuestionsHistory: []*Question{q},
	}

	if !stats.contains(q) {
		t.Errorf("Expected question to be found in history")
	}
}

func TestRemoveQuestion(t *testing.T) {
	q1 := &Question{Ask: "One"}
	q2 := &Question{Ask: "Two"}
	stats := &Statistics{
		wrongQuestionsHistory: []*Question{q1, q2},
	}

	stats.removeQuestion(q1)

	if len(stats.wrongQuestionsHistory) != 1 {
		t.Errorf("Expected 1 question in history, got %d", len(stats.wrongQuestionsHistory))
	}
	if stats.wrongQuestionsHistory[0] != q2 {
		t.Errorf("Expected remaining question to be q2")
	}
}
