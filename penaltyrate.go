package main

func CheckPenalityRateKwaeng(input int) bool {

	if input >= 0 && input < 37 {
		return true
	}
	return false
}

func CheckPenalityRateProvincial(input int) int {

	if input >= 0 && input <= 6 {
		return 1
	} else if input >= 7 && input <= 120 {
		return 2
	} else if input >= 121 {
		return 3
	}
	return 0
}
