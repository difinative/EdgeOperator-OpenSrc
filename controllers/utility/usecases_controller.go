package utility

import (
	operatorv1 "github.com/difinative/Edge-Operator/api/v1"
	"github.com/difinative/Edge-Operator/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
)

func CheckAndDeleteEmptyKeys(clt *dynamic.Interface, ucs *operatorv1.Usecases) {
	for k, v := range ucs.Spec.Usecases {
		if !utils.IsArrEmpty(v) {
			continue
		}
		ctrl.Log.Info("Following usecase has no edge", "Usecase", k)
		ctrl.Log.Info("Getting the list of user defined usecase")
		ucList, err := utils.GetListUc_vitals(clt)
		if err != nil {
			ctrl.Log.Error(err, "Error while trying to get the list fo usecase")
			delete(ucs.Spec.Usecases, k)
			continue
		}
		for _, u := range ucList.Items {
			if utils.IsStrEqual(k, u.Spec.Usecase) {
				ctrl.Log.Info("Found the usecase with the required type", "Usecase", k)
				err = utils.DeleteUc_vitals(clt, &u)
				if err != nil {
					ctrl.Log.Error(err, "Error while trying to delete the usecase")
				}
				break
			}
		}
		ctrl.Log.Info("Deleting the key from the Usecases map")
		delete(ucs.Spec.Usecases, k)
	}
	ctrl.Log.Info("Updating the usecases resource with new key-value pair")

	ucs.TypeMeta = v1.TypeMeta{
		Kind:       "Usecases",
		APIVersion: "operator.difinative/v1",
	}

	for i := 0; i < 10; i++ {
		err := utils.UpdateUsecasesCr(clt, ucs)
		if err != nil {
			ctrl.Log.Error(err, "Error while trying to update the usecases resource")
		} else {
			break
		}
	}

}
