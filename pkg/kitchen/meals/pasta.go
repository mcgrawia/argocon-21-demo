package meals

import (
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"ianmcgraw.com/m/v2/pkg/workflows"
)

type Pasta struct{ baseMeal }

func (p Pasta) Assemble() (v1alpha1.DAGTask, []v1alpha1.Template) {
	boilWater := p.step("boil-water")
	cookPasta := p.step("cook-pasta")
	heatSauce := p.step("heat-sauce")
	plate := p.step("plate")

	dagTasks := []v1alpha1.DAGTask{
		{
			Name:     boilWater.Name,
			Template: boilWater.Name,
		},
		{
			Name:     cookPasta.Name,
			Template: cookPasta.Name,
			Depends:  boilWater.Name,
		},
		{
			Name:     heatSauce.Name,
			Template: heatSauce.Name,
		},
		{
			Name:     plate.Name,
			Template: plate.Name,
			Depends:  workflows.DependsOnMultiple(heatSauce.Name, cookPasta.Name),
		},
	}
	templates := []v1alpha1.Template{boilWater, heatSauce, cookPasta, plate}
	return p.assembleDAG(dagTasks, templates)}
