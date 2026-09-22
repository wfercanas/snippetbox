package main

import (
	"net/http"
	"testing"

	"github.com/wfercanas/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	res := ts.get(t, "/ping")
	assert.Equal(t, res.status, http.StatusOK)
	assert.Equal(t, res.body, "OK")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name           string
		urlPath        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid ID",
			urlPath:        "/snippet/view/1",
			expectedStatus: http.StatusOK,
			expectedBody:   "An old silent pond...",
		},
		{
			name:           "Non-Existent ID",
			urlPath:        "/snippet/view/2",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Negative ID",
			urlPath:        "/snippet/view/-1",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Decimal ID",
			urlPath:        "/snippet/view/1.25",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "String ID",
			urlPath:        "/snippet/view/foo",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Empty ID",
			urlPath:        "/snipper/view/",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts.resetClientCookieJar(t)
			res := ts.get(t, tt.urlPath)
			assert.Equal(t, res.status, tt.expectedStatus)
			assert.StringContains(t, res.body, tt.expectedBody)
		})
	}
}
