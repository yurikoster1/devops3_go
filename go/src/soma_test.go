package main

import "testing"

func TestSoma(t *testing.T) {
	cases := []struct {
		number1, number2, result int
	}{
		{1, 2, 3},
		{5, 5 , 10},
		{40, 50, 90},
		{50,50, 100},
	}
	for _, c := range cases {
		got := Soma(c.number1, c.number2)
		if got != c.result {
			t.Errorf("Soma(%q, %q) == %q, want %q", c.number1,c.number2, got, c.result)
		}
	}
}