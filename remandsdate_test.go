package main
import "testing"

func Test_CheckRemandsDateProvincial_Input_Max10year_0imprisonDate_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := imprison_date(input)

	if expected != actual {
		t.Errorf("T44 :expected %s but it got %s", expected, actual)
	}
}

func Test_CheckRemandsDateProvincial_Input_Max10year_1imprisonDate_ShouldBe_True(t *testing.T) {
	input := 1
	expected := true
	actual := imprison_date(input)

	if expected != actual {
		t.Errorf("T45 :expected %s but it got %s", expected, actual)
	}
}

func Test_CheckRemandsDateProvincial_Input_Max10year_10imprisonDate_ShouldBe_True(t *testing.T) {
	input := 10
	expected := true
	actual := imprison_date(input)

	if expected != actual {
		t.Errorf("T46 :expected %s but it got %s", expected, actual)
	}
}

func Test_CheckRemandsDateProvincial_Input_Max10year_15imprisonDate_ShouldBe_True(t *testing.T) {
	input := 15
	expected := false
	actual := imprison_date(input)

	if expected != actual {
		t.Errorf("T47 :expected %s but it got %s", expected, actual)
	}
}