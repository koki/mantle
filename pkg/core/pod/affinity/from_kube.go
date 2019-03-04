package affinity

import (
	"fmt"
	"reflect"

	"k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAffinityFromKubeAffinity will create a new
// Affinity object with the data from a provided kubernetes
// Affinity object
func NewAffinityFromKubeAffinity(obj interface{}) (*Affinity, error) {
	switch reflect.TypeOf(obj) {
	case reflect.TypeOf(v1.Affinity{}):
		o := obj.(v1.Affinity)
		return fromKubeAffinityV1(&o)
	case reflect.TypeOf(&v1.Affinity{}):
		return fromKubeAffinityV1(obj.(*v1.Affinity))
	default:
		return nil, fmt.Errorf("unknown Affinity version: %s", reflect.TypeOf(obj))
	}
}

func fromKubeAffinityV1(kubeAffinity *v1.Affinity) (*Affinity, error) {
	if kubeAffinity == nil {
		return nil, nil
	}

	return &Affinity{
		NodeAffinity:    fromKubeNodeAffinityV1(kubeAffinity.NodeAffinity),
		PodAffinity:     fromKubePodAffinityV1(kubeAffinity.PodAffinity),
		PodAntiAffinity: fromKubePodAntiAffinityV1(kubeAffinity.PodAntiAffinity),
	}, nil
}

func fromKubeNodeAffinityV1(kubeNA *v1.NodeAffinity) map[AffinityType][]NodeTerm {
	var nodeAffinity map[AffinityType][]NodeTerm

	if kubeNA != nil {
		nodeAffinity = make(map[AffinityType][]NodeTerm)

		if kubeNA.RequiredDuringSchedulingIgnoredDuringExecution != nil {
			nodeAffinity[AffinityHard] = fromKubeNodeSelectorV1(kubeNA.RequiredDuringSchedulingIgnoredDuringExecution)
		}

		if len(kubeNA.PreferredDuringSchedulingIgnoredDuringExecution) > 0 {
			nodeAffinity[AffinitySoft] = fromKubePreferredSchedulingTermsV1(kubeNA.PreferredDuringSchedulingIgnoredDuringExecution)
		}
	}

	return nodeAffinity
}

func fromKubePodAffinityV1(pa *v1.PodAffinity) map[AffinityType][]PodTerm {
	var podAffinity map[AffinityType][]PodTerm

	if pa != nil {
		podAffinity = make(map[AffinityType][]PodTerm)

		if len(pa.RequiredDuringSchedulingIgnoredDuringExecution) > 0 {
			podAffinity[AffinityHard] = fromKubePodAffinityTermsV1(pa.RequiredDuringSchedulingIgnoredDuringExecution)
		}

		if len(pa.PreferredDuringSchedulingIgnoredDuringExecution) > 0 {
			podAffinity[AffinitySoft] = fromKubeWeightedPodAffinityTermsV1(pa.PreferredDuringSchedulingIgnoredDuringExecution)
		}
	}

	return podAffinity
}

func fromKubePodAntiAffinityV1(paa *v1.PodAntiAffinity) map[AffinityType][]PodTerm {
	var podAntiAffinity map[AffinityType][]PodTerm

	if paa != nil {
		podAntiAffinity = make(map[AffinityType][]PodTerm)

		if len(paa.RequiredDuringSchedulingIgnoredDuringExecution) > 0 {
			podAntiAffinity[AffinityHard] = fromKubePodAffinityTermsV1(paa.RequiredDuringSchedulingIgnoredDuringExecution)
		}

		if len(paa.PreferredDuringSchedulingIgnoredDuringExecution) > 0 {
			podAntiAffinity[AffinitySoft] = fromKubeWeightedPodAffinityTermsV1(paa.PreferredDuringSchedulingIgnoredDuringExecution)
		}
	}

	return podAntiAffinity
}

func fromKubeNodeSelectorV1(ns *v1.NodeSelector) []NodeTerm {
	terms := []NodeTerm{}

	if ns != nil {
		for _, nodeTerm := range ns.NodeSelectorTerms {
			term := fromKubeNodeSelectorTermV1(nodeTerm)
			terms = append(terms, term)
		}
	}

	return terms
}

func fromKubePodAffinityTermsV1(pat []v1.PodAffinityTerm) []PodTerm {
	terms := []PodTerm{}

	for _, podTerm := range pat {
		term := fromKubePodAffinityTermV1(podTerm)
		terms = append(terms, term)
	}

	return terms
}

func fromKubePreferredSchedulingTermsV1(prefered []v1.PreferredSchedulingTerm) []NodeTerm {
	nodeTerms := []NodeTerm{}

	for _, term := range prefered {
		nodeTerm := fromKubeNodeSelectorTermV1(term.Preference)
		nodeTerm.Weight = term.Weight
		nodeTerms = append(nodeTerms, nodeTerm)
	}

	return nodeTerms
}

func fromKubeWeightedPodAffinityTermsV1(wpat []v1.WeightedPodAffinityTerm) []PodTerm {
	terms := []PodTerm{}

	for _, podTerm := range wpat {
		term := fromKubePodAffinityTermV1(podTerm.PodAffinityTerm)
		term.Weight = podTerm.Weight
		terms = append(terms, term)
	}

	return terms
}

func fromKubeNodeSelectorTermV1(term v1.NodeSelectorTerm) NodeTerm {
	return NodeTerm{
		Expressions: fromKubeNodeSelectorRequirementV1(term.MatchExpressions),
		Fields:      fromKubeNodeSelectorRequirementV1(term.MatchFields),
	}
}

func fromKubePodAffinityTermV1(term v1.PodAffinityTerm) PodTerm {
	return PodTerm{
		Selector:   fromKubeLabelSelectorV1(term.LabelSelector),
		Topology:   term.TopologyKey,
		Namespaces: term.Namespaces,
	}
}

func fromKubeNodeSelectorRequirementV1(reqs []v1.NodeSelectorRequirement) []NodeExpression {
	expressions := []NodeExpression{}

	for _, req := range reqs {
		expression := NodeExpression{
			Key:    req.Key,
			Values: req.Values,
		}

		switch req.Operator {
		case v1.NodeSelectorOpIn:
			expression.Op = NodeOperatorIn
		case v1.NodeSelectorOpNotIn:
			expression.Op = NodeOperatorNotIn
		case v1.NodeSelectorOpExists:
			expression.Op = NodeOperatorExists
		case v1.NodeSelectorOpDoesNotExist:
			expression.Op = NodeOperatorDoesNotExist
		case v1.NodeSelectorOpGt:
			expression.Op = NodeOperatorGt
		case v1.NodeSelectorOpLt:
			expression.Op = NodeOperatorLt
		}

		expressions = append(expressions, expression)
	}

	return expressions
}

func fromKubeLabelSelectorV1(selector *metav1.LabelSelector) Selector {
	if selector != nil {
		return Selector{
			Labels:      selector.MatchLabels,
			Expressions: fromKubeLabelSelectorRequirementV1(selector.MatchExpressions),
		}
	}

	return Selector{}
}

func fromKubeLabelSelectorRequirementV1(lsr []metav1.LabelSelectorRequirement) []SelectorExpression {
	expressions := []SelectorExpression{}

	for _, req := range lsr {
		expression := SelectorExpression{
			Key:    req.Key,
			Values: req.Values,
		}

		switch req.Operator {
		case metav1.LabelSelectorOpIn:
			expression.Op = SelectorOperatorIn
		case metav1.LabelSelectorOpNotIn:
			expression.Op = SelectorOperatorNotIn
		case metav1.LabelSelectorOpExists:
			expression.Op = SelectorOperatorExists
		case metav1.LabelSelectorOpDoesNotExist:
			expression.Op = SelectorOperatorDoesNotExist
		}

		expressions = append(expressions, expression)
	}

	return expressions
}
