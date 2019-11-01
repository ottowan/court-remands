package main

import "testing"

//CheckRemandsDateKwaeng
func Test_CheckRemandsDateKwaeng_Input_0imprisonDate_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := CheckRemandsDateKwaeng(input)

	if expected != actual {
		t.Errorf("T13 :expected %v but it got %v", expected, actual)
	}
}

func Test_CheckRemandsDateKwaeng_Input_7imprisonDate_ShouldBe_True(t *testing.T) {
	input := 7
	expected := false
	actual := CheckRemandsDateKwaeng(input)

	if expected != actual {
		t.Errorf("T14 :expected %v but it got %v", expected, actual)
	}
}

func Test_CheckRemandsDateKwaeng_Input_5imprisonDate_ShouldBe_True(t *testing.T) {
	input := 5
	expected := false
	actual := CheckRemandsDateKwaeng(input)

	if expected != actual {
		t.Errorf("T15 :expected %v but it got %v", expected, actual)
	}
}

//CheckRemandsDateProvincial_6month
func Test_CheckRemandsDateProvincial_Input_6month_0imprisonDate_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T28 :expected %v but it got %v", expected, actual)
	}
}

func Test_CheckRemandsDateProvincial_Input_6month_3imprisonDate_ShouldBe_True(t *testing.T) {
	input := 3
	expected := true
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T29 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckRemandsDateProvincial_Input_6month_8imprisonDate_ShouldBe_True(t *testing.T) {
	input := 8
	expected := false
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T30 :expected %v but it got %v", expected, actual)
	}
}

//CheckRemandsDateProvincial_10year
func Test_CheckRemandsDateProvincial_Input_10year_0imprisonDate_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T36 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckRemandsDateProvincial_Input_10year_1imprisonDate_ShouldBe_True(t *testing.T) {
	input := 1
	expected := true
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T37 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckRemandsDateProvincial_Input_10year_11imprisonDate_ShouldBe_True(t *testing.T) {
	input := 11
	expected := true
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T38 :expected %v but it got %v", expected, actual)
	}
}
func Test_CheckRemandsDateProvincial_Input_10year_15imprisonDate_ShouldBe_True(t *testing.T) {
	input := 15
	expected := false
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T39 :expected %v but it got %v", expected, actual)
	}
}

//CheckRemandsDateProvincial_Max10year
func Test_CheckRemandsDateProvincial_Input_Max10year_0imprisonDate_ShouldBe_True(t *testing.T) {
	input := 0
	expected := false
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T44 :expected %v but it got %v", expected, actual)
	}
}

func Test_CheckRemandsDateProvincial_Input_Max10year_1imprisonDate_ShouldBe_True(t *testing.T) {
	input := 1
	expected := true
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T45 :expected %v but it got %v", expected, actual)
	}
}

func Test_CheckRemandsDateProvincial_Input_Max10year_10imprisonDate_ShouldBe_True(t *testing.T) {
	input := 10
	expected := true
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T46 :expected %v but it got %v", expected, actual)
	}
}

func Test_CheckRemandsDateProvincial_Input_Max10year_15imprisonDate_ShouldBe_True(t *testing.T) {
	input := 15
	expected := false
	actual := CheckRemandsDateProvincial(input)

	if expected != actual {
		t.Errorf("T47 :expected %v but it got %v", expected, actual)
	}
}
