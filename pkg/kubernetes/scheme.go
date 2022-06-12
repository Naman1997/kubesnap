package kubernetes

import (
	appsv1 "github.com/openshift/api/apps/v1"
	authorizationv1 "github.com/openshift/api/authorization/v1"
	buildv1 "github.com/openshift/api/build/v1"
	imagev1 "github.com/openshift/api/image/v1"
	networkv1 "github.com/openshift/api/network/v1"
	oauthv1 "github.com/openshift/api/oauth/v1"
	projectv1 "github.com/openshift/api/project/v1"
	quotav1 "github.com/openshift/api/quota/v1"
	routev1 "github.com/openshift/api/route/v1"
	securityv1 "github.com/openshift/api/security/v1"
	templatev1 "github.com/openshift/api/template/v1"
	userv1 "github.com/openshift/api/user/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/runtime/serializer/versioning"
	"k8s.io/client-go/kubernetes/scheme"
)

const version = "v1"

func GenerateCodec() runtime.Codec {
	scheme := scheme.Scheme
	AddToScheme(scheme)
	serializer := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme, scheme)
	return versioning.NewDefaultingCodecForScheme(
		scheme,
		serializer,
		serializer,
		schema.GroupVersion{Version: version},
		runtime.InternalGroupVersioner,
	)
}

func AddToScheme(scheme *runtime.Scheme) {
	appsv1.AddToScheme(scheme)
	authorizationv1.AddToScheme(scheme)
	buildv1.AddToScheme(scheme)
	imagev1.AddToScheme(scheme)
	networkv1.AddToScheme(scheme)
	oauthv1.AddToScheme(scheme)
	projectv1.AddToScheme(scheme)
	quotav1.AddToScheme(scheme)
	routev1.AddToScheme(scheme)
	securityv1.AddToScheme(scheme)
	templatev1.AddToScheme(scheme)
	userv1.AddToScheme(scheme)
	appsv1.AddToSchemeInCoreGroup(scheme)
	authorizationv1.AddToSchemeInCoreGroup(scheme)
	buildv1.AddToSchemeInCoreGroup(scheme)
	imagev1.AddToSchemeInCoreGroup(scheme)
	networkv1.AddToSchemeInCoreGroup(scheme)
	oauthv1.AddToSchemeInCoreGroup(scheme)
	projectv1.AddToSchemeInCoreGroup(scheme)
	quotav1.AddToSchemeInCoreGroup(scheme)
	routev1.AddToSchemeInCoreGroup(scheme)
	securityv1.AddToSchemeInCoreGroup(scheme)
	templatev1.AddToSchemeInCoreGroup(scheme)
	userv1.AddToSchemeInCoreGroup(scheme)
}