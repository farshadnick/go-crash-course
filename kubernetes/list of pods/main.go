package main

import (
	"context"
	"fmt"
	//	"os"
	"log"
	//	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Path to your kubeconfig file
	kubeconfig := "/home/farshad/.kube/config"
	// Build the client configuration from kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create a new clientset using the configuration
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating clientset: %v", err)
	}

	// Retrieve the list of pods in the default namespace
	podList, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error retrieving pod list: %v", err)
	}

	// Print the name of each pod
	fmt.Println("Pods:")
	for _, pod := range podList.Items {
		fmt.Printf("- %s\n", pod.GetName())
	}
}

