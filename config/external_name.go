/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// TerraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the no-fork
// architecture for this provider.
var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"azuredevops_agent_pool":                                        config.IdentifierFromProvider,
	}

var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := TerraformPluginSDKExternalNameConfigs[r.Name]; ok {

			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"

		i++
	}
	return l
}

// ResourceConfigurator applies all external name configs
// listed in the table TerraformPluginSDKExternalNameConfigs and
// CLIReconciledExternalNameConfigs and sets the version
// of those resources to v1beta1. For those resource in
// TerraformPluginSDKExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := TerraformPluginSDKExternalNameConfigs[r.Name]
		if !configured {
			e, configured = CLIReconciledExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}
		// r.Version = "v1alpha1"
		r.ExternalName = e
	}
}
