package networkPolicy

import (
	projectApiV1 "github.com/openshift/api/project/v1"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	networkingV1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/rest"
)

func CreateClient(c *rest.Config) (networkingV1.NetworkingV1Interface, error) {
	clientSetK8s, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	return clientSetK8s.NetworkingV1(), err
}

func GetAll(clientNetworkingSetOcp networkingV1.NetworkingV1Interface, project projectApiV1.Project) (*v1.NetworkPolicyList, error) {
	npList, err := clientNetworkingSetOcp.NetworkPolicies(project.Name).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return npList, nil
}
