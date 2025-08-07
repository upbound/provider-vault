/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	secretbackend "github.com/upbound/provider-vault/internal/controller/namespaced/ad/secretbackend"
	secretrole "github.com/upbound/provider-vault/internal/controller/namespaced/ad/secretrole"
	authbackendrole "github.com/upbound/provider-vault/internal/controller/namespaced/alicloud/authbackendrole"
	authbackendlogin "github.com/upbound/provider-vault/internal/controller/namespaced/approle/authbackendlogin"
	authbackendroleapprole "github.com/upbound/provider-vault/internal/controller/namespaced/approle/authbackendrole"
	authbackendrolesecretid "github.com/upbound/provider-vault/internal/controller/namespaced/approle/authbackendrolesecretid"
	requestheader "github.com/upbound/provider-vault/internal/controller/namespaced/audit/requestheader"
	backend "github.com/upbound/provider-vault/internal/controller/namespaced/auth/backend"
	authbackendcert "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendcert"
	authbackendclient "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendclient"
	authbackendconfigidentity "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendconfigidentity"
	authbackendidentitywhitelist "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendidentitywhitelist"
	authbackendloginaws "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendlogin"
	authbackendroleaws "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendrole"
	authbackendroletag "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendroletag"
	authbackendroletagblacklist "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendroletagblacklist"
	authbackendstsrole "github.com/upbound/provider-vault/internal/controller/namespaced/aws/authbackendstsrole"
	secretbackendaws "github.com/upbound/provider-vault/internal/controller/namespaced/aws/secretbackend"
	secretbackendrole "github.com/upbound/provider-vault/internal/controller/namespaced/aws/secretbackendrole"
	authbackendconfig "github.com/upbound/provider-vault/internal/controller/namespaced/azure/authbackendconfig"
	authbackendroleazure "github.com/upbound/provider-vault/internal/controller/namespaced/azure/authbackendrole"
	secretbackendazure "github.com/upbound/provider-vault/internal/controller/namespaced/azure/secretbackend"
	secretbackendroleazure "github.com/upbound/provider-vault/internal/controller/namespaced/azure/secretbackendrole"
	authbackendrolecert "github.com/upbound/provider-vault/internal/controller/namespaced/cert/authbackendrole"
	secretbackendconsul "github.com/upbound/provider-vault/internal/controller/namespaced/consul/secretbackend"
	secretbackendroleconsul "github.com/upbound/provider-vault/internal/controller/namespaced/consul/secretbackendrole"
	secretbackendconnection "github.com/upbound/provider-vault/internal/controller/namespaced/database/secretbackendconnection"
	secretbackendroledatabase "github.com/upbound/provider-vault/internal/controller/namespaced/database/secretbackendrole"
	secretbackendstaticrole "github.com/upbound/provider-vault/internal/controller/namespaced/database/secretbackendstaticrole"
	secretsmount "github.com/upbound/provider-vault/internal/controller/namespaced/database/secretsmount"
	policy "github.com/upbound/provider-vault/internal/controller/namespaced/egp/policy"
	authbackend "github.com/upbound/provider-vault/internal/controller/namespaced/gcp/authbackend"
	authbackendrolegcp "github.com/upbound/provider-vault/internal/controller/namespaced/gcp/authbackendrole"
	secretbackendgcp "github.com/upbound/provider-vault/internal/controller/namespaced/gcp/secretbackend"
	secretimpersonatedaccount "github.com/upbound/provider-vault/internal/controller/namespaced/gcp/secretimpersonatedaccount"
	secretroleset "github.com/upbound/provider-vault/internal/controller/namespaced/gcp/secretroleset"
	secretstaticaccount "github.com/upbound/provider-vault/internal/controller/namespaced/gcp/secretstaticaccount"
	endpoint "github.com/upbound/provider-vault/internal/controller/namespaced/generic/endpoint"
	secret "github.com/upbound/provider-vault/internal/controller/namespaced/generic/secret"
	authbackendgithub "github.com/upbound/provider-vault/internal/controller/namespaced/github/authbackend"
	team "github.com/upbound/provider-vault/internal/controller/namespaced/github/team"
	user "github.com/upbound/provider-vault/internal/controller/namespaced/github/user"
	entity "github.com/upbound/provider-vault/internal/controller/namespaced/identity/entity"
	entityalias "github.com/upbound/provider-vault/internal/controller/namespaced/identity/entityalias"
	entitypolicies "github.com/upbound/provider-vault/internal/controller/namespaced/identity/entitypolicies"
	group "github.com/upbound/provider-vault/internal/controller/namespaced/identity/group"
	groupalias "github.com/upbound/provider-vault/internal/controller/namespaced/identity/groupalias"
	groupmemberentityids "github.com/upbound/provider-vault/internal/controller/namespaced/identity/groupmemberentityids"
	groupmembergroupids "github.com/upbound/provider-vault/internal/controller/namespaced/identity/groupmembergroupids"
	grouppolicies "github.com/upbound/provider-vault/internal/controller/namespaced/identity/grouppolicies"
	mfaduo "github.com/upbound/provider-vault/internal/controller/namespaced/identity/mfaduo"
	mfaloginenforcement "github.com/upbound/provider-vault/internal/controller/namespaced/identity/mfaloginenforcement"
	mfaokta "github.com/upbound/provider-vault/internal/controller/namespaced/identity/mfaokta"
	mfapingid "github.com/upbound/provider-vault/internal/controller/namespaced/identity/mfapingid"
	mfatotp "github.com/upbound/provider-vault/internal/controller/namespaced/identity/mfatotp"
	oidc "github.com/upbound/provider-vault/internal/controller/namespaced/identity/oidc"
	oidcassignment "github.com/upbound/provider-vault/internal/controller/namespaced/identity/oidcassignment"
	oidcclient "github.com/upbound/provider-vault/internal/controller/namespaced/identity/oidcclient"
	oidckey "github.com/upbound/provider-vault/internal/controller/namespaced/identity/oidckey"
	oidckeyallowedclientid "github.com/upbound/provider-vault/internal/controller/namespaced/identity/oidckeyallowedclientid"
	oidcprovider "github.com/upbound/provider-vault/internal/controller/namespaced/identity/oidcprovider"
	oidcrole "github.com/upbound/provider-vault/internal/controller/namespaced/identity/oidcrole"
	oidcscope "github.com/upbound/provider-vault/internal/controller/namespaced/identity/oidcscope"
	authbackendjwt "github.com/upbound/provider-vault/internal/controller/namespaced/jwt/authbackend"
	authbackendrolejwt "github.com/upbound/provider-vault/internal/controller/namespaced/jwt/authbackendrole"
	secretbackendkmip "github.com/upbound/provider-vault/internal/controller/namespaced/kmip/secretbackend"
	secretrolekmip "github.com/upbound/provider-vault/internal/controller/namespaced/kmip/secretrole"
	secretscope "github.com/upbound/provider-vault/internal/controller/namespaced/kmip/secretscope"
	authbackendconfigkubernetes "github.com/upbound/provider-vault/internal/controller/namespaced/kubernetes/authbackendconfig"
	authbackendrolekubernetes "github.com/upbound/provider-vault/internal/controller/namespaced/kubernetes/authbackendrole"
	secretbackendkubernetes "github.com/upbound/provider-vault/internal/controller/namespaced/kubernetes/secretbackend"
	secretbackendrolekubernetes "github.com/upbound/provider-vault/internal/controller/namespaced/kubernetes/secretbackendrole"
	secretkv "github.com/upbound/provider-vault/internal/controller/namespaced/kv/secret"
	secretbackendv2 "github.com/upbound/provider-vault/internal/controller/namespaced/kv/secretbackendv2"
	secretv2 "github.com/upbound/provider-vault/internal/controller/namespaced/kv/secretv2"
	authbackendldap "github.com/upbound/provider-vault/internal/controller/namespaced/ldap/authbackend"
	authbackendgroup "github.com/upbound/provider-vault/internal/controller/namespaced/ldap/authbackendgroup"
	authbackenduser "github.com/upbound/provider-vault/internal/controller/namespaced/ldap/authbackenduser"
	keys "github.com/upbound/provider-vault/internal/controller/namespaced/managed/keys"
	duo "github.com/upbound/provider-vault/internal/controller/namespaced/mfa/duo"
	okta "github.com/upbound/provider-vault/internal/controller/namespaced/mfa/okta"
	pingid "github.com/upbound/provider-vault/internal/controller/namespaced/mfa/pingid"
	totp "github.com/upbound/provider-vault/internal/controller/namespaced/mfa/totp"
	secretbackendmongodbatlas "github.com/upbound/provider-vault/internal/controller/namespaced/mongodbatlas/secretbackend"
	secretrolemongodbatlas "github.com/upbound/provider-vault/internal/controller/namespaced/mongodbatlas/secretrole"
	secretbackendnomad "github.com/upbound/provider-vault/internal/controller/namespaced/nomad/secretbackend"
	secretrolenomad "github.com/upbound/provider-vault/internal/controller/namespaced/nomad/secretrole"
	authbackendokta "github.com/upbound/provider-vault/internal/controller/namespaced/okta/authbackend"
	authbackendgroupokta "github.com/upbound/provider-vault/internal/controller/namespaced/okta/authbackendgroup"
	authbackenduserokta "github.com/upbound/provider-vault/internal/controller/namespaced/okta/authbackenduser"
	policypassword "github.com/upbound/provider-vault/internal/controller/namespaced/password/policy"
	secretbackendcert "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendcert"
	secretbackendconfigca "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendconfigca"
	secretbackendconfigurls "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendconfigurls"
	secretbackendcrlconfig "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendcrlconfig"
	secretbackendintermediatecertrequest "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendintermediatecertrequest"
	secretbackendintermediatesetsigned "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendintermediatesetsigned"
	secretbackendrolepki "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendrole"
	secretbackendrootcert "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendrootcert"
	secretbackendrootsignintermediate "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendrootsignintermediate"
	secretbackendsign "github.com/upbound/provider-vault/internal/controller/namespaced/pki/secretbackendsign"
	providerconfig "github.com/upbound/provider-vault/internal/controller/namespaced/providerconfig"
	leasecount "github.com/upbound/provider-vault/internal/controller/namespaced/quota/leasecount"
	ratelimit "github.com/upbound/provider-vault/internal/controller/namespaced/quota/ratelimit"
	secretbackendrabbitmq "github.com/upbound/provider-vault/internal/controller/namespaced/rabbitmq/secretbackend"
	secretbackendrolerabbitmq "github.com/upbound/provider-vault/internal/controller/namespaced/rabbitmq/secretbackendrole"
	autopilot "github.com/upbound/provider-vault/internal/controller/namespaced/raft/autopilot"
	snapshotagentconfig "github.com/upbound/provider-vault/internal/controller/namespaced/raft/snapshotagentconfig"
	policyrgp "github.com/upbound/provider-vault/internal/controller/namespaced/rgp/policy"
	secretbackendca "github.com/upbound/provider-vault/internal/controller/namespaced/ssh/secretbackendca"
	secretbackendrolessh "github.com/upbound/provider-vault/internal/controller/namespaced/ssh/secretbackendrole"
	cloudsecretbackend "github.com/upbound/provider-vault/internal/controller/namespaced/terraform/cloudsecretbackend"
	cloudsecretcreds "github.com/upbound/provider-vault/internal/controller/namespaced/terraform/cloudsecretcreds"
	cloudsecretrole "github.com/upbound/provider-vault/internal/controller/namespaced/terraform/cloudsecretrole"
	authbackendroletoken "github.com/upbound/provider-vault/internal/controller/namespaced/token/authbackendrole"
	alphabet "github.com/upbound/provider-vault/internal/controller/namespaced/transform/alphabet"
	role "github.com/upbound/provider-vault/internal/controller/namespaced/transform/role"
	template "github.com/upbound/provider-vault/internal/controller/namespaced/transform/template"
	transformation "github.com/upbound/provider-vault/internal/controller/namespaced/transform/transformation"
	secretbackendkey "github.com/upbound/provider-vault/internal/controller/namespaced/transit/secretbackendkey"
	audit "github.com/upbound/provider-vault/internal/controller/namespaced/vault/audit"
	mount "github.com/upbound/provider-vault/internal/controller/namespaced/vault/mount"
	policyvault "github.com/upbound/provider-vault/internal/controller/namespaced/vault/policy"
	token "github.com/upbound/provider-vault/internal/controller/namespaced/vault/token"
	vaultnamespace "github.com/upbound/provider-vault/internal/controller/namespaced/vault/vaultnamespace"
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
