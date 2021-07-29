package core

import (
	"fmt"
	"k8s.io/api/apps/v1"
	"sync"
)

type DeploymentMap struct {
	data sync.Map
}

//添加
func (d *DeploymentMap) Add(dep *v1.Deployment) {
	if list, ok := d.data.Load(dep.Namespace); ok {
		list = append(list.([]*v1.Deployment), dep)
		d.data.Store(dep.Namespace, list)
	} else {
		d.data.Store(dep.Namespace, []*v1.Deployment{dep})
	}
}

func (d *DeploymentMap) Update(dep *v1.Deployment) error {
	if list, ok := d.data.Load(dep.Namespace); ok {
		for i, rangeDep := range list.([]*v1.Deployment) {
			if rangeDep.Name == dep.Name {
				list.([]*v1.Deployment)[i] = dep
			}
		}
		return nil
	}
	return fmt.Errorf("deployment-%s not found", dep.Name)
}

func (d *DeploymentMap) Delete(dep *v1.Deployment) {
	if list, ok := d.data.Load(dep.Namespace); ok {
		for i, rangeDep := range list.([]*v1.Deployment) {
			if rangeDep.Name == dep.Name {
				newList := append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				d.data.Store(dep.Namespace, newList)
				break
			}
		}
	}
}

func (d *DeploymentMap) ListByNs(ns string) ([]*v1.Deployment, error) {
	if list, ok := d.data.Load(ns); ok {
		return list.([]*v1.Deployment), nil
	}
	return nil, fmt.Errorf("record not found")
}

func (d *DeploymentMap) GetDeployment(ns string, depName string) (*v1.Deployment, error) {
	if list, ok := d.data.Load(ns); ok {
		for _, item := range list.([]*v1.Deployment) {
			if item.Name == depName {
				return item, nil
			}
		}
	}
	return nil, fmt.Errorf("record not found")
}
