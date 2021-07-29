package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-dev/core"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
	DepMap    *core.DeploymentMap   `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (c *DeploymentCtl) index(ctx *gin.Context) goft.Json {
	list, err := c.DepMap.ListByNs("myweb")
	goft.Error(err)
	return list
}

func (c *DeploymentCtl) Build(g *goft.Goft) {
	g.Handle("GET", "/deployments", c.index)
}

func (c *DeploymentCtl) Name() string {
	return "DeploymentCtl"
}
