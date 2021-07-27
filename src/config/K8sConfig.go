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
		Host: "139.198.109.248:8009",
		BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IlM4YnJPSWhLWFcwTDBHVmZQZWlCVVFXVDVvWFV1aThzWXBVdk5CVDVGd0kifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRlZmF1bHQtdG9rZW4tNW43a3EiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGVmYXVsdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjRkNjFlZWVjLWRmMjItNDNlNi1iMWVjLTkyMzg5ZTI0OWQyMCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.qauDckhyuH07GcBhVpfkNOpx91xcYkNCuueTs5HS1eEHzlNop51x0IbPIWN-j68F1mkPk_thyidjfLIjLKggAMDDqVaWhwAloMgpkcDbRhwd2g4x65aY8jk8oMduHzMnSCTRTpWojuYz0AtHrG1TKshi8jhwWukvW5GfqkNsT4WPPTyBsT73HRtIELU_Kd0kId05BD2XyD5vjSrdEjlAcGuHASSStliklT1ZhuZ3ZTa2m6S_218oJ71R1_cbd5ImWlUUAqcSp59grMqkZONK7BGc6JKMuL2oqNeDoyhhU_0n66l4rdY5esngVevt9Ji6d1kYMedqL-hS0r8f_5e6Zg",
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
