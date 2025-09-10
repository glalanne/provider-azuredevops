package feed_retention_policy

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_feed_retention_policy", func(r *config.Resource) {
		r.ShortGroup = "artifacts"
	})
}
