package meals

import (
	"fmt"
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

var All = []Meal{
	Omelette{baseMeal{name: "omelette"}},
	EggSandwich{baseMeal{name: "egg-sandwich"}},
	TurkeySandwich{baseMeal{name: "turkey-sandwich"}},
	Pasta{baseMeal{name: "pasta"}},
	Steak{baseMeal{name: "steak"}},
	Cake{baseMeal{name: "cake"}},
}

var ByName map[string]Meal

func init() {
	ByName = make(map[string]Meal, len(All))
	for _, meal := range All {
		ByName[meal.Name()] = meal
	}
}

type Meal interface {
	// Assemble returns a DAG task to make this meal and a slice of all (sub) templates used in the task
	Assemble() (v1alpha1.DAGTask, []v1alpha1.Template)
	// Name returns the meal name/identifier
	Name() string
}

type baseMeal struct {
	name string
}

func (b baseMeal) Name() string {
	return b.name
}

// assemble abstracts Meal boilerplate code
// It is responsible for assembling the meal tasks into a DAG
func (b baseMeal) assembleDAG(tasks []v1alpha1.DAGTask, templates []v1alpha1.Template) (v1alpha1.DAGTask, []v1alpha1.Template) {
	dagTemplate := v1alpha1.Template{
		Name: fmt.Sprintf("create-%s-meal", b.Name()),
		DAG: &v1alpha1.DAGTemplate{
			Tasks: tasks,
		},
	}

	task := v1alpha1.DAGTask{
		Name:     dagTemplate.Name,
		Template: dagTemplate.Name,
	}

	// add in encapsulating dag to list of templates
	templates = append(templates, dagTemplate)

	return task, templates
}

func (b baseMeal) step(name string) v1alpha1.Template {
	var containerSpec = corev1.Container{
		Name:    "step",
		Image:   "alpine",
		Command: []string{"true"},
	}

	return v1alpha1.Template{
		// prefix all steps with their meal name so they are unique
		Name:      fmt.Sprintf("%s-%s", b.Name(), name),
		Container: &containerSpec,
	}
}
