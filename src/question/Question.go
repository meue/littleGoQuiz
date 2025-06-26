package question

type Question struct {
	Ask        string            `json:"ask"`
	Answers    map[string]string `json:"answers"`
	Correct    string            `json:"correct"`
	RetryCount int
}
