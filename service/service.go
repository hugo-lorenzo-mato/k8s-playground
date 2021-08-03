package service

import (
	"errors"
	"fmt"
	projectApiV1 "github.com/openshift/api/project/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

func GetOne(clientCore *corev1client.CoreV1Client, p interface{}) (string, error) {
	var serviceToTest, namespace string

	switch v := p.(type) {
	case projectApiV1.Project:
		namespace = v.Namespace
	case string:
		namespace = v
	default:
		return "", errors.New(fmt.Sprintf("type is not covered: %T!\n", v))
	}

	serviceList, err := clientCore.Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}

	for _, service := range serviceList.Items {
		serviceToTest = fmt.Sprintf("%s.%s.%s:%s", service.Name, namespace, "svc.cluster.local", service.Spec.Ports[0].TargetPort.String())
		break
	}
	return serviceToTest, nil
}
