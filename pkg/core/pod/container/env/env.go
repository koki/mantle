package env

import (
	"fmt"
)

type EnvFromType string

const (
	EnvFromTypeSecret EnvFromType = "secret"
	EnvFromTypeConfig EnvFromType = "config"

	EnvFromTypeCPULimits              EnvFromType = "limits.cpu"
	EnvFromTypeMemLimits              EnvFromType = "limits.memory"
	EnvFromTypeEphemeralStorageLimits EnvFromType = "limits.ephemeral-storage"

	EnvFromTypeCPURequests              EnvFromType = "requests.cpu"
	EnvFromTypeMemRequests              EnvFromType = "requests.memory"
	EnvFromTypeEphemeralStorageRequests EnvFromType = "requests.ephemeral-storage"

	EnvFromTypeMetadataName       EnvFromType = "metadata.name"
	EnvFromTypeMetadataNamespace  EnvFromType = "metadata.namespace"
	EnvFromTypeMetadataLabels     EnvFromType = "metadata.labels"
	EnvFromTypeMetadataAnnotation EnvFromType = "metadata.annotations"

	EnvFromTypeSpecNodename           EnvFromType = "spec.nodeName"
	EnvFromTypeSpecServiceAccountName EnvFromType = "spec.serviceAccountName"

	EnvFromTypeStatusHostIP EnvFromType = "status.hostIP"
	EnvFromTypeStatusPodIP  EnvFromType = "status.podIP"
)

func NewEnv(key, val string) (Env, error) {
	if key == "" {
		return Env{}, fmt.Errorf("Env key cannnot be empty")
	}
	return Env{
		Type: EnvValEnvType,
		Val: &EnvVal{
			Key: key,
			Val: val,
		},
	}, nil
}

func NewEnvFrom(key string, from EnvFromType) (Env, error) {
	if key == "" {
		return Env{}, fmt.Errorf("Env key cannot be empty")
	}
	if from == EnvFromTypeConfig || from == EnvFromTypeSecret {
		return Env{}, fmt.Errorf("%s not supported. Use NewEnvFromSecret() or NewEnvFromConfig() for building new envs from Secret or ConfigMap resources", from)
	}
	required := false
	return Env{
		Type: EnvFromEnvType,
		From: &EnvFrom{
			VarNameOrPrefix: key,
			From:            from,
			Required:        &required,
		},
	}, nil
}

func NewEnvFromSecretOrConfig(resType EnvFromType, prefix, resName, resKey string) (Env, error) {
	if resType != EnvFromTypeSecret && resType != EnvFromTypeConfig {
		return Env{}, fmt.Errorf("%s not supported. Use NewEnvFrom() for building new envs from resources other than Secret or ConfigMap resources", resType)
	}

	required := true

	return Env{
		Type: EnvFromEnvType,
		From: &EnvFrom{
			VarNameOrPrefix:       prefix,
			From:                  resType,
			ConfigMapOrSecretName: resName,
			ConfigMapOrSecretKey:  resKey,
			Required:              &required,
		},
	}, nil
}

func NewEnvFromSecret(key, secretName, secretKey string) (Env, error) {
	return NewEnvFromSecretOrConfig(EnvFromTypeSecret, key, secretName, secretKey)
}

func NewEnvFromConfig(key, configName, configKey string) (Env, error) {
	return NewEnvFromSecretOrConfig(EnvFromTypeConfig, key, configName, configKey)
}

type EnvFrom struct {
	From                  EnvFromType
	VarNameOrPrefix       string `json:"varNameOrPrefix,omitempty"`
	ConfigMapOrSecretName string `json:"configMapOrSecretName,omitempty"`
	ConfigMapOrSecretKey  string `json:"configMapOrSecretKey,omitempty"`
	Required              *bool  `json:"required,omitempty"`
}

type EnvVal struct {
	Key string
	Val string
}

type Env struct {
	Type EnvType
	From *EnvFrom
	Val  *EnvVal
}

type EnvType int

const (
	EnvFromEnvType EnvType = iota
	EnvValEnvType
)

func (e EnvFrom) Optional() *bool {
	if e.Required == nil {
		return nil
	}

	optional := !(*e.Required)
	return &optional
}

func (e *Env) SetVal(val EnvVal) {
	e.Type = EnvValEnvType
	e.Val = &val
}

func (e *Env) SetFrom(from EnvFrom) {
	e.Type = EnvFromEnvType
	e.From = &from
}

func EnvWithVal(val EnvVal) Env {
	return Env{
		Type: EnvValEnvType,
		Val:  &val,
	}
}

func EnvWithFrom(from EnvFrom) Env {
	return Env{
		Type: EnvFromEnvType,
		From: &from,
	}
}
