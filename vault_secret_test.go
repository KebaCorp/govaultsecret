package vaultsecret

import (
	"fmt"
	"testing"
)

func TestSetParams(m *testing.T) {
	SetParams(New())
}

func TestGetSecret(m *testing.T) {
	s := GetSecret("test")
	fmt.Println(s)
}
