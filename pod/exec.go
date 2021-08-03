package pod

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"os"
)

func RemoteCommand(restConfig *rest.Config, clientCore *corev1client.CoreV1Client, pod *coreV1.Pod, command string) (string, string, error) {
	buf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	request := clientCore.RESTClient().
		Post().
		Namespace(pod.Namespace).
		Resource("pods").
		Name(pod.Name).
		SubResource("exec").
		VersionedParams(&coreV1.PodExecOptions{
			TypeMeta: metav1.TypeMeta{},
			Stdin:    true,
			Stdout:   true,
			Stderr:   true,
			TTY:      true,
			Command:  []string{"/bin/sh", "-c", command},
		}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(restConfig, "POST", request.URL())
	if err != nil {
		panic(err)
	}

	// Put the terminal into raw mode to prevent it echoing characters twice.
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatal(err)
	}
	defer func(fd int, state *terminal.State) {
		err := terminal.Restore(fd, state)
		if err != nil {
			log.Error(err)
		}
	}(0, oldState)

	// Connect this process' std{in,out,err} to the remote shell process.
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: buf,
		Stderr: errBuf,
		Tty:    false,
	})
	if err != nil {
		return "", "", err
	}
	return buf.String(), errBuf.String(), nil
}
