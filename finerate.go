package main

func CheckFineRateKwaeng(input int) bool {

	if input >= 0 && input <= 60000 {
		return true
	}
	return false
}

func CheckFineRateProvincial(input int) int {

	if input >= 0 && input <= 500 {
		return 1
	} else if input >= 501 {
		return 2
	}
	return 0
}
