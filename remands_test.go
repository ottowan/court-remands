// +build integration

package main

import "testing"

func Test_CheckRemandsKwaeng(t *testing.T) {
	t.Run("Input courtType:kwaeng penaltyRate:12month fineRate:30,000bath remandsTimes:1times remandsDate:5day Shoudbe true", func(t *testing.T) {

		courtType := "kwaeng"
		penaltyRate := 12
		fineRate := 30000
		remandsTimes := 1
		remandsDate := 5

		expected := true
		actual := CheckRemands(courtType, penaltyRate, fineRate, remandsTimes, remandsDate)

		if expected != actual {
			t.Errorf("expected %v but it got %v", expected, actual)
		}
	})

	t.Run("Input courtType:kwaeng penaltyRate:-1month fineRate:30,000bath remandsTimes:1times remandsDate:5day false", func(t *testing.T) {

		courtType := "kwaeng"
		penaltyRate := -1
		fineRate := 3000000
		remandsTimes := 1
		remandsDate := 5

		expected := false
		actual := CheckRemands(courtType, penaltyRate, fineRate, remandsTimes, remandsDate)

		if expected != actual {
			t.Errorf("expected %v but it got %v", expected, actual)
		}
	})
}

func Test_CheckRemandsProvincial(t *testing.T) {
	t.Run("(Less Penality)Input provincial 6month 500bath 1times 7day Shoudbe true", func(t *testing.T) {

		courtType := "provincial"
		penaltyRate := 6
		fineRate := 500
		remandsTimes := 1
		remandsDate := 7

		expected := true
		actual := CheckRemands(courtType, penaltyRate, fineRate, remandsTimes, remandsDate)

		if expected != actual {
			t.Errorf("expected %v but it got %v", expected, actual)
		}
	})

	t.Run("(Middle Penality)Input provincial 120month 501bath 4times 12day Shoudbe true", func(t *testing.T) {

		courtType := "provincial"
		penaltyRate := 120
		fineRate := 501
		remandsTimes := 4
		remandsDate := 12

		expected := true
		actual := CheckRemands(courtType, penaltyRate, fineRate, remandsTimes, remandsDate)

		if expected != actual {
			t.Errorf("expected %v but it got %v", expected, actual)
		}
	})

	t.Run("(Top Penality)Input provincial 121month 501bath 7times 12day Shoudbe true", func(t *testing.T) {

		courtType := "provincial"
		penaltyRate := 121
		fineRate := 501
		remandsTimes := 7
		remandsDate := 12

		expected := true
		actual := CheckRemands(courtType, penaltyRate, fineRate, remandsTimes, remandsDate)

		if expected != actual {
			t.Errorf("expected %v but it got %v", expected, actual)
		}
	})
}
