package main
import "testing"

func Test_CheckRemandsTimesProvincial_Input_Max10year_0imprisonNo_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T40 :expected %s but it got %s", expected, actual)
	}
}
func Test_CheckRemandsTimesProvincial_Input_Max10year_1imprisonNo_ShouldBe_True(t *testing.T) {
	input := 1
	expected := true
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T41 :expected %s but it got %s", expected, actual)
	}
}
func Test_CheckRemandsTimesProvincial_Input_Max10year_5imprisonNo_ShouldBe_True(t *testing.T) {
	input := 5
	expected := true
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T42 :expected %s but it got %s", expected, actual)
	}
}
func Test_CheckRemandsTimesProvincial_Input_Max10year_8imprisonNo_ShouldBe_True(t *testing.T) {
	input := 8
	expected := false
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T43 :expected %s but it got %s", expected, actual)
	}
}