package workflows

import (
	wfclientset "github.com/argoproj/argo/pkg/client/clientset/versioned"
	"github.com/argoproj/argo/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"ianmcgraw.com/m/v2/pkg/k8s"
)

var Client v1alpha1.WorkflowInterface

func init() {
	ns := k8s.GetNamespace()
	Client = wfclientset.NewForConfigOrDie(k8s.GetConfig()).ArgoprojV1alpha1().Workflows(ns)
}
