package connection

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	fileUtil "k8s-playground/util/file"
	"k8s.io/client-go/rest"
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
	json, err := fileUtil.YamlInput2Json(path)
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
