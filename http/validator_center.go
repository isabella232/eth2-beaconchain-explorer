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
	client              http.Client
	baseUrl             string
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
	resp, err := c.client.Get(fmt.Sprintf("%s/accounts/cached?network=%s", c.baseUrl, network))
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
	return accounts, nil
}
