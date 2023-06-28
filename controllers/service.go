package controllers

import (
	  corev1 "k8s.io/api/core/v1"
	  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createService(instanceName string) *corev1.Service {
	  return &corev1.Service{
			 ObjectMeta: metav1.ObjectMeta{
					Name:      getServiceName(instanceName),
					Namespace: "default",
			 },
			 Spec: corev1.ServiceSpec{
					Ports:     []corev1.ServicePort{},
					Selector:  map[string]string{"app": getPodName(instanceName)},
					ClusterIP: "None",
			 },
	  }
}
