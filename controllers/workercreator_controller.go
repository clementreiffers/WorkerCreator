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
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"k8s.io/apimachinery/pkg/runtime"
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

func searchWorkerDefinition(name string) unstructured.Unstructured {
	workerDef := unstructured.Unstructured{}
	workerDef.SetKind("WorkerDefinition")
	workerDef.SetName(name)
	workerDef.SetAPIVersion("api.worker-definition/v1alpha1")
	workerDef.SetNamespace("default")
	return workerDef
}

func searchWorkerDeployment(name string) unstructured.Unstructured {
	workerDeployment := unstructured.Unstructured{}
	workerDeployment.SetKind("WorkerDeployment")
	workerDeployment.SetName(name)
	workerDeployment.SetAPIVersion("api.worker-deployment/v1alpha1")
	workerDeployment.SetNamespace("default")
	return workerDeployment
}

func (r *WorkerCreatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.Log.WithValues("PodInstanciator", req.NamespacedName)

	instance := &apiv1alpha1.WorkerCreator{}
	err := r.Get(ctx, req.NamespacedName, instance)

	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	workerDef := searchWorkerDefinition(instance.Spec.WorkerDefinitionId)
	workerDepl := searchWorkerDeployment(instance.Spec.WorkerDeploymentId)

	workerDefNamespacedName := types.NamespacedName{Name: instance.Spec.WorkerDefinitionId, Namespace: "default"}
	workerDeplNamespacedName := types.NamespacedName{Name: instance.Spec.WorkerDeploymentId, Namespace: "default"}

	err = r.Get(ctx, workerDefNamespacedName, &workerDef)
	if err != nil {
		return ctrl.Result{}, err
	}
	if !workerDef.GetDeletionTimestamp().IsZero() {
		// Le CRD est en cours de suppression
		return ctrl.Result{}, fmt.Errorf("WorkerDefinition %s is being deleted", workerDef.GetName())
	}

	err = r.Get(ctx, workerDeplNamespacedName, &workerDepl)
	if err != nil {
		return ctrl.Result{}, err
	}
	if !workerDef.GetDeletionTimestamp().IsZero() {
		// Le CRD est en cours de suppression
		return ctrl.Result{}, fmt.Errorf("WorkerDeployment %s is being deleted", workerDepl.GetName())
	}

	logger.Info("i found all CRD!")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WorkerCreatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.WorkerCreator{}).
		Complete(r)
}
