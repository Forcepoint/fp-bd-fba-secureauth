package connectors

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func Send(method string, url string, data []byte, response interface{}, conf *viper.Viper) bool {

	// Build request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		logrus.Warn(fmt.Sprintf("Unable to build the %s request to %s.", method, url))
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	// Build client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.GetBool("ignore_ssl")},
	}
	client := &http.Client{Transport: tr}

	// Send request and handle response
	resp, err := client.Do(req)
	if err != nil {
		logrus.Warn(fmt.Sprintf("There was an issue sending the event to the endpoint: %s.", url))
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Warn(fmt.Sprintf("There was an error when requesting the endpoint: %s. The status code returned was: %d", url, resp.StatusCode))
		return false
	}

	// Handle return
	json.NewDecoder(resp.Body).Decode(response)

	return true

}
