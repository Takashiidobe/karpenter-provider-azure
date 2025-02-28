/*
Portions Copyright (c) Microsoft Corporation.

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

package test

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/samber/lo"

	azsettings "github.com/Azure/karpenter-provider-azure/pkg/apis/settings"
)

type SettingOptions struct {
	ClusterName                    *string
	ClusterEndpoint                *string
	ClusterID                      *string
	KubeletClientTLSBootstrapToken *string
	SSHPublicKey                   *string
	NetworkPlugin                  *string
	NetworkPolicy                  *string
	VMMemoryOverheadPercent        *float64
	NodeIdentities                 []string
	Tags                           map[string]string
}

func Settings(overrides ...SettingOptions) *azsettings.Settings {
	options := SettingOptions{}
	for _, override := range overrides {
		if err := mergo.Merge(&options, override, mergo.WithOverride); err != nil {
			panic(fmt.Sprintf("Failed to merge settings: %s", err))
		}
	}
	return &azsettings.Settings{
		ClusterName:                    lo.FromPtrOr(options.ClusterName, "test-cluster"),
		ClusterEndpoint:                lo.FromPtrOr(options.ClusterEndpoint, "https://test-cluster"),
		ClusterID:                      lo.FromPtrOr(options.ClusterID, "00000000"),
		KubeletClientTLSBootstrapToken: lo.FromPtrOr(options.KubeletClientTLSBootstrapToken, "test-token"),
		SSHPublicKey:                   lo.FromPtrOr(options.SSHPublicKey, "test-ssh-public-key"),
		NetworkPlugin:                  lo.FromPtrOr(options.NetworkPlugin, "kubenet"),
		NetworkPolicy:                  lo.FromPtrOr(options.NetworkPolicy, ""),
		VMMemoryOverheadPercent:        lo.FromPtrOr(options.VMMemoryOverheadPercent, 0.075),
		NodeIdentities:                 options.NodeIdentities,
		Tags:                           options.Tags,
	}
}
