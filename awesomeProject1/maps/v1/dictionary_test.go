package v1

import "testing"

func TestSearch(t *testing.T) {

	dictionary := map[string]string{"test": "ovo je test"}
	got := Search(dictionary, "test")
	want := "ovo je test"

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
