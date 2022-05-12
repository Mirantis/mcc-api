package keycloak

import (
	"context"
	"sync"
	"time"

	"github.com/Nerzal/gocloak/v7"
	"github.com/pkg/errors"
)

type Config struct {
	BasePath  string
	Username  string
	Password  string
	AuthRealm string
	Realm     string
	// These are IMC sync specific
	RoleTemplate   string
	RemoveOldRoles bool
	UsernameSuffix string
}

type Keycloak struct {
	cfg                   Config
	Client                *Client
	clientRedirectMutexes *sync.Map
}

func NewKeycloak(cfg Config) (*Keycloak, error) {
	keycloak := &Keycloak{
		cfg:                   cfg,
		clientRedirectMutexes: &sync.Map{},
	}

	// todo disable when we'll have good cert for Keycloak
	insecure := true
	var err error
	keycloak.Client, err = NewClient(cfg.BasePath, cfg.Username, cfg.Password, cfg.AuthRealm, insecure)
	if err != nil {
		return nil, errors.Wrap(err, "error creating client")
	}

	return keycloak, nil
}

func (keycloak *Keycloak) GetToken(clientID string) (*gocloak.JWT, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	return keycloak.Client.GetToken(ctx, keycloak.cfg.AuthRealm, gocloak.TokenOptions{
		ClientID:      gocloak.StringP(clientID),
		GrantType:     gocloak.StringP("password"),
		Scopes:        &[]string{"openid", "offline_access"},
		ResponseTypes: &[]string{"token", "id_token"},
		Username:      gocloak.StringP(keycloak.cfg.Username),
		Password:      gocloak.StringP(keycloak.cfg.Password),
	})
}
