package account_service

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"github.com/profzone/imblock/global"
	"os"
	"github.com/profzone/imblock/account_service/b58"
	"crypto/elliptic"
)

func init() {
	gob.Register(elliptic.P256())
}

type B58Bootstrap struct {
	Accounts map[string]*b58.B58AccountBootstrap
	Alias    map[string]*b58.B58AccountBootstrap
}

func NewB58Bootstrap() *B58Bootstrap {
	return &B58Bootstrap{
		Accounts: make(map[string]*b58.B58AccountBootstrap),
		Alias:    make(map[string]*b58.B58AccountBootstrap),
	}
}

func (b *B58Bootstrap) CreateAccount(alias string) Account {

	if _, exist := b.Alias[alias]; exist {
		return nil
	}

	account := b58.NewB58AccountBootstrap(alias)
	address := account.GetAddress()

	b.Accounts[string(address)] = account
	if alias != "" {
		b.Alias[alias] = account
	}

	b.Save(global.GetWalletFilePath())
	return account
}

func (b *B58Bootstrap) GetAddress() [][]byte {
	var addresses [][]byte

	for address := range b.Accounts {
		addresses = append(addresses, []byte(address))
	}

	return addresses
}

func (b *B58Bootstrap) GetAccount(address []byte) Account {
	return b.Accounts[string(address)]
}

func (b *B58Bootstrap) GetAccounts() []Account {
	result := make([]Account, 0)
	for _, account := range b.Accounts {
		result = append(result, account)
	}
	return result
}

func (b *B58Bootstrap) Save(filePath string) error {
	var buffer bytes.Buffer
	walletFileName := global.GetWalletFilePath()

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(b)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(walletFileName, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (b *B58Bootstrap) Load(filePath string) error {
	walletFileName := global.GetWalletFilePath()
	if _, err := os.Stat(walletFileName); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFileName)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(fileContent)
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(b)
	if err != nil {
		return err
	}

	return nil
}
