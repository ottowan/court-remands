package main

func CheckRemands(courtType string, penaltyRate int, fineRate int, remandsTimes int, remandsDate int) bool {

	//Validate kwaeng court
	if CheckCourtType(courtType) == 1 {
		if CheckPenalityRateKwaeng(penaltyRate) && CheckFineRateKwaeng(fineRate) {

			if CheckRemandsTimesKwaeng(remandsTimes) {
				if CheckRemandsDateKwaeng(remandsDate) {
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
	} else if CheckCourtType(courtType) == 2 {

		if CheckPenalityRateProvincial(penaltyRate) == 1 && CheckFineRateProvincial(fineRate) == 1 {
			if CheckRemandsTimesLessRatesProvincial(remandsTimes) {
				if CheckRemandsDateLessRateProvincial(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else if CheckPenalityRateProvincial(penaltyRate) == 2 && CheckFineRateProvincial(fineRate) == 2 {
			if CheckRemandsTimesMiddleRateProvincial(remandsTimes) {
				if CheckRemandsDateMiddleRateProvincial(remandsDate) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else if CheckPenalityRateProvincial(fineRate) == 3 {
			if CheckRemandsTimesTopRateProvincial(remandsTimes) {
				if CheckRemandsDateTopRateProvincial(remandsDate) {
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
