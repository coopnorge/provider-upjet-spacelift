package space

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("spacelift_space", func(r *config.Resource) {
		r.ShortGroup = "space"
	})
}
