package main

import "testing"

func TestRound(t *testing.T) {
	cases := []struct{
		i float64
		j int
	}{
		{
			1.15,
			1,
		},
		{
			10.3,
			10,
		},
		{
			-5.5,
			-6,
		},
	}

	for _, tCase := range cases {
		if tCase.j != round(tCase.i) {
			t.Errorf("wrong output, expected: %d, got: %d", tCase.j, round(tCase.i))
		}
	}
}

func TestToFixed(t *testing.T) {
	cases := []struct{
		i float64
		j float64
		p int
	}{
		{
			10.55,
			10.6,
			1,
		},
		{
			12.33,
			12.3,
			1,
		},
		{
			10.516,
			10.52,
			2,
		},
	}
	for _, tCase := range cases {
		if tCase.j != toFixed(tCase.i, tCase.p) {
			t.Errorf("wrong output, expected: %f, got: %d", tCase.j, round(tCase.i))
		}
	}
}

func TestRandomFloat(t *testing.T) {
	cases := []struct{
		min float64
		max float64
	}{
		{
			10.1,
			10.3,
		},
		{
			-10.1,
			-10.3,
		},
	}

	for _, tCase := range cases {
		number := randomFloat(tCase.min, tCase.max)
		t.Log(number)
		if number <= 0 {
			if number > tCase.min || number < tCase.max {
				t.Errorf("number out of range: %f, min: %f, max: %f", number, tCase.min, tCase.max)
			}
		} else {
			if number < tCase.min || number > tCase.max {
				t.Errorf("number out of range: %f, min: %f, max: %f", number, tCase.min, tCase.max)
			}
		}
	}
}