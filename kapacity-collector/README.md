# Kapacity-Collector
This serves as a basic Golang web application to collect data from the Kubernetes API against the cluster it's running within.
This data is intended to be forwarded to the Kapacity-API for storage.

## Run Locally

-- Running Rancher Desktop config
`export KUBECONFIG=$(echo "~/.kube/config")`

-- Running OpenLens or other K8s config on a MAC OS
`export KUBECONFIG=$(echo "~/Library/Application Support/OpenLens/kubeconfigs/kubeconfig.yaml")`

-- Other Dir On MAC OS (navigate to path for the things)
`export KUBECONFIG=$(echo "~/kubeconfigs/kubeconfig.yaml")`
