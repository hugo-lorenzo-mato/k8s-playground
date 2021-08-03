package connection

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"k8s.io/client-go/rest"
	"os"
)

type ConnectionInfo struct {
	Host        string
	Password    string
	Username    string
	BearerToken string
}

func createTlsConfig(insecure bool, caFilePath string) rest.TLSClientConfig {
	cfg := rest.TLSClientConfig{}
	if insecure {
		log.Debug("TLS certificate validation is off")
		cfg.Insecure = true
		return cfg
	} else {
		log.Debug("setting CA certificate path")
		cfg.CAFile = caFilePath
		return cfg
	}
}

func yamlInput2Json(fp string) ([]byte, error) {
	inputFile, err := os.Open(fp)
	defer inputFile.Close()
	if err != nil {
		return nil, fmt.Errorf("Error reading: %s. err: %s", fp, err)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(inputFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading input: %s. err: %s", fp, err)
	}
	content := buf.Bytes()
	converted2Json, err := yaml.YAMLToJSON(content)
	if err != nil {
		return nil, fmt.Errorf("Error in yaml conversion: %s", err)
	}
	return converted2Json, nil
}

func getConfig(ptd ConnectionInfo) (*rest.Config, error) {
	log.Debug("getConfig started")
	bearer, err := getBearerToken(ptd.Username, ptd.Password, ptd.Host)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error in get BearerToken function: %s", err))
	}
	configuration := rest.Config{
		Host:            ptd.Host,
		BearerToken:     bearer,
		UserAgent:       "oc",
		TLSClientConfig: createTlsConfig(true, ""),
	}
	log.Debug("getConfig finished")
	return &configuration, nil
}

func getInfo(path, platform, supraenv string) (ConnectionInfo, error) {
	log.Debug("- getInfo started -")
	json, err := yamlInput2Json(path)
	if err != nil {
		return ConnectionInfo{}, err
	}

	ci := ConnectionInfo{
		Host:     gjson.Get(string(json), fmt.Sprintf("openshift.%s.%s.url", supraenv, platform)).String(),
		Username: gjson.Get(string(json), fmt.Sprintf("openshift.%s.%s.user", supraenv, platform)).String(),
		Password: gjson.Get(string(json), fmt.Sprintf("openshift.%s.%s.password", supraenv, platform)).String(),
	}
	log.Debug("- getInfo finished -")
	return ci, err
}

func CreateNewConfig(platformsYaml, cluster, supraenv string) (*rest.Config, error) {
	log.Debug("- CreateNewConfig started -")
	platformToDeploy, err := getInfo(platformsYaml, cluster, supraenv)
	if err != nil {
		return nil, err
	}
	log.Debugf("\nPlatform connection retrieved: Host: %s Username: %s", platformToDeploy.Host, platformToDeploy.Username)
	c, err := getConfig(platformToDeploy)
	log.Debug("- CreateNewConfig finished -")
	return c, nil
}
