package v1alpha1

import (
	"encoding/json"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/apis/util/common/helmutil"
	"github.com/pkg/errors"
)

// +gocode:public-api=true
func MetallbAddressPoolFromValues(values interface{}) (kaasv1alpha1.MetallbAddressPool, error) {
	var addressPool kaasv1alpha1.MetallbAddressPool
	str, err := json.Marshal(values)
	if err != nil {
		return addressPool, errors.Wrap(err, "failed to json.Marshal")
	}
	err = json.Unmarshal(str, &addressPool)
	if err != nil {
		return addressPool, errors.Wrap(err, "failed to json.Unmarshal")
	}
	return addressPool, nil
}

// +gocode:public-api=true
func MetallbAddressPoolsFromChartValues(values helmutil.Values) ([]kaasv1alpha1.MetallbAddressPool, error) {
	pools := []kaasv1alpha1.MetallbAddressPool{}

	curAddressPools, exists, err := helmutil.NestedSlice(values, "configInline", "address-pools")
	if err != nil {
		return pools, err
	}
	if !exists {
		return pools, nil
	}

	for _, uCurAddressPool := range curAddressPools {
		curAddressPool, err := MetallbAddressPoolFromValues(uCurAddressPool)
		if err != nil {
			return pools, errors.Wrapf(err, "Failed to convert chart values %v to MetallbAddressPool", uCurAddressPool)
		}
		pools = append(pools, curAddressPool)
	}

	return pools, nil
}

// +gocode:public-api=true
func MetallbAddressPoolsCheckDuplicates(pools []kaasv1alpha1.MetallbAddressPool) (duplicatedNames []string, duplicatesFound bool) {
	seenNames := map[string]int{}
	for _, pool := range pools {
		if _, exists := seenNames[pool.Name]; exists {
			seenNames[pool.Name]++
			duplicatesFound = true
		} else {
			seenNames[pool.Name] = 1
		}
	}
	if duplicatesFound {
		for name, count := range seenNames {
			if count > 1 {
				duplicatedNames = append(duplicatedNames, name)
			}
		}
	}
	return duplicatedNames, duplicatesFound
}
