package config

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type K8sConfig struct {
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (c *K8sConfig) InitClient() *kubernetes.Clientset {
	config := &rest.Config{
		Host:        "139.198.109.248:8009",
		BearerToken: "e2a95216a52f30136808bf364810fbdd",
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
