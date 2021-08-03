package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"k8s-playground/connection"
	"k8s-playground/pod"
	"k8s-playground/service"
	printUtil "k8s-playground/util/print"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

func init() {

	clusterFlag := cli.StringFlag{
		Name:        "cluster",
		Aliases:     nil,
		Usage:       "Cluster to iterate through. E.G: micluster1",
		EnvVars:     []string{"CLUSTER"},
		Destination: &cluster,
		Required:    true,
	}

	supraenvironmentFlag := cli.StringFlag{
		Name:        "supraenvironment",
		Usage:       "Supraenvironment. E.G: des",
		EnvVars:     []string{"SUPRAENVIRONMENT"},
		Destination: &supraenvironment,
		Required:    true,
	}

	connectionDataYamlPathFlag := cli.StringFlag{
		Name:        "connectionDataYamlPath",
		Usage:       "Path to the yaml file containing the credentials and URIs to connect to the different clusters. E.G: example.yml",
		EnvVars:     []string{"CREDENTIALS_YAML_PATH"},
		Destination: &connectionDataYamlPath,
		Value:       "data/cluster.yml",
		Required:    false,
	}

	namespaceSourceFlag := cli.StringFlag{
		Name:        "namespaceSource",
		Usage:       "Namespace/project source of the connection request.",
		EnvVars:     []string{"NAMESPACE_SOURCE"},
		Destination: &nsSource,
		Required:    true,
	}

	namespaceTargetFlag := cli.StringFlag{
		Name:        "namespaceTarget",
		Usage:       "Namespace/project target of the connection request.",
		EnvVars:     []string{"NAMESPACE_TARGET"},
		Destination: &nsTarget,
		Required:    true,
	}

	checkConnectionFlags := []cli.Flag{
		&clusterFlag,
		&supraenvironmentFlag,
		&connectionDataYamlPathFlag,
		&namespaceSourceFlag,
		&namespaceTargetFlag,
	}

	checkConnection := &cli.Command{
		Name:   "checkConnection",
		Usage:  "Backup of all the networkpolicies of a cluster will be created.",
		Action: checkConnection,
		Flags:  checkConnectionFlags,
	}

	cmd := &cli.Command{
		Name:        "networking",
		Usage:       "networking related actions",
		Flags:       []cli.Flag{},
		Subcommands: []*cli.Command{checkConnection},
	}
	Cmds = append(Cmds, cmd)

}

func checkConnection(_ *cli.Context) error {

	printUtil.Flags(map[string]interface{}{"cluster": cluster, "supraenvironment": supraenvironment,
		"connectionDataYamlPath": connectionDataYamlPath, "nsSource": nsSource, "nsTarget": nsTarget})

	config, err := connection.CreateNewConfig(connectionDataYamlPath, cluster, supraenvironment)
	if err != nil {
		return err
	}

	clientCore, err := corev1client.NewForConfig(config)
	if err != nil {
		return err
	}

	serviceToRequest, err := service.GetOne(clientCore, nsTarget)
	if err != nil {
		return err
	}

	po, err := pod.GetOne(clientCore, nsSource)
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("curl -o /dev/null --max-time 3 -ks %s && echo Connection is available from "+
		"$(uname -n) in %s namespace to %s || echo Connection is NOT available from "+
		"$(uname -n) in %s namespace to %s", serviceToRequest, nsSource, serviceToRequest, nsSource, serviceToRequest)
	log.Infof("+ command: %s", cmd)
	output, _, err := pod.RemoteCommand(config, clientCore, po, cmd)
	if err != nil {
		return err
	} else {
		log.Infof("Output: %s", output)
	}

	return nil
}
