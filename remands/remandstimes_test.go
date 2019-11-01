package remands

import "testing"

func Test_CheckRemandsTimesKwaeng(t *testing.T) {

	t.Run("input 0 times ShouldBe false", func(t *testing.T) {
		input := 0
		expected := false
		actual := CheckRemandsTimesKwaeng(input)

		if expected != actual {
			t.Errorf("T13 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 2 times ShouldBe true", func(t *testing.T) {
		input := 2
		expected := true
		actual := CheckRemandsTimesKwaeng(input)

		if expected != actual {
			t.Errorf("T15 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 6 times ShouldBe false", func(t *testing.T) {
		input := 6
		expected := false
		actual := CheckRemandsTimesKwaeng(input)

		if expected != actual {
			t.Errorf("T14 : expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckRemandsTimesLessRatesProvincial(t *testing.T) {

	t.Run("input 0 times ShouldBe false", func(t *testing.T) {
		input := 0
		expected := false
		actual := CheckRemandsTimesLessRatesProvincial(input)

		if expected != actual {
			t.Errorf("T28 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 1 times ShouldBe true", func(t *testing.T) {
		input := 1
		expected := true
		actual := CheckRemandsTimesLessRatesProvincial(input)

		if expected != actual {
			t.Errorf("T29 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 2 times ShouldBe false", func(t *testing.T) {
		input := 2
		expected := false
		actual := CheckRemandsTimesLessRatesProvincial(input)

		if expected != actual {
			t.Errorf("T30 : expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckRemandsTimesMiddleRateProvincial(t *testing.T) {

	t.Run("input 0 times ShouldBe false", func(t *testing.T) {
		input := 0
		expected := false
		actual := CheckRemandsTimesMiddleRateProvincial(input)

		if expected != actual {
			t.Errorf("T36 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 2 times ShouldBe true", func(t *testing.T) {
		input := 2
		expected := true
		actual := CheckRemandsTimesMiddleRateProvincial(input)

		if expected != actual {
			t.Errorf("T38 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 5 times ShouldBe false", func(t *testing.T) {
		input := 5
		expected := false
		actual := CheckRemandsTimesMiddleRateProvincial(input)

		if expected != actual {
			t.Errorf("T39 : expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckRemandsTimesTopRateProvincial(t *testing.T) {

	t.Run("input 0 times ShouldBe false", func(t *testing.T) {
		input := 0
		expected := false
		actual := CheckRemandsTimesTopRateProvincial(input)

		if expected != actual {
			t.Errorf("T44 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 4 times ShouldBe true", func(t *testing.T) {
		input := 4
		expected := true
		actual := CheckRemandsTimesTopRateProvincial(input)

		if expected != actual {
			t.Errorf("T46 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 8 times ShouldBe false", func(t *testing.T) {
		input := 8
		expected := false
		actual := CheckRemandsTimesTopRateProvincial(input)

		if expected != actual {
			t.Errorf("T47 : expected %v but it got %v", expected, actual)
		}
	})
}
