# Terminal Quiz App in GO

A simple, interactive quiz application for the terminal – ideal for learning and practicing with custom question sets. Written in Go, this project is designed to be lightweight, extendable, and fun.

## Features

- Loads questions from a JSON file
- Asks random questions in a loop
- Tracks user performance (correct/wrong answers)
- Repeats incorrectly answered questions to reinforce learning
- Simple terminal UI with clear prompts and feedback

## Installation

```bash
git clone https://github.com/yourusername/terminal-quiz.git
cd terminal-quiz
go build -o quiz
```

## Usage

```bash
git clone https://github.com/yourusername/terminal-quiz.git
cd terminal-quiz
go build -o quiz
./quiz
```
By default, the app looks for a questions.json file in the current directory.

## Example questions.json

```json
{
  "questions": [
    {
      "ask": "What is the capital of France?",
      "answers": {
        "a": "Berlin",
        "b": "Paris",
        "c": "Rome"
      },
      "correct": "b"
    }
  ]
}
```