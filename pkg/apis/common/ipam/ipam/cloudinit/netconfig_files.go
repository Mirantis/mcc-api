/*
Copyright © 2021 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudinit

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sort"
	"strings"

	"sigs.k8s.io/yaml"

	"github.com/Mirantis/mcc-api/pkg/apis/common/ipam/config"
	k8types "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

type NetconfigFile struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// NetconfigFiles -- store of net config files. Sorted by Path.
type NetconfigFiles []NetconfigFile

// StoredNetplan -- how to Netplan config stored in file
type StoredNetplan struct {
	Network *UserDataNetworkV2 `json:"network"`
}

// ----------------------------------------------------------------------------

func encodeNetplan(np *UserDataNetworkV2) string {
	buff, _ := yaml.Marshal(StoredNetplan{Network: np})
	return base64.StdEncoding.EncodeToString(buff)
}

func (in NetconfigFiles) Sort() {
	sort.Slice(in, func(i, j int) bool { return strings.ToLower(in[i].Path) < strings.ToLower(in[j].Path) })
}

// IsDifferent -- add new file to store or replace existing
func (in NetconfigFiles) IsDifferent(path string, content []byte) bool {
	for i := range in {
		if strings.EqualFold(in[i].Path, path) {
			return in[i].Content != base64.StdEncoding.EncodeToString(content)
		}
	}
	return true
}

// GetIndexOf -- return index for given path, -1 if not found
func (in NetconfigFiles) GetIndexOf(path string) int {
	for i := range in {
		if strings.EqualFold(in[i].Path, path) {
			return i
		}
	}
	return -1
}

// GetConentBase64For -- return content encoded to base64 for givn path
func (in NetconfigFiles) GetConentBase64For(path string) string {
	i := in.GetIndexOf(path)
	if i == -1 {
		return ""
	}
	return in[i].Content
}

// AppendOrReplace -- add new file to store or replace existing
func (in NetconfigFiles) AppendOrReplace(path string, content []byte) NetconfigFiles {
	i := in.GetIndexOf(path)
	encodedContent := base64.StdEncoding.EncodeToString(content)
	if i == -1 {
		in = append(in, NetconfigFile{
			Path:    path,
			Content: encodedContent,
		})
	} else {
		in[i].Content = encodedContent
	}
	in.Sort()
	return in
}

func (in NetconfigFiles) String() string {
	in.Sort()
	rv, _ := yaml.Marshal(in)
	return string(rv)
}

// AppendOrReplaceNetplan -- store Netplan config into NetconfigFiles storage
func (in NetconfigFiles) AppendOrReplaceNetplan(netplan *UserDataNetworkV2) NetconfigFiles {
	if netplan == nil {
		return in
	}
	buff, _ := yaml.Marshal(StoredNetplan{Network: netplan}) // able to skip error, because *UserDataNetworkV2 always serializable
	return in.AppendOrReplace(config.NetconfigNetplanPath, buff)
}

// GetNetplanBase64 -- return Base64-encoded Netplan config from NetconfigFiles storage or "" if not found
func (in NetconfigFiles) GetNetplanBase64() string {
	return in.GetConentBase64For(config.NetconfigNetplanPath)
}

// GetNetplan -- return Netplan config from NetconfigFiles storage or error if not found
func (in NetconfigFiles) GetNetplan() (*UserDataNetworkV2, error) {
	b64 := in.GetNetplanBase64()
	if b64 == "" {
		return nil, fmt.Errorf("Netplan config %w", k8types.ErrorNotFound) //nolint:stylecheck
	}

	netplanYaml, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("unable to decode base64 to Netplan YAML: %w", err)
	}

	storedNetplan := &StoredNetplan{}
	err = yaml.Unmarshal(netplanYaml, storedNetplan)
	if err != nil {
		return nil, fmt.Errorf("unable to decode YAML to Netplan: %w", err)
	}

	return storedNetplan.Network, nil
}

// GetNetplanSectionByIfname -- return Netplan config section from NetconfigFiles storage or error if not found
func (in NetconfigFiles) GetNetplanSectionByIfname(ifname string) (*SCBase, error) {
	netplan, err := in.GetNetplan()
	if err != nil {
		return nil, err
	}

	masterSections := []NetplanIfaceSectioner{netplan.Ethernets, netplan.Bridges, netplan.Bonds, netplan.Vlans}
	for i := range masterSections {
		rv, err := masterSections[i].GetNetplanSectionByIfname(ifname)
		if err == nil || !errors.Is(err, k8types.ErrorNotFound) {
			return rv, err
		}
	}
	return nil, k8types.ErrorNotFound
}

// IsNetplanConfigBase64Different -- compare stored Netplan config with given base64 encoded Netplan
func (in NetconfigFiles) IsNetplanConfigBase64Different(b64 string) bool {
	return b64 != in.GetNetplanBase64()
}

// IsNetplanConfigDifferent -- compare stored Netplan config with given Netplan
func (in NetconfigFiles) IsNetplanConfigDifferent(netplan *UserDataNetworkV2) bool {
	return in.IsNetplanConfigBase64Different(encodeNetplan(netplan))
}
