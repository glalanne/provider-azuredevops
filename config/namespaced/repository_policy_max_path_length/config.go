package repository_policy_max_path_length

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_repository_policy_max_path_length", func(r *config.Resource) {
		r.ShortGroup = "repos"
	})
}
