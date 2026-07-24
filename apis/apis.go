/*
Copyright 2021 Upbound Inc.
*/

package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	clusterapis "github.com/coopnorge/provider-upjet-spacelift/apis/cluster"
	namespacedapis "github.com/coopnorge/provider-upjet-spacelift/apis/namespaced"
	"github.com/coopnorge/provider-upjet-spacelift/apis/v1alpha1"
	"github.com/coopnorge/provider-upjet-spacelift/apis/v1beta1"
)

func init() {
	AddToSchemes = append(AddToSchemes,
		clusterapis.AddToScheme,
		namespacedapis.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
