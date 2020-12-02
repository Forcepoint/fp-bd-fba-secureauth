package realms

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"secureauth/internal/structs"
)

func GetRealms(conf *viper.Viper) error {

	// Create client and authenticate
	client := createClient(conf)
	authenticateClient(client, conf)

	// Retrieve realms from endpoint
	realms := getRealms(client, conf)

	// Write realms to files
	writeRealms(realms)

	return nil

}

func createClient(conf *viper.Viper) *http.Client {

	// Create transport and cookie jar
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.GetBool("ignore_ssl")},
	}
	cookieJar, _ := cookiejar.New(nil)

	// Create client
	client := &http.Client{
		Jar: cookieJar,
		Transport: tr,
	}

	return client

}

func authenticateClient(client *http.Client, conf *viper.Viper) {

	// Build request for authentication
	url := fmt.Sprintf("%s/secureAuth0/localadmin.aspx", conf.GetString("admin_url"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Fatal("There was an error while generating the request to the admin realm to authenticate.")
	}

	// Send request for authentication
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatal("There was an issue with sending the request to the admin realm.")
	}
	defer resp.Body.Close()

	// Handle authentication request response
	if resp.StatusCode != http.StatusOK {
		logrus.Fatal(fmt.Sprintf("The status code returned when attempting to autheticate was: %d", resp.StatusCode))
	}

}

func getRealms(client *http.Client, conf *viper.Viper) []structs.Realm {

	// Build request for realm retrieval
	url := fmt.Sprintf("%s/httpproxy/api/v3/applications", conf.GetString("admin_url"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Fatal("There was an error while generating the request to retrieve realms from the application endpoint.")
	}

	// Send request and handle response
	resp, err := client.Do(req)
	if err != nil {
		logrus.Warn("There was an issue retrieving the required token.")
	}
	defer resp.Body.Close()

	// Handle returned data
	data := new(structs.Apps)
	realms := make([]structs.Realm, 0)

	if resp.StatusCode != http.StatusOK {
		logrus.Fatal(fmt.Sprintf("The status code returned was: %d", resp.StatusCode))
	}

	// Decode response data and format for realms files
	json.NewDecoder(resp.Body).Decode(data)
	for _, app := range data.Applications {
		realm := structs.Realm{
			Application: app.Name,
			File: fmt.Sprintf("D:/Secureauth/SecureAuth%d/AuditLogs/Audit.Log", app.Realm),
			LatestTime: "0000-00-00T00:00:00.000Z",
			Realm: fmt.Sprintf("SecureAuth%d", app.Realm),
		}
		realms = append(realms, realm)
	}

	return realms

}

func writeRealms(realms []structs.Realm)  {

	for _, realm := range realms {

		configFile := fmt.Sprintf("./config/realms/%s.yaml", realm.Realm)

		// If config file already exists don't overwrite
		if _, err := os.Stat(configFile); err == nil {
			continue
		}

		// Marshal data from struct
		data, err := yaml.Marshal(realm)
		if err != nil {
			logrus.Fatal("Unable to marshal realm data.")
		}

		// Write to config file
		err = ioutil.WriteFile(configFile, data, 0644)
		if err != nil {
			logrus.Fatal(fmt.Sprintf("Unable to write realm config file for realm %s. Error was: %s", realm.Realm, err))
		}

	}

}