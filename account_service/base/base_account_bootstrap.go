package base

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/sirupsen/logrus"
	"bytes"
	"math/big"
	"fmt"
	"errors"
)

const (
	Version            = byte(0x01)
	AddressChecksumLen = 4
)

type BaseAccountBootstrap struct {
	Alias      string
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

func NewBaseAccountBootstrap(alias string) *BaseAccountBootstrap {
	private, public := newKeyPair()
	boot := &BaseAccountBootstrap{
		PrivateKey: private,
		PublicKey:  public,
	}

	if alias != "" {
		boot.Alias = alias
	}

	return boot
}

func NewBaseAccountBootstrapByPubKey(pubKey []byte) *BaseAccountBootstrap {
	boot := &BaseAccountBootstrap{
		PublicKey: pubKey,
	}

	return boot
}

func newKeyPair() (*ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		logrus.Panicf("[BaseAccountBootstrap] newKeyPair ecdsa.GenerateKey err: %v", err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return private, pubKey
}

func (b *BaseAccountBootstrap) GetAlias() string {
	return b.Alias
}

func (b *BaseAccountBootstrap) GetAddress() []byte {
	pubKeyHash := HashPubKey(b.PublicKey)

	payload := bytes.Join([][]byte{{Version}, pubKeyHash}, []byte{})
	checksum := checkSum(payload)

	fullPayload := bytes.Join([][]byte{payload, checksum}, []byte{})

	return Base58Encode(fullPayload)
}

func (b *BaseAccountBootstrap) GetPubKey() []byte {
	return b.PublicKey
}

func (*BaseAccountBootstrap) Lock(content []byte) ([]byte, error) {
	panic("not implement")
}

func (*BaseAccountBootstrap) Unlock(content []byte) ([]byte, error) {
	panic("not implement")
}

func (b *BaseAccountBootstrap) Sign(hash []byte) ([]byte, error) {
	if b.PrivateKey == nil {
		return nil, errors.New("[BaseAccountBootstrap] Sign err: PrivateKey should not be null")
	}
	r, s, err := ecdsa.Sign(rand.Reader, b.PrivateKey, hash)
	if err != nil {
		logrus.Errorf("[BaseAccountBootstrap] Sign err: %v", err)
	}
	signature := bytes.Join([][]byte{r.Bytes(), s.Bytes()}, []byte{})
	return signature, nil
}

func (b *BaseAccountBootstrap) Verify(hash []byte, signature []byte) error {
	curve := elliptic.P256()

	r := big.Int{}
	s := big.Int{}
	sigLen := len(signature)
	r.SetBytes(signature[:sigLen/2])
	s.SetBytes(signature[sigLen/2:])

	var rawPublicKey *ecdsa.PublicKey
	if b.PrivateKey != nil {
		rawPublicKey = &b.PrivateKey.PublicKey
	} else {
		x := big.Int{}
		y := big.Int{}
		keyLen := len(b.PublicKey)
		x.SetBytes(b.PublicKey[:(keyLen / 2)])
		y.SetBytes(b.PublicKey[(keyLen / 2):])

		rawPublicKey = &ecdsa.PublicKey{
			Curve: curve,
			X:     &x,
			Y:     &y,
		}
	}

	if !ecdsa.Verify(rawPublicKey, hash, &r, &s) {
		return fmt.Errorf("[BaseAccountBootstrap] Verify failed")
	}

	return nil
}
