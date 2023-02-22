package utility

import (
	"strings"

	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	"github.com/difinative/Edge-Operator/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
)

func HandleCreateEvent(edge operatorv1.Edge) {
	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Info("Error while getting the rest config", "Error", err)
	} else {
		// edge := ce.Object.(*operatorv1.Edge)
		type_ := edge.Spec.Usecase
		clt, err := dynamic.NewForConfig(config)
		if err != nil {
			ctrl.Log.Info("Error while trying to get rest client", "Error", err)
		} else {
			uc, err := utils.GetUsescasesCr(&clt)

			if err != nil && errors.IsNotFound(err) {
				ctrl.Log.Info("No Usecases resource found, creating new resource")
				uc := operatorv1.Usecases{}
				uc.Spec.Usecases = make(map[string][]string)
				ctrl.Log.Info("Creating usecases resources with Spec", "Usecase", type_, "Edge", edge.Spec.Name)

				uc.Spec.Usecases[type_] = []string{edge.Spec.Name}
				err := utils.CreateUsecasesCr(&clt, &uc)
				if err != nil {
					ctrl.Log.Error(err, "Error while trying to create Usecases resource")
				}
			} else {
				ctrl.Log.Info("Usecases resource found, Updating the Usecases resource")
				ctrl.Log.Info("Checking the following usecase is there or not", "Usecase", type_)

				nameArr, isPresent := uc.Spec.Usecases[type_]
				if isPresent {
					ctrl.Log.Info("Usecase found in the resource")
					ctrl.Log.Info("Updating usecases resources with Spec", "Usecase", type_, "Edge", edge.Spec.Name)

					uc.Spec.Usecases[type_] = append(nameArr, edge.Spec.Name)
				} else {
					i := checkEdgeInArr(nameArr, edge.Spec.Name)
					if i == -1 {
						ctrl.Log.Info("Usecase found in the resource")
						ctrl.Log.Info("Updating usecases resources with Spec", "Usecase", type_, "Edge", edge.Spec.Name)
						if uc.Spec.Usecases != nil {
							uc.Spec.Usecases[type_] = []string{edge.Spec.Name}
						}
					}
				}
				err = utils.UpdateUsecasesCr(&clt, &uc)
				if err != nil {
					ctrl.Log.Error(err, "Error while trying to Update Usecases resource")
				}
			}
		}
	}
}

func HandleDeleteEvent(edge operatorv1.Edge) {
	ctrl.Log.Info("Deleting the following edge", "name", edge.Name)
	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Info("Error while getting the rest config", "Error", err)
	} else {
		// edge := ce.Object.(*operatorv1.Edge)
		name := edge.Name

		type_ := edge.Spec.Usecase
		clt, err := dynamic.NewForConfig(config)
		if err != nil {
			ctrl.Log.Info("Error while trying to get rest client", "Error", err)
		} else {
			ctrl.Log.Info("Getting usecases resource")
			uc, err := utils.GetUsescasesCr(&clt)
			if err != nil {
				ctrl.Log.Error(err, "Error while trying to get the Usecases")
				return
			}
			edgeArr := uc.Spec.Usecases[type_]
			ctrl.Log.Info("Iterating over the array of edge")
			i := checkEdgeInArr(edgeArr, name)
			if i != -1 {
				edgeArr = append(edgeArr[:i], edgeArr[i+1:]...)
				uc.Spec.Usecases[type_] = edgeArr
				err = utils.UpdateUsecasesCr(&clt, &uc)
				if err != nil {
					ctrl.Log.Error(err, "Error while trying to update the usecases resource")
				}
			}
		}
	}
}

func checkEdgeInArr(edgeArr []string, name string) int {
	for i, v := range edgeArr {
		if strings.EqualFold(name, v) {
			ctrl.Log.Info("Found the edge at following index", "Index", i)
			return i
		}
	}
	return -1
}
