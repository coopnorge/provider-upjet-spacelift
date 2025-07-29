// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	attachment "github.com/coopnorge/provider-upjet-spacelift/internal/controller/context/attachment"
	context "github.com/coopnorge/provider-upjet-spacelift/internal/controller/context/context"
	environmentvariable "github.com/coopnorge/provider-upjet-spacelift/internal/controller/environmentvariable/environmentvariable"
	gcpserviceaccount "github.com/coopnorge/provider-upjet-spacelift/internal/controller/gcpserviceaccount/gcpserviceaccount"
	module "github.com/coopnorge/provider-upjet-spacelift/internal/controller/module/module"
	providerconfig "github.com/coopnorge/provider-upjet-spacelift/internal/controller/providerconfig"
	space "github.com/coopnorge/provider-upjet-spacelift/internal/controller/space/space"
	stack "github.com/coopnorge/provider-upjet-spacelift/internal/controller/stack/stack"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachment.Setup,
		context.Setup,
		environmentvariable.Setup,
		gcpserviceaccount.Setup,
		module.Setup,
		providerconfig.Setup,
		space.Setup,
		stack.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
