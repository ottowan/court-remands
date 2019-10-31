package main

import "testing"

func Test_CheckPenalityRateKwaeng_Input_punishyear_1month_ShouldBe_True(t *testing.T) {
	input := -1
	expected := false
	actual := CheckPenalityRateKwaeng(input)

	if expected != actual {
		t.Errorf("T4 :expected %s but it got %s", expected, actual)
	}
}
func Test_CheckPenalityRateKwaeng_Input_punishyear_3year1month_ShouldBe_True(t *testing.T) {
	input := 37
	expected := false
	actual := CheckPenalityRateKwaeng(input)

	if expected != actual {
		t.Errorf("T5 : expected %s but it got %s", expected, actual)
	}
}
func Test_CheckPenalityRateKwaeng_Input_punishyear_1year_ShouldBe_True(t *testing.T) {
	input := 12
	expected := true
	actual := CheckPenalityRateKwaeng(input)

	if expected != actual {
		t.Errorf("T6 :expected %s but it got %s", expected, actual)
	}
}
