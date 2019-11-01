package main

import "testing"

func Test_CheckRemandsDateKwaeng(t *testing.T) {

	t.Run("input 0 day ShouldBe false", func(t *testing.T) {
		input := 0
		expected := false
		actual := CheckRemandsDateKwaeng(input)

		if expected != actual {
			t.Errorf("T13 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 5 day ShouldBe true", func(t *testing.T) {
		input := 5
		expected := true
		actual := CheckRemandsDateKwaeng(input)

		if expected != actual {
			t.Errorf("T15 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 7 day ShouldBe false", func(t *testing.T) {
		input := 7
		expected := false
		actual := CheckRemandsDateKwaeng(input)

		if expected != actual {
			t.Errorf("T14 : expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckRemandsDateLessRateProvincial(t *testing.T) {

	t.Run("input 0 day ShouldBe false", func(t *testing.T) {
		input := 0
		expected := false
		actual := CheckRemandsDateLessRateProvincial(input)

		if expected != actual {
			t.Errorf("T28 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 3 day ShouldBe true", func(t *testing.T) {
		input := 3
		expected := true
		actual := CheckRemandsDateLessRateProvincial(input)

		if expected != actual {
			t.Errorf("T29 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 8 day ShouldBe false", func(t *testing.T) {
		input := 8
		expected := false
		actual := CheckRemandsDateLessRateProvincial(input)

		if expected != actual {
			t.Errorf("T30 : expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckRemandsDateMiddleRateProvincial(t *testing.T) {

	t.Run("input 0 day ShouldBe false", func(t *testing.T) {
		input := 0
		expected := false
		actual := CheckRemandsDateMiddleRateProvincial(input)

		if expected != actual {
			t.Errorf("T36 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 8 day ShouldBe true", func(t *testing.T) {
		input := 8
		expected := true
		actual := CheckRemandsDateMiddleRateProvincial(input)

		if expected != actual {
			t.Errorf("T38 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 13 day ShouldBe false", func(t *testing.T) {
		input := 13
		expected := false
		actual := CheckRemandsDateMiddleRateProvincial(input)

		if expected != actual {
			t.Errorf("T39 : expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckRemandsDateTopRateProvincial(t *testing.T) {

	t.Run("input 0 day ShouldBe false", func(t *testing.T) {
		input := 0
		expected := false
		actual := CheckRemandsDateTopRateProvincial(input)

		if expected != actual {
			t.Errorf("T44 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 8 day ShouldBe true", func(t *testing.T) {
		input := 8
		expected := true
		actual := CheckRemandsDateTopRateProvincial(input)

		if expected != actual {
			t.Errorf("T46 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 13 day ShouldBe false", func(t *testing.T) {
		input := 13
		expected := false
		actual := CheckRemandsDateTopRateProvincial(input)

		if expected != actual {
			t.Errorf("T47 : expected %v but it got %v", expected, actual)
		}
	})
}
