/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	secretbackend "github.com/upbound/provider-vault/v3/internal/controller/cluster/ad/secretbackend"
	secretrole "github.com/upbound/provider-vault/v3/internal/controller/cluster/ad/secretrole"
	authbackendrole "github.com/upbound/provider-vault/v3/internal/controller/cluster/alicloud/authbackendrole"
	authbackendlogin "github.com/upbound/provider-vault/v3/internal/controller/cluster/approle/authbackendlogin"
	authbackendroleapprole "github.com/upbound/provider-vault/v3/internal/controller/cluster/approle/authbackendrole"
	authbackendrolesecretid "github.com/upbound/provider-vault/v3/internal/controller/cluster/approle/authbackendrolesecretid"
	requestheader "github.com/upbound/provider-vault/v3/internal/controller/cluster/audit/requestheader"
	backend "github.com/upbound/provider-vault/v3/internal/controller/cluster/auth/backend"
	authbackendcert "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendcert"
	authbackendclient "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendclient"
	authbackendconfigidentity "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendconfigidentity"
	authbackendidentitywhitelist "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendidentitywhitelist"
	authbackendloginaws "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendlogin"
	authbackendroleaws "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendrole"
	authbackendroletag "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendroletag"
	authbackendroletagblacklist "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendroletagblacklist"
	authbackendstsrole "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/authbackendstsrole"
	secretbackendaws "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/secretbackend"
	secretbackendrole "github.com/upbound/provider-vault/v3/internal/controller/cluster/aws/secretbackendrole"
	authbackendconfig "github.com/upbound/provider-vault/v3/internal/controller/cluster/azure/authbackendconfig"
	authbackendroleazure "github.com/upbound/provider-vault/v3/internal/controller/cluster/azure/authbackendrole"
	secretbackendazure "github.com/upbound/provider-vault/v3/internal/controller/cluster/azure/secretbackend"
	secretbackendroleazure "github.com/upbound/provider-vault/v3/internal/controller/cluster/azure/secretbackendrole"
	authbackendrolecert "github.com/upbound/provider-vault/v3/internal/controller/cluster/cert/authbackendrole"
	secretbackendconsul "github.com/upbound/provider-vault/v3/internal/controller/cluster/consul/secretbackend"
	secretbackendroleconsul "github.com/upbound/provider-vault/v3/internal/controller/cluster/consul/secretbackendrole"
	secretbackendconnection "github.com/upbound/provider-vault/v3/internal/controller/cluster/database/secretbackendconnection"
	secretbackendroledatabase "github.com/upbound/provider-vault/v3/internal/controller/cluster/database/secretbackendrole"
	secretbackendstaticrole "github.com/upbound/provider-vault/v3/internal/controller/cluster/database/secretbackendstaticrole"
	secretsmount "github.com/upbound/provider-vault/v3/internal/controller/cluster/database/secretsmount"
	policy "github.com/upbound/provider-vault/v3/internal/controller/cluster/egp/policy"
	authbackend "github.com/upbound/provider-vault/v3/internal/controller/cluster/gcp/authbackend"
	authbackendrolegcp "github.com/upbound/provider-vault/v3/internal/controller/cluster/gcp/authbackendrole"
	secretbackendgcp "github.com/upbound/provider-vault/v3/internal/controller/cluster/gcp/secretbackend"
	secretimpersonatedaccount "github.com/upbound/provider-vault/v3/internal/controller/cluster/gcp/secretimpersonatedaccount"
	secretroleset "github.com/upbound/provider-vault/v3/internal/controller/cluster/gcp/secretroleset"
	secretstaticaccount "github.com/upbound/provider-vault/v3/internal/controller/cluster/gcp/secretstaticaccount"
	endpoint "github.com/upbound/provider-vault/v3/internal/controller/cluster/generic/endpoint"
	secret "github.com/upbound/provider-vault/v3/internal/controller/cluster/generic/secret"
	authbackendgithub "github.com/upbound/provider-vault/v3/internal/controller/cluster/github/authbackend"
	team "github.com/upbound/provider-vault/v3/internal/controller/cluster/github/team"
	user "github.com/upbound/provider-vault/v3/internal/controller/cluster/github/user"
	entity "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/entity"
	entityalias "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/entityalias"
	entitypolicies "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/entitypolicies"
	group "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/group"
	groupalias "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/groupalias"
	groupmemberentityids "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/groupmemberentityids"
	groupmembergroupids "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/groupmembergroupids"
	grouppolicies "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/grouppolicies"
	mfaduo "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/mfaduo"
	mfaloginenforcement "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/mfaloginenforcement"
	mfaokta "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/mfaokta"
	mfapingid "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/mfapingid"
	mfatotp "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/mfatotp"
	oidc "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/oidc"
	oidcassignment "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/oidcassignment"
	oidcclient "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/oidcclient"
	oidckey "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/oidckey"
	oidckeyallowedclientid "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/oidckeyallowedclientid"
	oidcprovider "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/oidcprovider"
	oidcrole "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/oidcrole"
	oidcscope "github.com/upbound/provider-vault/v3/internal/controller/cluster/identity/oidcscope"
	authbackendjwt "github.com/upbound/provider-vault/v3/internal/controller/cluster/jwt/authbackend"
	authbackendrolejwt "github.com/upbound/provider-vault/v3/internal/controller/cluster/jwt/authbackendrole"
	secretbackendkmip "github.com/upbound/provider-vault/v3/internal/controller/cluster/kmip/secretbackend"
	secretrolekmip "github.com/upbound/provider-vault/v3/internal/controller/cluster/kmip/secretrole"
	secretscope "github.com/upbound/provider-vault/v3/internal/controller/cluster/kmip/secretscope"
	authbackendconfigkubernetes "github.com/upbound/provider-vault/v3/internal/controller/cluster/kubernetes/authbackendconfig"
	authbackendrolekubernetes "github.com/upbound/provider-vault/v3/internal/controller/cluster/kubernetes/authbackendrole"
	secretbackendkubernetes "github.com/upbound/provider-vault/v3/internal/controller/cluster/kubernetes/secretbackend"
	secretbackendrolekubernetes "github.com/upbound/provider-vault/v3/internal/controller/cluster/kubernetes/secretbackendrole"
	secretkv "github.com/upbound/provider-vault/v3/internal/controller/cluster/kv/secret"
	secretbackendv2 "github.com/upbound/provider-vault/v3/internal/controller/cluster/kv/secretbackendv2"
	secretv2 "github.com/upbound/provider-vault/v3/internal/controller/cluster/kv/secretv2"
	authbackendldap "github.com/upbound/provider-vault/v3/internal/controller/cluster/ldap/authbackend"
	authbackendgroup "github.com/upbound/provider-vault/v3/internal/controller/cluster/ldap/authbackendgroup"
	authbackenduser "github.com/upbound/provider-vault/v3/internal/controller/cluster/ldap/authbackenduser"
	keys "github.com/upbound/provider-vault/v3/internal/controller/cluster/managed/keys"
	duo "github.com/upbound/provider-vault/v3/internal/controller/cluster/mfa/duo"
	okta "github.com/upbound/provider-vault/v3/internal/controller/cluster/mfa/okta"
	pingid "github.com/upbound/provider-vault/v3/internal/controller/cluster/mfa/pingid"
	totp "github.com/upbound/provider-vault/v3/internal/controller/cluster/mfa/totp"
	secretbackendmongodbatlas "github.com/upbound/provider-vault/v3/internal/controller/cluster/mongodbatlas/secretbackend"
	secretrolemongodbatlas "github.com/upbound/provider-vault/v3/internal/controller/cluster/mongodbatlas/secretrole"
	secretbackendnomad "github.com/upbound/provider-vault/v3/internal/controller/cluster/nomad/secretbackend"
	secretrolenomad "github.com/upbound/provider-vault/v3/internal/controller/cluster/nomad/secretrole"
	authbackendokta "github.com/upbound/provider-vault/v3/internal/controller/cluster/okta/authbackend"
	authbackendgroupokta "github.com/upbound/provider-vault/v3/internal/controller/cluster/okta/authbackendgroup"
	authbackenduserokta "github.com/upbound/provider-vault/v3/internal/controller/cluster/okta/authbackenduser"
	policypassword "github.com/upbound/provider-vault/v3/internal/controller/cluster/password/policy"
	secretbackendcert "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendcert"
	secretbackendconfigca "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendconfigca"
	secretbackendconfigurls "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendconfigurls"
	secretbackendcrlconfig "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendcrlconfig"
	secretbackendintermediatecertrequest "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendintermediatecertrequest"
	secretbackendintermediatesetsigned "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendintermediatesetsigned"
	secretbackendrolepki "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendrole"
	secretbackendrootcert "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendrootcert"
	secretbackendrootsignintermediate "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendrootsignintermediate"
	secretbackendsign "github.com/upbound/provider-vault/v3/internal/controller/cluster/pki/secretbackendsign"
	providerconfig "github.com/upbound/provider-vault/v3/internal/controller/cluster/providerconfig"
	leasecount "github.com/upbound/provider-vault/v3/internal/controller/cluster/quota/leasecount"
	ratelimit "github.com/upbound/provider-vault/v3/internal/controller/cluster/quota/ratelimit"
	secretbackendrabbitmq "github.com/upbound/provider-vault/v3/internal/controller/cluster/rabbitmq/secretbackend"
	secretbackendrolerabbitmq "github.com/upbound/provider-vault/v3/internal/controller/cluster/rabbitmq/secretbackendrole"
	autopilot "github.com/upbound/provider-vault/v3/internal/controller/cluster/raft/autopilot"
	snapshotagentconfig "github.com/upbound/provider-vault/v3/internal/controller/cluster/raft/snapshotagentconfig"
	policyrgp "github.com/upbound/provider-vault/v3/internal/controller/cluster/rgp/policy"
	secretbackendca "github.com/upbound/provider-vault/v3/internal/controller/cluster/ssh/secretbackendca"
	secretbackendrolessh "github.com/upbound/provider-vault/v3/internal/controller/cluster/ssh/secretbackendrole"
	cloudsecretbackend "github.com/upbound/provider-vault/v3/internal/controller/cluster/terraform/cloudsecretbackend"
	cloudsecretcreds "github.com/upbound/provider-vault/v3/internal/controller/cluster/terraform/cloudsecretcreds"
	cloudsecretrole "github.com/upbound/provider-vault/v3/internal/controller/cluster/terraform/cloudsecretrole"
	authbackendroletoken "github.com/upbound/provider-vault/v3/internal/controller/cluster/token/authbackendrole"
	alphabet "github.com/upbound/provider-vault/v3/internal/controller/cluster/transform/alphabet"
	role "github.com/upbound/provider-vault/v3/internal/controller/cluster/transform/role"
	template "github.com/upbound/provider-vault/v3/internal/controller/cluster/transform/template"
	transformation "github.com/upbound/provider-vault/v3/internal/controller/cluster/transform/transformation"
	secretbackendkey "github.com/upbound/provider-vault/v3/internal/controller/cluster/transit/secretbackendkey"
	audit "github.com/upbound/provider-vault/v3/internal/controller/cluster/vault/audit"
	mount "github.com/upbound/provider-vault/v3/internal/controller/cluster/vault/mount"
	policyvault "github.com/upbound/provider-vault/v3/internal/controller/cluster/vault/policy"
	token "github.com/upbound/provider-vault/v3/internal/controller/cluster/vault/token"
	vaultnamespace "github.com/upbound/provider-vault/v3/internal/controller/cluster/vault/vaultnamespace"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secretbackend.Setup,
		secretrole.Setup,
		authbackendrole.Setup,
		authbackendlogin.Setup,
		authbackendroleapprole.Setup,
		authbackendrolesecretid.Setup,
		requestheader.Setup,
		backend.Setup,
		authbackendcert.Setup,
		authbackendclient.Setup,
		authbackendconfigidentity.Setup,
		authbackendidentitywhitelist.Setup,
		authbackendloginaws.Setup,
		authbackendroleaws.Setup,
		authbackendroletag.Setup,
		authbackendroletagblacklist.Setup,
		authbackendstsrole.Setup,
		secretbackendaws.Setup,
		secretbackendrole.Setup,
		authbackendconfig.Setup,
		authbackendroleazure.Setup,
		secretbackendazure.Setup,
		secretbackendroleazure.Setup,
		authbackendrolecert.Setup,
		secretbackendconsul.Setup,
		secretbackendroleconsul.Setup,
		secretbackendconnection.Setup,
		secretbackendroledatabase.Setup,
		secretbackendstaticrole.Setup,
		secretsmount.Setup,
		policy.Setup,
		authbackend.Setup,
		authbackendrolegcp.Setup,
		secretbackendgcp.Setup,
		secretimpersonatedaccount.Setup,
		secretroleset.Setup,
		secretstaticaccount.Setup,
		endpoint.Setup,
		secret.Setup,
		authbackendgithub.Setup,
		team.Setup,
		user.Setup,
		entity.Setup,
		entityalias.Setup,
		entitypolicies.Setup,
		group.Setup,
		groupalias.Setup,
		groupmemberentityids.Setup,
		groupmembergroupids.Setup,
		grouppolicies.Setup,
		mfaduo.Setup,
		mfaloginenforcement.Setup,
		mfaokta.Setup,
		mfapingid.Setup,
		mfatotp.Setup,
		oidc.Setup,
		oidcassignment.Setup,
		oidcclient.Setup,
		oidckey.Setup,
		oidckeyallowedclientid.Setup,
		oidcprovider.Setup,
		oidcrole.Setup,
		oidcscope.Setup,
		authbackendjwt.Setup,
		authbackendrolejwt.Setup,
		secretbackendkmip.Setup,
		secretrolekmip.Setup,
		secretscope.Setup,
		authbackendconfigkubernetes.Setup,
		authbackendrolekubernetes.Setup,
		secretbackendkubernetes.Setup,
		secretbackendrolekubernetes.Setup,
		secretkv.Setup,
		secretbackendv2.Setup,
		secretv2.Setup,
		authbackendldap.Setup,
		authbackendgroup.Setup,
		authbackenduser.Setup,
		keys.Setup,
		duo.Setup,
		okta.Setup,
		pingid.Setup,
		totp.Setup,
		secretbackendmongodbatlas.Setup,
		secretrolemongodbatlas.Setup,
		secretbackendnomad.Setup,
		secretrolenomad.Setup,
		authbackendokta.Setup,
		authbackendgroupokta.Setup,
		authbackenduserokta.Setup,
		policypassword.Setup,
		secretbackendcert.Setup,
		secretbackendconfigca.Setup,
		secretbackendconfigurls.Setup,
		secretbackendcrlconfig.Setup,
		secretbackendintermediatecertrequest.Setup,
		secretbackendintermediatesetsigned.Setup,
		secretbackendrolepki.Setup,
		secretbackendrootcert.Setup,
		secretbackendrootsignintermediate.Setup,
		secretbackendsign.Setup,
		providerconfig.Setup,
		leasecount.Setup,
		ratelimit.Setup,
		secretbackendrabbitmq.Setup,
		secretbackendrolerabbitmq.Setup,
		autopilot.Setup,
		snapshotagentconfig.Setup,
		policyrgp.Setup,
		secretbackendca.Setup,
		secretbackendrolessh.Setup,
		cloudsecretbackend.Setup,
		cloudsecretcreds.Setup,
		cloudsecretrole.Setup,
		authbackendroletoken.Setup,
		alphabet.Setup,
		role.Setup,
		template.Setup,
		transformation.Setup,
		secretbackendkey.Setup,
		audit.Setup,
		mount.Setup,
		policyvault.Setup,
		token.Setup,
		vaultnamespace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secretbackend.SetupGated,
		secretrole.SetupGated,
		authbackendrole.SetupGated,
		authbackendlogin.SetupGated,
		authbackendroleapprole.SetupGated,
		authbackendrolesecretid.SetupGated,
		requestheader.SetupGated,
		backend.SetupGated,
		authbackendcert.SetupGated,
		authbackendclient.SetupGated,
		authbackendconfigidentity.SetupGated,
		authbackendidentitywhitelist.SetupGated,
		authbackendloginaws.SetupGated,
		authbackendroleaws.SetupGated,
		authbackendroletag.SetupGated,
		authbackendroletagblacklist.SetupGated,
		authbackendstsrole.SetupGated,
		secretbackendaws.SetupGated,
		secretbackendrole.SetupGated,
		authbackendconfig.SetupGated,
		authbackendroleazure.SetupGated,
		secretbackendazure.SetupGated,
		secretbackendroleazure.SetupGated,
		authbackendrolecert.SetupGated,
		secretbackendconsul.SetupGated,
		secretbackendroleconsul.SetupGated,
		secretbackendconnection.SetupGated,
		secretbackendroledatabase.SetupGated,
		secretbackendstaticrole.SetupGated,
		secretsmount.SetupGated,
		policy.SetupGated,
		authbackend.SetupGated,
		authbackendrolegcp.SetupGated,
		secretbackendgcp.SetupGated,
		secretimpersonatedaccount.SetupGated,
		secretroleset.SetupGated,
		secretstaticaccount.SetupGated,
		endpoint.SetupGated,
		secret.SetupGated,
		authbackendgithub.SetupGated,
		team.SetupGated,
		user.SetupGated,
		entity.SetupGated,
		entityalias.SetupGated,
		entitypolicies.SetupGated,
		group.SetupGated,
		groupalias.SetupGated,
		groupmemberentityids.SetupGated,
		groupmembergroupids.SetupGated,
		grouppolicies.SetupGated,
		mfaduo.SetupGated,
		mfaloginenforcement.SetupGated,
		mfaokta.SetupGated,
		mfapingid.SetupGated,
		mfatotp.SetupGated,
		oidc.SetupGated,
		oidcassignment.SetupGated,
		oidcclient.SetupGated,
		oidckey.SetupGated,
		oidckeyallowedclientid.SetupGated,
		oidcprovider.SetupGated,
		oidcrole.SetupGated,
		oidcscope.SetupGated,
		authbackendjwt.SetupGated,
		authbackendrolejwt.SetupGated,
		secretbackendkmip.SetupGated,
		secretrolekmip.SetupGated,
		secretscope.SetupGated,
		authbackendconfigkubernetes.SetupGated,
		authbackendrolekubernetes.SetupGated,
		secretbackendkubernetes.SetupGated,
		secretbackendrolekubernetes.SetupGated,
		secretkv.SetupGated,
		secretbackendv2.SetupGated,
		secretv2.SetupGated,
		authbackendldap.SetupGated,
		authbackendgroup.SetupGated,
		authbackenduser.SetupGated,
		keys.SetupGated,
		duo.SetupGated,
		okta.SetupGated,
		pingid.SetupGated,
		totp.SetupGated,
		secretbackendmongodbatlas.SetupGated,
		secretrolemongodbatlas.SetupGated,
		secretbackendnomad.SetupGated,
		secretrolenomad.SetupGated,
		authbackendokta.SetupGated,
		authbackendgroupokta.SetupGated,
		authbackenduserokta.SetupGated,
		policypassword.SetupGated,
		secretbackendcert.SetupGated,
		secretbackendconfigca.SetupGated,
		secretbackendconfigurls.SetupGated,
		secretbackendcrlconfig.SetupGated,
		secretbackendintermediatecertrequest.SetupGated,
		secretbackendintermediatesetsigned.SetupGated,
		secretbackendrolepki.SetupGated,
		secretbackendrootcert.SetupGated,
		secretbackendrootsignintermediate.SetupGated,
		secretbackendsign.SetupGated,
		providerconfig.SetupGated,
		leasecount.SetupGated,
		ratelimit.SetupGated,
		secretbackendrabbitmq.SetupGated,
		secretbackendrolerabbitmq.SetupGated,
		autopilot.SetupGated,
		snapshotagentconfig.SetupGated,
		policyrgp.SetupGated,
		secretbackendca.SetupGated,
		secretbackendrolessh.SetupGated,
		cloudsecretbackend.SetupGated,
		cloudsecretcreds.SetupGated,
		cloudsecretrole.SetupGated,
		authbackendroletoken.SetupGated,
		alphabet.SetupGated,
		role.SetupGated,
		template.SetupGated,
		transformation.SetupGated,
		secretbackendkey.SetupGated,
		audit.SetupGated,
		mount.SetupGated,
		policyvault.SetupGated,
		token.SetupGated,
		vaultnamespace.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
