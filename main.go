package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)

var dictionary = []string{
	"calendar",
	"lucky",
	"degree",
	"bulldog",
	"Vietnam",
	"wrestling",
	"Pogramming",
	"butcher",
	"Mobile",
	"peach",
}

var isHintConsumed = false

func main() {

	targetWord := getRandomWord()

	guessedLetters := initializeGuessedWords(targetWord)

	fmt.Println("* Press # and hit enter to consume your free hint")

	hangmanState := 0

	for !isGameOver(targetWord, guessedLetters, hangmanState) {

		printGameState(targetWord, guessedLetters, hangmanState)

		input := readInput()

		if len(input) != 1 {
			fmt.Println("Invalid Input. Please use letters only...")
			continue
		}

		letter := rune(input[0])

		if letter == '#' && !isHintConsumed {
			fmt.Println(getRandomUnGuessedWord(targetWord, getWordGuessingProgress(targetWord, guessedLetters), guessedLetters))
			isHintConsumed = true
		}

		if guessedLetters[letter] {
			fmt.Println("You have already used this letter.")
			continue
		}

		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}

	}

	fmt.Print("Game Over...")

	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You loose!")
	} else {
		panic("Invalid State. Game is over and there is no winner")
	}
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}

	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return dictionary[rand.Intn(len(dictionary))]
}

func getRandomUnGuessedWord(targetWord string, WordBeingGuessed string, guessedLetters map[rune]bool) string {
	for index, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] && rune(WordBeingGuessed[index]) == '_' {
			return fmt.Sprintf("%c", ch)
		}
	}

	return fmt.Sprintf("%c", ' ')
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	gameState := getWordGuessingProgress(targetWord, guessedLetters)
	fmt.Print(gameState)
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {

	result := ""

	for _, ch := range targetWord {

		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "
	}

	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))

	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	fmt.Print("> ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
