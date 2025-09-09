// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cluster

import (
	"github.com/glalanne/provider-azuredevops/config/cluster/agent_pool"
)

func init() {
	ProviderConfiguration.AddConfig(agent_pool.Configure)
}
