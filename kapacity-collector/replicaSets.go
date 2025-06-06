package main

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getReplicaSet(namespaces *v1.NamespaceList, clientSet *kubernetes.Clientset) ([]TypeInfo, error) {
	var err error
	var allNamespaceInfo []TypeInfo

	for _, ns := range namespaces.Items {
		var replicaSets, err = clientSet.AppsV1().ReplicaSets(ns.Name).List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			return allNamespaceInfo, err
		}

		for _, deploy := range replicaSets.Items {
			owner := deploy.Labels["Owner"]
			for _, c := range deploy.Spec.Template.Spec.Containers {
				info := TypeInfo{
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
