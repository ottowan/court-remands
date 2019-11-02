package remands

func CheckCourtType(input string) int {
	if input == "kwaeng" {
		//fmt.Println("Input : " + input)
		return 1
	} else if input == "provincial" {

		//fmt.Println("Input : " + input)
		return 2
	}

	//fmt.Println("Input : fail")
	return 0
}
