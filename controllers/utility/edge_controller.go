package utility

import (
	"context"
	"fmt"
	"strings"
	"time"

	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	"github.com/difinative/Edge-Operator/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func HandleCreateEvent(edge operatorv1.Edge) {
	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Info("Error while getting the rest config", "Error", err)
		return
	}
	// edge := ce.Object.(*operatorv1.Edge)
	type_ := edge.Spec.Usecase
	clt, err := dynamic.NewForConfig(config)
	if err != nil {
		ctrl.Log.Info("Error while trying to get rest client", "Error", err)
		return
	}

	uc, err := utils.GetUsescasesCr(&clt)
	if err != nil && errors.IsNotFound(err) {
		ctrl.Log.Info("No Usecases resource found, creating new resource")
		uc := operatorv1.Usecases{}
		uc.Spec.Usecases = make(map[string][]string)
		ctrl.Log.Info("Creating usecases resources with Spec", "Usecase", type_, "Edge", edge.Name)

		uc.Spec.Usecases[type_] = []string{edge.Name}
		err := utils.CreateUsecasesCr(&clt, &uc)
		if err != nil {
			ctrl.Log.Error(err, "Error while trying to create Usecases resource")
		}
		return
	}
	ctrl.Log.Info("Usecases resource found, Updating the Usecases resource")
	ctrl.Log.Info("Checking the following usecase is there or not", "Usecase", type_)

	nameArr, isPresent := uc.Spec.Usecases[type_]
	if isPresent {
		i := checkEdgeInArr(nameArr, edge.Name)
		if i == -1 {
			ctrl.Log.Info("Usecase found in the resource")
			ctrl.Log.Info("Updating usecases resources with Spec", "Usecase", type_, "Edge", edge.Name)

			uc.Spec.Usecases[type_] = append(nameArr, edge.Name)
		}
	} else {
		ctrl.Log.Info("Usecase not found in the resource")
		ctrl.Log.Info("Updating usecases resources with Spec", "Usecase", type_, "Edge", edge.Name)
		if uc.Spec.Usecases != nil {
			uc.Spec.Usecases[type_] = []string{edge.Name}
		}
	}
	err = utils.UpdateUsecasesCr(&clt, &uc)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to Update Usecases resource")
	}
}

func HandleDeleteEvent(edge operatorv1.Edge) {
	ctrl.Log.Info("Deleting the following edge", "name", edge.Name)
	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Info("Error while getting the rest config", "Error", err)
		return
	}
	// edge := ce.Object.(*operatorv1.Edge)
	name := edge.Name

	type_ := edge.Spec.Usecase
	clt, err := dynamic.NewForConfig(config)
	if err != nil {
		ctrl.Log.Info("Error while trying to get rest client", "Error", err)
		return
	}
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

func UpdateEdgeUc(e operatorv1.Edge, prevUc string) {
	ctrl.Log.Info("Deleting the following edge", "name", e.Name)
	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Info("Error while getting the rest config", "Error", err)
		return
	}
	// edge := ce.Object.(*operatorv1.Edge)
	name := e.Name
	type_ := e.Spec.Usecase

	clt, err := dynamic.NewForConfig(config)
	if err != nil {
		ctrl.Log.Info("Error while trying to get rest client", "Error", err)
		return
	}

	ctrl.Log.Info("Getting usecases resource")
	uc, err := utils.GetUsescasesCr(&clt)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to get the Usecases")
		return
	}

	edgeArr, isPresent := uc.Spec.Usecases[type_]
	if isPresent {
		edgeArr = append(edgeArr, name)
		uc.Spec.Usecases[type_] = edgeArr
	} else {
		if uc.Spec.Usecases != nil {
			uc.Spec.Usecases[type_] = []string{name}
		}
	}
	prevEdgeArr := uc.Spec.Usecases[prevUc]
	i := checkEdgeInArr(prevEdgeArr, name)
	if i != -1 {
		prevEdgeArr = append(prevEdgeArr[:i], prevEdgeArr[i+1:]...)
	}
	uc.Spec.Usecases[prevUc] = prevEdgeArr

	utils.UpdateUsecasesCr(&clt, &uc)

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

func CheckLTU(edgeList operatorv1.EdgeList, clt client.Client) {

	edges := edgeList.Items
	now := time.Now().UTC()
	for _, se := range edges {
		if se.Status.LUT == "" {
			continue
		}
		ltu, err := time.Parse(time.RFC850, se.Status.LUT)
		if err != nil {
			ctrl.Log.Error(err, "Error while trying to parse the LTU time", "edge name", se.Name, " LTU", se.Status.LUT)
		}
		fmt.Println("Edge name: ", se.Name)
		fmt.Println("Edge LTU :", se.Status.LUT)
		fmt.Println("TIME NOW :", now)
		fmt.Println("DIFFERENCE :", now.Sub(ltu).Minutes())
		if now.Sub(ltu).Minutes() > 20 {
			se.Status.HealthVitals.UpOrDown = utils.DOWN
			se.Status.HealthVitals.SqNet = utils.INACTIVE
			err := clt.Status().Update(context.TODO(), &se, &client.UpdateOptions{})
			for err != nil && errors.IsConflict(err) {
				err = clt.Update(context.TODO(), &se, &client.UpdateOptions{})
			}
		} else if strings.EqualFold(strings.ToLower(se.Status.HealthVitals.UpOrDown), strings.ToLower(utils.DOWN)) {
			se.Status.HealthVitals.UpOrDown = utils.UP
			se.Status.HealthVitals.SqNet = utils.ACTIVE
			err := clt.Status().Update(context.TODO(), &se, &client.UpdateOptions{})
			for err != nil && errors.IsConflict(err) {
				err = clt.Update(context.TODO(), &se, &client.UpdateOptions{})
			}
		}
	}
}
