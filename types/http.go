package types



// Account is a struct to hold blox account data
type Account struct {
	PublicKey string `json:"publicKey"`
	ActivationEpoch uint64 `json:"activationEpoch"`
	Status string `json:"status"`
	IsSlashed bool `json:"is_slashed"`
	//index     uint64
}

type Accounts []Account
