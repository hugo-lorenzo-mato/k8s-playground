package app

import (
	appsApiV1 "github.com/openshift/api/apps/v1"
	appsV1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func CreateClient(c *rest.Config) (*appsV1.AppsV1Client, error) {
	clientSetOcp, err := appsV1.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	return clientSetOcp, err
}

func GetDeploymentConfig(clientAppsSetOcp *appsV1.AppsV1Client, ns, projectName string) (*appsApiV1.DeploymentConfig, error) {
	deploymentConfig, err := clientAppsSetOcp.DeploymentConfigs(ns).Get(projectName, metaV1.GetOptions{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ResourceVersion: "",
	})
	if err != nil {
		return deploymentConfig, err
	}
	return deploymentConfig, nil
}
