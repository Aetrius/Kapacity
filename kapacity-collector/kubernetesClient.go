package main

import (
	"context"
	"fmt"
	"os"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

func getNamespaceData(clientSet *kubernetes.Clientset) (*v1.NamespaceList, error) {
	namespaces, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	return namespaces, err

}

func getDeploymentData(namespaces *v1.NamespaceList, clientSet *kubernetes.Clientset) ([]PodInfo, error) {
	var err error
	var allNamespaceInfo []PodInfo

	for _, ns := range namespaces.Items {
		var deployments, err = clientSet.AppsV1().Deployments(ns.Name).List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			return allNamespaceInfo, err
		}

		for _, deploy := range deployments.Items {
			owner := deploy.Labels["Owner"]
			for _, c := range deploy.Spec.Template.Spec.Containers {
				info := PodInfo{
					Namespace:     ns.Name,
					Type:          "Deployment",
					Name:          deploy.Name,
					Owner:         owner,
					PodCount:      *deploy.Spec.Replicas,
					Container:     c.Name,
					CPURequest:    c.Resources.Requests.Cpu().String(),
					CPULimit:      c.Resources.Limits.Cpu().String(),
					MemoryRequest: c.Resources.Requests.Memory().String(),
					MemoryLimit:   c.Resources.Limits.Memory().String(),
				}
				allNamespaceInfo = append(allNamespaceInfo, info)
			}
		}
	}

	return allNamespaceInfo, err

}

func getStatefulsetData(namespaces *v1.NamespaceList, clientSet *kubernetes.Clientset) ([]PodInfo, error) {
	var err error
	var allNamespaceInfo []PodInfo

	for _, ns := range namespaces.Items {
		var statefulsets, err = clientSet.AppsV1().StatefulSets(ns.Name).List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			return allNamespaceInfo, err
		}

		for _, deploy := range statefulsets.Items {
			owner := deploy.Labels["Owner"]
			for _, c := range deploy.Spec.Template.Spec.Containers {
				info := PodInfo{
					Namespace:     ns.Name,
					Type:          "Statefulset",
					Name:          deploy.Name,
					Owner:         owner,
					PodCount:      *deploy.Spec.Replicas,
					Container:     c.Name,
					CPURequest:    c.Resources.Requests.Cpu().String(),
					CPULimit:      c.Resources.Limits.Cpu().String(),
					MemoryRequest: c.Resources.Requests.Memory().String(),
					MemoryLimit:   c.Resources.Limits.Memory().String(),
				}
				allNamespaceInfo = append(allNamespaceInfo, info)
			}
		}
	}

	return allNamespaceInfo, err

}

func getReplicaSet(namespaces *v1.NamespaceList, clientSet *kubernetes.Clientset) ([]PodInfo, error) {
	var err error
	var allNamespaceInfo []PodInfo

	for _, ns := range namespaces.Items {
		var replicaSets, err = clientSet.AppsV1().ReplicaSets(ns.Name).List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			return allNamespaceInfo, err
		}

		for _, deploy := range replicaSets.Items {
			owner := deploy.Labels["Owner"]
			for _, c := range deploy.Spec.Template.Spec.Containers {
				info := PodInfo{
					Namespace:     ns.Name,
					Type:          "ReplicaSet",
					Name:          deploy.Name,
					Owner:         owner,
					PodCount:      *deploy.Spec.Replicas,
					Container:     c.Name,
					CPURequest:    c.Resources.Requests.Cpu().String(),
					CPULimit:      c.Resources.Limits.Cpu().String(),
					MemoryRequest: c.Resources.Requests.Memory().String(),
					MemoryLimit:   c.Resources.Limits.Memory().String(),
				}
				allNamespaceInfo = append(allNamespaceInfo, info)
			}
		}
	}

	return allNamespaceInfo, err

}

func getDaemonSets(namespaces *v1.NamespaceList, clientSet *kubernetes.Clientset) ([]PodInfo, error) {
	var err error
	var allNamespaceInfo []PodInfo

	for _, ns := range namespaces.Items {
		var daemonSets, err = clientSet.AppsV1().DaemonSets(ns.Name).List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			return allNamespaceInfo, err
		}

		for _, deploy := range daemonSets.Items {
			owner := deploy.Labels["Owner"]
			for _, c := range deploy.Spec.Template.Spec.Containers {
				info := PodInfo{
					Namespace:     ns.Name,
					Type:          "DaemonSets",
					Name:          deploy.Name,
					Owner:         owner,
					PodCount:      *&deploy.Status.NumberAvailable,
					Container:     c.Name,
					CPURequest:    c.Resources.Requests.Cpu().String(),
					CPULimit:      c.Resources.Limits.Cpu().String(),
					MemoryRequest: c.Resources.Requests.Memory().String(),
					MemoryLimit:   c.Resources.Limits.Memory().String(),
				}
				allNamespaceInfo = append(allNamespaceInfo, info)
			}
		}
	}

	return allNamespaceInfo, err

}

func gatherKubernetesInfo() ([]PodInfo, error) {
	var allContainerInfo []PodInfo
	var err error
	clientSet, err := kubernetesClientConnection()

	if err != nil {
		fmt.Println("Error: ", err)

	}

	namespaces, _ := getNamespaceData(clientSet)
	deployments, _ := getDeploymentData(namespaces, clientSet)
	statefulsets, _ := getStatefulsetData(namespaces, clientSet)
	replicasets, _ := getReplicaSet(namespaces, clientSet)
	daemonsets, _ := getDaemonSets(namespaces, clientSet)

	allContainerInfo = append(allContainerInfo, deployments...)
	allContainerInfo = append(allContainerInfo, statefulsets...)
	allContainerInfo = append(allContainerInfo, replicasets...)
	allContainerInfo = append(allContainerInfo, daemonsets...)

	return allContainerInfo, err

}
