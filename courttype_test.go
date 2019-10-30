package main

import "testing"

func Test_CheckCourtType_Input_1KweangCourt_ShouldBe_True(t *testing.T) {
	input := 1
	expected := true
	actual := CheckCourtType(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_CheckCourtType_Input_2ProvincialCourt_ShouldBe_True(t *testing.T) {
	input := 2
	expected := true
	actual := CheckCourtType(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_CheckCourtType_Input_2ProvincialCourt_ShouldBe_False(t *testing.T) {
	input := 0
	expected := false
	actual := CheckCourtType(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_CheckCourtType_Input_3ProvincialCourt_ShouldBe_False(t *testing.T) {
	input := 3
	expected := false
	actual := CheckCourtType(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}
