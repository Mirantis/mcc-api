package v1alpha1

const (
	KaaSReleaseActiveLabel = "kaas.mirantis.com/active"
)

func (r *KaaSRelease) IsActive() bool {
	return r.Labels[KaaSReleaseActiveLabel] != ""
}

func (r *KaaSRelease) SetActive() {
	if r.Labels == nil {
		r.Labels = map[string]string{
			KaaSReleaseActiveLabel: "true",
		}
	} else {
		r.Labels[KaaSReleaseActiveLabel] = "true"
	}
}

func (r *KaaSRelease) UnsetActive() {
	if r.Labels != nil {
		delete(r.Labels, KaaSReleaseActiveLabel)
	}
}

func (r *KaaSRelease) IsClusterReleaseSupported(name, provider string) bool {
	for _, supportedRelease := range r.Spec.SupportedClusterReleases {
		if supportedRelease.Name == name {
			return supportedRelease.Providers.Contain(provider)
		}
	}
	return false
}

func (r *KaaSRelease) GetAvailableUpgrades(release string) []AvailableUpgrade {
	for _, supportedRelease := range r.Spec.SupportedClusterReleases {
		if supportedRelease.Name == release {
			return supportedRelease.AvailableUpgrades
		}
	}
	return nil
}

func (r *KaaSRelease) GetClusterReleaseTag(release string) string {
	for _, supportedRelease := range r.Spec.SupportedClusterReleases {
		if supportedRelease.Name == release {
			return supportedRelease.Tag
		}
	}
	return ""
}
