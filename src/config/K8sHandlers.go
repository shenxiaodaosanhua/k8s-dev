package config

import "k8s-dev/core"

type K8sHandler struct {
}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}

func (h *K8sHandler) DepHandlers() *core.DepHandler {
	return &core.DepHandler{}
}
