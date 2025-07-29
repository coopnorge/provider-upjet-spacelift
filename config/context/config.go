package context

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("spacelift_context", func(r *config.Resource) {
		r.ShortGroup = "context"
		r.References["space_id"] = config.Reference{
			TerraformName: "spacelift_space",
		}
	})

	p.AddResourceConfigurator("spacelift_context_attachment", func(r *config.Resource) {
		r.ShortGroup = "context"
		r.References["stack_id"] = config.Reference{
			TerraformName: "spacelift_stack",
		}
		r.References["context_id"] = config.Reference{
			TerraformName: "spacelift_context",
		}
	})
}
