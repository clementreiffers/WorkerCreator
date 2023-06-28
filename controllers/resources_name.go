package controllers

import apiv1alpha1 "operators/WorkerCreator/api/v1alpha1"

func getPodName(instance string) string {
	return instance + "-pod"
}

func getServiceName(instance string) string {
	return instance + "-svc"
}

func getIngressName(instance string) string {
	return instance + "-ingress"
}

func getIngressPathName(port apiv1alpha1.Port) string {
	return "/" + port.PortName
}
