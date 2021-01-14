package base

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"
)

//AMP TODO NOTE
type AMP struct {
	BaseURL  string
	ClientID string
	APIKey   string
	Client   CustomClient
}

//CustomClient provides the interface for custom Do request
//Do not need to custom the whole http.Client, just Do to add basic Auth
type CustomClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//AMPBaseURL will expose the base URL
//the purpose of Exposing publicly in case will be changed in the future
//or if the consumer wants to access v0
const AMPBaseURL = "https://api.eu.amp.cisco.com/v1"

//validResource contain the regular expression for the valid resource
//resource must begin with "/"
const validResource = `\/[-a-zA-Z0-9@:%_\+.~#?&//=]*`

//NewAMP return the construct for  AMP.
//Take clientID and apiKey as arguments
// this SDK support v1
//
// more info can be found: https://api-docs.amp.cisco.com/api_resources?api_host=api.eu.amp.cisco.com&api_version=v1
func NewAMP(clientID, apiKey string) *AMP {
	return &AMP{
		BaseURL:  AMPBaseURL,
		ClientID: clientID,
		APIKey:   apiKey,
	}
}

//Do implements the CustomClient interface to invoke a request.
//all sub requests that go through this custom Do do not need
//to add headers or Auth
//
//
//Do will implement basic auth as well as the required headers for the API calls
func (a *AMP) Do(r *http.Request) (*http.Response, error) {
	baseURL := *r.URL
	newreq, err := http.NewRequest(r.Method, baseURL.String(), r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("could not invoke the http request check the Error Message")
		return nil, err
	}
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

//GenericReq acts as generic HTTP reqquest method
//simplify the Do Method and checking for errors
func (a *AMP) GenericReq(res, method string) ([]byte, error) {
	url := a.BaseURL + res
	var resource = regexp.MustCompile(validResource)
	if !resource.MatchString(res) {
		logrus.WithFields(logrus.Fields{
			"Error": "not correct resource patter.",
			"Info:": "make sure that you have '/' for the resource. check https://api-docs.amp.cisco.com/api_resources?api_host=api.eu.amp.cisco.com&api_version=v1",
		}).Errorf("please check the resources. Check Info for more information")
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
			"Url":   url,
		}).Errorf("could not invoke the http request check the Error Message")
		return nil, err
	}
	result, err := a.Client.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error":     err,
			"ErrorCode": result.StatusCode,
		}).Error("could not invoke the http request check the Error Message")
		return nil, err
	}
	defer result.Body.Close()

	b, err := ioutil.ReadAll(result.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error":        err,
			"ErrorCode":    result.StatusCode,
			"ErrorMessage": "", //TODO
		}).Error("could not read the http response, check the Error Message")
		return nil, err
	}
	return b, nil

}
