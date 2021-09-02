package http

import (
	"encoding/json"
	"eth2-exporter/types"
	"eth2-exporter/utils"
	"fmt"
	"net/http"
	"time"
)

type VCClient struct {
	client  http.Client
	baseUrl string
}

// NewVCClient is used for a new VC client connection
func NewVCClient(baseUrl string) (*VCClient, error) {
	timeout := 60 * time.Second
	httpClient := http.Client{
		Timeout: timeout,
	}

	logger.Printf("http client established")

	client := &VCClient{
		client:  httpClient,
		baseUrl: baseUrl,
	}
	return client, nil
}

//// Close will close a Prysm client connection
//func (resp http.Response) Close() {
//	defer resp.Body.Close()
//}

func (c VCClient) GetAccounts() (types.Accounts, error) {
	logger.Infof("getting accounts...")
	start := time.Now()
	network := utils.Config.Indexer.ValidatorCenter.Network
	ssvAccounts := utils.Config.Indexer.ValidatorCenter.SsvAccounts
	resp, err := c.client.Get(fmt.Sprintf("%s/accounts/cached?network=%s&ssv_accounts=%s", c.baseUrl, network, ssvAccounts))
	if err != nil {
		return types.Accounts{}, err
	}
	defer resp.Body.Close()
	var accounts types.Accounts
	err = json.NewDecoder(resp.Body).Decode(&accounts)
	if err != nil {
		return types.Accounts{}, err
	}
	logger.Infof("Got %v accounts in %v", len(accounts), time.Since(start))

	return append(accounts, types.Account{
		PublicKey:       "0x8399365b797c1df7eafce19067f16ec00e21c554166d2a2d726f613167078a99b7284bb0534a0354cf421ea53126491b",
		ActivationEpoch: 33461,
		Status:          "active",
		IsSlashed:       false,
	},
	types.Account{
		PublicKey:       "0xb95e35e5cb5502e8c5eb4345c82016b870c26341bcd952a7662309c24dffc83052d4902891011cf806ec34f853fe1e18",
		ActivationEpoch: 33461,
		Status:          "active",
		IsSlashed:       false,
	},
		types.Account{
			PublicKey:       "0x80bced0450cfad50f38b856d06c3678e5436f6d29f793e83962b961e21dd7c77b9a94b9ae5a5fa6cc7bd0e7d22abee8f",
			ActivationEpoch: 33461,
			Status:          "active",
			IsSlashed:       false,
		}), nil
	//return accounts, nil
}
