package vaultsecret

import "fmt"

func SetParams(p VaultSecretParams) {
	fmt.Println("SetParams success" + p.Source)
}

func GetSecret(key string) string {
	fmt.Println("GetSecret success " + key)

	return "GetSecret success"
}
