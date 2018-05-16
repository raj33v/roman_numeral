package romannumerals

import (
	"errors"
	"math"
	"strconv"
)

var romanMap = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	20:   "XX",
	30:   "XXX",
	40:   "XL",
	50:   "L",
	60:   "LX",
	70:   "LXX",
	80:   "LXXX",
	90:   "XC",
	100:  "C",
	200:  "CC",
	300:  "CCC",
	400:  "CD",
	500:  "D",
	600:  "DC",
	700:  "DCC",
	800:  "DCCC",
	900:  "CM",
	1000: "M",
	2000: "MM",
	3000: "MMM",
}

// ToRomanNumeral convert decimal number to roman number
func ToRomanNumeral(n int) (string, error) {
	if n > 3000 || n <= 0 {
		return "", errors.New("Invalid number")
	}
	str := ""
	s := strconv.Itoa(n)
	l := len(s)
	m := int(math.Pow(10.0, float64(l-1)))
	for i := 0; i < l; i++ {
		if s[i] != '0' {
			x := int(s[i]-'0') * m
			str += romanMap[x]
		}
		m /= 10
	}
	return str, nil
}
