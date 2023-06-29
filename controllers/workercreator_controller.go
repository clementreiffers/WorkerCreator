/*
Copyright 2023 clementreiffers.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	apiv1alpha1 "operators/WorkerCreator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// WorkerCreatorReconciler reconciles a WorkerCreator object
type WorkerCreatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=api.worker-creator,resources=workercreators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=api.worker-creator,resources=workercreators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=api.worker-creator,resources=workercreators/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the WorkerCreator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile

func applyResource(r *WorkerCreatorReconciler, ctx context.Context, resource client.Object, foundResource client.Object) error {
	err := r.Get(ctx, types.NamespacedName{Name: resource.GetName(), Namespace: resource.GetNamespace()}, foundResource)
	if err != nil && errors.IsNotFound(err) {
		err = r.Create(ctx, resource)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

func (r *WorkerCreatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//logger := log.Log.WithValues("PodInstanciator", req.NamespacedName)

	instance := &apiv1alpha1.WorkerCreator{}
	err := r.Get(ctx, req.NamespacedName, instance)

	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	//accounts := instance.Spec.Accounts
	//scripts := instance.Spec.Scripts
	//
	//instanceName := fmt.Sprintf("worker-%s-%s", accounts, scripts)
	//pod := createPod(instanceName, fmt.Sprintf("%s", image), ports)
	//svc := createService(instanceName)
	//ing := createIngress(ports, instanceName)
	//
	//err = applyResource(r, ctx, pod, &corev1.Pod{})
	//if err != nil {
	//	logger.Error(err, "unable to create Pod")
	//	return ctrl.Result{}, err
	//}
	//err = applyResource(r, ctx, svc, &corev1.Service{})
	//if err != nil {
	//	logger.Error(err, "unable to create Service")
	//	return ctrl.Result{}, err
	//}
	//err = applyResource(r, ctx, ing, &networkingv1.Ingress{})
	//if err != nil {
	//	logger.Error(err, "unable to create Ingress")
	//	return ctrl.Result{}, err
	//}
	//logger.Info("all resources created!")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WorkerCreatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.WorkerCreator{}).
		Complete(r)
}
