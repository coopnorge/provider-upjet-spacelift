/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/aksel-allas-org/provider-spacelift/config/context"
	"github.com/aksel-allas-org/provider-spacelift/config/environmentvariable"
	"github.com/aksel-allas-org/provider-spacelift/config/gcpserviceaccount"
	"github.com/aksel-allas-org/provider-spacelift/config/module"
	"github.com/aksel-allas-org/provider-spacelift/config/space"
	"github.com/aksel-allas-org/provider-spacelift/config/stack"
)

const (
	resourcePrefix = "spacelift"
	modulePath     = "github.com/aksel-allas-org/provider-spacelift"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("spacelift.upbound.io"),
		ujconfig.WithShortName("spacelift"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		module.Configure,
		stack.Configure,
		space.Configure,
		context.Configure,
		environmentvariable.Configure,
		gcpserviceaccount.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
