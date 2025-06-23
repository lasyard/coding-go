package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func main() {
	// Use in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating in-cluster config: %v\n", err)
		os.Exit(1)
	}

	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating dynamic client: %v\n", err)
		os.Exit(1)
	}

	// Replace these with your CRD's group, version, resource
	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "configmaps", // plural!
	}

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name": "test-dyn",
			},
			"data": map[string]interface{}{
				"name": "Betty",
				"age":  "20",
			},
		},
	}

	result, err := dynClient.Resource(gvr).Namespace("default").Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create CRD: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created: %s\n", result.GetName())
}
