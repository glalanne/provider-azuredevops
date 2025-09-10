package check_rest_api

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_check_rest_api", func(r *config.Resource) {
		r.ShortGroup = "repos"
		r.References["project_id"] = config.Reference{
			TerraformName: "azuredevops_project",
		}
		delete(r.References, "connected_service_name")

	})
}
