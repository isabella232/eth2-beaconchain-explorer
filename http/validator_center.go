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

	return []types.Account{{
		PublicKey:       "0x966ce696114658a88feb21ada9f3af1deff762ae03f1b41a30c68c6abb0f9d410e6f93228039bb32e14b2ffc99532f33",
		ActivationEpoch: 33461,
		Status:          "active",
		IsSlashed:       false,
		}}, nil
	//return accounts, nil
}
