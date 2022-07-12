package helm

type ApplicationChart struct {
	values    values
	templates map[string]string
}

type values struct {
	App               *app                   `json:"app"`
	IngressController *IngressControllerSpec `json:"ingressController,omitempty1"`
}

type IngressControllerType string

func (t IngressControllerType) String() string { return string(t) }

const (
	TraefikIngressControllerType IngressControllerType = "traefik"
	IstioIngressControllerType   IngressControllerType = "istio"
	NginxIngressControllerType   IngressControllerType = "nginx"
)

type IngressControllerSpec struct {
	ClassName       string                `json:"className,omitempty"`
	ServiceEndpoint string                `json:"serviceEndpoint,omitempty"`
	IngressType     IngressControllerType `json:"type"`
	ClusterIssuer   string                `json:"clusterIssuer,omitempty"`
}

type app struct {
	//ID   string `json:"id"`
	Name string `json:"name"`
	/*
		Deployments []deployment  `json:"deployments"`
		Env         []ketchv1.Env `json:"env"`
		Ingress     ingress       `json:"ingress"`
		// IsAccessible if not set, ketch won't create kubernetes objects like Ingress/Gateway to handle incoming request.
		// These objects could be broken without valid routes to the application.
		// For example, "spec.rules" of an Ingress object must contain at least one rule.
		IsAccessible bool   `json:"isAccessible"`
		Group        string `json:"group"`
		Service      *gatewayService
		// MetadataLabels is a list of labels to be added to k8s resources.
		MetadataLabels []ketchv1.MetadataItem
		// MetadataAnnotations is a list of labels to be added to k8s resources.
		MetadataAnnotations []ketchv1.MetadataItem `json:"metadataAnnotations"`
		// ServiceAccountName specifies a service account name to be used for this application.
		// SA should exist.
		ServiceAccountName string `json:"serviceAccountName"`
		// SecurityContext specifies security settings for a pod/app, which get applied to all containers.
		SecurityContext *v1.PodSecurityContext `json:"securityContext,omitempty"`
		// VolumeClaimTemplates is a list of an app's volumeClaimTemplates
		VolumeClaimTemplates []ketchv1.PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"`
	*/
	// Type specifies whether the app should be a deployment or a statefulset
	Type AppType `json:"type"`
}

type AppType string

const (
	DeploymentAppType  AppType = "Deployment"
	StatefulSetAppType AppType = "StatefulSet"
)
