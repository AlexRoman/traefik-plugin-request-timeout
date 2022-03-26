package requesttimeout_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	requesttimeout "github.com/AlexRoman/traefik-plugin-request-timeout"
)

func TestRequestTimeout(t *testing.T) {
	cfg := requesttimeout.CreateConfig()
	cfg.ResponseTimeout = time.Second * 5

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := requesttimeout.New(ctx, next, cfg, "test-timeout")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)
}
