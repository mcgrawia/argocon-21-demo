package meals

import (
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"ianmcgraw.com/m/v2/pkg/workflows"
)

type Cake struct{ baseMeal }

func (c Cake) Assemble() (v1alpha1.DAGTask, []v1alpha1.Template) {
	addIngredients := c.step("add-ingredients")
	crackEggs := c.step("crack-eggs")
	mix := c.step("mix")
	bake := c.step("bake")
	frost := c.step("frost")
	slice := c.step("slice")
	plate := c.step("plate")

	dagTasks := []v1alpha1.DAGTask{
		{
			Name:     addIngredients.Name,
			Template: addIngredients.Name,
		},
		{
			Name:     crackEggs.Name,
			Template: crackEggs.Name,
		},
		{
			Name:     mix.Name,
			Template: mix.Name,
			Depends:  workflows.DependsOnMultiple(crackEggs.Name, addIngredients.Name),
		},
		{
			Name:     bake.Name,
			Template: bake.Name,
			Depends:  mix.Name,
		},
		{
			Name:     frost.Name,
			Template: frost.Name,
			Depends:  bake.Name,
		},
		{
			Name:     slice.Name,
			Template: slice.Name,
			Depends:  frost.Name,
		},
		{
			Name:     plate.Name,
			Template: plate.Name,
			Depends:  slice.Name,
		},
	}
	templates := []v1alpha1.Template{addIngredients, crackEggs, mix, bake, frost, slice, plate}
	return c.assembleDAG(dagTasks, templates)
}
