package main

import (
	"fmt"
	"time"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"os"
	"k8s.io/client-go/rest"
)

func main() {
	//!!!!!!! dev !!!!!!!!!!
	//kubeconfig := flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
	//flag.Parse()
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	//!!!!!!! prod !!!!!!!!!!
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var image = os.Getenv("KSD_CONTAINER_IMAGE")
	var name = os.Getenv("KSD_DEPLOYMENT_NAME")
	var namespace = os.Getenv("KSD_NAMESPACE")
	var containerName = os.Getenv("KSD_CONTAINER_NAME")

	var tag = "latest"

	dp, err := clientset.Deployments(namespace).Get(name)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Deployment started for " + name)
	for idx,c := range dp.Spec.Template.Spec.Containers {
		if c.Name == containerName {
			dp.Spec.Template.Spec.Containers[idx].Image = fmt.Sprintf("%s:%s", image, tag)
			fmt.Println("Image Found, image has been set: " + dp.Spec.Template.Spec.Containers[idx].Image)
		}
	}

	dp, err = clientset.Deployments(namespace).Update(dp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Deployment updated")

	for {
		pods, err := clientset.CoreV1().Pods("").List(v1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
