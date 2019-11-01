package main

import "testing"

func Test_CheckCourtType_Input_1KweangCourt_ShouldBe_True(t *testing.T) {
	input := "kwaeng"
	expected := true
	actual := CheckCourtType(input)

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}

func Test_CheckCourtType_Input_2ProvincialCourt_ShouldBe_True(t *testing.T) {
	input := "provincial"
	expected := true
	actual := CheckCourtType(input)

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}

func Test_CheckCourtType_Input_2ProvincialCourt_ShouldBe_False(t *testing.T) {
	input := ""
	expected := false
	actual := CheckCourtType(input)

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}
