package snippets

import "testing"

//TestSnippetCreation tests that a snippet is created properly
func TestSnippetCreation(t *testing.T) {
	got := "{}"
	want := "{User}"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
