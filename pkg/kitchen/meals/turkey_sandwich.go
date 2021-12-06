package meals

import (
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

type TurkeySandwich struct{ baseMeal }

func (t TurkeySandwich) Assemble() (v1alpha1.DAGTask, []v1alpha1.Template) {
	sliceBread := t.step("slice-bread")
	butterBread := t.step("butter-bread")
	assembleIngredients := t.step("assemble-ingredients")
	toast := t.step("toast")
	cutInHalf := t.step("cut-in-half")
	plate := t.step("plate")

	dagTasks := []v1alpha1.DAGTask{
		{
			Name:     sliceBread.Name,
			Template: sliceBread.Name,
		},
		{
			Name:     butterBread.Name,
			Template: butterBread.Name,
			Depends:  sliceBread.Name,
		},
		{
			Name:     assembleIngredients.Name,
			Template: assembleIngredients.Name,
			Depends:  butterBread.Name,
		},
		{
			Name:     toast.Name,
			Template: toast.Name,
			Depends:  assembleIngredients.Name,
		},
		{
			Name:     cutInHalf.Name,
			Template: cutInHalf.Name,
			Depends:  toast.Name,
		},
		{
			Name:     plate.Name,
			Template: plate.Name,
			Depends:  cutInHalf.Name,
		},
	}
	templates := []v1alpha1.Template{sliceBread, butterBread, assembleIngredients, toast, cutInHalf, plate}
	return t.assembleDAG(dagTasks, templates)
}
