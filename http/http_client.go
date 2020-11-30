package http

import (
	"eth2-exporter/types"
	"github.com/sirupsen/logrus"
)

// Client provides an interface for RPC clients
type Client interface {
	GetAccounts() (types.Accounts, error)
}

var logger = logrus.New().WithField("module", "http")
