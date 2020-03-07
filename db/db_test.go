package db

import "testing"

func TestDbConnection(t *testing.T) {

}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
