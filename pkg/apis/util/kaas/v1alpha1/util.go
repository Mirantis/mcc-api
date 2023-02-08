package v1alpha1

import (
	"context"
	"encoding/base32"

	"sigs.k8s.io/controller-runtime/pkg/client"

	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/kaas/v1alpha1"
)

var (
	Base32Encoder = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").WithPadding(base32.NoPadding)
)

func GetPublicKey(crclient client.Client, namespace, name string) (*kaasv1alpha1.PublicKey, error) {
	pubKey := &kaasv1alpha1.PublicKey{}
	objKey := client.ObjectKey{
		Name:      name,
		Namespace: namespace,
	}
	err := crclient.Get(context.TODO(), objKey, pubKey)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}
