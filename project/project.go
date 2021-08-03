package project

import (
	projectApiV1 "github.com/openshift/api/project/v1"
	projectsV1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func CreateClient(c *rest.Config) (*projectsV1.ProjectV1Client, error) {
	clientSetOcp, err := projectsV1.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	return clientSetOcp, err
}

func GetAll(clientProjectSetOcp *projectsV1.ProjectV1Client) (*projectApiV1.ProjectList, error) {
	projects, err := clientProjectSetOcp.Projects().List(metaV1.ListOptions{})
	if err != nil {
		return projects, err
	}
	return projects, nil
}
