package gutil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type RSAKeyFormat int

const (
	RSAKeyFormatPKCS1 = 1
	RSAKeyFormatPKCS8 = 2
)

// ParsePrivateKeyFromPem parses private key from the pem value.
func ParseRSAPrivateKey(pemBytes []byte, keyFormat RSAKeyFormat) (*rsa.PrivateKey, error) {
	key, err := ParsePrivateKey(pemBytes, keyFormat)

	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("key type is not RSA")
	}
	return rsaKey, nil
}

func ParsePrivateKey(pemBytes []byte, keyFormat RSAKeyFormat) (key interface{}, err error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	switch keyFormat {
	case RSAKeyFormatPKCS1:
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case RSAKeyFormatPKCS8:
		key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	}

	return
}

func ParseRSAPublicKey(pemBytes []byte) (*rsa.PublicKey, error) {
	pub, err := ParsePublicKey(pemBytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("key type is not RSA")
	}

	return publicKey, nil
}

func ParsePublicKey(pemBytes []byte) (key interface{}, err error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	key, err = x509.ParsePKIXPublicKey(block.Bytes)
	return
}
