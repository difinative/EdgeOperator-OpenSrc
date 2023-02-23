package utils

import (
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
)

func GetDynClt() (dynamic.Interface, error) {
	config, err := ctrl.GetConfig()
	var clt dynamic.Interface
	if err != nil {
		ctrl.Log.Info("Error while getting the rest config", "Error", err)
		return clt, err
	} else {
		clt, err := dynamic.NewForConfig(config)
		return clt, err
	}
}
