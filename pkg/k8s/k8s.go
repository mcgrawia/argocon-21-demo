package k8s

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	restClient "k8s.io/client-go/rest"
	cmdClient "k8s.io/client-go/tools/clientcmd"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const defaultNS = "default"

func GetConfig() *restClient.Config {
	_, inCluster := os.LookupEnv("KUBERNETES_SERVICE_HOST")
	switch inCluster {
	case true:
		log.Debugf("Program running inside k8s cluster, reading the in-cluster configuration")
		config, err := restClient.InClusterConfig()
		if err != nil {
			log.Fatalf("could not read in-cluster config: %s", err)
		}
		return config
	default:
		// get current user to determine home directory
		usr, err := user.Current()
		if err != nil {
			log.Fatalf("could not find user: %s", err)
		}

		kubeConfPath := filepath.Join(usr.HomeDir, ".kube", "config")
		log.Debugf("running outside k8s, reading kubeconfig at: %s", kubeConfPath)
		config, err := cmdClient.BuildConfigFromFlags("", kubeConfPath)
		if err != nil {
			log.Fatalf("error reading kubeconfig: %s", err)
		}
		return config
	}
}

// GetNamespace will return the current namespace for the running program
// Checks for the user passed ENV variable POD_NAMESPACE if not available
// pulls the namespace from pod, if not returns "default"
func GetNamespace() string {
	if ns := os.Getenv("POD_NAMESPACE"); ns != "" {
		return ns
	}
	if data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns
		}
	}
	return defaultNS
}
