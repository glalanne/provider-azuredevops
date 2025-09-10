package git_repository_file

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_git_repository_file", func(r *config.Resource) {
		r.ShortGroup = "repos"
	})
}
