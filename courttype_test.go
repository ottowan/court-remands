package main

import "testing"

func Test_CheckCourtType(t *testing.T) {

	t.Run("input kweang ShouldBe true", func(t *testing.T) {
		input := "kwaeng"
		expected := true
		actual := CheckCourtType(input)

		if expected != actual {
			t.Errorf("T1 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input provincial ShouldBe true", func(t *testing.T) {
		input := "provincial"
		expected := true
		actual := CheckCourtType(input)

		if expected != actual {
			t.Errorf("T2 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input empty ShouldBe false", func(t *testing.T) {
		input := ""
		expected := false
		actual := CheckCourtType(input)

		if expected != actual {
			t.Errorf("T3 : expected %v but it got %v", expected, actual)
		}
	})
}
