package pokeapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "test")
	}))
	defer ts.Close()

	res, err := get(ts.URL)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	if string(res) != "test" {
		t.Errorf("response did not match. expected: test, got: %s", err)
	}
}
