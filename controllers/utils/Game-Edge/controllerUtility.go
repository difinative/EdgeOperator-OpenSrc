package controllerutils

import (
	"strconv"
	"strings"

	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/event"
)

func InitGameEdge(ce event.CreateEvent) {
	gameEdgeobj := ce.Object
	gameEdge := gameEdgeobj.(*operatorv1.GameEdge)

	if gameEdge.Status.Vitals.UpOrDown == "" {
		ctrl.Log.Info("Creating a new smart chillar edge", "EDge Name", gameEdge.Spec.Edgename)
		// e.Spec.Vitals.FreeMemory = "5G"

		// go utils.SetupEdge(os.Getenv("teleportUrl"), os.Getenv("identityKey"), os.Getenv("teleportUser"), os.Getenv("inferencServer"), e.Spec.Edgename)
	}
}

func DeleteGameEdge(de event.DeleteEvent) {
	gameEdgeobj := de.Object
	gameEdge := gameEdgeobj.(*operatorv1.GameEdge)
	ctrl.Log.Info("Deleting a smart chillar edge", "Edge Name", gameEdge.Spec.Edgename)

}

func UpdateForGameEdge(ue event.UpdateEvent) {
	if ue.ObjectNew != ue.ObjectOld {
		edge := ue.ObjectNew.(*operatorv1.GameEdge)
		if edge.Status.Vitals.UpOrDown == "Down" {
			ctrl.Log.Info("Edge is Down >>>", "edge name", edge.Name, " status", edge.Status.Vitals.UpOrDown)
		} else if edge.Status.Vitals.FreeMemory != "" {
			availableMemory := string(edge.Status.Vitals.FreeMemory[:len(edge.Status.Vitals.FreeMemory)-1])
			am, err := strconv.Atoi(availableMemory)
			if err != nil {
				ctrl.Log.Info("Erro while tring to convert available memory from string to int")
				ctrl.Log.Info(">>>", "Error", err)
			}
			availableMemoryInSpec := string(edge.Spec.Vitals.FreeMemory[:len(edge.Spec.Vitals.FreeMemory)-1])
			amspec, err := strconv.Atoi(availableMemoryInSpec)
			if err != nil {
				ctrl.Log.Info("Erro while tring to convert available memory from string to int")
				ctrl.Log.Info(">>>", "Error", err)
			}
			if am < amspec {
				ctrl.Log.Info("Edge has available memory less than expected >>>", "edge name", edge.Name, "expected", edge.Spec.Vitals.FreeMemory, "actual", edge.Status.Vitals.FreeMemory)
			}
		}
		CameraCheck(&edge.Spec.Cameras, &edge.Status.Cameras, edge.Spec.Edgename)
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
		if strings.EqualFold(strings.ToLower(c.UpOrDown), strings.ToLower("Down")) {
			ctrl.Log.Info("Camera is not working", "Edge", name, "Camera", cname)
		}
	}
}
