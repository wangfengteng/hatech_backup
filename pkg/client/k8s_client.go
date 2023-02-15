package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var defaultK8sClient *kubernetes.Clientset

func GetK8sClient() *kubernetes.Clientset {
	if defaultK8sClient == nil {
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		defaultK8sClient = clientset
	}

	return defaultK8sClient
}
