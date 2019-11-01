package remandsintegration

import (
	"court-remands/remands"
	"fmt"
)

func CheckRemands(courtType string, penaltyRate int, fineRate int, remandsTimes int, remandsDate int) bool {

	//Validate kwaeng court
	if remands.CheckCourtType(courtType) == 1 {
		if remands.CheckPenalityRateKwaeng(penaltyRate) && remands.CheckFineRateKwaeng(fineRate) {

			if remands.CheckRemandsTimesKwaeng(remandsTimes) {
				if remands.CheckRemandsDateKwaeng(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	} else if remands.CheckCourtType(courtType) == 2 {

		if remands.CheckPenalityRateProvincial(penaltyRate) == 1 && remands.CheckFineRateProvincial(fineRate) == 1 {
			if remands.CheckRemandsTimesLessRatesProvincial(remandsTimes) {
				if remands.CheckRemandsDateLessRateProvincial(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else if remands.CheckPenalityRateProvincial(penaltyRate) == 2 && remands.CheckFineRateProvincial(fineRate) == 2 {
			if remands.CheckRemandsTimesMiddleRateProvincial(remandsTimes) {
				if remands.CheckRemandsDateMiddleRateProvincial(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else if remands.CheckPenalityRateProvincial(fineRate) == 3 {
			if remands.CheckRemandsTimesTopRateProvincial(remandsTimes) {
				if remands.CheckRemandsDateTopRateProvincial(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		}

		return false
	}

	return false
}
func CheckRemandsAPI(courtType int, penaltyRate int, fineRate int, remandsTimes int, remandsDate int) bool {

	fmt.Println(courtType)
	fmt.Println(penaltyRate)
	fmt.Println(fineRate)
	fmt.Println(remandsTimes)
	fmt.Println(remandsDate)

	//Validate kwaeng court
	//fmt.Println("CALL : " + strconv.Itoa(remands.CheckCourtType(courtType)))
	if courtType == 1 {
		if remands.CheckPenalityRateKwaeng(penaltyRate) && remands.CheckFineRateKwaeng(fineRate) {

			if remands.CheckRemandsTimesKwaeng(remandsTimes) {
				if remands.CheckRemandsDateKwaeng(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	} else if courtType == 2 {

		if remands.CheckPenalityRateProvincial(penaltyRate) == 1 && remands.CheckFineRateProvincial(fineRate) == 1 {
			if remands.CheckRemandsTimesLessRatesProvincial(remandsTimes) {
				if remands.CheckRemandsDateLessRateProvincial(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else if remands.CheckPenalityRateProvincial(penaltyRate) == 2 && remands.CheckFineRateProvincial(fineRate) == 2 {
			if remands.CheckRemandsTimesMiddleRateProvincial(remandsTimes) {
				if remands.CheckRemandsDateMiddleRateProvincial(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else if remands.CheckPenalityRateProvincial(fineRate) == 3 {
			if remands.CheckRemandsTimesTopRateProvincial(remandsTimes) {
				if remands.CheckRemandsDateTopRateProvincial(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		}

		return false
	}

	return false
}
