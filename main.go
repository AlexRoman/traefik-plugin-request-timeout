// Package requesttimeout a plugin to set a timeout for a request.
package traefik_plugin_request_timeout

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// Config the plugin configuration.
type Config struct {
	ResponseTimeout string `json:"responseTimeout,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		ResponseTimeout: "30s",
	}
}

// ResponseTimeout a ResponseTimeout plugin.
type ResponseTimeout struct {
	next            http.Handler
	responseTimeout time.Duration
	name            string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	responseTimeout, err := time.ParseDuration(config.ResponseTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse responseTimeout")
	}
	return &ResponseTimeout{
		responseTimeout: responseTimeout,
		next:            next,
		name:            name,
	}, nil
}

func (a *ResponseTimeout) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	http.TimeoutHandler(a.next, a.responseTimeout, "request timeout by custom plugin").ServeHTTP(rw, req)
}
