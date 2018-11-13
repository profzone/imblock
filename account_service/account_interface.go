package account_service

type Account interface {
	GetAlias() string
	GetAddress() []byte
	GetPubKey() []byte

	// 对内容使用公钥进行锁定
	Lock(content []byte) ([]byte, error)
	// 对内容使用私钥解锁
	Unlock(content []byte) ([]byte, error)

	// 对内容使用私钥进行签名
	Sign([]byte) ([]byte, error)
	// 对内容使用公钥进行验签
	Verify(hash []byte, signature []byte) error
}

type AccountService interface {
	CreateAccount(alias string) Account
	GetAccount(address []byte) Account
	GetAccounts() []Account

	AccountStorage
}

type AccountStorage interface {
	Save(filePath string) error
	Load(filePath string) error
}