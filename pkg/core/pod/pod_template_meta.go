package pod

import (
	"fmt"
	"reflect"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodTemplateMeta defines the metadata for a pod
type PodTemplateMeta struct {
	Cluster     string            `json:"cluster,omitempty"`
	Name        string            `json:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// ToKube converts the PodTemplateMeta object to a
// kubernetes meta object of the version specified
func (ptm *PodTemplateMeta) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return ptm.toKubeV1()
	case "":
		return ptm.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for PodTemplateMeta: %s", version)
	}
}

func (ptm *PodTemplateMeta) toKubeV1() (*metav1.ObjectMeta, error) {
	var labels map[string]string
	var annotations map[string]string

	if len(ptm.Labels) > 0 {
		labels = ptm.Labels
	}

	if len(ptm.Annotations) > 0 {
		annotations = ptm.Annotations
	}

	return &metav1.ObjectMeta{
		Name:        ptm.Name,
		Namespace:   ptm.Namespace,
		ClusterName: ptm.Cluster,
		Labels:      labels,
		Annotations: annotations,
	}, nil
}

// NewPodTemplateMetaFromKubeObjectMeta will create a new
// PodTemplateMeta object with the data from a provided kubernetes
// ObjectMeta object
func NewPodTemplateMetaFromKubeObjectMeta(obj interface{}) (*PodTemplateMeta, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(metav1.ObjectMeta{}):
		return fromKubeObjectMetaV1(obj.(metav1.ObjectMeta))
	case reflect.TypeOf(&metav1.ObjectMeta{}):
		o := obj.(*metav1.ObjectMeta)
		return fromKubeObjectMetaV1(*o)
	default:
		return nil, fmt.Errorf("unknown ObjectMeta version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeObjectMetaV1(meta metav1.ObjectMeta) (*PodTemplateMeta, error) {
	return &PodTemplateMeta{
		Name:        meta.Name,
		Namespace:   meta.Namespace,
		Cluster:     meta.ClusterName,
		Labels:      meta.Labels,
		Annotations: meta.Annotations,
	}, nil
}
