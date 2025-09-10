package group_entitlement

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_group_entitlement", func(r *config.Resource) {
		r.ShortGroup = "security"
	})
}
