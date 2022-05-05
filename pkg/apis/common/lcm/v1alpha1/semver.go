package v1alpha1

import (
	"strings"

	"github.com/Masterminds/semver"
)

var (
	agentInMCC215 = semver.MustParse("0.3.0-132-g83a348fa")
	AgentInMCC216 = semver.MustParse("0.3.0-187-gba894556")
)

func AgentGreater132(lcmVersion string) bool {
	if lcmVersion == "" { //tests
		return true
	}
	agentVersion := semver.MustParse(strings.TrimLeft(lcmVersion, "v"))
	return agentVersion.GreaterThan(agentInMCC215)
}

func AgentGreater187(lcmVersion string) bool {
	if lcmVersion == "" { //tests
		return true
	}
	agentVersion := semver.MustParse(strings.TrimLeft(lcmVersion, "v"))
	return agentVersion.GreaterThan(AgentInMCC216)
}
