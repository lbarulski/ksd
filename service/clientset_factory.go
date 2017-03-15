package service

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewKubernetesClientset() *kubernetes.Clientset {
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

	return clientset
}
