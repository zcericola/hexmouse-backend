package server

import "testing"

func TestServer(t *testing.T) {
	got := ""
	want := "ok"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
