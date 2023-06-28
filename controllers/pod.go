package controllers

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1alpha1 "operators/WorkerCreator/api/v1alpha1"
)

func createPodPorts(ports []apiv1alpha1.Port) []corev1.ContainerPort {
	podPorts := make([]corev1.ContainerPort, len(ports))
	for i, port := range ports {
		podPorts[i] = corev1.ContainerPort{
			Name:          port.PortName,
			ContainerPort: port.PortNumber,
		}
	}
	return podPorts
}

func createPod(instanceName string, image string, ports []apiv1alpha1.Port) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      getPodName(instanceName),
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  getPodName(instanceName),
					Image: image,
					Ports: createPodPorts(ports),
				},
			},
		},
	}
}
