package b58

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/sirupsen/logrus"
	"crypto/sha256"
	"bytes"
	"math/big"
	"fmt"
	"errors"
)

const (
	Version            = byte(0x01)
	AddressChecksumLen = 4
)

type B58AccountBootstrap struct {
	Alias      string
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

func NewB58AccountBootstrap(alias string) *B58AccountBootstrap {
	private, public := newKeyPair()
	boot := &B58AccountBootstrap{
		PrivateKey: private,
		PublicKey:  public,
	}

	if alias != "" {
		boot.Alias = alias
	}

	return boot
}

func NewB58AccountBootstrapByPubKey(pubKey []byte) *B58AccountBootstrap {
	boot := &B58AccountBootstrap{
		PublicKey: pubKey,
	}

	return boot
}

func newKeyPair() (*ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		logrus.Panicf("[B58AccountBootstrap] newKeyPair ecdsa.GenerateKey err: %v", err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return private, pubKey
}

func (b *B58AccountBootstrap) GetAlias() string {
	return b.Alias
}

func (b *B58AccountBootstrap) GetAddress() []byte {
	pubKeyHash := HashPubKey(b.PublicKey)

	payload := bytes.Join([][]byte{{Version}, pubKeyHash}, []byte{})
	checksum := checkSum(payload)

	fullPayload := bytes.Join([][]byte{payload, checksum}, []byte{})

	return Base58Encode(fullPayload)
}

func checkSum(payload []byte) []byte {
	firstSum := sha256.Sum256(payload)
	secondSum := sha256.Sum256(firstSum[:])
	return secondSum[:AddressChecksumLen]
}

func (b *B58AccountBootstrap) GetPubKey() []byte {
	return b.PublicKey
}

func (*B58AccountBootstrap) Lock(content []byte) ([]byte, error) {
	panic("not implement")
}

func (*B58AccountBootstrap) Unlock(content []byte) ([]byte, error) {
	panic("not implement")
}

func (b *B58AccountBootstrap) Sign(hash []byte) ([]byte, error) {
	if b.PrivateKey == nil {
		return nil, errors.New("[B58AccountBootstrap] Sign err: PrivateKey should not be null")
	}
	r, s, err := ecdsa.Sign(rand.Reader, b.PrivateKey, hash)
	if err != nil {
		logrus.Errorf("[B58AccountBootstrap] Sign err: %v", err)
	}
	signature := bytes.Join([][]byte{r.Bytes(), s.Bytes()}, []byte{})
	return signature, nil
}

func (b *B58AccountBootstrap) Verify(hash []byte, signature []byte) error {
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
		return fmt.Errorf("[B58AccountBootstrap] Verify failed")
	}

	return nil
}
