package utils

import (
	"strconv"
	"strings"

	ctrl "sigs.k8s.io/controller-runtime"
)

func IsStrEquals(str1, str2 string) bool {
	return strings.EqualFold(strings.ToLower(str1), strings.ToLower(str2))
}

func CheckFreeMemory(statusMem, specMem, edgeName string) {
	availableMemory := string(statusMem[:len(statusMem)-1])
	am, err := strconv.Atoi(availableMemory)
	if err != nil {
		ctrl.Log.Info("Erro while tring to convert available memory from string to int")
		ctrl.Log.Info(">>>", "Error", err)
	}
	availableMemoryInSpec := string(specMem[:len(specMem)-1])
	amspec, err := strconv.Atoi(availableMemoryInSpec)
	if err != nil {
		ctrl.Log.Info("Erro while tring to convert available memory from string to int")
		ctrl.Log.Info(">>>", "Error", err)
	}
	if am < amspec {
		ctrl.Log.Info("Edge has available memory less than expected >>>", "edge name", edgeName, "expected", specMem, "actual", statusMem)
	}
}

func CheckTemperature(statusTemp, specTemp int, edgeName string) {
	if statusTemp > specTemp {
		ctrl.Log.Info("Edge is heating >>>", "edge name", edgeName, "standard temperature", specTemp, "actual", statusTemp)
	}
}

func CheckTeleport(statusTeleport, specTeleport, edgeName string) {
	status := ACTIVE
	if specTeleport != "" {
		status = specTeleport
	}
	if !IsStrEquals(statusTeleport, status) {
		ctrl.Log.Info("Teleport is not running on the following edge >>>", "edge name", edgeName)
	}
}
