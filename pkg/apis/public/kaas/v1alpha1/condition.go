package v1alpha1

type ConditionType string

type Condition struct {
	Type    ConditionType `json:"type"`
	Ready   bool          `json:"ready"`
	Message string        `json:"message"`
}

type ConditionsSummary struct {
	Ready      bool        `json:"ready"`
	Conditions []Condition `json:"conditions,omitempty"`
}

func GetConditionsSummary(conditions []Condition) ConditionsSummary {
	cs := ConditionsSummary{
		Conditions: conditions,
	}
	for _, condition := range cs.Conditions {
		if !condition.Ready {
			return cs
		}
	}
	cs.Ready = true
	return cs
}

func GetConditionFromSummary(cs ConditionsSummary, ctype ConditionType) *Condition {
	for i := range cs.Conditions {
		if cs.Conditions[i].Type == ctype {
			return &cs.Conditions[i]
		}
	}
	return nil
}
