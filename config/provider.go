/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/coopnorge/provider-upjet-spacelift/config/context"
	"github.com/coopnorge/provider-upjet-spacelift/config/environmentvariable"
	"github.com/coopnorge/provider-upjet-spacelift/config/gcpserviceaccount"
	"github.com/coopnorge/provider-upjet-spacelift/config/module"
	"github.com/coopnorge/provider-upjet-spacelift/config/space"
	"github.com/coopnorge/provider-upjet-spacelift/config/stack"
)

const (
	resourcePrefix = "spacelift"
	modulePath     = "github.com/coopnorge/provider-upjet-spacelift"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("spacelift.crossplane.io"),
		ujconfig.WithShortName("spacelift"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithBasePackages(ujconfig.BasePackages{
			APIVersion: []string{
				"../v1alpha1",
				"../v1beta1",
			},
			Controller: []string{
				"../providerconfig",
			},
			ControllerMap: map[string]string{
				"../providerconfig": ujconfig.PackageNameConfig,
			},
		}),
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

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("spacelift.m.crossplane.io"),
		ujconfig.WithShortName("spacelift"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithBasePackages(ujconfig.BasePackages{
			APIVersion: []string{
				"../v1alpha1",
				"../v1beta1",
			},
			Controller: []string{
				"../providerconfig",
			},
			ControllerMap: map[string]string{
				"../providerconfig": ujconfig.PackageNameConfig,
			},
		}),
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
