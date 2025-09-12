package serviceendpoint_jfrog_distribution_v2

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/glalanne/provider-azuredevops/config/namespaced/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_serviceendpoint_jfrog_distribution_v2", func(r *config.Resource) {
		r.ShortGroup = "serviceendpoint"
		r.ExternalName.GetExternalNameFn = common.GetNameFromProjectID
	})
}
