package converterutils

import (
	"mantle/pkg/util"

	"k8s.io/api/core/v1"
)

func NewKubeV1LocalObjectRef(ref string) *v1.LocalObjectReference {
	if len(ref) == 0 {
		return nil
	}
	return &v1.LocalObjectReference{
		Name: ref,
	}
}

func RequiredToOptional(required *bool) *bool {
	if required == nil {
		return nil
	}

	return util.BoolPtr(!*required)
}

func FromKubeLocalObjectReferenceV1(kubeRef *v1.LocalObjectReference) string {
	if kubeRef == nil {
		return ""
	}

	return kubeRef.Name
}

func OptionalToRequired(optional *bool) *bool {
	if optional == nil {
		return nil
	}

	return util.BoolPtr(!*optional)
}
