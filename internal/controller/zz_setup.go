/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	secretbackend "github.com/upbound/provider-vault/internal/controller/ad/secretbackend"
	secretrole "github.com/upbound/provider-vault/internal/controller/ad/secretrole"
	authbackendrole "github.com/upbound/provider-vault/internal/controller/alicloud/authbackendrole"
	authbackendlogin "github.com/upbound/provider-vault/internal/controller/approle/authbackendlogin"
	providerconfig "github.com/upbound/provider-vault/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secretbackend.Setup,
		secretrole.Setup,
		authbackendrole.Setup,
		authbackendlogin.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
