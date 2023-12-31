/*
Copyright 2022.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	"github.com/difinative/Edge-Operator/controllers/utility"
	"github.com/difinative/Edge-Operator/utils"
)

// UsecasesReconciler reconciles a Usecases object
type UsecasesReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=operator.difinative,resources=usecases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=operator.difinative,resources=usecases/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=operator.difinative,resources=usecases/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Usecases object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *UsecasesReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UsecasesReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1.Usecases{}).
		WithEventFilter(
			predicate.Funcs{
				CreateFunc: func(ce event.CreateEvent) bool {
					ctrl.Log.Info(">>>>>>>> Create event called for ", "Usecase", ce.Object.GetName())

					ucs := ce.Object.(*operatorv1.Usecases)

					clt, err := utils.GetDynClt()
					if err != nil {
						return false
					}
					ucList, _ := utils.GetListUcs(&clt)
					if len(ucList.Items) > 1 {
						utils.DeleteUcs(&clt, ucs)
					}
					return false
				},

				DeleteFunc: func(ce event.DeleteEvent) bool { return false },

				UpdateFunc: func(ue event.UpdateEvent) bool {
					ctrl.Log.Info(">>>>>>>> Update event called for ", "Usecase", ue.ObjectOld.GetName())

					ucs := ue.ObjectNew.(*operatorv1.Usecases)
					clt, err := utils.GetDynClt()
					if err != nil {
						return false
					}

					utility.CheckAndDeleteEmptyKeys(&clt, ucs)
					// if err != nil {
					// 	return false
					// }
					return true
				},
			},
		).
		Complete(r)
}
