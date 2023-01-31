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

func InitScEdge(ce event.CreateEvent) {
	scEdgeobj := ce.Object
	scEdge := scEdgeobj.(*operatorv1.ScEdge)

	if scEdge.Status.Vitals.UpOrDown == "" {
		ctrl.Log.Info("Creating a new smart chillar edge", "EDge Name", scEdge.Spec.Edgename)
		// e.Spec.Vitals.FreeMemory = "5G"

		// go utils.SetupEdge(os.Getenv("teleportUrl"), os.Getenv("identityKey"), os.Getenv("teleportUser"), os.Getenv("inferencServer"), e.Spec.Edgename)
	}
}

func DeleteScEdge(de event.DeleteEvent) {
	scEdgeobj := de.Object
	scEdge := scEdgeobj.(*operatorv1.ScEdge)
	ctrl.Log.Info("Deleting a smart chillar edge", "Edge Name", scEdge.Spec.Edgename)

}

func UpdateForScEdge(ue event.UpdateEvent) {
	if ue.ObjectNew != ue.ObjectOld {
		edge := ue.ObjectNew.(*operatorv1.ScEdge)
		old_edge := ue.ObjectOld.(*operatorv1.ScEdge)
		if !reflect.DeepEqual(edge.Status, old_edge.Status) {
			if strings.EqualFold(strings.ToLower(edge.Status.Vitals.UpOrDown), strings.ToLower(utils.DOWN)) {
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
			CameraCheck(&edge.Spec.Cameras, &edge.Status.Cameras, edge.Spec.Edgename)
		}
	}
}

func CameraCheck(camerasSpec *map[string]operatorv1.Camera, cameraStatus *map[string]operatorv1.Camera, name string) {
	// camerasSpec := edge.Spec.Cameras
	// cameraStatus := edge.Status.Cameras

	for cname, c := range *cameraStatus {
		cspec, isPresent := (*camerasSpec)[cname]
		if isPresent {
			if c.Resolution != cspec.Resolution {
				ctrl.Log.Info("Camera resolution did not match", "Edge", name, "Camera", cname, "Resolution", c.Resolution)
			}
		} else {
			ctrl.Log.Info("Following camera is not present in edge spec", "Camera", cname, "Resolution", c.Resolution, "Edge", name)
			ctrl.Log.Info("Please add the camera spec in edge CR")

			if c.Resolution != "1600x1200" {
				ctrl.Log.Info("Camera resolution is not 1600x1200", "Edge", name, "Camera", cname, "Resolution", c.Resolution)
			}
		}
		if strings.EqualFold(strings.ToLower(c.UpOrDown), strings.ToLower(utils.DOWN)) {
			ctrl.Log.Info("Camera is not working", "Edge", name, "Camera", cname)
		}
	}
}

func CheckLTU(edgeList operatorv1.ScEdgeList, clt client.Client) {

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
