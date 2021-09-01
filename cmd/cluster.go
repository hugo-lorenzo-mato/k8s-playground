package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	fileUtil "k8s-playground/util/file"
	"os"
	"strings"
)

func init() {

	connectionDataYamlPathFlag := cli.StringFlag{
		Name:        "connectionDataYamlPath",
		Usage:       "Path to the yaml file containing the credentials and URIs to connect to the different clusters. E.G: example.yml",
		EnvVars:     []string{"CREDENTIALS_YAML_PATH"},
		Destination: &connectionDataYamlPath,
		Value:       "data/cluster.yml",
		Required:    false,
	}

	checkConnectionFlags := []cli.Flag{
		&connectionDataYamlPathFlag,
	}

	login := &cli.Command{
		Name:   "login",
		Usage:  "Retrieves cluster.yml data to build login command.",
		Action: login,
		Flags:  checkConnectionFlags,
	}

	cmd := &cli.Command{
		Name:        "cluster",
		Usage:       "Retrieves cluster.yml data to build some utilities.",
		Flags:       []cli.Flag{},
		Subcommands: []*cli.Command{login},
	}
	Cmds = append(Cmds, cmd)

}

type Supraenvironment struct {
	Clusters []struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		URL      string `json:"url"`
		User     string `json:"user"`
	} `json:"clusters"`
	Name string `json:"name"`
}

func login(_ *cli.Context) error {

	data, err := builData()
	if err != nil {
		return err
	}

	renderTable(data)

	return nil
}

func builData() ([][]string, error) {
	var sArr []Supraenvironment
	var data [][]string
	jsonFile, err := fileUtil.YamlInput2Json(connectionDataYamlPath)
	if err != nil {
		return nil, err
	}

	result := gjson.Get(string(jsonFile), "openshift")
	result.ForEach(func(key, value gjson.Result) bool {
		s := Supraenvironment{}
		if err := json.Unmarshal([]byte(value.String()), &s); err != nil {
			return false
		}
		sArr = append(sArr, s)
		return true // keep iterating
	})

	for _, supra := range sArr {
		for _, c := range supra.Clusters {
			if strings.Contains(c.Name, "aks") {
				continue
			}
			var arrStr = []string{supra.Name, c.Name, fmt.Sprintf("oc login --insecure-skip-tls-verify --username %s --password %s %s", c.User, c.Password, c.URL)}
			data = append(data, arrStr)
		}
	}
	return data, nil
}

func renderTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Supraenvironment", "Cluster", "Login"})
	// additional options
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetRowLine(true)
	table.SetHeaderLine(true)
	table.SetBorder(true)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.SetAutoMergeCells(true)
	// end additional options
	table.AppendBulk(data)
	table.Render()
}
