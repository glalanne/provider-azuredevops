package repository_policy_case_enforcement

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_repository_policy_case_enforcement", func(r *config.Resource) {
		r.ShortGroup = "repos"
	})
}
