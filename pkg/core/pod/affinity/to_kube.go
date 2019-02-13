package affinity

import (
	"fmt"
	"strings"

	"k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ToKube will return a kubernetes container object of the api version provided
func (a *Affinity) ToKube(version string) (interface{}, error) {
	switch strings.ToLower(version) {
	case "v1":
		return a.toKubeV1()
	case "":
		return a.toKubeV1()
	default:
		return nil, fmt.Errorf("unsupported api version for Affinity: %s", version)
	}
}

func (a *Affinity) toKubeV1() (*v1.Affinity, error) {
	nodeAffinity := a.toKubeNodeAffinityV1(a.NodeAffinity)
	podAffinity := a.toKubePodAffinityV1(a.PodAffinity)
	podAntiAffinity := a.toKubePodAntiAffinityV1(a.PodAntiAffinity)

	if podAffinity != nil || podAntiAffinity != nil || nodeAffinity != nil {
		return &v1.Affinity{
			NodeAffinity:    nodeAffinity,
			PodAffinity:     podAffinity,
			PodAntiAffinity: podAntiAffinity,
		}, nil
	}

	return nil, nil
}

func (a *Affinity) toKubePodAffinityV1(affinities map[AffinityType][]PodTerm) *v1.PodAffinity {
	hard, soft := a.toKubeHardSoftPodAffinityTermsV1(affinities)

	if len(hard) == 0 {
		hard = nil
	}

	if len(soft) == 0 {
		soft = nil
	}

	if len(hard) > 0 || len(soft) > 0 {
		return &v1.PodAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution:  hard,
			PreferredDuringSchedulingIgnoredDuringExecution: soft,
		}
	}

	return nil
}

func (a *Affinity) toKubePodAntiAffinityV1(affinities map[AffinityType][]PodTerm) *v1.PodAntiAffinity {
	hard, soft := a.toKubeHardSoftPodAffinityTermsV1(affinities)

	if len(hard) == 0 {
		hard = nil
	}

	if len(soft) == 0 {
		soft = nil
	}

	if len(hard) > 0 || len(soft) > 0 {
		return &v1.PodAntiAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution:  hard,
			PreferredDuringSchedulingIgnoredDuringExecution: soft,
		}
	}

	return nil
}

func (a *Affinity) toKubeNodeAffinityV1(affinities map[AffinityType][]NodeTerm) *v1.NodeAffinity {
	hard, soft := a.toKubeHardSoftNodeAffinityTermsV1(affinities)

	if len(soft) == 0 {
		soft = nil
	}

	if hard != nil || len(soft) > 0 {
		return &v1.NodeAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution:  hard,
			PreferredDuringSchedulingIgnoredDuringExecution: soft,
		}
	}

	return nil
}

func (a *Affinity) toKubeHardSoftPodAffinityTermsV1(affinities map[AffinityType][]PodTerm) ([]v1.PodAffinityTerm, []v1.WeightedPodAffinityTerm) {
	var hard []v1.PodAffinityTerm
	var soft []v1.WeightedPodAffinityTerm

	for t, list := range affinities {
		switch t {
		case AffinityHard:
			hard = a.toKubeHardPodAffinityTermsV1(list)

		case AffinitySoft:
			soft = a.toKubeSoftPodAffinityTermsV1(list)
		}
	}

	return hard, soft
}

func (a *Affinity) toKubeHardSoftNodeAffinityTermsV1(affinities map[AffinityType][]NodeTerm) (*v1.NodeSelector, []v1.PreferredSchedulingTerm) {
	var hard *v1.NodeSelector
	var soft []v1.PreferredSchedulingTerm

	for t, list := range affinities {
		switch t {
		case AffinityHard:
			hard = a.toKubeHardNodeAffinityTermsV1(list)

		case AffinitySoft:
			soft = a.toKubeSoftNodeAffinityTermsV1(list)
		}
	}

	return hard, soft
}

func (a *Affinity) toKubeHardPodAffinityTermsV1(terms []PodTerm) []v1.PodAffinityTerm {
	var hard []v1.PodAffinityTerm

	for _, term := range terms {
		hardTerm := a.toKubePodAffinityTermV1(term)
		hard = append(hard, hardTerm)
	}

	return hard
}

