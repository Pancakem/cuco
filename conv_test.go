package main

import (
	"testing"
)

func TestParseCurrencyCode(t *testing.T) {
	tst := "KSH5000"

	code := ParseCurrencyCode(tst)

	if code != "KSH" {
		t.Errorf("result incorrect, got: %s, want %s", code, "KSH")
	}

}

func TestIsSupported(t *testing.T) {
	shouldPass := []string{"KSH", "NgN", "ghs", "ksh"}
	shouldFail := []string{"kes", "ush", "tsh"}
	for _, v := range shouldPass {
		if val := IsSupported(v); !val {
			t.Errorf("result incorrect, got: %v, want: %v", val, true)
		}
	}

	for _, v := range shouldFail {
		if val := IsSupported(v); val {
			t.Errorf("result incorrect, got: %v, want: %v", val, false)
		}
	}

	

}

func TestConvert(t *testing.T) {
	res := Convert(50, "ksh", "ghs")
	if !relativelyEqual(res, 2.75) {
		t.Errorf("result incorrect, got: %v, want: %v", res, 2.75)
	}

	res = Convert(50, "ghs", "ksh")
	if !relativelyEqual(res, 908.47) {
		t.Errorf("result incorrect, got: %v, want: %v", res, 908.47)
	}

	res = Convert(50, "ksh", "ngn")
	if !relativelyEqual(res, 190.20) {
		t.Errorf("result incorrect, got: %v, want: %v", res, 190.20)
	}

	res = Convert(50, "ngn", "ksh")
	if !relativelyEqual(res, 13.18) {
		t.Errorf("result incorrect, got: %v, want: %v", res, 13.18)
	}
	
	res = Convert(50, "ngn", "ghs")
	if !relativelyEqual(res, 0.72) {
		t.Errorf("result incorrect, got: %v, want: %v", res, 0.72)
	}

	res = Convert(50, "ghs", "ngn")
	if !relativelyEqual(res, 3455.67) {
		t.Errorf("result incorrect, got: %v, want: %v", res, 3455.67)
	}
}
