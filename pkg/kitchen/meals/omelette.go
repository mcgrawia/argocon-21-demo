package meals

import (
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"ianmcgraw.com/m/v2/pkg/workflows"
)

type Omelette struct{ baseMeal }

func (o Omelette) Assemble() (v1alpha1.DAGTask, []v1alpha1.Template) {
	scrambleEggs := o.step("scramble-eggs")
	fryVeg := o.step("fry-veggies")
	meltCheese := o.step("melt-cheese")
	fold := o.step("fold")
	plate := o.step("plate")

	dagTasks := []v1alpha1.DAGTask{
		{
			Name:     scrambleEggs.Name,
			Template: scrambleEggs.Name,
		},
		{
			Name:     fryVeg.Name,
			Template: fryVeg.Name,
		},
		{
			Name:     meltCheese.Name,
			Template: meltCheese.Name,
			Depends:  workflows.DependsOnMultiple(fryVeg.Name, scrambleEggs.Name),
		},
		{
			Name:     fold.Name,
			Template: fold.Name,
			Depends:  meltCheese.Name,
		},
		{
			Name:     plate.Name,
			Template: plate.Name,
			Depends:  fold.Name,
		},
	}
	templates := []v1alpha1.Template{scrambleEggs, fryVeg, meltCheese, fold, plate}
	return o.assembleDAG(dagTasks, templates)
}
