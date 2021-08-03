package connection

import (
	"crypto/tls"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strings"
)

func getBearerFromUrl(s, host string) (string, error) {
	var re = regexp.MustCompile(`(?m)access_token=(?P<bearer>[a-zA-Z0-9_\/\\\-*~.+]*)&`)
	match := re.FindStringSubmatch(s)
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	if result["bearer"] != "" {
		return result["bearer"], nil
	}
	return "", errors.New(fmt.Sprint("access_token not found in URL %s", s))
}

func getCleanUrl(host string) (string, error) {
	log.Debug("- getCleanUrl started -")
	var re = regexp.MustCompile(`(?m)(?m)api.(?P<cleanurl>[a-zA-Z0-9.]*):[0-9]*`)
	match := re.FindStringSubmatch(host)
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	if result["cleanurl"] != "" {
		log.Debugf("Clean URL: %s", result["cleanurl"])
		log.Debug("- getCleanUrl Finished -")
		return result["cleanurl"], nil
	}
	return "", errors.New(fmt.Sprint("getCleanUrl failed - API URL to clean %s", host))
}

func getBearerToken(username, password, host string) (string, error) {
	log.Debug("getBearerToken started")
	var req *http.Request
	var errReq error
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	if strings.Contains(host, "api.") {
		hostFixed, err := getCleanUrl(host)
		if err != nil {
			return "", err
		}
		log.Debugf(fmt.Sprintf("https://oauth-openshift.apps.%s/oauth/authorize?client_id=openshift-challenging-client&response_type=token", hostFixed))
		req, errReq = http.NewRequest("GET", fmt.Sprintf("https://oauth-openshift.apps.%s/oauth/authorize?client_id=openshift-challenging-client&response_type=token", hostFixed), nil)
	} else {
		log.Debugf(fmt.Sprintf("https://%s/oauth/authorize?client_id=openshift-challenging-client&response_type=token", host))
		req, errReq = http.NewRequest("GET", fmt.Sprintf("https://%s/oauth/authorize?client_id=openshift-challenging-client&response_type=token", host), nil)
	}
	if errReq != nil {
		return "", errors.New(fmt.Sprintf("Error in http.NewRequest: %s", errReq))
	}
	req.SetBasicAuth(username, password)
	// Set X-CSRF-Token header to a non-empty value (https://docs.openshift.com/container-platform/3.5/architecture/additional_concepts/authentication.html#obtaining-oauth-tokens)
	req.Header.Set("X-Csrf-Token", "xxx")

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error in client.Do request: %s", err))
	}
	defer resp.Body.Close()

	s, err := getBearerFromUrl(resp.Request.URL.String(), host)
	if err != nil {
		return "", err
	}
	log.Debug("getBearerToken finished")
	return s, nil
}
