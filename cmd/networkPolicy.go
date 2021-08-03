package cmd

import (
	"fmt"
	projectsV1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s-playground/connection"
	"k8s-playground/networkPolicy"
	"k8s-playground/project"
	fileUtil "k8s-playground/util/file"
	printUtil "k8s-playground/util/print"
	networkingV1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	"os"
	"path/filepath"
	"time"
)

func init() {

	clusterFlag := cli.StringFlag{
		Name:        "cluster, c",
		Aliases:     nil,
		Usage:       "Cluster to iterate through. E.G: ssccarteixo",
		EnvVars:     []string{"CLUSTER"},
		Destination: &cluster,
		Required:    true,
	}

	supraenvironmentFlag := cli.StringFlag{
		Name:        "supraenvironment, s",
		Usage:       "Supraenvironment. E.G: des",
		EnvVars:     []string{"SUPRAENVIRONMENT"},
		Destination: &supraenvironment,
		Required:    true,
	}

	connectionDataYamlPathFlag := cli.StringFlag{
		Name:        "connectionDataYamlPath, y",
		Usage:       "Path to the yaml file containing the credentials and URIs to connect to the different clusters. E.G: example.yml",
		EnvVars:     []string{"CREDENTIALS_YAML_PATH"},
		Destination: &connectionDataYamlPath,
		Value:       "data/cluster.yml",
		Required:    false,
	}

	dryRunflag := cli.BoolFlag{
		Name:        "dryRun",
		Aliases:     []string{"d"},
		Usage:       "dryRun execution",
		EnvVars:     []string{"DRY_RUN"},
		Destination: &dryRun,
		Required:    true,
		Hidden:      false,
	}

	checkFlags := []cli.Flag{
		&clusterFlag,
		&connectionDataYamlPathFlag,
		&dryRunflag,
		&supraenvironmentFlag,
	}

	backupAll := &cli.Command{
		Name:   "backupAll",
		Usage:  "Backup of all the networkpolicies of a cluster will be created.",
		Action: backupAll,
		Flags:  checkFlags,
	}

	cmd := &cli.Command{
		Name:        "networkpolicy",
		Usage:       "networkpolicy related actions",
		Flags:       []cli.Flag{},
		Subcommands: []*cli.Command{backupAll},
	}
	Cmds = append(Cmds, cmd)
}

func backupAll(_ *cli.Context) error {

	printUtil.Flags(map[string]interface{}{"cluster": cluster, "supraenvironment": supraenvironment,
		"connectionDataYamlPath": connectionDataYamlPath, "dryRun": dryRun})

	clientProjectSetOcp, clientNetworkingSetOcp, err := createClients(connectionDataYamlPath, cluster, supraenvironment)
	if err != nil {
		return err
	}

	projectsList, err := project.GetAll(clientProjectSetOcp)
	if err != nil {
		return err
	}

	projectCounter := 0
	npCounter := 0

	for _, pj := range projectsList.Items {

		projectCounter += 1

		npList, err := networkPolicy.GetAll(clientNetworkingSetOcp, pj)
		if err != nil {
			return err
		}

		dir := filepath.Join(TARGETFOLDER, pj.Name)
		if !fileUtil.Exists(dir) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return err
			}
		}

		for _, np := range npList.Items {

			npCounter += 1

			data, err := yaml.Marshal(np)
			if err != nil {
				return err
			}

			output := filepath.Join(dir, fmt.Sprintf("%s-%d.yml", np.Name, time.Now().Unix()))
			log.Infof("+ Backing up network policy: %s - project: %s in %s dir", np.Name, pj.Name, output)
			if err = ioutil.WriteFile(output, data, os.ModePerm); err != nil {
				return err
			}
		}
	}
	log.Infof("-> Number of projects analyzed: %d. Total number of Network Policies: %d.", projectCounter, npCounter)
	return nil
}

func createClients(path, cluster, supraenv string) (*projectsV1.ProjectV1Client, networkingV1.NetworkingV1Interface, error) {

	config, err := connection.CreateNewConfig(path, cluster, supraenv)
	if err != nil {
		return nil, nil, err
	}

	clientProjectSetOcp, err := project.CreateClient(config)
	if err != nil {
		return nil, nil, err
	}

	clientNetworkingSetOcp, err := networkPolicy.CreateClient(config)
	if err != nil {
		return nil, nil, err
	}
	return clientProjectSetOcp, clientNetworkingSetOcp, nil
}
