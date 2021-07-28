package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (c *DeploymentCtl) index(ctx *gin.Context) goft.Json {
	list, err := c.K8sClient.
		AppsV1().
		Deployments("myweb").
		List(ctx, metav1.ListOptions{})
	goft.Error(err)
	return list
}

func (c *DeploymentCtl) Build(g *goft.Goft) {
	g.Handle("GET", "/deployments", c.index)
}

func (c *DeploymentCtl) Name() string {
	return "DeploymentCtl"
}
