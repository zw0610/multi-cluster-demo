/*
Copyright 2021.

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
	kubefloworgv1 "github.com/zw0610/multi-cluster-demo/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

// KubeflowJobReconciler reconciles a KubeflowJob object
type KubeflowJobReconciler struct {
	client.Client
	clientMapper map[string]client.Client
	Scheme       *runtime.Scheme
}

const (
	ClusterKey = "cluster"
)

//+kubebuilder:rbac:groups=kubeflow.org,resources=kubeflowjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubeflow.org,resources=kubeflowjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kubeflow.org,resources=kubeflowjobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *KubeflowJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	adminCtx := context.WithValue(ctx, ClusterKey, "lucas-admin")

	kfjob := &kubefloworgv1.KubeflowJob{}
	err := r.Get(adminCtx, req.NamespacedName, kfjob)

	if err != nil {
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, err
	}

	logger.Info("KubeflowJob Got:", "metadata", kfjob.ObjectMeta)

	workerCtx := context.WithValue(ctx, ClusterKey, "lucas-worker")
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: kfjob.Name,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				corev1.Container{
					Name:  "test",
					Image: "python:latest",
				},
			},
		},
	}

	err = r.Create(workerCtx, pod, nil)
	if err != nil {
		logger.Info("Failed to create pod on", ClusterKey, workerCtx.Value(ClusterKey))
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KubeflowJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubefloworgv1.KubeflowJob{}).
		Complete(r)
}
