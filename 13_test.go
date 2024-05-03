package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	if romanToInt("III") != 3 {
		t.Errorf(`TestRomanToInt("III") == 3`)
	}
	if romanToInt("IV") != 4 {
		t.Errorf(`TestRomanToInt("IV") == 4`)
	}
}
func TestNotRomanToInt(t *testing.T) {
	if romanToInt("III") == 4 {
		t.Errorf(`TestRomanToInt("III") == 4`)
	}

}
