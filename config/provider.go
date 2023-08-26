/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/spot/provider-spot/config/aksnp"
	"github.com/spot/provider-spot/config/aksnpvng"
	"github.com/spot/provider-spot/config/oceanaws"
	"github.com/spot/provider-spot/config/oceanawsvng"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "spot"
	modulePath     = "github.com/spot/provider-spot"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		//null.Configure,
		oceanaws.Configure,
		aksnp.Configure,
		aksnpvng.Configure,
		oceanawsvng.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
