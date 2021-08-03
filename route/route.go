package route

import (
	v1 "github.com/openshift/api/route/v1"
	routesV1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func CreateClient(c *rest.Config) (*routesV1.RouteV1Client, error) {
	clientSetOcp, err := routesV1.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	return clientSetOcp, err
}

func GetAll(client *routesV1.RouteV1Client, ns string) (*v1.RouteList, error) {
	routeList, err := client.Routes(ns).List(metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return routeList, nil
}
