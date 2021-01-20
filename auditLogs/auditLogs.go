package auditLogs

import (
	"encoding/json"
	"errors"
	"github.com/abdelmmu/amp/base"
	"github.com/abdelmmu/amp/model"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
)

//AuditLogs will implement the method to interact with AuditLog endpoint
type AuditLogs struct {
	Proxy base.AMP
	Model model.AuditLogs
}
//GUID is regexp for the Valid GUID
const GUID =`^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$`

//NewAuditLogs will take valid Auth and return
// an initialized AuditLogs
func NewAuditLogs(auth string) *AuditLogs{
	amp := base.NewAMP(auth)
	if amp == nil {
		return nil
	}
	return &AuditLogs{
		Proxy: *amp,
		Model: model.AuditLogs{},
	}
}

//GetAuditLogs supports request to the audit_logs.
//supports following params: audit_log_type,  event, start_time and end_time
//
//
//you set the Limit to true if you want set the limit value.
//just the first  limit value will be accepted.
func (a *AuditLogs) GetAuditLogs(params , value string, limit bool, limitValue ...int) (*model.ListAuditLogs, error){
	var queryParameters string
	var err error
	switch strings.ToLower(params) {
	case "audit_log_type":
		queryParameters = "?audit_log_type=" + value
	case "event":
		queryParameters = "?event=" + value
	case "start_time":
		queryParameters = "?start_time=" + value
	case "end_time":
		queryParameters = "?end_time=" + value
	default:
		err = errors.New("the Params should be onf of the following:" +
			"audit_log_type, audit_log_id, event, start_time and end_time")
		return nil, err
	}
	var limitParameters string
	if limit{
		//index the first value as is not allowed to provide more than one int
		if limitValue == nil{
			err = errors.New("you need to provide the limit value, or change limit to be false")
			return nil, err
		}
		limitParameters = "&limit=" + strconv.Itoa(limitValue[0])
	}
	parameters := queryParameters + limitParameters
	statusCode, body, err := a.Proxy.GenericReq("GET","/audit_logs", parameters)
	if err != nil{
		return nil,err
	}
	if statusCode != 200 {
		err = errors.New("unauthorized")
		logrus.WithFields(logrus.Fields{
			"error": "check the Auth and the API",
		}).WithError(err).Errorf("expected 200 ok, got %v", statusCode)
		return nil, err
	}

/*	jsonFile, err := os.Open("AuditLogs.json")
	if err != nil{
		return nil, err
	}
	defer jsonFile.Close()
	body, _ := ioutil.ReadAll(jsonFile)

*/
	//data := &model.BaseMoodle{}
	data := a.Model.Data
	err = json.Unmarshal(body, &data)
	if err != nil{
		logrus.Errorf("body %v", string(body))
		return nil, err
	}

	
	return &data, nil
}

//GetAuditLogsByID will take the audit_log_id as params.
//the ID type is GUID
//valid GUID is `^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$`
func (a *AuditLogs) GetAuditLogsByID(id string, limit bool, limitValue ...int)(*model.ListAuditLogsByID, error){
	validGUID :=  regexp.MustCompile(GUID)
	if id == "" || !validGUID.MatchString(id){
		nilError := errors.New("id is empty or is not valid GUID")
		logrus.WithFields(logrus.Fields{
			"error": nilError,
			"info":"if you need Generic request to API you can use {GetAuditLogs}",
			"validID": GUID,
		}).WithError(nilError).Error("check if ID is provided and it's valid")
		return nil, nilError
	}
	var limitParameters string
	if limit{
		//index the first value as is not allowed to provide more than one int
		if limitValue == nil{
			err := errors.New("you need to provide the limit value, or change limit to be false")
			return nil, err
		}
		limitParameters = "&limit=" + strconv.Itoa(limitValue[0])
	}
	parameters := id + limitParameters
	statusCode, body, err := a.Proxy.GenericReq("GET","/audit_logs", parameters)
	if err != nil{
		return nil,err
	}
	if statusCode != 200 {
		err = errors.New("unauthorized")
		logrus.WithFields(logrus.Fields{
			"error": "check the Auth and the API",
		}).WithError(err).Errorf("expected 200 ok, got %v", statusCode)
		return nil, err
	}

	data := a.Model.DataLogByID
	err = json.Unmarshal(body, &data)
	if err != nil{
		logrus.Errorf("body %v", string(body))
		return nil, err
	}

	return &data, nil
}


//GetAuditLogsByUser will take the audit_log_user as params.
//you can provide the Username as it's if contains '@'
//will be replaced by the function
func (a *AuditLogs) GetAuditLogsByUser(user string, limit bool, limitValue ...int)(*model.ListAuditLogsByUser, error){
	if user == "" {
		nilError := errors.New("id is empty")
		logrus.WithFields(logrus.Fields{
			"error": nilError,
			"info":"if you need Generic request to API you can use {GetAuditLogs}",
		}).WithError(nilError).Error("id is not optional you need to provide an ID to use this method")
		return nil, nilError
	}
	//check if the user contains @
	if strings.Contains(user,"@"){
		user = strings.Replace(user,"@","%40",1)
	}

	var limitParameters string
	if limit{
		//index the first value as is not allowed to provide more than one int
		if limitValue == nil{
			err := errors.New("you need to provide the limit value, or change limit to be false")
			return nil, err
		}
		limitParameters = "&limit=" + strconv.Itoa(limitValue[0])
	}

	parameters := user + limitParameters
	statusCode, body, err := a.Proxy.GenericReq("GET","/audit_logs", parameters)
	if err != nil{
		return nil,err
	}
	if statusCode != 200 {
		err = errors.New("unauthorized")
		logrus.WithFields(logrus.Fields{
			"error": "check the Auth and the API",
		}).WithError(err).Errorf("expected 200 ok, got %v", statusCode)
		return nil, err
	}

	data := a.Model.DataLogByUser
	err = json.Unmarshal(body, &data)
	if err != nil{
		logrus.Errorf("body %v", string(body))
		return nil, err
	}

	return &data, nil
}

//GetAuditLogsTypes will provide the request to /audit_log_types
func (a *AuditLogs) GetAuditLogsTypes()(*model.ListAuditLogsTypes, error){

	statusCode, body, err := a.Proxy.GenericReq("GET","/audit_log_types","")
	if err != nil{
		return nil,err
	}
	if statusCode != 200 {
		err = errors.New("unauthorized")
		logrus.WithFields(logrus.Fields{
			"error": "check the Auth and the API",
		}).WithError(err).Errorf("expected 200 ok, got %v", statusCode)
		return nil, err
	}
	data := a.Model.DataLogTypes
	err = json.Unmarshal(body, &data)
	if err != nil{
		logrus.Errorf("body %v", string(body))
		return nil, err
	}
	return &data, nil
}