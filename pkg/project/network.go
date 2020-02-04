package project

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eclipse/codewind-installer/pkg/connections"
	"github.com/eclipse/codewind-installer/pkg/sechttp"
	"github.com/eclipse/codewind-installer/pkg/utils"
)

type (
	ProjectNetwork struct {
		ProjectID     string `json:"projectID"`
		ProjectName   string `json:"projectName"`
		ProjectURL    string `json:"projectURL"`
		ConnectionID  string `json:"connectionID"`
		ConnectionURL string `json:"connectionURL"`
		Env           string `json:"env"`
	}

	ProjectNetworkList map[string]ProjectNetwork
)

// GetProject : Get project details from Codewind
func GetProjectNetwork(httpClient utils.HTTPClient, connection *connections.Connection, url, projectID string) (ProjectNetworkList, *ProjectError) {
	req, requestErr := http.NewRequest("GET", url+"/api/v1/projects/"+projectID+"/network/", nil)
	if requestErr != nil {
		return nil, &ProjectError{errOpRequest, requestErr, requestErr.Error()}
	}

	// send request
	resp, httpSecError := sechttp.DispatchHTTPRequest(httpClient, req, connection)
	if httpSecError != nil {
		return nil, &ProjectError{errOpRequest, httpSecError, httpSecError.Desc}
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		respErr := errors.New(textAPINotFound)
		return nil, &ProjectError{errOpNotFound, respErr, textAPINotFound}
	}

	byteArray, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, &ProjectError{errOpRequest, readErr, readErr.Error()}
	}
	var projectNetworkList ProjectNetworkList
	jsonErr := json.Unmarshal(byteArray, &projectNetworkList)
	if jsonErr != nil {
		return nil, &ProjectError{errOpRequest, jsonErr, jsonErr.Error()}
	}
	return projectNetworkList, nil
}

// GetProject : Get project details from Codewind
func CreateProjectNetwork(httpClient utils.HTTPClient, connection *connections.Connection, url, projectID string, newProjectNetwork ProjectNetwork) (*ProjectNetworkList, *ProjectError) {
	jsonPayload, _ := json.Marshal(newProjectNetwork)

	req, requestErr := http.NewRequest("POST", url+"/api/v1/projects/"+projectID+"/network/", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	if requestErr != nil {
		return nil, &ProjectError{errOpRequest, requestErr, requestErr.Error()}
	}

	// send request
	resp, httpSecError := sechttp.DispatchHTTPRequest(httpClient, req, connection)
	if httpSecError != nil {
		return nil, &ProjectError{errOpRequest, httpSecError, httpSecError.Desc}
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		respErr := errors.New(textAPINotFound)
		return nil, &ProjectError{errOpNotFound, respErr, textAPINotFound}
	}
	fmt.Println(resp.StatusCode)

	return nil, nil
}
