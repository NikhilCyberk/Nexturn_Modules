/*
Exercise 4: Online Examination System

Topics Covered: Go Loops, Go Break and Continue, Go Constants, Go Strings, Go Functions, Go Errors

Case Study:

Develop an online examination system where users can take a quiz.

1. Question Bank: Define a slice of structs to store questions. Each question should have a question string, options (array), and the correct answer.

2. Take Quiz: Use loops to iterate over questions and display them one by one. Allow the user to select an answer by entering the option number.

	o Use continue to skip invalid inputs and prompt the user again.
	o Use break to exit the quiz early if the user enters a specific command (e.g., "exit").

3. Score Calculation: After the quiz, calculate the user's score and display it. Use conditions to classify performance (e.g., "Excellent", "Good", "Needs Improvement").

4. Error Handling: Handle errors like invalid input during the quiz (e.g., entering a non-integer value for an option).

Bonus:
• Add a timer for the quiz, limiting each question to a fixed amount of time.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Constants for quiz configuration
const (
	TimePerQuestion = 30 * time.Second
	ExitCommand     = "exit"
)

type Question struct {
	Question string
	Options  []string
	Answer   string
}

// validateAnswer checks if the given answer is valid
func validateAnswer(answer string, options []string) error {
	answer = strings.TrimSpace(strings.ToLower(answer))
	if answer == ExitCommand {
		return nil
	}

	// Check if answer exists in options (case-insensitive)
	for _, opt := range options {
		if strings.ToLower(opt) == answer {
			return nil
		}
	}
	return fmt.Errorf("invalid answer: must be one of the given options")
}

// displayQuestion shows the question and its options
func displayQuestion(q Question, qNum int) {
	fmt.Printf("\nQuestion %d:\n", qNum)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(q.Question)
	fmt.Println(strings.Repeat("-", 30))
	
	for i, opt := range q.Options {
		fmt.Printf("%c) %s\n", rune('A'+i), opt)
	}
	fmt.Println(strings.Repeat("-", 30))
}

// calculatePerformance determines the performance level based on score
func calculatePerformance(score, totalQuestions int) (string, string) {
	percentage := float64(score) / float64(totalQuestions) * 100
	
	var performance, emoji string
	switch {
	case percentage >= 80:
		performance = "Excellent"
		emoji = "🌟"
	case percentage >= 60:
		performance = "Good"
		emoji = "👍"
	default:
		performance = "Needs Improvement"
		emoji = "📚"
	}
	
	return performance, emoji
}

// readAnswer reads and processes user input
func readAnswer(reader *bufio.Reader) (string, error) {
	fmt.Print("Your answer (or 'exit' to quit): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}
	
	return strings.TrimSpace(input), nil
}

func main() {
	questions := []Question{
		{
			Question: "What is the capital of France?",
			Options:  []string{"Paris", "London", "Berlin", "Madrid"},
			Answer:   "Paris",
		},
		{
			Question: "Which programming language is this quiz written in?",
			Options:  []string{"Go", "Python", "Java", "C++"},
			Answer:   "Go",
		},
		{
			Question: "What is the result of 2 + 2?",
			Options:  []string{"3", "4", "5", "6"},
			Answer:   "4",
		},
	}

	fmt.Println("\n🎓 Welcome to the Online Examination System!")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("• You have %v seconds per question\n", TimePerQuestion/time.Second)
	fmt.Printf("• Enter '%s' to quit the quiz early\n", ExitCommand)
	fmt.Println("• Answer by typing the full option (case-insensitive)")
	fmt.Println(strings.Repeat("=", 50))

	reader := bufio.NewReader(os.Stdin)
	score := 0
	totalQuestions := len(questions)

	// Start the quiz
	for i, question := range questions {
		displayQuestion(question, i+1)
		
		// Create timer for the question
		timer := time.NewTimer(TimePerQuestion)
		answerChan := make(chan string)
		errChan := make(chan error)

		// Goroutine to handle user input
		go func() {
			answer, err := readAnswer(reader)
			if err != nil {
				errChan <- err
				return
			}
			answerChan <- answer
		}()

		// Wait for either answer or timeout
		select {
		case <-timer.C:
			fmt.Println("\n⏰ Time's up for this question!")
			continue

		case err := <-errChan:
			fmt.Printf("\n❌ Error: %v\n", err)
			continue

		case answer := <-answerChan:
			timer.Stop()
			
			if answer == ExitCommand {
				fmt.Println("\n👋 Exiting quiz early...")
				goto QuizEnd
			}

			if err := validateAnswer(answer, question.Options); err != nil {
				fmt.Printf("\n❌ %v\n", err)
				continue
			}

			if strings.EqualFold(answer, question.Answer) {
				score++
				fmt.Println("\n✅ Correct!")
			} else {
				fmt.Printf("\n❌ Incorrect. The correct answer was: %s\n", question.Answer)
			}
		}
	}

QuizEnd:
	fmt.Println(strings.Repeat("=", 50))
	performance, emoji := calculatePerformance(score, totalQuestions)
	fmt.Printf("\nFinal Score: %d/%d %s\n", score, totalQuestions, emoji)
	fmt.Printf("Performance: %s\n", performance)
	fmt.Println(strings.Repeat("=", 50))
}