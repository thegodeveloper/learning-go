package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/api/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pod, err := clientset.CoreV1().Pods("book").Get("example", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("value of pod: %v\n", pod)
}
