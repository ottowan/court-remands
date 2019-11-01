package remands

import "testing"

func Test_CheckFineRateKwaeng(t *testing.T) {

	t.Run("input -1 bath ShouldBe false", func(t *testing.T) {
		input := -1
		expected := false
		actual := CheckFineRateKwaeng(input)

		if expected != actual {
			t.Errorf("T7 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 30000 bath ShouldBe true", func(t *testing.T) {
		input := 30000
		expected := true
		actual := CheckFineRateKwaeng(input)

		if expected != actual {
			t.Errorf("T8 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 60001 bath ShouldBe false", func(t *testing.T) {
		input := 60001
		expected := false
		actual := CheckFineRateKwaeng(input)

		if expected != actual {
			t.Errorf("T9 : expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckFineRateProvincial(t *testing.T) {

	t.Run("input -1 bath ShouldBe 0", func(t *testing.T) {
		input := -1
		expected := 0
		actual := CheckFineRateProvincial(input)

		if expected != actual {
			t.Errorf("T24 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 300 bath ShouldBe 1", func(t *testing.T) {
		input := 300
		expected := 1
		actual := CheckFineRateProvincial(input)

		if expected != actual {
			t.Errorf("T26 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input 501 bath ShouldBe 2", func(t *testing.T) {
		input := 501
		expected := 2
		actual := CheckFineRateProvincial(input)

		if expected != actual {
			t.Errorf("T27 : expected %v but it got %v", expected, actual)
		}
	})
}
