package repository_policy_reserved_names

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_repository_policy_reserved_names", func(r *config.Resource) {
		r.ShortGroup = "repos"
	})
}
