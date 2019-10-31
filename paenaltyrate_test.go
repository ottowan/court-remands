
package main

import "testing"
func Test_CheckFineRateKwaeng_Input_punishyear_1month_ShouldBe_True(t *testing.T) {
	input := -1
	expected := false
	actual := Punish(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}
func Test_CheckFineRateKwaeng_Input_punishyear_3year1month_ShouldBe_True(t *testing.T) {
	input := 3.1
	expected := false
	actual := Punish(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}
func Test_CheckFineRateKwaeng_Input_punishyear_1year_ShouldBe_True(t *testing.T) {
	input := 1
	expected := True
	actual := Punish(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}