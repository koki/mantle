package projected

import (
	"github.com/koki/json"
	serrors "github.com/koki/structurederrors"
)

type VolumeProjection struct {
	Secret      *SecretProjection      `json:"-"`
	DownwardAPI *DownwardAPIProjection `json:"-"`
	ConfigMap   *ConfigMapProjection   `json:"-"`
}

func (p *VolumeProjection) UnmarshalJSON(data []byte) error {
	obj := map[string]interface{}{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return serrors.ContextualizeErrorf(err, "projected volume source item")
	}

	if _, ok := obj["secret"]; ok {
		p.Secret = &SecretProjection{}
		return json.Unmarshal(data, p.Secret)
	}
	if _, ok := obj["config"]; ok {
		p.ConfigMap = &ConfigMapProjection{}
		return json.Unmarshal(data, p.ConfigMap)
	}
	p.DownwardAPI = &DownwardAPIProjection{}
	return json.Unmarshal(data, p.DownwardAPI)
}

func (p VolumeProjection) MarshalJSON() ([]byte, error) {
	if p.Secret != nil {
		return json.Marshal(p.Secret)
	}

	if p.DownwardAPI != nil {
		return json.Marshal(p.DownwardAPI)
	}

	if p.ConfigMap != nil {
		return json.Marshal(p.ConfigMap)
	}

	return nil, serrors.InvalidInstanceErrorf(p, "empty volume projection")
}
