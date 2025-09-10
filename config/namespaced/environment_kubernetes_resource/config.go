package environment_kubernetes_resource

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_environment_kubernetes_resource", func(r *config.Resource) {
		r.ShortGroup = "pipelines"
	})
}
