package datastore

import (
	"go-boilerplate/config"
	"testing"
)

func TestConnect(t *testing.T) {
	config.InitConfigs("./../.env")
	Connect()
}
