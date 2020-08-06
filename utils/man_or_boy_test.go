package utils

import "testing"

func TestManOrBoy(t *testing.T) {
	want := -67
	if got := ManOrBoy(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
