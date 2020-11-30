package types



// Account is a struct to hold blox account data
type Account struct {
	PublicKey string `json:"publicKey"`
	//index     uint64
}

type Accounts []Account
