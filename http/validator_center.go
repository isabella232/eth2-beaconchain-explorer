package http

import (
	"encoding/json"
	"eth2-exporter/types"
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
	timeout := 10 * time.Second
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
	resp, err := c.client.Get(fmt.Sprintf("%s/accounts?orgId=237", c.baseUrl))  // TODO remove orgid!!
	if err != nil {
		return types.Accounts{}, err
	}
	defer resp.Body.Close()
	var accounts types.Accounts
	err = json.NewDecoder(resp.Body).Decode(&accounts)
	if err != nil {
		return types.Accounts{}, err
	}
	return accounts, nil
}