func (a *Affinity) toKubeHardNodeAffinityTermsV1(terms []NodeTerm) *v1.NodeSelector {
	var nstList []v1.NodeSelectorTerm

	for _, term := range terms {
		nst := a.toKubeNodeSelectorTermV1(term)
		nstList = append(nstList, nst)
	}

	if len(nstList) > 0 {
		return &v1.NodeSelector{
			NodeSelectorTerms: nstList,
		}
	}

	return nil
}

func (a *Affinity) toKubeSoftPodAffinityTermsV1(terms []PodTerm) []v1.WeightedPodAffinityTerm {
	var soft []v1.WeightedPodAffinityTerm

	for _, term := range terms {
		softTerm := v1.WeightedPodAffinityTerm{}
		affinityTerm := a.toKubePodAffinityTermV1(term)
		softTerm.Weight = term.Weight
		softTerm.PodAffinityTerm = affinityTerm
		soft = append(soft, softTerm)
	}

	return soft
}

func (a *Affinity) toKubeSoftNodeAffinityTermsV1(terms []NodeTerm) []v1.PreferredSchedulingTerm {
	var soft []v1.PreferredSchedulingTerm

	for _, term := range terms {
		softTerm := v1.PreferredSchedulingTerm{}
		nodeSelectorTerm := a.toKubeNodeSelectorTermV1(term)
		softTerm.Weight = term.Weight
		softTerm.Preference = nodeSelectorTerm
		soft = append(soft, softTerm)
	}

	return soft
}

func (a *Affinity) toKubePodAffinityTermV1(term PodTerm) v1.PodAffinityTerm {
	selector := metav1.LabelSelector{
		MatchLabels:      term.Selector.Labels,
		MatchExpressions: a.toKubeLabelSelectorRequirementV1(term.Selector.Expressions),
	}

	podAffinityTerm := v1.PodAffinityTerm{
		Namespaces:    term.Namespaces,
		TopologyKey:   term.Topology,
		LabelSelector: &selector,
	}

	return podAffinityTerm
}

func (a *Affinity) toKubeNodeSelectorTermV1(term NodeTerm) v1.NodeSelectorTerm {
	nodeSelectorTerm := v1.NodeSelectorTerm{
		MatchExpressions: a.toKubeNodeSelectorRequirementV1(term.Expressions),
		MatchFields:      a.toKubeNodeSelectorRequirementV1(term.Fields),
	}

	return nodeSelectorTerm
}

func (a *Affinity) toKubeLabelSelectorRequirementV1(expressions []SelectorExpressions) []metav1.LabelSelectorRequirement {
	requirements := []metav1.LabelSelectorRequirement{}

	for _, e := range expressions {
		expression := metav1.LabelSelectorRequirement{
			Key:    e.Key,
			Values: e.Values,
		}

		switch e.Op {
		case SelectorOperatorIn:
			expression.Operator = metav1.LabelSelectorOpIn
		case SelectorOperatorNotIn:
			expression.Operator = metav1.LabelSelectorOpNotIn
		case SelectorOperatorExists:
			expression.Operator = metav1.LabelSelectorOpExists
		case SelectorOperatorDoesNotExist:
			expression.Operator = metav1.LabelSelectorOpDoesNotExist
		}

		requirements = append(requirements, expression)
	}

	return requirements
}

func (a *Affinity) toKubeNodeSelectorRequirementV1(expressions []NodeExpressions) []v1.NodeSelectorRequirement {
	requirements := []v1.NodeSelectorRequirement{}

	for _, e := range expressions {
		expression := v1.NodeSelectorRequirement{
			Key:    e.Key,
			Values: e.Values,
		}

		switch e.Op {
		case NodeOperatorIn:
			expression.Operator = v1.NodeSelectorOpIn
		case NodeOperatorNotIn:
			expression.Operator = v1.NodeSelectorOpNotIn
		case NodeOperatorExists:
			expression.Operator = v1.NodeSelectorOpExists
		case NodeOperatorDoesNotExist:
			expression.Operator = v1.NodeSelectorOpDoesNotExist
		case NodeOperatorGt:
			expression.Operator = v1.NodeSelectorOpGt
		case NodeOperatorLt:
			expression.Operator = v1.NodeSelectorOpLt
		}

		requirements = append(requirements, expression)
	}

	return requirements
}
