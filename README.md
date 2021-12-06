# argocon-21-demo
Demo for my talk at ArgoCon '21 showing how to use Go to create and submit dynamic Argo Workflows.
This repo implements a Go-based CLI called "feedme" that will accept one or more of the following 
meals to prepare using Argo Workflows:
```shell
omelette
egg-sandwich
turkey-sandwich
pasta
steak
cake
```

Usage:
```shell
./feedme omlette steak cake
```

# Setup
1. Have a Kubernetes cluster available, I used [Docker for Mac's](https://docs.docker.com/desktop/kubernetes/).
2. Install Argo:
    ```shell
    kubectl create namespace argo
    kubectl apply -n argo -f https://github.com/argoproj/argo-workflows/releases/download/v3.2.4/install.yaml
    ```
3. Disable Argo Auth (for local installs only!):
   ```shell
   kubectl patch deployment -n argo argo-server -p '{"spec":{"template":{"spec":{"containers":[{"name":"argo-server","args":["server", "--auth-mode", "server"]}]}}}}'
   ```
4. Make Argo accessible outside the cluster (may differ depending on your K8s setup):
   ```shell
   kubectl patch service -n argo argo-server -p '{"spec":{"type":"NodePort"}}'
   # Get the Argo UI port:
   kubectl get service -n argo argo-server -oyaml | grep nodePort:
   # open the UI
   open https://localhost:<PORT>
   ```
5. Build this program (tested using Go 1.17 originally)
    ```shell
    go build -o feedme ./cmd
    ```
6. Run
   ```shell
   ./feedme omlette steak cake
   ```