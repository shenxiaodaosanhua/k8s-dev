package core

import (
	"k8s.io/api/apps/v1"
	"log"
)

type DepHandler struct {
	DepMap *DeploymentMap `inject:"-"`
}

func (d *DepHandler) OnAdd(obj interface{}) {
	d.DepMap.Add(obj.(*v1.Deployment))
}
func (d *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	err := d.DepMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	}
}
func (d *DepHandler) OnDelete(obj interface{}) {
	if dep, ok := obj.(*v1.Deployment); ok {
		d.DepMap.Delete(dep)
	}
}
