package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"os"
	"net/http"
	"ksd/controller"
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

	var KSDListenPort = os.Getenv("KSD_PORT")
	if len(KSDListenPort) < 1 {
		KSDListenPort = "8080"
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

	http.HandleFunc("/deploy", controller.Deploy)
	err = http.ListenAndServe(fmt.Sprintf(":%s", KSDListenPort), nil)
	if nil != err {
		panic(err)
	}
}
