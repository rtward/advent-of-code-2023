package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getNumAt(input string, index int) (string, bool) {
	subStr := input[index:]

	// Handle the digit case
	chars := strings.Split(subStr, "")
	numberChars := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if slices.Contains(numberChars, chars[0]) {
		return chars[0], true
	}

	// Handle the word case
	numWords := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for word, num := range numWords {
		if strings.HasPrefix(subStr, word) {
			return num, true
		}
	}

	return "", false
}

func getNum(input string) int {
	fmt.Println(input)

	nums := make([]string, 0)
	for i := 0; i < len(input); i++ {
		num, found := getNumAt(input, i)
		if found {
			nums = append(nums, num)
		}
	}

	fmt.Println(strings.Join(nums, " "))

	first := nums[0]
	last := nums[len(nums)-1]
	numStr := strings.Join([]string{first, last}, "")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numStr)
	return num
}

func Run() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		num := getNum(scanner.Text())
		sum += num
	}

	fmt.Printf("Sum: %d\n", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
