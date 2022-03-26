package traefik_plugin_request_timeout_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	traefik_plugin_request_timeout "github.com/AlexRoman/traefik-plugin-request-timeout"
)

func TestRequestTimeout(t *testing.T) {
	cfg := traefik_plugin_request_timeout.CreateConfig()
	cfg.ResponseTimeout = "5s"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefik_plugin_request_timeout.New(ctx, next, cfg, "test-timeout")
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
