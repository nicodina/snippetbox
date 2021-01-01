package main

import (
	"net/http"
	"bytes"
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

func TestShowSnippet(t *testing.T) {

	app := newTestApplication(t)
	server := newTestServer(t, app.routes())
	defer server.Close()

	tests := []struct {
		name string
		urlPath string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("A mocked snippet content")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			status, _, body := server.get(t, test.urlPath)

			if status != test.wantCode {
				t.Errorf("want %d; got %d", test.wantCode, status)
			}
			if !bytes.Contains(body, test.wantBody) {
				t.Errorf("want body to contain %q", test.wantBody)
			}
		})
	}
}