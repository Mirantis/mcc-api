package lib

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gophercloud/utils/openstack/clientconfig"
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"

	"github.com/Mirantis/mcc-api/pkg/errors"
)

func GetOsCloudConfig(osConfigPath, osCloud string) (*clientconfig.Cloud, error) {
	if osConfigPath == "" {
		klog.Info("os-config-path is not specified, using default clouds.yaml")
		cloudsMap, err := clientconfig.LoadCloudsYAML()
		if err != nil {
			klog.Error("clouds.yaml not found")
			return nil, errors.Wrap(err, "clouds.yaml not found")
		}
		cloud, ok := cloudsMap[osCloud]
		if !ok {
			return nil, errors.Errorf("failed to find %v cloud in OS config", osCloud)
		}
		return &cloud, nil
	}
	if _, err := os.Stat(osConfigPath); err != nil {
		return nil, errors.Wrapf(err, "error get cloud config path %v", osConfigPath)
	}
	absOsConfigPath, err := filepath.Abs(osConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "error get cloud config abs path %v", osConfigPath)
	}
	osConfigData, err := ioutil.ReadFile(absOsConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "error read cloud config abs path %v file", osConfigPath)
	}
	var clouds clientconfig.Clouds
	err = yaml.Unmarshal(osConfigData, &clouds)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal OS config yaml")
	}
	cloud, ok := clouds.Clouds[osCloud]
	if !ok {
		return nil, errors.Errorf("failed to find %v cloud in OS config", osCloud)
	}
	return &cloud, nil
}
