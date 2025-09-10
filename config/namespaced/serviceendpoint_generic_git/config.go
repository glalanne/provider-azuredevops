package serviceendpoint_generic_git

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_serviceendpoint_generic_git", func(r *config.Resource) {
		r.ShortGroup = "serviceendpoint"
	})
}
