package serializable

type Account struct {
	Alias string `json:"alias"`
	Address string `json:"address"`
	PubKey  string `json:"pubKey"`
}
