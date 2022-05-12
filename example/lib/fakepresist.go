package lib

import (
	"k8s.io/client-go/rest"
)

type FakePersister struct{}

var _ rest.AuthProviderConfigPersister = &FakePersister{}

func NewFakePersister() rest.AuthProviderConfigPersister {
	return &FakePersister{}
}

func (*FakePersister) Persist(map[string]string) error {
	return nil
}
