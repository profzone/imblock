package base

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"github.com/profzone/imblock/global"
	"os"
	"crypto/elliptic"
	"github.com/profzone/imblock/core/model"
)

func init() {
	gob.Register(elliptic.P256())
}

var base *BaseBootstrap

type BaseBootstrap struct {
	Accounts map[string]*BaseAccountBootstrap
	Alias    map[string]*BaseAccountBootstrap
}

func GetBase() *BaseBootstrap {
	return base
}

func NewBaseBootstrap() *BaseBootstrap {
	if base != nil {
		return base
	}
	base = &BaseBootstrap{
		Accounts: make(map[string]*BaseAccountBootstrap),
		Alias:    make(map[string]*BaseAccountBootstrap),
	}
	return base
}

func (b *BaseBootstrap) CreateAccount(alias string) model.Account {

	if _, exist := b.Alias[alias]; exist {
		return nil
	}

	account := NewBaseAccountBootstrap(alias)
	address := account.GetAddress()

	b.Accounts[string(address)] = account
	if alias != "" {
		b.Alias[alias] = account
	}

	b.Save(global.GetWalletFilePath())
	return account
}

func (b *BaseBootstrap) GetAddress() [][]byte {
	var addresses [][]byte

	for address := range b.Accounts {
		addresses = append(addresses, []byte(address))
	}

	return addresses
}

func (b *BaseBootstrap) GetAccount(address []byte) model.Account {
	return b.Accounts[string(address)]
}

func (b *BaseBootstrap) GetAccounts() []model.Account {
	result := make([]model.Account, 0)
	for _, account := range b.Accounts {
		result = append(result, account)
	}
	return result
}

func (b *BaseBootstrap) Save(filePath string) error {
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

func (b *BaseBootstrap) Load(filePath string) error {
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
