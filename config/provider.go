/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-vault/config/adsecretbackend"
	"github.com/upbound/provider-vault/config/adsecretbackendlibrary"
	"github.com/upbound/provider-vault/config/adsecretrole"
	"github.com/upbound/provider-vault/config/alicloudauthbackendrole"
	"github.com/upbound/provider-vault/config/approleauthbackendlogin"
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
		approleauthbackendrole.Configure,
		approleauthbackendrolesecretid.Configure,
		audit.Configure,
		auditrequestheader.Configure,
		authbackend.Configure,
		awsauthbackendcert.Configure,
		awsauthbackendclient.Configure,
		awsauthbackendconfigidentity.Configure,
		awsauthbackendidentitywhitelist.Configure,
		awsauthbackendlogin.Configure,
		awsauthbackendrole.Configure,
		awsauthbackendroletag.Configure,
		awsauthbackendroletagblacklist.Configure,
		awsauthbackendstsrole.Configure,
		awssecretbackend.Configure,
		awssecretbackendrole.Configure,
		azureauthbackendconfig.Configure,
		azureauthbackendrole.Configure,
		azuresecretbackend.Configure,
		azuresecretbackendrole.Configure,
		certauthbackendrole.Configure,
		consulsecretbackend.Configure,
		consulsecretbackendrole.Configure,
		databasesecretbackendconnection.Configure,
		databasesecretbackendrole.Configure,
		databasesecretbackendstaticrole.Configure,
		databasesecretsmount.Configure,
		egppolicy.Configure,
		gcpauthbackend.Configure,
		gcpauthbackendrole.Configure,
		gcpsecretbackend.Configure,
		gcpsecretimpersonatedaccount.Configure,
		gcpsecretroleset.Configure,
		gcpsecretstaticaccount.Configure,
		genericendpoint.Configure,
		genericsecret.Configure,
		githubauthbackend.Configure,
		githubteam.Configure,
		githubuser.Configure,
		identityentity.Configure,
		identityentityalias.Configure,
		identityentitypolicies.Configure,
		identitygroup.Configure,
		identitygroupalias.Configure,
		identitygroupmemberentityids.Configure,
		identitygroupmembergroupids.Configure,
		identitygrouppolicies.Configure,
		identitymfaduo.Configure,
		identitymfaloginenforcement.Configure,
		identitymfaokta.Configure,
		identitymfapingid.Configure,
		identitymfatotp.Configure,
		identityoidc.Configure,
		identityoidcassignment.Configure,
		identityoidcclient.Configure,
		identityoidckey.Configure,
		identityoidckeyallowedclientid.Configure,
		identityoidcprovider.Configure,
		identityoidcrole.Configure,
		identityoidcscope.Configure,
		jwtauthbackend.Configure,
		jwtauthbackendrole.Configure,
		kmipsecretbackend.Configure,
		kmipsecretrole.Configure,
		kmipsecretscope.Configure,
		kubernetesauthbackendconfig.Configure,
		kubernetesauthbackendrole.Configure,
		kubernetessecretbackend.Configure,
		kubernetessecretbackendrole.Configure,
		kvsecret.Configure,
		kvsecretbackendv2.Configure,
		kvsecretv2.Configure,
		ldapauthbackend.Configure,
		ldapauthbackendgroup.Configure,
		ldapauthbackenduser.Configure,
		managedkeys.Configure,
		mfaduo.Configure,
		mfaokta.Configure,
		mfapingid.Configure,
		mfatotp.Configure,
		mongodbatlassecretbackend.Configure,
		mongodbatlassecretrole.Configure,
		mount.Configure,
		namespace.Configure,
		nomadsecretbackend.Configure,
		nomadsecretrole.Configure,
		oktaauthbackend.Configure,
		oktaauthbackendgroup.Configure,
		oktaauthbackenduser.Configure,
		passwordpolicy.Configure,
		pkisecretbackendcert.Configure,
		pkisecretbackendconfigca.Configure,
		pkisecretbackendconfigurls.Configure,
		pkisecretbackendcrlconfig.Configure,
		pkisecretbackendintermediatecertrequest.Configure,
		pkisecretbackendintermediatesetsigned.Configure,
		pkisecretbackendrole.Configure,
		pkisecretbackendrootcert.Configure,
		pkisecretbackendrootsignintermediate.Configure,
		pkisecretbackendsign.Configure,
		policy.Configure,
		quotaleasecount.Configure,
		quotaratelimit.Configure,
		rabbitmqsecretbackend.Configure,
		rabbitmqsecretbackendrole.Configure,
		raftautopilot.Configure,
		raftsnapshotagentconfig.Configure,
		rgppolicy.Configure,
		sshsecretbackendca.Configure,
		sshsecretbackendrole.Configure,
		terraformcloudsecretbackend.Configure,
		terraformcloudsecretcreds.Configure,
		terraformcloudsecretrole.Configure,
		token.Configure,
		tokenauthbackendrole.Configure,
		transformalphabet.Configure,
		transformrole.Configure,
		transformtemplate.Configure,
		transformtransformation.Configure,
		transitsecretbackendcacheconfig.Configure,
		transitsecretbackendkey.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
