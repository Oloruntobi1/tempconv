package tconv

import "testing"

func TestCelsiusToKelvin(t *testing.T){

	const value = 273.15

	temp := 50.0 // this allows the go to infer this as a floating point number 

	got := CelsiusToKelvin(temp)

	want := 50 + value

	if got != want {
		t.Errorf("Error. got %.2f, wanted %.2f", got, want)
	}
}