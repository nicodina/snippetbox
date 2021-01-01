package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/nicodina/snippetbox/pkg/models/mock"
)

func newTestApplication(t *testing.T) *application {

	templateCache, err := newTemplateCache("./../../ui/html")
	if err != nil {
		t.Fatal(err)
	}

	session := sessions.New([]byte("s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	return &application{
		errLog:        log.New(ioutil.Discard, "", 0),
		infoLog:       log.New(ioutil.Discard, "", 0),
		session:       session,
		templateCache: templateCache,
		users:         &mock.UsersService{},
		snippets:      &mock.SnippetService{},
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, handler http.Handler) *testServer {
	server := httptest.NewTLSServer(handler)

	// Let's attach a jar to store response cookies
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	server.Client().Jar = jar

	// Stop redirections
	server.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{server}
}

func (ts *testServer) get(t *testing.T, path string) (int, http.Header, []byte) {
	resp, err := ts.Client().Get(ts.URL + path)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, resp.Header, body
}
