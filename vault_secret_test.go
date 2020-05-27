package vaultsecret_test

import (
	"testing"

	"github.com/KebaCorp/go-vault-secret/vaultsecret"
)

var (
	databaseURL string
)

func TestSetParams(m *testing.M) {
	vaultsecret.SetParams()
}
