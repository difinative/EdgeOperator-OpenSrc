package utils

import (
	"context"

	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	ctrl "sigs.k8s.io/controller-runtime"
)

var resource schema.GroupVersionResource = schema.GroupVersionResource{
	Group:    "operator.difinative",
	Version:  "v1",
	Resource: "",
}

func GetUsescasesCr(dynClient *dynamic.Interface) (operatorv1.Usecases, error) {
	resource.Resource = "usecases"
	uc := operatorv1.Usecases{}
	uc_obj, err := (*dynClient).Resource(resource).Namespace("default").Get(context.TODO(), "usecases", v1.GetOptions{})
	if err != nil {
		ctrl.Log.Info("Error while trying to get usecase", "Error", err)
		return uc, err
	}
	runtime.DefaultUnstructuredConverter.FromUnstructured(uc_obj.UnstructuredContent(), &uc)
	return uc, err
}

func UpdateUsecasesCr(dynClient *dynamic.Interface, uc *operatorv1.Usecases) error {
	resource.Resource = "usecases"
	uc_obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(uc)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to convert the Usecases object to unstructured object")
		return err
	}
	unstruct_obj := unstructured.Unstructured{Object: uc_obj}
	_, err = (*dynClient).Resource(resource).Namespace("default").Update(context.TODO(), &unstruct_obj, v1.UpdateOptions{})
	return err
}

func CreateUsecasesCr(dynClient *dynamic.Interface, uc *operatorv1.Usecases) error {
	resource.Resource = "usecases"
	uc.TypeMeta = v1.TypeMeta{
		Kind:       "Usecases",
		APIVersion: "operator.difinative/v1",
	}
	uc.ObjectMeta = v1.ObjectMeta{
		Name:      "usecases",
		Namespace: "default",
	}

	uc_obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(uc)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to convert the Usecases object to unstructured object")
		return err
	}
	unstruct_obj := unstructured.Unstructured{Object: uc_obj}
	_, err = (*dynClient).Resource(resource).Namespace("default").Create(context.TODO(), &unstruct_obj, v1.CreateOptions{})
	return err
}

func DeleteUcs(dynClient *dynamic.Interface, uc *operatorv1.Usecases) error {
	resource.Resource = "usecases"
	err := (*dynClient).Resource(resource).Namespace("default").Delete(context.TODO(), uc.Name, v1.DeleteOptions{})
	return err
}

func GetListUcs(dynClient *dynamic.Interface) (operatorv1.UsecasesList, error) {
	resource.Resource = "usecases"
	ucList := operatorv1.UsecasesList{}
	uc_obj, err := (*dynClient).Resource(resource).Namespace("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		ctrl.Log.Info("Error while trying to get usecase", "Error", err)
		return ucList, err
	}
	runtime.DefaultUnstructuredConverter.FromUnstructured(uc_obj.UnstructuredContent(), &ucList)
	return ucList, err
}

func DeleteUc_vitals(dynClient *dynamic.Interface, uc *operatorv1.UsecaseVitals) error {
	resource.Resource = "usecasevitals"
	err := (*dynClient).Resource(resource).Namespace("default").Delete(context.TODO(), uc.Name, v1.DeleteOptions{})
	return err
}

func GetListUc_vitals(dynClient *dynamic.Interface) (operatorv1.UsecaseVitalsList, error) {
	resource.Resource = "usecasevitals"
	ucList := operatorv1.UsecaseVitalsList{}
	uc_obj, err := (*dynClient).Resource(resource).Namespace("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		ctrl.Log.Info("Error while trying to get usecase", "Error", err)
		return ucList, err
	}
	runtime.DefaultUnstructuredConverter.FromUnstructured(uc_obj.UnstructuredContent(), &ucList)
	return ucList, err
}

func UpdateEdgeCr(dynClient *dynamic.Interface, e *operatorv1.Edge) error {
	resource.Resource = "edges"
	edge_obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(e)
	if err != nil {
		ctrl.Log.Error(err, "Error while trying to convert the Usecases object to unstructured object")
		return err
	}
	unstruct_obj := unstructured.Unstructured{Object: edge_obj}
	_, err = (*dynClient).Resource(resource).Namespace("default").Update(context.TODO(), &unstruct_obj, v1.UpdateOptions{})
	return err
}
