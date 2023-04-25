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
	"github.com/upbound/provider-vault/config/approleauthbackendrole"
	"github.com/upbound/provider-vault/config/approleauthbackendrolesecretid"
	"github.com/upbound/provider-vault/config/audit"
	"github.com/upbound/provider-vault/config/auditrequestheader"
	"github.com/upbound/provider-vault/config/authbackend"
	"github.com/upbound/provider-vault/config/awsauthbackendcert"
	"github.com/upbound/provider-vault/config/awsauthbackendclient"
	"github.com/upbound/provider-vault/config/awsauthbackendconfigidentity"
	"github.com/upbound/provider-vault/config/awsauthbackendidentitywhitelist"
	"github.com/upbound/provider-vault/config/awsauthbackendlogin"
	"github.com/upbound/provider-vault/config/awsauthbackendrole"
	"github.com/upbound/provider-vault/config/awsauthbackendroletag"
	"github.com/upbound/provider-vault/config/awsauthbackendroletagblacklist"
	"github.com/upbound/provider-vault/config/awsauthbackendstsrole"
	"github.com/upbound/provider-vault/config/awssecretbackend"
	"github.com/upbound/provider-vault/config/awssecretbackendrole"
	"github.com/upbound/provider-vault/config/azureauthbackendconfig"
	"github.com/upbound/provider-vault/config/azureauthbackendrole"
	"github.com/upbound/provider-vault/config/azuresecretbackend"
	"github.com/upbound/provider-vault/config/azuresecretbackendrole"
	"github.com/upbound/provider-vault/config/certauthbackendrole"
	"github.com/upbound/provider-vault/config/consulsecretbackend"
	"github.com/upbound/provider-vault/config/consulsecretbackendrole"
	"github.com/upbound/provider-vault/config/databasesecretbackendconnection"
	"github.com/upbound/provider-vault/config/databasesecretbackendrole"
	"github.com/upbound/provider-vault/config/databasesecretbackendstaticrole"
	"github.com/upbound/provider-vault/config/databasesecretsmount"
	"github.com/upbound/provider-vault/config/egppolicy"
	"github.com/upbound/provider-vault/config/gcpauthbackend"
	"github.com/upbound/provider-vault/config/gcpauthbackendrole"
	"github.com/upbound/provider-vault/config/gcpsecretbackend"
	"github.com/upbound/provider-vault/config/gcpsecretimpersonatedaccount"
	"github.com/upbound/provider-vault/config/gcpsecretroleset"
	"github.com/upbound/provider-vault/config/gcpsecretstaticaccount"
	"github.com/upbound/provider-vault/config/genericendpoint"
	"github.com/upbound/provider-vault/config/genericsecret"
	"github.com/upbound/provider-vault/config/githubauthbackend"
	"github.com/upbound/provider-vault/config/githubteam"
	"github.com/upbound/provider-vault/config/githubuser"
	"github.com/upbound/provider-vault/config/identityentity"
	"github.com/upbound/provider-vault/config/identityentityalias"
	"github.com/upbound/provider-vault/config/identityentitypolicies"
	"github.com/upbound/provider-vault/config/identitygroup"
	"github.com/upbound/provider-vault/config/identitygroupalias"
	"github.com/upbound/provider-vault/config/identitygroupmemberentityids"
	"github.com/upbound/provider-vault/config/identitygroupmembergroupids"
	"github.com/upbound/provider-vault/config/identitygrouppolicies"
	"github.com/upbound/provider-vault/config/identitymfaduo"
	"github.com/upbound/provider-vault/config/identitymfaloginenforcement"
	"github.com/upbound/provider-vault/config/identitymfaokta"
	"github.com/upbound/provider-vault/config/identitymfapingid"
	"github.com/upbound/provider-vault/config/identitymfatotp"
	"github.com/upbound/provider-vault/config/identityoidc"
	"github.com/upbound/provider-vault/config/identityoidcassignment"
	"github.com/upbound/provider-vault/config/identityoidcclient"
	"github.com/upbound/provider-vault/config/identityoidckey"
	"github.com/upbound/provider-vault/config/identityoidckeyallowedclientid"
	"github.com/upbound/provider-vault/config/identityoidcprovider"
	"github.com/upbound/provider-vault/config/identityoidcrole"
	"github.com/upbound/provider-vault/config/identityoidcscope"
	"github.com/upbound/provider-vault/config/jwtauthbackend"
	"github.com/upbound/provider-vault/config/jwtauthbackendrole"
	"github.com/upbound/provider-vault/config/kmipsecretbackend"
	"github.com/upbound/provider-vault/config/kmipsecretrole"
	"github.com/upbound/provider-vault/config/kmipsecretscope"
	"github.com/upbound/provider-vault/config/kubernetesauthbackendconfig"
	"github.com/upbound/provider-vault/config/kubernetesauthbackendrole"
	"github.com/upbound/provider-vault/config/kubernetessecretbackend"
	"github.com/upbound/provider-vault/config/kubernetessecretbackendrole"
	"github.com/upbound/provider-vault/config/kvsecret"
	"github.com/upbound/provider-vault/config/kvsecretbackendv2"
	"github.com/upbound/provider-vault/config/kvsecretv2"
	"github.com/upbound/provider-vault/config/ldapauthbackend"
	"github.com/upbound/provider-vault/config/ldapauthbackendgroup"
	"github.com/upbound/provider-vault/config/ldapauthbackenduser"
	"github.com/upbound/provider-vault/config/managedkeys"
	"github.com/upbound/provider-vault/config/mfaduo"
	"github.com/upbound/provider-vault/config/mfaokta"
	"github.com/upbound/provider-vault/config/mfapingid"
	"github.com/upbound/provider-vault/config/mfatotp"
	"github.com/upbound/provider-vault/config/mongodbatlassecretbackend"
	"github.com/upbound/provider-vault/config/mongodbatlassecretrole"
	"github.com/upbound/provider-vault/config/mount"
	"github.com/upbound/provider-vault/config/namespace"
	"github.com/upbound/provider-vault/config/nomadsecretbackend"
	"github.com/upbound/provider-vault/config/nomadsecretrole"
	"github.com/upbound/provider-vault/config/oktaauthbackend"
	"github.com/upbound/provider-vault/config/oktaauthbackendgroup"
	"github.com/upbound/provider-vault/config/oktaauthbackenduser"
	"github.com/upbound/provider-vault/config/passwordpolicy"
	"github.com/upbound/provider-vault/config/pkisecretbackendcert"
	"github.com/upbound/provider-vault/config/pkisecretbackendconfigca"
	"github.com/upbound/provider-vault/config/pkisecretbackendconfigurls"
	"github.com/upbound/provider-vault/config/pkisecretbackendcrlconfig"
	"github.com/upbound/provider-vault/config/pkisecretbackendintermediatecertrequest"
	"github.com/upbound/provider-vault/config/pkisecretbackendintermediatesetsigned"
	"github.com/upbound/provider-vault/config/pkisecretbackendrole"
	"github.com/upbound/provider-vault/config/pkisecretbackendrootcert"
	"github.com/upbound/provider-vault/config/pkisecretbackendrootsignintermediate"
	"github.com/upbound/provider-vault/config/pkisecretbackendsign"
	"github.com/upbound/provider-vault/config/policy"
	"github.com/upbound/provider-vault/config/quotaleasecount"
	"github.com/upbound/provider-vault/config/quotaratelimit"
	"github.com/upbound/provider-vault/config/rabbitmqsecretbackend"
	"github.com/upbound/provider-vault/config/rabbitmqsecretbackendrole"
	"github.com/upbound/provider-vault/config/raftautopilot"
	"github.com/upbound/provider-vault/config/raftsnapshotagentconfig"
	"github.com/upbound/provider-vault/config/rgppolicy"
	"github.com/upbound/provider-vault/config/sshsecretbackendca"
	"github.com/upbound/provider-vault/config/sshsecretbackendrole"
	"github.com/upbound/provider-vault/config/terraformcloudsecretbackend"
	"github.com/upbound/provider-vault/config/terraformcloudsecretcreds"
	"github.com/upbound/provider-vault/config/terraformcloudsecretrole"
	"github.com/upbound/provider-vault/config/token"
	"github.com/upbound/provider-vault/config/tokenauthbackendrole"
	"github.com/upbound/provider-vault/config/transformalphabet"
	"github.com/upbound/provider-vault/config/transformrole"
	"github.com/upbound/provider-vault/config/transformtemplate"
	"github.com/upbound/provider-vault/config/transformtransformation"
	"github.com/upbound/provider-vault/config/transitsecretbackendcacheconfig"
	"github.com/upbound/provider-vault/config/transitsecretbackendkey"
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
