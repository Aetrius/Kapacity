package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var config *rest.Config
	var err error

	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig() // if env var isn't present use the rbac from kubernetes service account
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch all namespaces
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var builder strings.Builder
	builder.WriteString("<html><head><title>Kubernetes Info</title></head><body>")
	builder.WriteString("<h1>Kubernetes Deployment and StatefulSet Information</h1>")
	builder.WriteString("<table border='1'><tr><th>Namespace</th><th>Type</th><th>Name</th><th>Owner</th><th>Pod Count</th><th>Resources</th></tr>")

	for _, ns := range namespaces.Items {
		// Get Deployments
		deployments, _ := clientset.AppsV1().Deployments(ns.Name).List(context.TODO(), metav1.ListOptions{})
		for _, deploy := range deployments.Items {
			owner := deploy.Labels["Owner"]
			builder.WriteString(fmt.Sprintf("<tr><td>%s</td><td>Deployment</td><td>%s</td><td>%s</td><td>%d</td><td>", ns.Name, deploy.Name, owner, *deploy.Spec.Replicas))
			for _, c := range deploy.Spec.Template.Spec.Containers {
				builder.WriteString(fmt.Sprintf("%s: CPU Requests: %s, CPU Limits: %s, Memory Requests: %s, Memory Limits: %s<br>",
					c.Name,
					c.Resources.Requests.Cpu().String(),
					c.Resources.Limits.Cpu().String(),
					c.Resources.Requests.Memory().String(),
					c.Resources.Limits.Memory().String(),
				))
			}
			builder.WriteString("</td></tr>")
		}

		// Get StatefulSets
		statefulSets, _ := clientset.AppsV1().StatefulSets(ns.Name).List(context.TODO(), metav1.ListOptions{})
		for _, ss := range statefulSets.Items {
			owner := ss.Annotations["owner"]
			builder.WriteString(fmt.Sprintf("<tr><td>%s</td><td>StatefulSet</td><td>%s</td><td>%s</td><td>%d</td><td>", ns.Name, ss.Name, owner, *ss.Spec.Replicas))
			for _, c := range ss.Spec.Template.Spec.Containers {
				builder.WriteString(fmt.Sprintf("%s: CPU Requests: %s, CPU Limits: %s, Memory Requests: %s, Memory Limits: %s<br>",
					c.Name,
					c.Resources.Requests.Cpu().String(),
					c.Resources.Limits.Cpu().String(),
					c.Resources.Requests.Memory().String(),
					c.Resources.Limits.Memory().String(),
				))
			}
			builder.WriteString("</td></tr>")
		}
	}

	builder.WriteString("</table></body></html>")
	w.Write([]byte(builder.String()))
}
