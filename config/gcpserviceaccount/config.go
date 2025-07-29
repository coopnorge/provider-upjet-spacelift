package gcpserviceaccount

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("spacelift_gcp_service_account", func(r *config.Resource) {
		r.ShortGroup = "gcpserviceaccount"
		r.Kind = "GcpServiceAccount"
		r.References["stack_id"] = config.Reference{
			TerraformName: "spacelift_stack",
		}
	})
}
