package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)

var dic = []string{
	"Zombie",
	"Gopher",
	"United States of America",
	"Indonisia",
	"Nazism",
	"Apple",
	"Programming",
}

func main() {
	targetWord := getRandomName()
	fmt.Println(targetWord)

	guessedLetters := initGuessedWords(targetWord)
	hangmanState := 0

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid. please input one letter only...")
			continue
		}

		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}
	
	printGameState(targetWord, guessedLetters, hangmanState)

	fmt.Print("Game Over... ")
	if isWordGussed(targetWord, guessedLetters) {
		fmt.Println("You Won!")
	} else if isHangmanCompleted(hangmanState) {
		fmt.Println("You Lose!")
	} else {
		panic("Invalid state. Game is over and there is no winner!")
	}
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGussed(targetWord, guessedLetters) || isHangmanCompleted(hangmanState)
}

func isHangmanCompleted(hangmanState int) bool {
	return hangmanState >= 9
}

func isWordGussed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if ch != ' ' && !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}

func readInput() string {
	fmt.Print("> ")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getHangmanDrawing(hangmanState int) string {
	data, err := os.ReadFile(fmt.Sprintf("states/h%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func initGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[rune(targetWord[len(targetWord)-1])] = true

	return guessedLetters
}

func getRandomName() string {
	return dic[rand.Intn(6)]
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	res := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			res += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
			res += fmt.Sprintf("%c", ch)
		} else {
			res += "_"
		}

		res += " "
	}

	return res
}
