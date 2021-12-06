package kitchen

import (
	"errors"
	"fmt"
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"ianmcgraw.com/m/v2/pkg/kitchen/meals"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ChefArgo struct {
	meals []meals.Meal
}

func NewChefArgo(mealNames []string) (ChefArgo, error) {
	if len(mealNames) == 0 {
		return ChefArgo{}, errors.New("cannot build workflow, at least one meal must be provided")
	}
	var mealsToMake []meals.Meal
	for _, name := range mealNames {
		if meal, found := meals.ByName[name]; found {
			mealsToMake = append(mealsToMake, meal)
		} else {
			return ChefArgo{}, fmt.Errorf("meal %s not found", name)
		}
	}
	return ChefArgo{meals: mealsToMake}, nil
}

func (w ChefArgo) Cook() *v1alpha1.Workflow {
	baseDAG := v1alpha1.Template{Name: "make-meals-dag", DAG: &v1alpha1.DAGTemplate{}}
	allTemplates := []v1alpha1.Template{baseDAG}

	// modify base dag with chosen meals
	for _, meal := range w.meals {
		task, templates := meal.Assemble()
		baseDAG.DAG.Tasks = append(baseDAG.DAG.Tasks, task)
		allTemplates = append(allTemplates, templates...)
	}

	return &v1alpha1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "argo-kitchen-",
		},
		Spec: v1alpha1.WorkflowSpec{
			Entrypoint: baseDAG.Name,
			Templates:  allTemplates,
		},
	}
}
