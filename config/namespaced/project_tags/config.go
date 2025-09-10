package project_tags

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_project_tags", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
}
