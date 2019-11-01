package main

func CheckCourtType(input string) int {
	if input == "kwaeng" {
		return 1
	} else if input == "provincial" {

		return 2
	}
	return 0
}
