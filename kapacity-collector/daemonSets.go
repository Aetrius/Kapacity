package main

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

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
