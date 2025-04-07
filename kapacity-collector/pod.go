package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

type Pod struct {
	Namespace     string `json:"namespace"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Owner         string `json:"owner"`
	Container     string `json:"container"`
	CPURequest    string `json:"cpu_request"`
	CPULimit      string `json:"cpu_limit"`
	CPUUsage      string `json:"cpu_usage"`
	MemoryRequest string `json:"memory_request"`
	MemoryLimit   string `json:"memory_limit"`
	MemoryUsage   string `json:"memory_usage"`
	Status        string `json:"status"`
}

func getPodData(clientSet *kubernetes.Clientset, metricsClient *metricsv.Clientset) ([]Pod, error) {
	var podList []Pod

	// Get the list of all pods in the cluster
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error listing pods: %v", err)
	}

	// Get the metrics for all pods
	podMetricsList, err := metricsClient.MetricsV1beta1().PodMetricses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error listing pod metrics: %v", err)
	}

	podMetricsMap := make(map[string]corev1.ResourceList)
	for _, podMetrics := range podMetricsList.Items {
		for _, container := range podMetrics.Containers {
			podMetricsMap[podMetrics.Namespace+"/"+podMetrics.Name+"/"+container.Name] = container.Usage
		}
	}

	// Iterate through the pods and populate the podList
	for _, pod := range pods.Items {
		ownerType := getOwnerType(pod.OwnerReferences)
		ownerName := getOwnerName(pod.OwnerReferences)

		for _, container := range pod.Spec.Containers {
			metrics, exists := podMetricsMap[pod.Namespace+"/"+pod.Name+"/"+container.Name]
			var cpuUsage resource.Quantity
			var memoryUsage resource.Quantity
			if exists {
				cpuUsage = metrics[corev1.ResourceCPU]
				memoryUsage = metrics[corev1.ResourceMemory]
			}

			cpuRequest := container.Resources.Requests[corev1.ResourceCPU]
			cpuLimit := container.Resources.Limits[corev1.ResourceCPU]
			memoryRequest := container.Resources.Requests[corev1.ResourceMemory]
			memoryLimit := container.Resources.Limits[corev1.ResourceMemory]

			podDetails := Pod{
				Namespace:     pod.Namespace,
				Type:          ownerType,
				Name:          pod.Name,
				Owner:         ownerName,
				Container:     container.Name,
				CPURequest:    cpuRequest.String(),
				CPULimit:      cpuLimit.String(),
				MemoryRequest: memoryRequest.String(),
				MemoryLimit:   memoryLimit.String(),
				CPUUsage:      cpuUsage.String(),
				MemoryUsage:   memoryUsage.String(),
				Status:        string(pod.Status.Phase),
			}
			podList = append(podList, podDetails)
		}
	}

	return podList, nil
}

func getOwnerType(ownerReferences []metav1.OwnerReference) string {
	if len(ownerReferences) > 0 {
		return ownerReferences[0].Kind
	}
	return "Unknown"
}

func getOwnerName(ownerReferences []metav1.OwnerReference) string {
	if len(ownerReferences) > 0 {
		return ownerReferences[0].Name
	}
	return "Unknown"
}
