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

package v1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var appconfiglog = logf.Log.WithName("appconfig-resource")

func (r *AppConfig) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-stable-example-com-v1-appconfig,mutating=true,failurePolicy=fail,groups=stable.example.com,resources=appconfigs,verbs=create;update,versions=v1,name=mappconfig.kb.io

var _ webhook.Defaulter = &AppConfig{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *AppConfig) Default() {
	appconfiglog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-stable-example-com-v1-appconfig,mutating=false,failurePolicy=fail,groups=stable.example.com,resources=appconfigs,versions=v1,name=vappconfig.kb.io

var _ webhook.Validator = &AppConfig{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AppConfig) ValidateCreate() error {
	appconfiglog.Info("validate create", "name", r.Name)

	return r.ValidateAppConfig()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AppConfig) ValidateUpdate(old runtime.Object) error {
	appconfiglog.Info("validate update", "name", r.Name)

	return r.ValidateAppConfig()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AppConfig) ValidateDelete() error {
	appconfiglog.Info("validate delete", "name", r.Name)

	return r.ValidateAppConfig()
}

func (r *AppConfig) ValidateAppConfig() error {
	appconfiglog.Info("validate delete", "name", r.Name)

	if r.Spec.DesiredCount > 5 {
		return fmt.Errorf("Spec.DesiredCount should not be more than 5")
	}

	if r.Spec.Image == "" {
		return fmt.Errorf("Spec.Image can not be empty")
	}

	return nil
}
