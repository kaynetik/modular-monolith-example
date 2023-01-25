package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JWTSuite struct {
	suite.Suite

	privKeyPEM []byte
	pubKeyPEM  []byte
}

func (s *JWTSuite) SetupSuite() {}

func TestSuite(t *testing.T) {
	suite.Run(t, new(JWTSuite))
}

func (s *JWTSuite) TearDownSuite() {}

// SetupTest - defines method that will be called before each test in the suite.
func (s *JWTSuite) SetupTest() {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		s.T().Fatal(err)
	}

	privKeyPEM, pubKeyPEM, err := pemKeyPair(key)
	if err != nil {
		s.T().Fatal(err)
	}

	s.privKeyPEM = privKeyPEM
	s.pubKeyPEM = pubKeyPEM
}

// pemKeyPair - stores the generated key for later use with jwt.ParseECPrivateKeyFromPEM and jwt.ParseECPublicKeyFromPEM.
func pemKeyPair(key *ecdsa.PrivateKey) (privKeyPEM, pubKeyPEM []byte, err error) {
	der, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, nil, err
	}

	privKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der,
	})

	der, err = x509.MarshalPKIXPublicKey(key.Public())
	if err != nil {
		return nil, nil, err
	}

	pubKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: der,
	})

	return privKeyPEM, pubKeyPEM, err
}
