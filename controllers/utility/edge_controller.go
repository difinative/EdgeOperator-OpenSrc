package utility

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	"github.com/difinative/Edge-Operator/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func HandleCreateEvent(edge operatorv1.Edge) {

	ctrl.Log.Info(">>>>>>>> Handle create event called")
	type_ := edge.Spec.Usecase

	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Error(err, "!!!!! Error while getting the rest config !!!!!")
		return
	}

	clt, err := dynamic.NewForConfig(config)
	if err != nil {
		ctrl.Log.Error(err, "!!!!! Error while trying to get rest client !!!!!")
		return
	}

	uc, err := utils.GetUsescasesCr(&clt)
	if err != nil && errors.IsNotFound(err) {
		ctrl.Log.Info(">>>>No Usecases resource found, creating new resource<<<<")
		uc := operatorv1.Usecases{}
		uc.Spec.Usecases = make(map[string][]string)
		ctrl.Log.Info("Creating usecases resources with Spec", "Usecase", type_, "Edge", edge.Name)

		uc.Spec.Usecases[type_] = []string{edge.Name}
		err := utils.CreateUsecasesCr(&clt, &uc)
		if err != nil {
			ctrl.Log.Error(err, "!!!!! Error while trying to create Usecases resource !!!!!")
		}
		edge.Status.SqNet = utils.NA
		edge.Status.UpOrDown = utils.NA

		for i := 0; i < 10; i++ {
			err = utils.UpdateEdgeStatusCr(&clt, &edge)
			if err != nil {
				ctrl.Log.Error(err, "!!!!! Error while trying to Update Edge resource !!!!!")
			} else {
				break
			}
		}
		return
	}

	ctrl.Log.Info(">>>> Usecases resource found, Updating the Usecases resource <<<<")
	ctrl.Log.Info("Checking the following usecase is there or not", "Usecase", type_)

	nameArr, isPresent := uc.Spec.Usecases[type_]
	if isPresent {
		i := checkEdgeInArr(nameArr, edge.Name)
		if i == -1 {
			ctrl.Log.Info(">>>> Usecase found in the resource <<<<")
			ctrl.Log.Info("Updating usecases resources with Spec", "Usecase", type_, "Edge", edge.Name)

			uc.Spec.Usecases[type_] = append(nameArr, edge.Name)
		}
	} else {
		ctrl.Log.Info(">>>> Usecase not found in the resource <<<<")
		ctrl.Log.Info("Updating usecases resources with Spec", "Usecase", type_, "Edge", edge.Name)
		if uc.Spec.Usecases != nil {
			uc.Spec.Usecases[type_] = []string{edge.Name}
		} else {
			uc.Spec.Usecases = map[string][]string{type_: {edge.Name}}
		}
	}

	for i := 0; i < 10; i++ {
		err = utils.UpdateUsecasesCr(&clt, &uc)
		if err != nil {
			ctrl.Log.Error(err, "!!!!! Error while trying to Update Usecases resource !!!!!")
		} else {
			break
		}
	}

	edge.Status.SqNet = utils.NA
	edge.Status.UpOrDown = utils.NA

	for i := 0; i < 10; i++ {
		err = utils.UpdateEdgeStatusCr(&clt, &edge)
		if err != nil {
			ctrl.Log.Error(err, "Error while trying to Update Edge resource")
		} else {
			break
		}
	}

}

