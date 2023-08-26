/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	aksnp "github.com/spot/provider-spot/internal/controller/aksnp/aksnp"
	aksnpvirtualnodegroup "github.com/spot/provider-spot/internal/controller/aksnpvng/aksnpvirtualnodegroup"
	aws "github.com/spot/provider-spot/internal/controller/oceanaws/aws"
	awslaunchspec "github.com/spot/provider-spot/internal/controller/oceanawsvng/awslaunchspec"
	providerconfig "github.com/spot/provider-spot/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aksnp.Setup,
		aksnpvirtualnodegroup.Setup,
		aws.Setup,
		awslaunchspec.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
