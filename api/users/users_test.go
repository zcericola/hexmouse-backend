package users

import "testing"

func TestUserCreation(t *testing.T) {
	got := "{}"
	want := "{User}"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
