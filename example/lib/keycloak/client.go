package keycloak

import (
	"context"
	"crypto/tls"
	"net/http"
	"sync"
	"time"

	"github.com/Nerzal/gocloak/v7"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pkg/errors"
)

const (
	adminClientID string = "admin-cli"
)

type Client struct {
	gocloak.GoCloak
	username          string
	password          string
	authRealm         string
	token             *gocloak.JWT
	tokenRefreshMutex sync.Mutex
}

func NewClient(basePath, username, password, authRealm string, insecure bool) (*Client, error) {
	client := &Client{
		username:  username,
		password:  password,
		authRealm: authRealm,
	}

	client.GoCloak = gocloak.NewClient(basePath)
	// keycloak.client.RestyClient().SetDebug(true)
	var transport http.RoundTripper
	if insecure {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		transport = http.DefaultTransport
	}
	client.GoCloak.RestyClient().SetTransport(NewTokenRefresher(transport, client.RefreshToken))
	token, err := client.GoCloak.LoginAdmin(context.TODO(), username, password, authRealm)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to Keycloak")
	}
	client.token = token

	return client, nil
}

func (client *Client) RefreshToken() (string, error) {
	client.tokenRefreshMutex.Lock()
	defer client.tokenRefreshMutex.Unlock()
	_, err := jwt.ParseWithClaims(
		client.token.AccessToken,
		jwt.StandardClaims{},
		// We don't care about further token validation, just standard claims check
		func(t *jwt.Token) (interface{}, error) { return nil, nil },
	)
	if err == nil {
		// no need to refresh, token is valid
		return client.token.AccessToken, nil
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	token, err := client.GoCloak.RefreshToken(ctx, client.token.RefreshToken, adminClientID, "", client.authRealm)
	cancel()
	if err != nil {
		// refresh token expired, need to re-login
		ctx, cancel = context.WithTimeout(context.TODO(), time.Minute)
		token, err = client.LoginAdmin(ctx, client.username, client.password, client.authRealm)
		cancel()
		if err != nil {
			return "", errors.Wrap(err, "error refreshing token")
		}
	}
	client.token = token
	return token.AccessToken, nil
}

func (client *Client) GetAccessToken() string {
	return client.token.AccessToken
}
