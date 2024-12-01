package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mmhand123/advent-of-code-2024/internal/math_custom"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	p1 := partOne(string(input))
	p2 := partTwo(string(input))

	log.Printf("part one: %d\n", p1)
	log.Printf("part two: %d\n", p2)

}

func partOne(input string) int {
	lists := generateLists(input)

	var sum int = 0

	for i, leftNum := range lists.left {
		rightNum := lists.right[i]
		distance := math_custom.Abs(leftNum - rightNum)

		sum += distance

	}

	return sum
}

func partTwo(input string) int {
	lists := generateLists(input)
	similarityMap := make(map[int]int)

	for _, rightNum := range lists.right {
		similarityMap[rightNum] = similarityMap[rightNum] + 1
	}

	var sum int = 0
	for _, leftNum := range lists.left {

		adjustedNum := leftNum * similarityMap[leftNum]

		sum += adjustedNum
	}

	return sum
}

type Lists struct {
	left  []int
	right []int
}

func generateLists(input string) Lists {

	var leftList []int
	var rightList []int
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, "   ")

		if len(parts) != 2 {
			continue
		}

		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)

	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return Lists{left: leftList, right: rightList}
}
