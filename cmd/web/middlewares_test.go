package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	// Dummy request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Fake handler
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, req)

	resp := rr.Result()

	frameOptions := resp.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("Want %q; got %q", "deny", frameOptions)
	}

	xssProtection := resp.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("Want %q; got %q", "1; mode=block", xssProtection)
	}

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