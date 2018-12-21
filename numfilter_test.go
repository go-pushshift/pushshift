package pushshift

import (
	"testing"
)

func TestNumFilter_Matches(t *testing.T) {
	// Direct match tests
	if !NumFilter(42).Matches(42) {
		t.Error("Failed to match positive constant")
	}
	if !NumFilter(-42).Matches(-42) {
		t.Error("Failed to match negative constant")
	}
	if NumFilter(42).Matches(420) {
		t.Error("Sanity check")
	}

	if GreaterThan(42).Matches(42) {
		t.Error("Positive gt off-by-one")
	}
	if !GreaterThan(42).Matches(420) {
		t.Errorf("Positive gt failed")
	}
	if GreaterThan(-42).Matches(-42) {
		t.Error("Negative gt off-by-one")
	}
	if !GreaterThan(-42).Matches(-24) {
		t.Errorf("Negative gt failed")
	}

	if LessThan(42).Matches(42) {
		t.Error("Positive lt off-by-one")
	}
	if !LessThan(42).Matches(24) {
		t.Errorf("Positive lt failed")
	}
	if LessThan(-42).Matches(-42) {
		t.Error("Negative lt off-by-one")
	}
	if !LessThan(-42).Matches(-420) {
		t.Errorf("Negative lt failed")
	}
}

func TestNumFilter_String(t *testing.T) {
	var tests = []struct {
		f   NumFilter
		exp string
	}{
		{NumFilter(42), "42"},
		{NumFilter(-42), "-42"},
		{GreaterThan(42), ">42"},
		{LessThan(42), "<42"},
		{GreaterThan(-42), ">-42"},
		{LessThan(-42), "<-42"},
	}

	for i, test := range tests {
		if test.f.String() != test.exp {
			t.Errorf("#%d: expected %s got %s",
				i+1, test.exp, test.f.String())
		}
	}
}
