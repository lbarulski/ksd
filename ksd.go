package main

import (
	"fmt"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/api"
	"os"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var image = os.Getenv("KSD_DOCKER_IMAGE")
	var name = os.Getenv("KSD_DEPLOYMENT_NAME")
	var namespace = os.Getenv("KSD_NAMESPACE")

	var tag = "latest"

	var data = fmt.Sprintf(`{"spec":{"template":{"spec":{"containers":[{"image":"%s:%s"}]}}}}`, image, tag);
	_, err = clientset.Deployments(namespace).Patch(name, api.JSONPatchType, []byte (data));


	if err != nil {
		panic(err.Error())
	}

	for {
		pods, err := clientset.CoreV1().Pods("").List(v1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
