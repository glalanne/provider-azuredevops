package wiki_page

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_wiki_page", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
}