func HandleDeleteEvent(edge operatorv1.Edge) {
	ctrl.Log.Info(">>>>>>>> Handle delete event called")
	name := edge.Name
	type_ := edge.Spec.Usecase

	ctrl.Log.Info("!!! Deleting the following edge", "name", edge.Name)
	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Error(err, "!!!! Error while getting the rest config !!!!")
		return
	}

	clt, err := dynamic.NewForConfig(config)
	if err != nil {
		ctrl.Log.Error(err, "!!!! Error while trying to get rest client !!!!")
		return
	}

	ctrl.Log.Info("Getting usecases resource")
	uc, err := utils.GetUsescasesCr(&clt)
	if err != nil {
		ctrl.Log.Error(err, "!!!! Error while trying to get the Usecases !!!!")
		return
	}
	edgeArr := uc.Spec.Usecases[type_]
	ctrl.Log.Info("Iterating over the array of edge ...")
	i := checkEdgeInArr(edgeArr, name)
	if i != -1 {
		edgeArr = append(edgeArr[:i], edgeArr[i+1:]...)
		uc.Spec.Usecases[type_] = edgeArr

		for i := 0; i < 10; i++ {
			err = utils.UpdateUsecasesCr(&clt, &uc)
			if err != nil {
				ctrl.Log.Error(err, "!!!! Error while trying to update the usecases resource !!!!")
			} else {
				break
			}
		}

	}
}

func UpdateEdgeUc(e operatorv1.Edge, prevUc string) {

	ctrl.Log.Info(">>>> Handling the 'USECASE' update event for edge", "name", e.Name)
	name := e.Name
	type_ := e.Spec.Usecase

	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Error(err, "!!!! Error while getting the rest config !!!!")
		return
	}

	clt, err := dynamic.NewForConfig(config)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to get rest client")
		return
	}

	ctrl.Log.Info(">>> Getting usecases resource <<<")
	uc, err := utils.GetUsescasesCr(&clt)
	if err != nil {
		ctrl.Log.Error(err, "!!!! Error while trying to get the Usecases !!!!")
		return
	}

	edgeArr, isPresent := uc.Spec.Usecases[type_]
	ctrl.Log.Info("Checking if the new usecase is present in usecase resource")
	if isPresent {
		ctrl.Log.Info("Usecase found")
		ctrl.Log.Info("Placing the edge under its new usecase")
		edgeArr = append(edgeArr, name)
		uc.Spec.Usecases[type_] = edgeArr
	} else {
		ctrl.Log.Info("Usecase not found")
		if uc.Spec.Usecases != nil {
			ctrl.Log.Info("Making a new usecase section with edge name")
			uc.Spec.Usecases[type_] = []string{name}
		}
	}

	prevEdgeArr := uc.Spec.Usecases[prevUc]
	i := checkEdgeInArr(prevEdgeArr, name)
	if i != -1 {
		prevEdgeArr = append(prevEdgeArr[:i], prevEdgeArr[i+1:]...)
	}
	uc.Spec.Usecases[prevUc] = prevEdgeArr

	for i := 0; i < 10; i++ {
		err = utils.UpdateUsecasesCr(&clt, &uc)
		if err == nil {
			break
		}
		ctrl.Log.Error(err, "Error while trying to update the usecases resources")
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

func CheckLTU(edgeList operatorv1.EdgeList, clt client.Client) {
	eUp := []string{}
	eDown := []string{}
	edges := edgeList.Items
	now := time.Now().UTC()
	down := 0
	for _, se := range edges {
		if se.Status.LUT == "" {
			continue
		}
		ltu, err := time.Parse(time.RFC850, se.Status.LUT)
		if err != nil {
			ctrl.Log.Error(err, "!!!! Error while trying to parse the LUT time !!!!", "edge name", se.Name, " LTU", se.Status.LUT)
		}
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("Edge name: ", se.Name)
		fmt.Println("Edge LTU :", se.Status.LUT)
		fmt.Println("TIME NOW :", now)
		fmt.Println("DIFFERENCE :", now.Sub(ltu).Minutes())
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		if now.Sub(ltu).Minutes() > utils.LUT_TIME {
			if utils.IsStrEqual(se.Status.UpOrDown, utils.UP) {
				se.Status.UpOrDown = utils.DOWN
				se.Status.SqNet = utils.INACTIVE
				se.Status.HealthPercentage = "0"
				se.Status.VitalsStatsPercentage = "0"
				se.Status.Uptime = 0
				err := clt.Status().Update(context.TODO(), &se, &client.UpdateOptions{})
				for err != nil && errors.IsConflict(err) {
					err = clt.Update(context.TODO(), &se, &client.UpdateOptions{})
				}
				eDown = append(eDown, se.Name)
			}
		} else if strings.EqualFold(strings.ToLower(se.Status.UpOrDown), strings.ToLower(utils.DOWN)) {
			se.Status.UpOrDown = utils.UP
			se.Status.SqNet = utils.ACTIVE
			err := clt.Status().Update(context.TODO(), &se, &client.UpdateOptions{})
			for err != nil && errors.IsConflict(err) {
				err = clt.Update(context.TODO(), &se, &client.UpdateOptions{})
			}
			eUp = append(eUp, se.Name)
		}

		if strings.EqualFold(strings.ToLower(se.Status.UpOrDown), strings.ToLower(utils.DOWN)) || utils.IsStrEqual(se.Status.UpOrDown, utils.NA) {
			down++
		}

	}
	EdgeUp.Set(float64(len(edges) - down))
	EdgeDown.Set(float64(down))
	fmt.Println("Edge up gauge:", EdgeUp.Desc().String())
	fmt.Println("Edge down gauge:", EdgeDown.Desc().String())
	fmt.Println("Edge UP :", len(edges)-down)
	fmt.Println("Edge DOWN :", down)
	fmt.Println("----------------------------------------------------------------")
	if len(eDown) <= 0 && len(eUp) <= 0 {
		return
	}
	var wg sync.WaitGroup
	r := utils.WebHookReqBody{
		Down: make(map[string][]string),
		Up:   make(map[string][]string),
	}
	r.Down["The following edges are down"] = eDown
	r.Up["The following edge status updated to UP/active"] = eUp

	body, err := json.Marshal(r)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to marshal struct to json")
		return
	}
	wg.Add(1)
	go utils.Http_(utils.IFTTT_WEBHOOK, "POST", body, nil, &wg)

	if len(eDown) > 0 {
		wg.Add(1)
		go LogIncident(fmt.Sprintf("Following edges are 'DOWN': %s", eDown), "Edges are Down", &wg)
	}

}

