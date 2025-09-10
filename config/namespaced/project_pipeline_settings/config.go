package project_pipeline_settings

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_project_pipeline_settings", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
}
