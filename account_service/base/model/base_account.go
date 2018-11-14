package model

type BaseAccount struct {
	Alias   string `json:"alias"`
	Address string `json:"address"`
	PubKey  string `json:"pubKey"`
}
