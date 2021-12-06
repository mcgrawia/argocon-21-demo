package meals

import (
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"ianmcgraw.com/m/v2/pkg/workflows"
)

type EggSandwich struct{ baseMeal }

func (e EggSandwich) Assemble() (v1alpha1.DAGTask, []v1alpha1.Template) {
	fryEggs := e.step("fry-eggs")
	toastBread := e.step("toast-bread")
	meltCheese := e.step("melt-cheese")
	butterBread := e.step("butter-bread")
	plate := e.step("plate")

	dagTasks := []v1alpha1.DAGTask{
		{
			Name:     fryEggs.Name,
			Template: fryEggs.Name,
		},
		{
			Name:     toastBread.Name,
			Template: toastBread.Name,
		},
		{
			Name:     meltCheese.Name,
			Template: meltCheese.Name,
			Depends:  fryEggs.Name,
		},
		{
			Name:     butterBread.Name,
			Template: butterBread.Name,
			Depends:  toastBread.Name,
		},
		{
			Name:     plate.Name,
			Template: plate.Name,
			Depends:  workflows.DependsOnMultiple(butterBread.Name, meltCheese.Name),
		},
	}
	templates := []v1alpha1.Template{fryEggs, toastBread, meltCheese, butterBread, plate}
	return e.assembleDAG(dagTasks, templates)
}
