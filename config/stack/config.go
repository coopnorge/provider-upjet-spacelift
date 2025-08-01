package stack

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("spacelift_stack", func(r *config.Resource) {
		r.ShortGroup = "stack"
		r.References["space_id"] = config.Reference{
			TerraformName: "spacelift_space",
		}
	})
}
