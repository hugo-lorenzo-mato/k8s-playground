package cmd

import (
	projectApiV1 "github.com/openshift/api/project/v1"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"k8s-playground/connection"
	"k8s-playground/project"
	route "k8s-playground/route"
	printUtil "k8s-playground/util/print"
	"k8s.io/client-go/rest"
)

func init() {

	clusterFlag := cli.StringFlag{
		Name:        "cluster",
		Aliases:     nil,
		Usage:       "Cluster to iterate through. E.G: ssccarteixo",
		EnvVars:     []string{"CLUSTER"},
		Destination: &cluster,
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

	hostFlag := cli.StringFlag{
		Name:        "host",
		Usage:       "Route host without protocol. E.G: miruta.ruta.com",
		EnvVars:     []string{"host"},
		Destination: &host,
		Required:    true,
	}

	pathFlag := cli.StringFlag{
		Name:        "path",
		Usage:       "Route path. E.G. /custom-path",
		EnvVars:     []string{"path"},
		Destination: &path,
		Required:    true,
	}

	supraenvironmentFlag := cli.StringFlag{
		Name:        "supraenvironment",
		Usage:       "Supraenvironment. E.G: des",
		EnvVars:     []string{"SUPRAENVIRONMENT"},
		Destination: &supraenvironment,
		Required:    true,
	}

	getNamespaceFlags := []cli.Flag{
		&clusterFlag,
		&connectionDataYamlPathFlag,
		&hostFlag,
		&pathFlag,
		&supraenvironmentFlag,
	}

	getNamespace := &cli.Command{
		Name:   "getNamespace",
		Usage:  "Check for duplicates for a specific host and path.",
		Action: getNamespace,
		Flags:  getNamespaceFlags,
	}

	cmd := &cli.Command{
		Name:        "route",
		Usage:       "Route related actions",
		Flags:       []cli.Flag{},
		Subcommands: []*cli.Command{getNamespace},
	}
	Cmds = append(Cmds, cmd)
}

func getNamespace(_ *cli.Context) error {

	printUtil.Flags(map[string]interface{}{"cluster": cluster, "connectionDataYamlPath": connectionDataYamlPath,
		"host": host, "path": path, "supraenvironment": supraenvironment})

	config, err := connection.CreateNewConfig(connectionDataYamlPath, cluster, supraenvironment)
	if err != nil {
		return err
	}

	projectsList, err := listProjects(err, config)
	if err != nil {
		return err
	}

	clientRouteSetOcp, err := route.CreateClient(config)
	if err != nil {
		return err
	}

	projectCounter := 0
	routeCounter := 0

	for _, pj := range projectsList.Items {
		projectCounter += 1
		routeList, err := route.GetAll(clientRouteSetOcp, pj.Name)
		if err != nil {
			return err
		}

		for _, r := range routeList.Items {
			routeCounter += 1
			if host == r.Spec.Host && path == r.Spec.Path {
				log.Infof("+ Host and path found in %s project: Host: %s Path: %s Status: %s", pj.Name, r.Spec.Host, r.Spec.Path, r.Status.String())
			}
		}
	}
	log.Infof("-> Number of projects analyzed: %d. Total number of Routes processed: %d.", projectCounter, routeCounter)
	return nil
}

func listProjects(err error, config *rest.Config) (*projectApiV1.ProjectList, error) {
	clientProjectSetOcp, err := project.CreateClient(config)
	if err != nil {
		return nil, err
	}

	projectsList, err := project.GetAll(clientProjectSetOcp)
	if err != nil {
		return nil, err
	}
	return projectsList, nil
}
