package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	// Dummy request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, req)
	
	resp := rr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Want %q; got %q", http.StatusOK, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if string(body) != "OK" {
		t.Errorf("Want %q; got %q", "OK", string(body))
	}
}