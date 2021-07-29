package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-dev/src/config"
	"k8s-dev/src/controller"
)

func main() {
	goft.Ignite().Config(
		config.NewK8sHandler(),
		config.NewK8sConfig(),
		config.NewK8sMaps(),
	).Mount(
		"v1",
		controller.NewDeploymentCtl(),
	).Launch()
}
