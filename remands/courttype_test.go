package remands

import "testing"

func Test_CheckCourtType(t *testing.T) {

	t.Run("input kweang ShouldBe 1", func(t *testing.T) {
		input := "kwaeng"
		expected := 1
		actual := CheckCourtType(input)

		if expected != actual {
			t.Errorf("T1 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input provincial ShouldBe 2", func(t *testing.T) {
		input := "provincial"
		expected := 2
		actual := CheckCourtType(input)

		if expected != actual {
			t.Errorf("T2 : expected %v but it got %v", expected, actual)
		}
	})

	t.Run("input empty ShouldBe 0", func(t *testing.T) {
		input := ""
		expected := 0
		actual := CheckCourtType(input)

		if expected != actual {
			t.Errorf("T3 : expected %v but it got %v", expected, actual)
		}
	})
}
