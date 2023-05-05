package utility

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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
		} else {
			uc.Spec.Usecases = map[string][]string{type_: {edge.Name}}
		}
	}

	for i := 0; i < 10; i++ {
		err = utils.UpdateUsecasesCr(&clt, &uc)
		if err != nil {
			ctrl.Log.Error(err, "Error while trying to Update Usecases resource")
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

		for i := 0; i < 10; i++ {
			err = utils.UpdateUsecasesCr(&clt, &uc)
			if err != nil {
				ctrl.Log.Error(err, "Error while trying to update the usecases resource")
			} else {
				break
			}
		}

	}
}

func UpdateEdgeUc(e operatorv1.Edge, prevUc string) {
	ctrl.Log.Info("Handling the update event for edge", "name", e.Name)
	config, err := ctrl.GetConfig()
	if err != nil {
		ctrl.Log.Info("Error while getting the rest config", "Error", err)
		return
	}
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

	for i := 0; i < 10; i++ {
		err = utils.UpdateUsecasesCr(&clt, &uc)
		if err == nil {
			break
		}
		fmt.Println("Error while trying to update the usecases resources")
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
		if now.Sub(ltu).Minutes() > utils.LUT_TIME {
			if !utils.IsStrEqual(se.Status.UpOrDown, utils.DOWN) {
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
	}

	if len(eDown) <= 0 && len(eUp) <= 0 {
		return
	}

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
	go utils.Http_(utils.IFTTT_WEBHOOK, "POST", body, nil)

	EdgeUp.Set(float64(len(eUp)))
	EdgeDown.Set(float64(len(eDown)))

	if len(eDown) > 0 {
		go LogIncident(fmt.Sprintf("Following edges are 'DOWN': %s", eDown), "Edges are Down")
	}

}

func HandleEdgeUpdateEvent(e operatorv1.Edge) {
	if e.Status.HealthPercentage == "" {
		ctrl.Log.Info("Health percentage is empty!!!!!, for following edge", "edge", e.Name)
		return
	}
	per, err := strconv.ParseFloat(e.Status.HealthPercentage, 64)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to parse the health percentage")
		return
	}
	if per < float64(e.Spec.HealthPercentage) {
		ctrl.Log.Info("Following edge health is below expected threshold", "edge", e.Name, "health", e.Status.HealthPercentage)
		body := []byte(fmt.Sprint("Following edge health is below expected threshold", "edge", e.Name, "health", e.Status.HealthPercentage))
		go utils.Http_(utils.IFTTT_WEBHOOK, "POST", body, nil)
		go LogIncident(fmt.Sprint("Following edge health is below expected threshold", "edge", e.Name, "health", e.Status.HealthPercentage), fmt.Sprintf("HP: ", e.Name))
	}
}

func LogIncident(msg, title string) {
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
	utils.Http_(utils.INCIDENT_LOG_API, "POST", body, map[string]string{"Authorization": utils.BEARER_TOKEN})
}
