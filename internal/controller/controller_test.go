/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	"testing"
	"time"

	"github.com/go-logr/logr"

	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/coopnorge/provider-upjet-spacelift/apis"
	"github.com/coopnorge/provider-upjet-spacelift/config"
	"github.com/coopnorge/provider-upjet-spacelift/internal/clients"
)

// TestSetupDoesNotPanic is a generic test that boots the provider's
// controller the same way the real binary does and asserts it completes
// without erroring.

func TestSetupDoesNotPanic(t *testing.T) {
	mgr, err := ctrl.NewManager(&rest.Config{Host: "http://127.0.0.1:1", QPS: 10, Burst: 20}, ctrl.Options{})
	if err != nil {
		t.Fatalf("ctrl.NewManager: %v", err)
	}
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		t.Fatalf("apis.AddToScheme: %v", err)
	}

	log := logging.NewLogrLogger(logr.Discard())

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()

	o := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(10),
			PollInterval:            10 * time.Minute,
			MaxConcurrentReconciles: 10,
			Features:                &feature.Flags{},
			MetricOptions: &xpcontroller.MetricOptions{
				PollStateMetricInterval: 5 * time.Second,
				MRMetrics:               metricRecorder,
				MRStateMetrics:          stateMetrics,
			},
		},
		Provider:       config.GetProvider(),
		WorkspaceStore: terraform.NewWorkspaceStore(log),
		SetupFn:        clients.TerraformSetupBuilder("1.7.0", "spacelift-io/terraform-provider-spacelift", "v1.53.5"),
	}

	if err := Setup(mgr, o); err != nil {
		t.Fatalf("controller.Setup (main.go boot wiring) failed: %v", err)
	}
}
