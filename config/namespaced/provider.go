// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package namespaced

import (
	"github.com/glalanne/provider-azuredevops/config/namespaced/agent_pool"
)

func init() {
	ProviderConfiguration.AddConfig(agent_pool.Configure)
}
