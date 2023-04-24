/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/humoflife/provider-vault/config/adsecretbackend"
	"github.com/humoflife/provider-vault/config/adsecretbackendlibrary"
	"github.com/humoflife/provider-vault/config/adsecretrole"
	"github.com/humoflife/provider-vault/config/alicloudauthbackendrole"
	"github.com/humoflife/provider-vault/config/approleauthbackendlogin"
)

const (
	resourcePrefix = "vault"
	modulePath     = "github.com/upbound/provider-vault"
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
		adsecretbackend.Configure,
		adsecretbackendlibrary.Configure,
		adsecretrole.Configure,
		alicloudauthbackendrole.Configure,
		approleauthbackendlogin.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
