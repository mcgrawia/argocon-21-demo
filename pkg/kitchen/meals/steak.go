package meals

import (
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

type Steak struct{ baseMeal }

func (s Steak) Assemble() (v1alpha1.DAGTask, []v1alpha1.Template) {
	grill := s.step("grill")
	plate := s.step("plate")

	dagTasks := []v1alpha1.DAGTask{
		{
			Name:     grill.Name,
			Template: grill.Name,
		},
		{
			Name:     plate.Name,
			Template: plate.Name,
			Depends:  grill.Name,
		},
	}
	templates := []v1alpha1.Template{grill, plate}
	return s.assembleDAG(dagTasks, templates)
}
