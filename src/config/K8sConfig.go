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
		BearerToken: "kubeconfig-user-q2kd8:69ssrzn44svv5sq2tll4tn44lc2t6bdspmjccdq2fnrl4rvrkdzz9w",
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
