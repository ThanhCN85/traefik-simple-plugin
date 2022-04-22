package headerverify

import (
	"context"
	"net/http"
)

type Config struct {}

func CreateConfig() *Config {
	return &Config{}
}

type HeaderVerify struct {
	next	http.Handler
	name	string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HeaderVerify{
		next: next,
		name: name,
	}, nil
}

func (h *HeaderVerify) ServeHTTP(rw http.ResponseWriter, req *http.Request){
	//rw.Write([]byte("Hello world\n"))

	if req.Header.Get("gatewaysecret") != "optiq-key" {
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	

	//req.Header.Set("X-Traefik-Request", "12345780")

	h.next.ServeHTTP(rw, req)
}