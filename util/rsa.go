package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}

//加密
func Encrypt(key *rsa.PublicKey, input []string) []string {
	result := []string{}
	for _, s := range input {
		if s == "" {
			//no message no encrypt
			result = append(result, "")
			continue
		}
		if len(s) > 124 {
			s = s[len(s)-124:]
		}

		encryptedBytes, err := rsa.EncryptOAEP(
			sha512.New(),
			rand.Reader,
			key,
			[]byte(s),
			nil)
		if err != nil {
			println(input)
			//log.Fatalln(err)
		}
		str1 := base64.RawURLEncoding.EncodeToString(encryptedBytes) + "tx2022"
		result = append(result, str1)
	}
	return result
}

//解密
func Decrypt(key *rsa.PrivateKey, input []string) []string {
	result := []string{}
	for _, s := range input {
		if s == "" {
			result = append(result, "")
			continue
		}
		if len(s) <= 50 { //标识性主键
			result = append(result, s)
			continue
		}
		str2, _ := base64.RawURLEncoding.DecodeString(strings.Split(s, "tx2022")[0])
		encryptedBytes := []byte(str2)
		decryptedBytes, err := key.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA512})
		if err != nil {
			//log.Fatalln("信息异常"+s, err)
			result = append(result, s)
			continue
		}
		result = append(result, string(decryptedBytes))
	}
	return result
}
