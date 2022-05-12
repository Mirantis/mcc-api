package bootstrap

import (
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"

	"github.com/Mirantis/mcc-api/pkg/errors"
)

const (
	// RSAPrivateKeyBlockType is a possible value for pem.Block.Type.
	RSAPrivateKeyBlockType = "RSA PRIVATE KEY"
)

func CreatePrivateKey() (*rsa.PrivateKey, error) {
	return NewRSA(4096)
}

func ParsePrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, err := DecodePEM(pemBytes)
	if err != nil {
		return nil, errors.Wrap(err, "ssh: failed to parse ssh key")
	}

	if strings.Contains(block.Headers["Proc-Type"], "ENCRYPTED") {
		return nil, errors.New("ssh: cannot decode encrypted private keys")
	}

	switch block.Type {
	case RSAPrivateKeyBlockType:
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	default:
		return nil, fmt.Errorf("ssh: unsupported key type: %q", block.Type)
	}
}

func getKeyFromFile(keyPath string) (*rsa.PrivateKey, error) {
	_, err := os.Stat(keyPath)
	if err != nil {
		return nil, errors.Errorf("get key path fails: %w", err)
	}

	privateKeyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, errors.Errorf("read key path fails: %w", err)
	}

	privateKey, err := ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func createPrivateKeyFile(keyPath string) (*rsa.PrivateKey, error) {
	err := os.MkdirAll(path.Dir(keyPath), 0700)
	if err != nil {
		return nil, errors.Errorf("make key dir fails: %w", err)
	}

	privateKey, err := CreatePrivateKey()
	if err != nil {
		return nil, err
	}

	privateKeyData := RSAToPEM(privateKey)

	err = ioutil.WriteFile(keyPath, privateKeyData, 0600)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func getOrCreatePrivateKey(desiredKeyPath string) (*rsa.PrivateKey, error) {
	key, err := getKeyFromFile(desiredKeyPath)
	if err == nil {
		klog.Infof("ssh: using existing SSH key %s", desiredKeyPath)
		return key, nil
	}
	if !os.IsNotExist(err) {
		return nil, err
	}

	klog.Infof("ssh: generate SSH key %s", desiredKeyPath)
	return createPrivateKeyFile(desiredKeyPath)
}

func GetSSHKey(privateKeyPath string) (*rsa.PrivateKey, ssh.PublicKey, error) {
	privateKey, err := getOrCreatePrivateKey(privateKeyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("ssh: failed to load private key: %w", err)
	}

	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("ssh: failed to get public key: %w", err)
	}
	return privateKey, publicKey, nil
}

func DecodePEM(pemData []byte) (*pem.Block, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, errors.New("decode PEM")
	}

	return block, nil
}

func NewRSA(size int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(cryptorand.Reader, size)
	if err != nil {
		return nil, errors.Wrap(err, "create RSA key")
	}

	return privateKey, privateKey.Validate()
}

func RSAToPEM(key *rsa.PrivateKey) []byte {
	return NewPEM(RSAPrivateKeyBlockType, x509.MarshalPKCS1PrivateKey(key))
}

func NewPEM(blockType string, data []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  blockType,
		Bytes: data,
	})
}
