package main

func CheckRemandsDateKwaeng(input int) bool {
	if input >= 1 && input <= 6 {
		return true
	}
	return false
}

func CheckRemandsDateLessRateProvincial(input int) bool {
	if input >= 1 && input <= 7 {
		return true
	}
	return false
}

func CheckRemandsDateMiddleRateProvincial(input int) bool {
	if input >= 1 && input <= 12 {
		return true
	}
	return false
}

func CheckRemandsDateTopRateProvincial(input int) bool {
	if input >= 1 && input <= 12 {
		return true
	}
	return false
}
