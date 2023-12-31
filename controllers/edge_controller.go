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

// EdgeReconciler reconciles a Edge object
type EdgeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=operator.difinative,resources=edges,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=operator.difinative,resources=edges/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=operator.difinative,resources=edges/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Edge object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *EdgeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	ctrl.Log.Info("Running the reconcile loop")
	edgeList := operatorv1.EdgeList{}
	error := r.Client.List(ctx, &edgeList, &client.ListOptions{})
	if error != nil {
		ctrl.Log.Error(error, "!!!! Error while trying to get the edge list !!!!")
	}
	utility.CheckLTU(edgeList, r.Client)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EdgeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1.Edge{}).
		WithEventFilter(
			predicate.Funcs{
				CreateFunc: func(ce event.CreateEvent) bool {
					ctrl.Log.Info(">>>>>>>> Create event called for ", "Edge", ce.Object.GetName())
					edge := ce.Object.(*operatorv1.Edge)
					utility.HandleCreateEvent(*edge)
					return false
				},
				DeleteFunc: func(de event.DeleteEvent) bool {

					ctrl.Log.Info(">>>>>>>> Create event called for ", "Edge", de.Object.GetName())

					edge := de.Object.(*operatorv1.Edge)
					utility.HandleDeleteEvent(*edge)
					return false
				},

				UpdateFunc: func(ue event.UpdateEvent) bool {
					ctrl.Log.Info(">>>>>>>> Create event called for ", "Edge", ue.ObjectOld.GetName())

					oldEdge := ue.ObjectOld.(*operatorv1.Edge)
					newEdge := ue.ObjectNew.(*operatorv1.Edge)

					if !utils.IsStrEqual(oldEdge.Spec.Usecase, newEdge.Spec.Usecase) {
						utility.UpdateEdgeUc(*newEdge, oldEdge.Spec.Usecase)
						return false
					}
					utility.HandleEdgeUpdateEvent(*newEdge)
					return true
				},
			},
		).
		Complete(r)
}
