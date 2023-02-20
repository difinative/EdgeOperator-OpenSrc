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
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	v1 "github.com/difinative/Edge-Operator/api/v1"
	controllerutils "github.com/difinative/Edge-Operator/controllers/utils/Sc-Edge"
)

// ScEdgeReconciler reconciles a ScEdge object
type ScEdgeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=operator.difinative,resources=scedges,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=operator.difinative,resources=scedges/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=operator.difinative,resources=scedges/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ScEdge object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ScEdgeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	fmt.Println(">>Reconcile loop called<<<")

	// TODO(user): your logic here
	edgeList := v1.ScEdgeList{}
	err := r.Client.List(ctx, &edgeList, &client.ListOptions{})
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to get the list of standard edge")
	}
	controllerutils.CheckLTU(edgeList, r.Client)
	fmt.Println("")
	fmt.Println("")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ScEdgeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.ScEdge{}).
		WithEventFilter(predicate.Funcs{

			//Handle create event
			CreateFunc: func(ce event.CreateEvent) bool {
				controllerutils.InitScEdge(ce)
				return false
			},

			//Handle delete event
			DeleteFunc: func(de event.DeleteEvent) bool {
				controllerutils.DeleteScEdge(de)

				return false
			},

			//Handle update event
			UpdateFunc: func(ue event.UpdateEvent) bool {
				fmt.Println("***************************", ue.ObjectNew.GetName(), "******************************")
				if ue.ObjectNew == ue.ObjectOld {
					return true
				}
				fmt.Println(">>> Calling update event <<<")

				controllerutils.UpdateForScEdge(ue, r.Client)
				fmt.Println("")
				fmt.Println("")
				return false
			},
		}).Complete(r)
}
