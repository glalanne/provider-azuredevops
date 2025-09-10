package serviceendpoint_checkmarx_sca

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_serviceendpoint_checkmarx_sca", func(r *config.Resource) {
		r.ShortGroup = "serviceendpoint"
	})
}
