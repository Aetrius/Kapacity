package main

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func kubernetesClientConnection() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	var clientset *kubernetes.Clientset

	config, err = getClientSet()

	if err != nil {
		return clientset, err
	}

	clientset, err = kubernetes.NewForConfig(config)

	return clientset, err

}

func kubernetesMetricClientConnection() (*metricsv.Clientset, error) {
	var config *rest.Config
	var err error
	var clientset *metricsv.Clientset

	config, err = getClientSet()

	if err != nil {
		return clientset, err
	}

	clientset, err = metricsv.NewForConfig(config)

	return clientset, err

}

func getClientSet() (*rest.Config, error) {
	var config *rest.Config
	var err error

	kubeconfig := os.Getenv("KUBECONFIG")

	// Determine the configuration based on environment.
	if kubeconfig != "" {
		fmt.Println("running config from env var")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		fmt.Println("running config from within k8s - pod")
		config, err = rest.InClusterConfig()
	}

	return config, err

}

func gatherKubernetesInfo() ([]TypeInfo, error) {
	var allContainerInfo []TypeInfo
	var err error
	clientSet, err := kubernetesClientConnection()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	namespaces, err := getNamespaceData(clientSet)
	deployments, err := getDeploymentData(namespaces, clientSet)
	statefulsets, err := getStatefulsetData(namespaces, clientSet)
	replicasets, err := getReplicaSet(namespaces, clientSet)
	daemonsets, err := getDaemonSets(namespaces, clientSet)

	allContainerInfo = append(allContainerInfo, deployments...)
	allContainerInfo = append(allContainerInfo, statefulsets...)
	allContainerInfo = append(allContainerInfo, replicasets...)
	allContainerInfo = append(allContainerInfo, daemonsets...)

	return allContainerInfo, err
}

func gatherKubernetesPods() ([]Pod, error) {
	var allContainerInfo []Pod
	var err error
	clientSet, err := kubernetesClientConnection()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	metricsClientSet, err := kubernetesMetricClientConnection()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	pods, err := getPodData(clientSet, metricsClientSet)

	allContainerInfo = append(allContainerInfo, pods...)

	return allContainerInfo, err

}
