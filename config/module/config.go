package module

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("spacelift_module", func(r *config.Resource) {
		r.ShortGroup = "module"
		r.References["space_id"] = config.Reference{
			TerraformName: "spacelift_space",
		}
	})
}
