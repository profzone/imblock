package b58

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"github.com/sirupsen/logrus"
	"bytes"
)

func HashPubKey(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)
	hasher := ripemd160.New()

	_, err := hasher.Write(pubHash[:])
	if err != nil {
		logrus.Errorf("[B58AccountBootstrap] HashPubKey err: %v", err)
		return nil
	}

	return hasher.Sum(nil)
}

func ValidateAddress(address []byte) bool {
	publicKeyHash := Base58Decode(address)
	originCheckSum := publicKeyHash[len(publicKeyHash)-AddressChecksumLen:]
	version := publicKeyHash[0]
	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-AddressChecksumLen]
	targetCheckSum := checkSum(bytes.Join([][]byte{{version}, publicKeyHash}, []byte{}))

	return bytes.Compare(originCheckSum, targetCheckSum) == 0
}
