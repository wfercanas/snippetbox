package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wfercanas/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	staticDir := "./ui/static/"

	app := &application{
		logger:    slog.New(slog.DiscardHandler),
		staticDir: &staticDir,
	}

	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
