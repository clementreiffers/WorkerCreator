package controllers

import (
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1alpha1 "operators/WorkerCreator/api/v1alpha1"
)

func createIngressPaths(ports []apiv1alpha1.Port, instanceName string) []networkingv1.HTTPIngressPath {
	paths := make([]networkingv1.HTTPIngressPath, len(ports))
	pathType := networkingv1.PathTypePrefix
	for i, port := range ports {
		paths[i] = networkingv1.HTTPIngressPath{
			Path:     getIngressPathName(port),
			PathType: &pathType,
			Backend: networkingv1.IngressBackend{
				Service: &networkingv1.IngressServiceBackend{
					Name: getServiceName(instanceName),
					Port: networkingv1.ServiceBackendPort{
						Number: port.PortNumber,
					},
				},
			},
		}
	}
	return paths
}

func createIngress(ports []apiv1alpha1.Port, instanceName string) *networkingv1.Ingress {
	return &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      getIngressName(instanceName),
			Namespace: "default",
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/rewrite-target": "/",
			},
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: "worker.127.0.0.1.sslip.io",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: createIngressPaths(ports, instanceName),
						},
					},
				},
			},
		},
	}
}
