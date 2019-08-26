package controllers

import (
	"../httputil"
	"../models"
	"../models/azure"
	"../utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

//const (
//	// 1 - organization name
//	azureProjectsUrl = "https://dev.azure.com/%s/_apis/projects?api-version=5.1"
//)

func RegisterApiRouters(apiRouter *gin.RouterGroup) {
	apiRouter.GET("v1/validatePhone", validatePhone) // validatePhone?phone=79162795609
	apiRouter.GET("v1/getProjects", getProjects)
	apiRouter.GET("v1/getProjectMembers")
}

func validatePhone(c *gin.Context) {
	phone := c.Query("phone")

	//time.Sleep(time.Duration(Utils.RandomIntInRange(1, 3)) * time.Second)

	if phone == "" {
		c.String(http.StatusBadRequest, "You must provide a phone number")
		return
	}

	var numverify = verifyPhone(phone)
	var response = models.Response{StatusCode: 200, Result: numverify}

	if numverify == nil {
		c.String(http.StatusBadRequest, "Error while")
		return
	}

	c.JSON(http.StatusOK, response)
}

func verifyPhone(phone string) *models.Numverify {
	safePhone := url.QueryEscape(phone)

	apiUrl := fmt.Sprintf("http://apilayer.net/api/validate?access_key=1af9e2ba48e1e24dbf7f71ae28a4918b&number=%s", safePhone)

	// Build the request
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record models.Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return &record
}

func getProjects(c *gin.Context) {
	var projects *azure.ProjectsList
	utils.GetFromAzure(fmt.Sprintf(azureProjectsUrl, "databriz"), "5tnv6tlqvvnctftdpssymz7fgy7y47rht3oucr3bnbsyme6hagba", &projects)
	//var projects = fetchProjectsList("databriz", "5tnv6tlqvvnctftdpssymz7fgy7y47rht3oucr3bnbsyme6hagba")

	if projects == nil {
		code := http.StatusInternalServerError
		c.JSON(code, httputil.HTTPError{StatusCode: code, Message: "Error while requesting data from Azure"})
		return
	}

	c.JSON(http.StatusOK, projects.Projects)
}

/*func fetchProjectsList(organizationName, token string) *utils.AzureProjectsList {
	apiUrl := fmt.Sprintf(azureProjectsUrl, organizationName) // TODO provide organization name

	// Build the request
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	// Add authorization header to request
	req.Header.Add("Authorization", "Basic "+basicAuth("", token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record utils.AzureProjectsList

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	// Use json.Decode for reading streams of JSON data
	//if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
	//	log.Println(err)
	//}

	if err := json.Unmarshal([]byte(body), &record); err != nil {
		log.Println(err)
	}

	return &record

}*/
