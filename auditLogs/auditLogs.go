package auditLogs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abdelmmu/amp/base"
	"github.com/abdelmmu/amp/model"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type AuditLogs struct {
	Proxy base.AMP
	Model model.AuditLogs
}

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

//Get will make ar request without params to the base url.
//TODO Tested on file not on the model itself
func (a *AuditLogs) GetLogs(params , value string, limit bool, limitValue ...int) (*model.ListAuditLogs, error){
	var queryParameters string
	var err error

	switch strings.ToLower(params) {
	case "audit_log_type":
		queryParameters = "?audit_log_type=" + value
	case "audit_log_id":
		queryParameters = "?audit_log_id=" + value
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
	fmt.Println(parameters)
	statusCode, body, err := a.Proxy.GenericReq("GET","/audit_logs")
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
	fmt.Println(a.Model.Data)
	
	return &data, nil
}