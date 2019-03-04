package env

import (
	"fmt"
	"reflect"

	"mantle/internal/converterutils"

	"k8s.io/api/core/v1"
)

// NewEnvFromKubeEnvVar will create a new
// Env object with the data from a provided kubernetes
// EnvVar object
func NewEnvFromKubeEnvVar(obj interface{}) (*Env, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.EnvVar{}):
		return fromKubeEnvVarV1(obj.(v1.EnvVar))
	case reflect.TypeOf(&v1.EnvVar{}):
		o := obj.(*v1.EnvVar)
		return fromKubeEnvVarV1(*o)
	default:
		return nil, fmt.Errorf("unknown EnvVar version: %s", reflect.TypeOf(obj))
	}
}

// NewEnvFromKubeEnvFromSource will create a new
// Env object with the data from a provided kubernetes
// EnvFromSource object
func NewEnvFromKubeEnvFromSource(obj interface{}) (*Env, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.EnvFromSource{}):
		return fromKubeEnvFromSourceV1(obj.(v1.EnvFromSource))
	case reflect.TypeOf(&v1.EnvFromSource{}):
		o := obj.(*v1.EnvFromSource)
		return fromKubeEnvFromSourceV1(*o)
	default:
		return nil, fmt.Errorf("unknown EnvFromSource version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeEnvFromSourceV1(envFromSource v1.EnvFromSource) (*Env, error) {
	e := EnvFrom{
		VarNameOrPrefix: envFromSource.Prefix,
	}

	if envFromSource.ConfigMapRef != nil {
		e.From = EnvFromTypeConfig
		e.ConfigMapOrSecretName = envFromSource.ConfigMapRef.Name
		e.Required = converterutils.OptionalToRequired(envFromSource.ConfigMapRef.Optional)
	}
	if envFromSource.SecretRef != nil {
		e.From = EnvFromTypeSecret
		e.ConfigMapOrSecretName = envFromSource.SecretRef.Name
		e.Required = converterutils.OptionalToRequired(envFromSource.SecretRef.Optional)
	}

	env := EnvWithFrom(e)
	return &env, nil
}

func fromKubeEnvVarV1(envVar v1.EnvVar) (*Env, error) {
	var env Env

	if envVar.ValueFrom == nil {
		env = EnvWithVal(EnvVal{
			Key: envVar.Name,
			Val: envVar.Value,
		})
	} else {
		e := EnvFrom{
			VarNameOrPrefix: envVar.Name,
		}

		if envVar.ValueFrom.FieldRef != nil {
			e.From = EnvFromType(envVar.ValueFrom.FieldRef.FieldPath)
		}

		if envVar.ValueFrom.ResourceFieldRef != nil {
			//This might be losing some information
			e.From = EnvFromType(envVar.ValueFrom.ResourceFieldRef.Resource)
		}

		if envVar.ValueFrom.ConfigMapKeyRef != nil {
			e.From = EnvFromTypeConfig
			e.ConfigMapOrSecretName = envVar.ValueFrom.ConfigMapKeyRef.Name
			e.ConfigMapOrSecretKey = envVar.ValueFrom.ConfigMapKeyRef.Key
			e.Required = converterutils.OptionalToRequired(envVar.ValueFrom.ConfigMapKeyRef.Optional)
		}

		if envVar.ValueFrom.SecretKeyRef != nil {
			e.From = EnvFromTypeSecret
			e.ConfigMapOrSecretName = envVar.ValueFrom.SecretKeyRef.Name
			e.ConfigMapOrSecretKey = envVar.ValueFrom.SecretKeyRef.Key
			e.Required = converterutils.OptionalToRequired(envVar.ValueFrom.SecretKeyRef.Optional)
		}

		env = EnvWithFrom(e)
	}

	return &env, nil
}
