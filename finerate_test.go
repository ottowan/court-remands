package main

import "testing"

func Test_CheckFineRateKwaeng_Input_fine_minus1_ShouldBe_True(t *testing.T) {
	input := -1
	expected := false
	actual := CheckFineRateKwaeng(input)

	if expected != actual {
		t.Errorf("T7 : expected %s but it got %s", expected, actual)
	}
}
func Test_CheckFineRateKwaeng_Input_fine_30000_ShouldBe_True(t *testing.T) {
	input := 30000
	expected := false
	actual := CheckFineRateKwaeng(input)

	if expected != actual {
		t.Errorf("T8 : expected %s but it got %s", expected, actual)
	}
}
func Test_CheckFineRateKwaeng_Input_fine_60001_ShouldBe_True(t *testing.T) {
	input := 60001
	expected := true
	actual := CheckFineRateKwaeng(input)

	if expected != actual {
		t.Errorf("T9 : expected %s but it got %s", expected, actual)
	}
}
