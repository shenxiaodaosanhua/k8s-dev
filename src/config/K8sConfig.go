package config

import (
	"k8s-dev/core"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type K8sConfig struct {
	DepHandler *core.DepHandler `inject:"-"`
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

func (c *K8sConfig) InitInformer() informers.SharedInformerFactory {
	fact := informers.NewSharedInformerFactory(c.InitClient(), 0)

	depInformer := fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(c.DepHandler)

	fact.Start(wait.NeverStop)

	return fact
}
