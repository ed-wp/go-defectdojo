package defectdojo_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockAPIServer struct {
	t          *testing.T
	server     *httptest.Server
	apiToken   string
	statusCode int
	response   []byte
}

func NewMockAPIServer(t *testing.T, apiToken string, statusCode int, response string) *httptest.Server {
	m := &mockAPIServer{
		t:          t,
		apiToken:   apiToken,
		statusCode: statusCode,
		response:   []byte(response),
	}
	s := httptest.NewServer(m)
	m.server = s
	return s
}

func (m *mockAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Token "+m.apiToken {
		m.t.Fatal("mock server internal error: missing api token/authorization header")
	}

	w.WriteHeader(m.statusCode)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		m.t.Fatalf("mock server internal error: %s", err)
	}
	r.ParseMultipartForm(10 * 1024 * 1024 * 1024)
	for k, v := range r.Header {
		m.t.Logf("[mockAPIServer] %s=%s", k, strings.Join(v, ""))
	}
	m.t.Logf("[mockAPIServer] %s", b)

	_, err = w.Write(m.response)
	if err != nil {
		m.t.Fatalf("mock server internal error: %s", err)
	}
}

func (m *mockAPIServer) Close() {
	m.server.Close()
}
