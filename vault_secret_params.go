package vaultsecret

// VaultSecretParams
type VaultSecretParams struct {
	Source               string
	Cache                string
	Token                string
	IsSaveTemplate       bool
	SaveTemplateFilename string
}

// New VaultSecretParams
func New() VaultSecretParams {
	v := VaultSecretParams{
		Source:               "string",
		Cache:                "string",
		Token:                "string",
		IsSaveTemplate:       true,
		SaveTemplateFilename: "secretsTemplate.json",
	}

	return v
}
