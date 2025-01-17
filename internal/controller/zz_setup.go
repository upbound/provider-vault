/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	secretbackend "github.com/upbound/provider-vault/internal/controller/ad/secretbackend"
	secretrole "github.com/upbound/provider-vault/internal/controller/ad/secretrole"
	authbackendrole "github.com/upbound/provider-vault/internal/controller/alicloud/authbackendrole"
	authbackendlogin "github.com/upbound/provider-vault/internal/controller/approle/authbackendlogin"
	authbackendroleapprole "github.com/upbound/provider-vault/internal/controller/approle/authbackendrole"
	authbackendrolesecretid "github.com/upbound/provider-vault/internal/controller/approle/authbackendrolesecretid"
	requestheader "github.com/upbound/provider-vault/internal/controller/audit/requestheader"
	backend "github.com/upbound/provider-vault/internal/controller/auth/backend"
	authbackendcert "github.com/upbound/provider-vault/internal/controller/aws/authbackendcert"
	authbackendclient "github.com/upbound/provider-vault/internal/controller/aws/authbackendclient"
	authbackendconfigidentity "github.com/upbound/provider-vault/internal/controller/aws/authbackendconfigidentity"
	authbackendidentitywhitelist "github.com/upbound/provider-vault/internal/controller/aws/authbackendidentitywhitelist"
	authbackendloginaws "github.com/upbound/provider-vault/internal/controller/aws/authbackendlogin"
	authbackendroleaws "github.com/upbound/provider-vault/internal/controller/aws/authbackendrole"
	authbackendroletag "github.com/upbound/provider-vault/internal/controller/aws/authbackendroletag"
	authbackendroletagblacklist "github.com/upbound/provider-vault/internal/controller/aws/authbackendroletagblacklist"
	authbackendstsrole "github.com/upbound/provider-vault/internal/controller/aws/authbackendstsrole"
	secretbackendaws "github.com/upbound/provider-vault/internal/controller/aws/secretbackend"
	secretbackendrole "github.com/upbound/provider-vault/internal/controller/aws/secretbackendrole"
	authbackendconfig "github.com/upbound/provider-vault/internal/controller/azure/authbackendconfig"
	authbackendroleazure "github.com/upbound/provider-vault/internal/controller/azure/authbackendrole"
	secretbackendazure "github.com/upbound/provider-vault/internal/controller/azure/secretbackend"
	secretbackendroleazure "github.com/upbound/provider-vault/internal/controller/azure/secretbackendrole"
	authbackendrolecert "github.com/upbound/provider-vault/internal/controller/cert/authbackendrole"
	secretbackendconsul "github.com/upbound/provider-vault/internal/controller/consul/secretbackend"
	secretbackendroleconsul "github.com/upbound/provider-vault/internal/controller/consul/secretbackendrole"
	secretbackendconnection "github.com/upbound/provider-vault/internal/controller/database/secretbackendconnection"
	secretbackendroledatabase "github.com/upbound/provider-vault/internal/controller/database/secretbackendrole"
	secretbackendstaticrole "github.com/upbound/provider-vault/internal/controller/database/secretbackendstaticrole"
	secretsmount "github.com/upbound/provider-vault/internal/controller/database/secretsmount"
	policy "github.com/upbound/provider-vault/internal/controller/egp/policy"
	authbackend "github.com/upbound/provider-vault/internal/controller/gcp/authbackend"
	authbackendrolegcp "github.com/upbound/provider-vault/internal/controller/gcp/authbackendrole"
	secretbackendgcp "github.com/upbound/provider-vault/internal/controller/gcp/secretbackend"
	secretimpersonatedaccount "github.com/upbound/provider-vault/internal/controller/gcp/secretimpersonatedaccount"
	secretroleset "github.com/upbound/provider-vault/internal/controller/gcp/secretroleset"
	secretstaticaccount "github.com/upbound/provider-vault/internal/controller/gcp/secretstaticaccount"
	endpoint "github.com/upbound/provider-vault/internal/controller/generic/endpoint"
	secret "github.com/upbound/provider-vault/internal/controller/generic/secret"
	authbackendgithub "github.com/upbound/provider-vault/internal/controller/github/authbackend"
	team "github.com/upbound/provider-vault/internal/controller/github/team"
	user "github.com/upbound/provider-vault/internal/controller/github/user"
	entity "github.com/upbound/provider-vault/internal/controller/identity/entity"
	entityalias "github.com/upbound/provider-vault/internal/controller/identity/entityalias"
	entitypolicies "github.com/upbound/provider-vault/internal/controller/identity/entitypolicies"
	group "github.com/upbound/provider-vault/internal/controller/identity/group"
	groupalias "github.com/upbound/provider-vault/internal/controller/identity/groupalias"
	groupmemberentityids "github.com/upbound/provider-vault/internal/controller/identity/groupmemberentityids"
	groupmembergroupids "github.com/upbound/provider-vault/internal/controller/identity/groupmembergroupids"
	grouppolicies "github.com/upbound/provider-vault/internal/controller/identity/grouppolicies"
	mfaduo "github.com/upbound/provider-vault/internal/controller/identity/mfaduo"
	mfaloginenforcement "github.com/upbound/provider-vault/internal/controller/identity/mfaloginenforcement"
	mfaokta "github.com/upbound/provider-vault/internal/controller/identity/mfaokta"
	mfapingid "github.com/upbound/provider-vault/internal/controller/identity/mfapingid"
	mfatotp "github.com/upbound/provider-vault/internal/controller/identity/mfatotp"
	oidc "github.com/upbound/provider-vault/internal/controller/identity/oidc"
	oidcassignment "github.com/upbound/provider-vault/internal/controller/identity/oidcassignment"
	oidcclient "github.com/upbound/provider-vault/internal/controller/identity/oidcclient"
	oidckey "github.com/upbound/provider-vault/internal/controller/identity/oidckey"
	oidckeyallowedclientid "github.com/upbound/provider-vault/internal/controller/identity/oidckeyallowedclientid"
	oidcprovider "github.com/upbound/provider-vault/internal/controller/identity/oidcprovider"
	oidcrole "github.com/upbound/provider-vault/internal/controller/identity/oidcrole"
	oidcscope "github.com/upbound/provider-vault/internal/controller/identity/oidcscope"
	authbackendjwt "github.com/upbound/provider-vault/internal/controller/jwt/authbackend"
	authbackendrolejwt "github.com/upbound/provider-vault/internal/controller/jwt/authbackendrole"
	secretbackendkmip "github.com/upbound/provider-vault/internal/controller/kmip/secretbackend"
	secretrolekmip "github.com/upbound/provider-vault/internal/controller/kmip/secretrole"
	secretscope "github.com/upbound/provider-vault/internal/controller/kmip/secretscope"
	authbackendconfigkubernetes "github.com/upbound/provider-vault/internal/controller/kubernetes/authbackendconfig"
	authbackendrolekubernetes "github.com/upbound/provider-vault/internal/controller/kubernetes/authbackendrole"
	secretbackendkubernetes "github.com/upbound/provider-vault/internal/controller/kubernetes/secretbackend"
	secretbackendrolekubernetes "github.com/upbound/provider-vault/internal/controller/kubernetes/secretbackendrole"
	secretkv "github.com/upbound/provider-vault/internal/controller/kv/secret"
	secretbackendv2 "github.com/upbound/provider-vault/internal/controller/kv/secretbackendv2"
	secretv2 "github.com/upbound/provider-vault/internal/controller/kv/secretv2"
	authbackendldap "github.com/upbound/provider-vault/internal/controller/ldap/authbackend"
	authbackendgroup "github.com/upbound/provider-vault/internal/controller/ldap/authbackendgroup"
	authbackenduser "github.com/upbound/provider-vault/internal/controller/ldap/authbackenduser"
	keys "github.com/upbound/provider-vault/internal/controller/managed/keys"
	duo "github.com/upbound/provider-vault/internal/controller/mfa/duo"
	okta "github.com/upbound/provider-vault/internal/controller/mfa/okta"
	pingid "github.com/upbound/provider-vault/internal/controller/mfa/pingid"
	totp "github.com/upbound/provider-vault/internal/controller/mfa/totp"
	secretbackendmongodbatlas "github.com/upbound/provider-vault/internal/controller/mongodbatlas/secretbackend"
	secretrolemongodbatlas "github.com/upbound/provider-vault/internal/controller/mongodbatlas/secretrole"
	secretbackendnomad "github.com/upbound/provider-vault/internal/controller/nomad/secretbackend"
	secretrolenomad "github.com/upbound/provider-vault/internal/controller/nomad/secretrole"
	authbackendokta "github.com/upbound/provider-vault/internal/controller/okta/authbackend"
	authbackendgroupokta "github.com/upbound/provider-vault/internal/controller/okta/authbackendgroup"
	authbackenduserokta "github.com/upbound/provider-vault/internal/controller/okta/authbackenduser"
	policypassword "github.com/upbound/provider-vault/internal/controller/password/policy"
	secretbackendcert "github.com/upbound/provider-vault/internal/controller/pki/secretbackendcert"
	secretbackendconfigca "github.com/upbound/provider-vault/internal/controller/pki/secretbackendconfigca"
	secretbackendconfigurls "github.com/upbound/provider-vault/internal/controller/pki/secretbackendconfigurls"
	secretbackendcrlconfig "github.com/upbound/provider-vault/internal/controller/pki/secretbackendcrlconfig"
	secretbackendintermediatecertrequest "github.com/upbound/provider-vault/internal/controller/pki/secretbackendintermediatecertrequest"
	secretbackendintermediatesetsigned "github.com/upbound/provider-vault/internal/controller/pki/secretbackendintermediatesetsigned"
	secretbackendrolepki "github.com/upbound/provider-vault/internal/controller/pki/secretbackendrole"
	secretbackendrootcert "github.com/upbound/provider-vault/internal/controller/pki/secretbackendrootcert"
	secretbackendrootsignintermediate "github.com/upbound/provider-vault/internal/controller/pki/secretbackendrootsignintermediate"
	secretbackendsign "github.com/upbound/provider-vault/internal/controller/pki/secretbackendsign"
	providerconfig "github.com/upbound/provider-vault/internal/controller/providerconfig"
	leasecount "github.com/upbound/provider-vault/internal/controller/quota/leasecount"
	ratelimit "github.com/upbound/provider-vault/internal/controller/quota/ratelimit"
	secretbackendrabbitmq "github.com/upbound/provider-vault/internal/controller/rabbitmq/secretbackend"
	secretbackendrolerabbitmq "github.com/upbound/provider-vault/internal/controller/rabbitmq/secretbackendrole"
	autopilot "github.com/upbound/provider-vault/internal/controller/raft/autopilot"
	snapshotagentconfig "github.com/upbound/provider-vault/internal/controller/raft/snapshotagentconfig"
	policyrgp "github.com/upbound/provider-vault/internal/controller/rgp/policy"
	secretbackendca "github.com/upbound/provider-vault/internal/controller/ssh/secretbackendca"
	secretbackendrolessh "github.com/upbound/provider-vault/internal/controller/ssh/secretbackendrole"
	cloudsecretbackend "github.com/upbound/provider-vault/internal/controller/terraform/cloudsecretbackend"
	cloudsecretcreds "github.com/upbound/provider-vault/internal/controller/terraform/cloudsecretcreds"
	cloudsecretrole "github.com/upbound/provider-vault/internal/controller/terraform/cloudsecretrole"
	authbackendroletoken "github.com/upbound/provider-vault/internal/controller/token/authbackendrole"
	alphabet "github.com/upbound/provider-vault/internal/controller/transform/alphabet"
	role "github.com/upbound/provider-vault/internal/controller/transform/role"
	template "github.com/upbound/provider-vault/internal/controller/transform/template"
	transformation "github.com/upbound/provider-vault/internal/controller/transform/transformation"
	secretbackendkey "github.com/upbound/provider-vault/internal/controller/transit/secretbackendkey"
	audit "github.com/upbound/provider-vault/internal/controller/vault/audit"
	mount "github.com/upbound/provider-vault/internal/controller/vault/mount"
	policyvault "github.com/upbound/provider-vault/internal/controller/vault/policy"
	token "github.com/upbound/provider-vault/internal/controller/vault/token"
	vaultnamespace "github.com/upbound/provider-vault/internal/controller/vault/vaultnamespace"
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
