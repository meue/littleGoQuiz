package question

import (
	"os"
	"testing"
)

func TestLoadQuizReturnsData(t *testing.T) {
	const fileName = "test_questions.json"
	content := `{
		"questions": [
			{
				"ask": "2+2?",
				"answers": {
					"a": "3",
					"b": "4"
				},
				"correct": "b"
			}
		]
	}`

	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	data := loadQuiz(fileName)

	if len(data.Questions) != 1 {
		t.Errorf("Expected 1 question, got %d", len(data.Questions))
	}
	if data.Questions[0].Correct != "b" {
		t.Errorf("Expected correct answer to be 'b'")
	}
}
