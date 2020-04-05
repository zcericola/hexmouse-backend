package users

import "testing"

//TestUserCreation tests that a user is created properly
func TestUserCreation(t *testing.T) {
	got := "{}"
	want := "{User}"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
