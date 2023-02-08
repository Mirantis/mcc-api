package v1alpha1

// +gocode:public-api=true
type Condition struct {
	Type    ConditionType `json:"type"`
	Ready   bool          `json:"ready"`
	Message string        `json:"message"`
}
type ConditionType string

// +gocode:public-api=true
type ConditionsSummary struct {
	Ready      bool        `json:"ready"`
	Conditions []Condition `json:"conditions,omitempty"`
}

func (cs ConditionsSummary) IsEqualTo(cs2 ConditionsSummary) bool {
	if cs.Ready != cs2.Ready || len(cs.Conditions) != len(cs2.Conditions) {
		return false
	}
	for i, condition := range cs.Conditions {
		if condition != cs2.Conditions[i] {
			return false
		}
	}
	return true
}

// +gocode:public-api=true
func GetConditionsSummary(conditions []Condition) ConditionsSummary {
	cs := ConditionsSummary{
		Conditions: conditions,
	}
	for _, condition := range cs.Conditions {
		isRebootCondition := condition.Type == RebootCondition ||
			condition.Type == RebootMachinesCondition

		if !condition.Ready && !isRebootCondition {
			return cs
		}
	}
	cs.Ready = true
	return cs
}

// +gocode:public-api=true
func GetConditionFromSummary(cs ConditionsSummary, ctype ConditionType) *Condition {
	for i := range cs.Conditions {
		if cs.Conditions[i].Type == ctype {
			return &cs.Conditions[i]
		}
	}
	return nil
}
