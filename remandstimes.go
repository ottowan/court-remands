package main

func CheckRemandsTimesKwaeng(input int) bool {
	if input >= 1 && input <= 5 {
		return true
	}
	return false
}

func CheckRemandsTimesLessRatesProvincial(input int) bool {
	if input == 1 {
		return true
	}
	return false
}

func CheckRemandsTimesMiddleRateProvincial(input int) bool {
	if input >= 1 && input <= 4 {
		return true
	}
	return false
}

func CheckRemandsTimesTopRateProvincial(input int) bool {
	if input >= 1 && input <= 7 {
		return true
	}
	return false
}
