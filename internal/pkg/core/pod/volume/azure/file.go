package azure

type AzureFileVolume struct {
	SecretName string `json:"-"`
	ShareName  string `json:"-"`
	ReadOnly   bool   `json:"-"`
}
