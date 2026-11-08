package textutil

import "testing"

func TestReverse(t *testing.T) {
	if got := Reverse("hello"); got != "olleh" {
		t.Errorf("Reverse(hello) = %q, want olleh", got)
	}
}

func TestShout(t *testing.T) {
	if got := Shout("hi"); got != "HI!" {
		t.Errorf("Shout(hi) = %q, want HI!", got)
	}
}