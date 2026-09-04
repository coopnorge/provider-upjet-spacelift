// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0


package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	environmentvariable "github.com/coopnorge/provider-upjet-spacelift/internal/controller/cluster/environmentvariable/environmentvariable"
gcpserviceaccount "github.com/coopnorge/provider-upjet-spacelift/internal/controller/cluster/gcpserviceaccount/gcpserviceaccount"
module "github.com/coopnorge/provider-upjet-spacelift/internal/controller/cluster/module/module"
space "github.com/coopnorge/provider-upjet-spacelift/internal/controller/cluster/space/space"
stack "github.com/coopnorge/provider-upjet-spacelift/internal/controller/cluster/stack/stack"
providerconfig "github.com/coopnorge/provider-upjet-spacelift/internal/controller/providerconfig"
attachment "github.com/coopnorge/provider-upjet-spacelift/internal/controller/cluster/context/attachment"
context "github.com/coopnorge/provider-upjet-spacelift/internal/controller/cluster/context/context"

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
		space.Setup,
		stack.Setup,
		providerconfig.Setup,
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
		attachment.SetupGated,
		context.SetupGated,
		environmentvariable.SetupGated,
		gcpserviceaccount.SetupGated,
		module.SetupGated,
		space.SetupGated,
		stack.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		attachment.SetupWebhookWithManager,
		context.SetupWebhookWithManager,
		environmentvariable.SetupWebhookWithManager,
		gcpserviceaccount.SetupWebhookWithManager,
		module.SetupWebhookWithManager,
		space.SetupWebhookWithManager,
		stack.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}