package controllerutils

import (
	"context"
	"reflect"
	"strings"
	"time"

	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	"github.com/difinative/Edge-Operator/controllers/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func InitStEdge(ce event.CreateEvent) {
	stEdgeobj := ce.Object
	stEdge := stEdgeobj.(*operatorv1.StandardEdge)

	if stEdge.Status.Vitals.UpOrDown == "" {
		ctrl.Log.Info("A new standard edge is created", "Edge Name", stEdge.Spec.Edgename)
		// e.Spec.Vitals.FreeMemory = "5G"

		// go utils.SetupEdge(os.Getenv("teleportUrl"), os.Getenv("identityKey"), os.Getenv("teleportUser"), os.Getenv("inferencServer"), e.Spec.Edgename)
	}
}

func DeleteScEdge(de event.DeleteEvent) {
	scEdgeobj := de.Object
	scEdge := scEdgeobj.(*operatorv1.StandardEdge)
	ctrl.Log.Info("Deleting a standard edge", "Edge Name", scEdge.Spec.Edgename)

}

func UpdateForScEdge(ue event.UpdateEvent) {
	if ue.ObjectNew != ue.ObjectOld {
		edge := ue.ObjectNew.(*operatorv1.ScEdge)
		old_edge := ue.ObjectOld.(*operatorv1.ScEdge)
		if !reflect.DeepEqual(edge.Status, old_edge.Status) {
			if utils.IsStrEquals(edge.Status.Vitals.TeleportStatus, utils.INACTIVE) {
				ctrl.Log.Info("Edge is Down >>>", "edge name", edge.Name, " status", edge.Status.Vitals.UpOrDown)
			}
			if edge.Status.Vitals.FreeMemory != "" {
				utils.CheckFreeMemory(edge.Status.Vitals.FreeMemory, edge.Spec.Vitals.FreeMemory, edge.Name)
			}
			if edge.Status.Vitals.Temperature != 0 {
				utils.CheckTemperature(edge.Status.Vitals.Temperature, edge.Spec.Vitals.Temperature, edge.Name)
			}
			if edge.Status.Vitals.TeleportStatus != "" {
				utils.CheckTeleport(edge.Status.Vitals.TeleportStatus, edge.Spec.Vitals.TeleportStatus, edge.Name)
			}
			if edge.Status.Vitals.UpOrDown == utils.DOWN {
				ctrl.Log.Info("Following edge is down>>>", "Name", edge.Name)
			}
			if edge.Status.Vitals.SqNet == utils.INACTIVE {
				ctrl.Log.Info("Following edge is not connected to sqnet>>>", "Name", edge.Name)
			}
		}
	}
}

func CheckLTU(edgeList operatorv1.StandardEdgeList, clt client.Client) {

	edges := edgeList.Items
	now := time.Now()
	for _, se := range edges {
		if se.Status.LTU == "" {
			continue
		}
		ltu, err := time.Parse(time.RFC850, se.Status.LTU)
		if err != nil {
			ctrl.Log.Error(err, "Error while trying to parse the LTU time", "edge name", se.Name, " LTU", se.Status.LTU)
		}
		if now.Sub(ltu).Minutes() > 3 {
			se.Status.Vitals.UpOrDown = utils.DOWN
			se.Status.Vitals.SqNet = utils.INACTIVE
			err := clt.Status().Update(context.TODO(), &se, &client.UpdateOptions{})
			for err != nil && errors.IsConflict(err) {
				err = clt.Update(context.TODO(), &se, &client.UpdateOptions{})
			}
		} else if strings.EqualFold(strings.ToLower(se.Status.Vitals.UpOrDown), strings.ToLower(utils.DOWN)) {
			se.Status.Vitals.UpOrDown = utils.UP
			se.Status.Vitals.SqNet = utils.ACTIVE
			err := clt.Status().Update(context.TODO(), &se, &client.UpdateOptions{})
			for err != nil && errors.IsConflict(err) {
				err = clt.Update(context.TODO(), &se, &client.UpdateOptions{})
			}
		}
	}
}
