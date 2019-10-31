package main

import "testing"

func Test_CheckRemandsTimesKwaeng_Input_0imprisonNo_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := CheckRemandsTimesKwaeng(input)

	if expected != actual {
		t.Errorf("T10 :expected %s but it got %s", expected, actual)
	}
}

func Test_CheckRemandsTimesKwaeng_Input_6imprisonNo_ShouldBe_True(t *testing.T) {
	input := 6
	expected := false
	actual := CheckRemandsTimesKwaeng(input)

	if expected != actual {
		t.Errorf("T11 :expected %s but it got %s", expected, actual)
	}
}

func Test_CheckRemandsTimesKwaeng_Input_3imprisonNo_ShouldBe_True(t *testing.T) {
	input := 3
	expected := true
	actual := CheckRemandsTimesKwaeng(input)

	if expected != actual {
		t.Errorf("T12 :expected %s but it got %s", expected, actual)
	}
}

//CheckRemandsTimesProvincial_6month
func Test_CheckRemandsTimesProvincial_Input_6month_0imprisonNo_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T31 :expected %s but it got %s", expected, actual)
	}
}
func Test_CheckRemandsTimesProvincial_Input_6month_1imprisonNo_ShouldBe_True(t *testing.T) {
	input := 1
	expected := true
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T32 :expected %s but it got %s", expected, actual)
	}
}
func Test_CheckRemandsTimesProvincial_Input_Max10year_2imprisonNo_ShouldBe_True(t *testing.T) {
	input := 2
	expected := false
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T33 :expected %s but it got %s", expected, actual)
	}
}

//CheckRemandsTimesProvincial_10year
func Test_CheckRemandsTimesProvincial_Input_10year_0imprisonNo_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T34 :expected %s but it got %s", expected, actual)
	}
}
func Test_CheckRemandsTimesProvincial_Input_10year_1imprisonNo_ShouldBe_True(t *testing.T) {
	input := 1
	expected := true
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T35 :expected %s but it got %s", expected, actual)
	}
}
func Test_CheckRemandsTimesProvincial_Input_10year_5imprisonNo_ShouldBe_True(t *testing.T) {
	input := 5
	expected := false
	actual := CheckRemandsTimesProvincial(input)

	if expected != actual {
		t.Errorf("T36 :expected %s but it got %s", expected, actual)
	}
}

//CheckRemandsTimesProvincial_Max10year
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
