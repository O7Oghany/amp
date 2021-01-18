package base

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
)

type AMP struct {
	BaseURL string
	Auth string
	Client   Client
}
//Client provides the interface for custom Do request
//Do not need to custom the whole http.Client, just Do to add basic Auth
type Client interface {
	Do(req *http.Request)(*http.Response, error)
}

//AMPBaseURL will expose the base URL
//the purpose of Exposing publicly in case will be changed in the future
//or if the consumer wants to access v0
const AMPBaseURL = "https://api.eu.amp.cisco.com/v1"

//validResource contain the regular expression for the valid resource
//resource must begin with "/"
const validResource =`\/[-a-zA-Z0-9@:%_\+.~#?&//=]*`

//validAuth will validate that the user provide
//valid "client_id:api_key"
const validAuth = `(?P<key>\w.+):(?P<value>\w.+)$`

//NewAMP return the construct for  AMP.
//Take clientID and apiKey as arguments
//this SDK support v1
//
//implements the default HTTPClient.
//more info can be found: https://api-docs.amp.cisco.com/api_resources?api_host=api.eu.amp.cisco.com&api_version=v1
func NewAMP( auth string) *AMP {
	pattern := regexp.MustCompile(validAuth)
	if !pattern.MatchString(auth){
		logrus.WithFields(logrus.Fields{
			"Error": "not correct auth patter.",
			"Info:":"make sure that you have provided zhe auth in following format 'client_id:api_key' ",
		}).Errorf("please check the auth Format. Provide both in format 'client_id:api_key' ")
		return nil
	}
	return &AMP{
		BaseURL: AMPBaseURL,
		Auth: auth,
		Client: &http.Client{},
	}
}

//Do implements the CustomClient interface to invoke a request.
//all sub requests that go through this custom Do do not need
//to add headers or Auth
//
//
//Do will implement basic auth as well as the required headers for the API calls
func(a *AMP) Do (r *http.Request)(*http.Response, error){
	baseURL :=  *r.URL
	newreq, err := http.NewRequest(r.Method,baseURL.String(),r.Body)
	if err != nil{
		logrus.WithFields(logrus.Fields{
			"Error":err,
		}).Error("could not invoke the http request check the Error Message")
		return nil, err
	}
	auth := fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(a.Auth)))
	newreq.Header.Set("Authorization", auth)
	newreq.Header.Set("Accept", "application/json")
	newreq.Header.Set("content-type", "application/json")

	res, err := a.Client.Do(newreq)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("could not invoke the http request check the Error Message")
		return nil, err
	}
	return res, nil
}
//GenericReq acts as generic HTTP request method
//simplify the Do Method and checking for errors
//
//
//returning byte, all sub request can implement encoder
func(a *AMP) GenericReq(method, res string)([]byte, error){
	url := a.BaseURL + res
	var resource = regexp.MustCompile(validResource)
	if !resource.MatchString(res){
		err := errors.New("not correct resource patter")
		logrus.WithError(err)
		logrus.WithFields(logrus.Fields{
			"Error": "not correct resource patter.",
			"Info:":"make sure that you have '/' for the resource. check https://api-docs.amp.cisco.com/api_resources?api_host=api.eu.amp.cisco.com&api_version=v1",
		}).WithError(err)
		return nil, err
	}
	req, err := http.NewRequest(method,url,nil)
	if err != nil{
		logrus.WithFields(logrus.Fields{
			"Error":err,
			"Url": url,
		}).Errorf("could not invoke the http request check the Error Message")
		return nil, err
	}
	r,err := a.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("could not invoke the http request check the Error Message")
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("could not read the http response")
		return nil, err
	}
	return body, nil

}