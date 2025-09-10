package extension

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_extension", func(r *config.Resource) {
		r.ShortGroup = "core"
	})
}
