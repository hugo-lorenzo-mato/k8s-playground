package pod

import (
	"errors"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"strings"
)

func GetAll(clientCore *corev1client.CoreV1Client, ns string) (*coreV1.PodList, error) {
	listPods, err := clientCore.Pods(ns).List(metav1.ListOptions{
		TypeMeta:        metav1.TypeMeta{},
		LabelSelector:   "",
		FieldSelector:   "",
		Watch:           false,
		ResourceVersion: "",
		TimeoutSeconds:  nil,
		Limit:           0,
		Continue:        "",
	})
	if err != nil {
		return listPods, err
	}
	return listPods, err
}


func GetOne(clientCore *corev1client.CoreV1Client, ns string) (*coreV1.Pod, error) {
	var pod *coreV1.Pod

	listPods, err := clientCore.Pods(ns).List(metav1.ListOptions{
		TypeMeta:        metav1.TypeMeta{},
		LabelSelector:   "",
		FieldSelector:   "",
		Watch:           false,
		ResourceVersion: "",
		TimeoutSeconds:  nil,
		Limit:           0,
		Continue:        "",
	})
	if err != nil {
		return nil, err
	}

	for _, pd := range listPods.Items{
		if strings.ToUpper(string(pd.Status.Phase)) == "RUNNING"{
			pod = &pd
			break
		}
	}
	if pod == nil{
		return nil, errors.New("not pod running found")
	}
	return pod, err
}