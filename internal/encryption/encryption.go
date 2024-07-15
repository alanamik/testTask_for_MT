package encryption

import (
	"crypto"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

var (
	ErrUndefinedAlgorithm = errors.New("available algorithms are SHA-256 or MD5")
	ErrEmptyString        = errors.New("string is empty")
)

type EncryptionService struct {
}

func New() *EncryptionService {
	return &EncryptionService{}
}
func (e *EncryptionService) Encrypt(str string, alg string) (string, error) {
	if len(str) == 0 {
		return "", ErrEmptyString
	}
	var cypher string
	var err error
	switch alg {
	case crypto.SHA256.String():
		cypher, err = e.EncryptSHA256(str)
	case crypto.MD5.String():
		cypher, err = e.EncryptMD5(str)
	default:
		return "", ErrUndefinedAlgorithm
	}

	return cypher, err
}

func (e *EncryptionService) EncryptSHA256(str string) (string, error) {
	h := sha256.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		return "", err
	}
	cypher := fmt.Sprintf("%x", h.Sum(nil))

	return cypher, nil
}

func (e *EncryptionService) EncryptMD5(str string) (string, error) {
	h := md5.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		return "", err
	}
	cypher := fmt.Sprintf("%x", h.Sum(nil))

	return cypher, nil
}
