package keycloak

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Mirantis/mcc-api/pkg/errors"
)

type tokenRefresher struct {
	upstream    http.RoundTripper
	refreshFunc func() (string, error)
}

var _ http.RoundTripper = &tokenRefresher{}

func NewTokenRefresher(t http.RoundTripper, refreshFunc func() (string, error)) http.RoundTripper {
	return &tokenRefresher{
		upstream:    t,
		refreshFunc: refreshFunc,
	}
}

func (r *tokenRefresher) RoundTrip(req *http.Request) (*http.Response, error) {
	auth := req.Header.Get("Authorization")
	hasBearerAuth := strings.HasPrefix(auth, "Bearer ")
	resp, err := r.upstream.RoundTrip(req)
	if err != nil || !hasBearerAuth || resp.StatusCode != http.StatusUnauthorized {
		return resp, err
	}
	fmt.Println("Refreshing token...")
	token, err := r.refreshFunc()
	if err != nil {
		return resp, errors.Errorf("refresh func fails: %w", err)
	}
	newreq := new(http.Request)
	*newreq = *req
	if req.URL != nil {
		newreq.URL = new(url.URL)
		*newreq.URL = *req.URL
	}
	newreq.Header.Set("Authorization", "Bearer "+token)
	return r.upstream.RoundTrip(newreq)
}
