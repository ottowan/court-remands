package main

import "testing"

func Test_CheckPenalityRateKwaeng(t *testing.T) {

	t.Run("input -1 month ShouldBe false", func(t *testing.T) {
		input := -1
		expected := false
		actual := CheckPenalityRateKwaeng(input)

		if expected != actual {
			t.Errorf("T4 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 37 month (3y1m) ShouldBe false", func(t *testing.T) {
		input := 37
		expected := false
		actual := CheckPenalityRateKwaeng(input)

		if expected != actual {
			t.Errorf("T5 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 12 month (1y) ShouldBe false", func(t *testing.T) {
		input := 12
		expected := true
		actual := CheckPenalityRateKwaeng(input)

		if expected != actual {
			t.Errorf("T6 : expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckPenalityRateProvincial(t *testing.T) {

	t.Run("input -1 month ShouldBe 0", func(t *testing.T) {
		input := -1
		expected := 0
		actual := CheckPenalityRateProvincial(input)

		if expected != actual {
			t.Errorf("T16 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 0 month ShouldBe 1", func(t *testing.T) {
		input := 0
		expected := 1
		actual := CheckPenalityRateProvincial(input)

		if expected != actual {
			t.Errorf("T17 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 3 month bath ShouldBe 1", func(t *testing.T) {
		input := 3
		expected := 1
		actual := CheckPenalityRateProvincial(input)

		if expected != actual {
			t.Errorf("T18 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 6 month ShouldBe 1", func(t *testing.T) {
		input := 6
		expected := 1
		actual := CheckPenalityRateProvincial(input)

		if expected != actual {
			t.Errorf("T19 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 7 month ShouldBe 2", func(t *testing.T) {
		input := 7
		expected := 2
		actual := CheckPenalityRateProvincial(input)

		if expected != actual {
			t.Errorf("T20 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 60 month (5y) ShouldBe 2", func(t *testing.T) {
		input := 60
		expected := 2
		actual := CheckPenalityRateProvincial(input)

		if expected != actual {
			t.Errorf("T21 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 120 month (10y) ShouldBe 2", func(t *testing.T) {
		input := 120
		expected := 2
		actual := CheckPenalityRateProvincial(input)

		if expected != actual {
			t.Errorf("T22 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 121 month (10y1m) ShouldBe 3", func(t *testing.T) {
		input := 121
		expected := 3
		actual := CheckPenalityRateProvincial(input)

		if expected != actual {
			t.Errorf("T22 : expected %v but it got %v", expected, actual)
		}
	})
}
