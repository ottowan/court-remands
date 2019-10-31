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

//Test CheckFineRateProvincial()
func Test_CheckFineRateProvincial_Input_fine_minus1_ShouldBe_False(t *testing.T) {
	input := -1
	expected := false
	actual := CheckFineRateProvincial(input)

	if expected != actual {
		t.Errorf("T24 : expected %s but it got %s", expected, actual)
	}
}
func Test_CheckFineRateProvincial_Input_fine_0_ShouldBe_True(t *testing.T) {
	input := 0
	expected := true
	actual := CheckFineRateProvincial(input)

	if expected != actual {
		t.Errorf("T25 : expected %s but it got %s", expected, actual)
	}
}
func Test_CheckFineRateProvincial_Input_fine_300_ShouldBe_True(t *testing.T) {
	input := 300
	expected := true
	actual := CheckFineRateProvincial(input)

	if expected != actual {
		t.Errorf("T26 : expected %s but it got %s", expected, actual)
	}
}

func Test_CheckFineRateProvincial_Input_fine_600_ShouldBe_True(t *testing.T) {
	input := 600
	expected := true
	actual := CheckFineRateProvincial(input)

	if expected != actual {
		t.Errorf("T27 : expected %s but it got %s", expected, actual)
	}
}
