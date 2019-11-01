package main

func CheckFineRateKwaeng(input int) bool {

	if input > -1 && input < 60001 {
		return true
	}
	return false
}

func CheckFineRateProvincial(input int) int {

	if input > -1 && input <= 500 {
		return 1
	} else if input > 500 {
		return 2
	}
	return 0
}
