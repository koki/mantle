package env

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"
)

// ToKube will return kubernetes EnvVar and EnvVarFromSource objects of the api version provided
func (e *Env) ToKube(version string) (interface{}, interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return e.toKubeV1()
	case "":
		return e.toKubeV1()
	default:
		return nil, nil, fmt.Errorf("unsupported api version for Env: %s", version)
	}
}

func (e *Env) toKubeV1() (*v1.EnvVar, *v1.EnvFromSource, error) {
	var envVar *v1.EnvVar
	var envVarFromSrc *v1.EnvFromSource

	if e.Type == EnvValEnvType {
		envVar = &v1.EnvVar{
			Name:  e.Val.Key,
			Value: e.Val.Val,
		}
	} else {
		from := e.From

		switch from.From {
		case EnvFromTypeConfig:
			if len(from.ConfigMapOrSecretKey) > 0 {
				envVar = &v1.EnvVar{
					Name: from.VarNameOrPrefix,
					ValueFrom: &v1.EnvVarSource{
						ConfigMapKeyRef: &v1.ConfigMapKeySelector{
							LocalObjectReference: v1.LocalObjectReference{
								Name: from.ConfigMapOrSecretName,
							},
							Key:      from.ConfigMapOrSecretKey,
							Optional: from.Optional(),
						},
					},
				}
			} else {
				envVarFromSrc = &v1.EnvFromSource{
					Prefix: from.VarNameOrPrefix,
					ConfigMapRef: &v1.ConfigMapEnvSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: from.ConfigMapOrSecretName,
						},
						Optional: from.Optional(),
					},
				}
			}

		case EnvFromTypeSecret:
			if len(from.ConfigMapOrSecretKey) > 0 {
				envVar = &v1.EnvVar{
					Name: from.VarNameOrPrefix,
					ValueFrom: &v1.EnvVarSource{
						SecretKeyRef: &v1.SecretKeySelector{
							LocalObjectReference: v1.LocalObjectReference{
								Name: from.ConfigMapOrSecretName,
							},
							Key:      from.ConfigMapOrSecretKey,
							Optional: from.Optional(),
						},
					},
				}
			} else {
				envVarFromSrc = &v1.EnvFromSource{
					Prefix: from.VarNameOrPrefix,
					SecretRef: &v1.SecretEnvSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: from.ConfigMapOrSecretName,
						},
						Optional: from.Optional(),
					},
				}
			}

		case EnvFromTypeCPULimits, EnvFromTypeMemLimits, EnvFromTypeEphemeralStorageLimits,
			EnvFromTypeCPURequests, EnvFromTypeMemRequests, EnvFromTypeEphemeralStorageRequests:
			// ResourceFieldRef
			envVar = &v1.EnvVar{
				Name: from.VarNameOrPrefix,
				ValueFrom: &v1.EnvVarSource{
					ResourceFieldRef: &v1.ResourceFieldSelector{
						Resource: string(from.From),
					},
				},
			}

		case EnvFromTypeMetadataName, EnvFromTypeMetadataNamespace, EnvFromTypeMetadataLabels, EnvFromTypeMetadataAnnotation,
			EnvFromTypeSpecNodename, EnvFromTypeSpecServiceAccountName,
			EnvFromTypeStatusHostIP, EnvFromTypeStatusPodIP:
			// FieldRef
			envVar = &v1.EnvVar{
				Name: from.VarNameOrPrefix,
				ValueFrom: &v1.EnvVarSource{
					FieldRef: &v1.ObjectFieldSelector{
						FieldPath: string(from.From),
					},
				},
			}

		default:
			return nil, nil, fmt.Errorf("unrecognized EnvFromType: %s", from.From)
		}

	}

	return envVar, envVarFromSrc, nil
}
