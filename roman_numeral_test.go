package romannumeral

import "testing"

func TestRomanNumeral(t *testing.T) {
	testcases := []struct {
		Number  int
		Roman   string
		IsError bool
	}{
		{1, "I", false},
		{2, "II", false},
		{3, "III", false},
		{4, "IV", false},
		{5, "V", false},
		{6, "VI", false},
		{9, "IX", false},
		{27, "XXVII", false},
		{48, "XLVIII", false},
		{59, "LIX", false},
		{93, "XCIII", false},
		{141, "CXLI", false},
		{163, "CLXIII", false},
		{402, "CDII", false},
		{575, "DLXXV", false},
		{911, "CMXI", false},
		{1024, "MXXIV", false},
		{3000, "MMM", false},
		{0, "", true},
		{-1, "", true},
		{3001, "", true},
	}

	for _, test := range testcases {
		val, err := ToRomanNumeral(test.Number)
		if err != nil {
			if !test.IsError {
				t.Fail()
				t.Errorf("Input %d Unexpected Error : ", test.Number)
			}
			continue
		}
		if test.IsError {
			t.Errorf("Input %d Expected Error : ", test.Number)
			continue
		}
		if val != test.Roman {
			t.Fail()
			t.Errorf("Input %d: Got %s, Expected %s", test.Number, val, test.Roman)
		}
	}

}
