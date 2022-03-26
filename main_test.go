package traefikpluginrequesttimeout_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexRoman/traefikpluginrequesttimeout"
)

func TestRequestTimeout(t *testing.T) {
	cfg := traefikpluginrequesttimeout.CreateConfig()
	cfg.ResponseTimeout = "5s"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefikpluginrequesttimeout.New(ctx, next, cfg, "test-timeout")
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
