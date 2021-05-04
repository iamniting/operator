/*


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

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	stablev1 "github.com/iamniting/operator/api/v1"
)

// AppConfigReconciler reconciles a AppConfig object
type AppConfigReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=stable.example.com,resources=appconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=stable.example.com,resources=appconfigs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

func (r *AppConfigReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("appconfig", req.NamespacedName)

	instance := &stablev1.AppConfig{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Info("appconfig resource not found")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	logger.Info("appconfig resource found")

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": instance.Name,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:            instance.Name,
					Image:           instance.Spec.Image,
					ImagePullPolicy: corev1.PullIfNotPresent,
				},
			},
		},
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: func(n int32) *int32 { return &n }(int32(instance.Spec.DesiredCount)),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": instance.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: pod.ObjectMeta,
				Spec:       pod.Spec,
			},
		},
	}

	err = r.Client.Create(context.TODO(), deployment)

	return ctrl.Result{}, err
}

func (r *AppConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&stablev1.AppConfig{}).
		Complete(r)
}
