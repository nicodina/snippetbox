package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	
	app := newTestApplication(t)

	server := newTestServer(t, app.routes())
	defer server.Close()

	status, _, body := server.get(t, "/ping")
	
	if status != http.StatusOK {
		t.Errorf("Want %q; got %q", http.StatusOK, status)
	}

	if string(body) != "OK" {
		t.Errorf("Want %q; got %q", "OK", string(body))
	}
}