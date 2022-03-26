// Package requesttimeout a plugin to set a timeout for a request.
package requesttimeout

import (
	"context"
	"net/http"
	"time"
)

// Config the plugin configuration.
type Config struct {
	ResponseTimeout time.Duration `json:"responseTimeout,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		ResponseTimeout: time.Second * 30,
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
	return &ResponseTimeout{
		responseTimeout: config.ResponseTimeout,
		next:            next,
		name:            name,
	}, nil
}

func (a *ResponseTimeout) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx, cancelFunc := context.WithTimeout(req.Context(), a.responseTimeout)
	defer cancelFunc()
	a.next.ServeHTTP(rw, req.WithContext(ctx))
}
