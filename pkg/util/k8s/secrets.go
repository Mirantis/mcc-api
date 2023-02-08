package k8s

import (
	"context"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// +gocode:public-api=true
func GetData(secret *corev1.Secret, key string) ([]byte, error) {
	if secret.Data == nil {
		return nil, errors.New("secret contains no data")
	}
	if v, ok := secret.Data[key]; ok {
		return v, nil
	}
	return nil, errors.Errorf("%s key is missing", key)
}

// +gocode:public-api=true
func GetDataFromSecret(client crclient.Client, namespace, name, key string) ([]byte, error) {
	secret, err := GetSecret(client, namespace, name)
	if err != nil {
		return nil, err
	}
	return GetData(secret, key)
}

// +gocode:public-api=true
func GetSecret(client crclient.Client, namespace, name string) (*corev1.Secret, error) {
	var secret corev1.Secret
	err := client.Get(context.TODO(), crclient.ObjectKey{
		Namespace: namespace,
		Name:      name,
	}, &secret)
	if err != nil {
		return &secret, errors.Wrapf(err, "failed to get secret %v/%v", namespace, name)
	}
	return &secret, nil
}

// +gocode:public-api=true
func CreateSecret(client crclient.Client, namespace, name string, data map[string][]byte) (*corev1.Secret, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace},
		Type:       corev1.SecretTypeOpaque,
		Data:       data,
	}
	err := client.Create(context.TODO(), secret)
	if err != nil {
		return secret, errors.Wrapf(err, "failed to create secret %v/%v", namespace, name)
	}
	return secret, nil
}

// +gocode:public-api=true
func DeleteSecret(client crclient.Client, namespace, name string) error {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace},
	}
	err := client.Delete(context.TODO(), secret)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return errors.Wrapf(err, "failed to delete secret %v/%v", namespace, name)
		}
	}
	return nil
}
