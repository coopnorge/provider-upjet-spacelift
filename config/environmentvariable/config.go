package environmentvariable

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("spacelift_environment_variable", func(r *config.Resource) {
		r.ShortGroup = "environmentvariable"
		r.References["context_id"] = config.Reference{
			TerraformName: "spacelift_context",
		}
	})
}
