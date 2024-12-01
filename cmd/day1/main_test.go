package main

import (
	"os"
	"testing"
)

func Test_PartOnePartial(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := 11
	actual := partOne(input)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

}

func Test_PartOneReal(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	expected := 3714264
	actual := partOne(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

}

func Test_PartTwoPartial(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := 31
	actual := partTwo(input)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

}
