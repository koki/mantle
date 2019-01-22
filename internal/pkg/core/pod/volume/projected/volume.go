package projected

type VolumeProjection struct {
	Secret      *SecretProjection      `json:"-"`
	DownwardAPI *DownwardAPIProjection `json:"-"`
	ConfigMap   *ConfigMapProjection   `json:"-"`
}