func HandleEdgeUpdateEvent(e operatorv1.Edge) {
	ctrl.Log.Info(">>>> Handling update event for edge", "name", e.Name)
	ctrl.Log.Info("Checking if health percentage is present")
	if e.Status.HealthPercentage == "" {
		ctrl.Log.Info("Health percentage is empty!!!!!, for following edge", "edge", e.Name)
		return
	}
	ctrl.Log.Info("Parsing the helath persentage froms string to float")

	per, err := strconv.ParseFloat(e.Status.HealthPercentage, 64)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to parse the health percentage")
		return
	}
	if per < float64(e.Spec.HealthPercentage) {
		// var wg sync.WaitGroup

		ctrl.Log.Info("Following edge health is below expected threshold", "edge", e.Name, "health", e.Status.HealthPercentage)
		// body := []byte(fmt.Sprint("Following edge health is below expected threshold", "edge", e.Name, "health", e.Status.HealthPercentage))
		// wg.Add(1)
		// go utils.Http_(utils.IFTTT_WEBHOOK, "POST", body, nil, &wg)
		// wg.Add(1)
		// go LogIncident(fmt.Sprint("Following edge health is below expected threshold", "edge", e.Name, "health", e.Status.HealthPercentage), fmt.Sprint("HP: ", e.Name), &wg)
	}
}

func LogIncident(msg, title string, wg *sync.WaitGroup) {
	ctrl.Log.Info("Creating an incident log in squirrel ui")
	incidentLogBody := utils.IncidentDbBody{
		Title:        title,
		Description:  msg,
		Category:     "squirrel.assets.edge.down",
		SeverityType: "CRITICAL",
		IncidenType:  "ASSET",
		StatusType:   "OPEN",
		StatusEvents: "DEFAULT",
		InfoChannel:  "Squirrel Operator",
		ExternalLink: "eg.com",
	}
	body, err := json.Marshal(incidentLogBody)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to marshal struct to json")
		return
	}
	utils.Http_(utils.INCIDENT_LOG_API, "POST", body, map[string]string{"Authorization": utils.BEARER_TOKEN}, nil)
	wg.Done()
}
