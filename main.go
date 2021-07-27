package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-dev/src/config"
	"k8s-dev/src/controller"
)

func main() {
	goft.Ignite().Config(
		config.NewK8sConfig(),
	).Mount(
		"v1",
		controller.NewDeploymentCtl(),
	).Launch()
}
