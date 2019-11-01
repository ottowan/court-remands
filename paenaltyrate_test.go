package main

import "testing"

func Test_CheckPenalityRateKwaeng_Input_punishyear_1month_ShouldBe_True(t *testing.T) {
	input := -1
	expected := false
	actual := CheckPenalityRateKwaeng(input)

	if expected != actual {
		t.Errorf("T4 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckPenalityRateKwaeng_Input_punishyear_3year1month_ShouldBe_True(t *testing.T) {
	input := 37
	expected := false
	actual := CheckPenalityRateKwaeng(input)

	if expected != actual {
		t.Errorf("T5 : expected %v but it got %v", expected, actual)
	}
}
func Test_CheckPenalityRateKwaeng_Input_punishyear_1year_ShouldBe_True(t *testing.T) {
	input := 12
	expected := true
	actual := CheckPenalityRateKwaeng(input)

	if expected != actual {
		t.Errorf("T6 :expected %v but it got %v", expected, actual)
	}
}

//CheckPenalityRateProvincial
func Test_CheckPenalityRateProvincial_Input_punishyear_1month_ShouldBe_True(t *testing.T) {
	input := -1
	expected := false
	actual := CheckPenalityRateProvincial(input)

	if expected != actual {
		t.Errorf("T16 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckPenalityRateProvincial_Input_punishyear_7month_ShouldBe_True(t *testing.T) {
	input := 7
	expected := false
	actual := CheckPenalityRateProvincial(input)

	if expected != actual {
		t.Errorf("T17 :expected %v but it got %v", expected, actual)
	}
}

func Test_CheckPenalityRateProvincial_Input_punishyear_3month_ShouldBe_True(t *testing.T) {
	input := 3
	expected := true
	actual := CheckPenalityRateProvincial(input)

	if expected != actual {
		t.Errorf("T18 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckPenalityRateProvincial_Input_punishyear_5month_ShouldBe_True(t *testing.T) {
	input := 5
	expected := false
	actual := CheckPenalityRateProvincial(input)

	if expected != actual {
		t.Errorf("T19 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckPenalityRateProvincial_Input_punishyear_10year1month_ShouldBe_True(t *testing.T) {
	input := 120
	expected := false
	actual := CheckPenalityRateProvincial(input)

	if expected != actual {
		t.Errorf("T20 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckPenalityRateProvincial_Input_punishyear_5year_ShouldBe_True(t *testing.T) {
	input := 60
	expected := true
	actual := CheckPenalityRateProvincial(input)

	if expected != actual {
		t.Errorf("T21 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckPenalityRateProvincial_Input_punishyear_9year_ShouldBe_True(t *testing.T) {
	input := 108
	expected := false
	actual := CheckPenalityRateProvincial(input)

	if expected != actual {
		t.Errorf("T22 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckPenalityRateProvincial_Input_punishyear_11year_ShouldBe_True(t *testing.T) {
	input := 110
	expected := true
	actual := CheckPenalityRateProvincial(input)

	if expected != actual {
		t.Errorf("T23 :expected %v but it got %v", expected, actual)
	}
}
