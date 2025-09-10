package branch_policy_status_check

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_branch_policy_status_check", func(r *config.Resource) {
		r.ShortGroup = "repos"
	})
}
