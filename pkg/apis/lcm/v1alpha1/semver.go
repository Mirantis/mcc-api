package v1alpha1

import (
	"github.com/Masterminds/semver"
	"strings"
)

var (
	// +gocode:public-api=true
	agentInMCC215 = semver.MustParse("0.3.0-132-g83a348fa")
	// +gocode:public-api=true
	AgentInMCC216 = semver.MustParse("0.3.0-187-gba894556")
	// +gocode:public-api=true
	AgentInMCC219 = semver.MustParse("v0.3.0-257-ga93244da")
)

// +gocode:public-api=true
func AgentGreater132(lcmVersion string) bool {
	if lcmVersion == "" {
		return true
	}
	agentVersion := semver.MustParse(strings.TrimLeft(lcmVersion, "v"))
	return agentVersion.GreaterThan(agentInMCC215)
}

// AgentGreater187 MCC ver greater than 2.16
// +gocode:public-api=true
func AgentGreater187(lcmVersion string) bool {
	if lcmVersion == "" {
		return true
	}
	agentVersion := semver.MustParse(strings.TrimLeft(lcmVersion, "v"))
	return agentVersion.GreaterThan(AgentInMCC216)
}

// +gocode:public-api=true
func AgentGreater257(lcmVersion string) bool {
	if lcmVersion == "" {
		return true
	}
	agentVersion := semver.MustParse(strings.TrimLeft(lcmVersion, "v"))
	return agentVersion.GreaterThan(AgentInMCC219)
}
