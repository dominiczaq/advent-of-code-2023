package main

import (
	"reflect"
	"testing"
)

func TestGetAllStringAndNumericNumbersFromString(t *testing.T) {

	got := GetAllStringAndNumericNumbersFromString("jsbdh16snnllpvvgnggfive5nhjpgdzh")
	want := []string{"1", "6", "five", "5"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = GetAllStringAndNumericNumbersFromString("84gnprjhr3eightsixsix")
	want = []string{"8", "4", "3", "eight", "six", "six"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = GetAllStringAndNumericNumbersFromString("two1nine")
	want = []string{"two", "1", "nine"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got5 := ReplaceAllStringNumericNumbers("oneight")
	want5 := "1eight"

	if !reflect.DeepEqual(got5, want5) {
		t.Errorf("got %q, wanted %q", got5, want5)
	}

	got2 := ParseNumberString("two")
	want2 := "2"

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got3 := Part2("input-test.txt")
	want3 := 281

	if !reflect.DeepEqual(got3, want3) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
