package check_business_hours

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_check_business_hours", func(r *config.Resource) {
		r.ShortGroup = "repos"
	})
}
