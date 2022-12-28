package controllerutils

import (
	"strconv"

	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/event"
)

func InitStEdge(ce event.CreateEvent) {
	stEdgeobj := ce.Object
	stEdge := stEdgeobj.(*operatorv1.ScEdge)

	if stEdge.Status.Vitals.UpOrDown == "" {
		ctrl.Log.Info("A new standard edge is created", "Edge Name", stEdge.Spec.Edgename)
		// e.Spec.Vitals.FreeMemory = "5G"

		// go utils.SetupEdge(os.Getenv("teleportUrl"), os.Getenv("identityKey"), os.Getenv("teleportUser"), os.Getenv("inferencServer"), e.Spec.Edgename)
	}
}

func DeleteScEdge(de event.DeleteEvent) {
	scEdgeobj := de.Object
	scEdge := scEdgeobj.(*operatorv1.ScEdge)
	ctrl.Log.Info("Deleting a standard edge", "Edge Name", scEdge.Spec.Edgename)

}

func UpdateForScEdge(ue event.UpdateEvent) {
	if ue.ObjectNew != ue.ObjectOld {
		edge := ue.ObjectNew.(*operatorv1.ScEdge)
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
	}
}
