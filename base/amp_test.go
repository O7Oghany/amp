package base

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewAMP(t *testing.T) {
	nilAmp:= NewAMP("tetet")
	if nilAmp != nil {
		t.Errorf("error: %v", nilAmp)
	}

	validAMP := NewAMP("client_id:api_key")
	if validAMP == nil {
		t.Errorf("error: %v", validAMP)
	}

}

func TestAMP_Do(t *testing.T) {
	//not checking for "client_id:api_key" because it's already checked in TestNewAMP Function
	validmp:= NewAMP("client_id:api_key")
	//ignoring the Error from the new request
	//test will return unauthorized
	req, _ := http.NewRequest("GET", validmp.BaseURL+"/computers", nil)
	resp, err := validmp.Do(req)
	if err != nil  {
		t.Errorf("error: %v", err)
		fmt.Println(resp)
	}
	//check that the validation was incorrect
	if resp.StatusCode != 401{
		t.Errorf("expected 401 and got : %v", resp.StatusCode)
	}
}

func TestAMP_GenericReq(t *testing.T) {
	//not checking for "client_id:api_key" because it's already checked in TestNewAMP Function
	validmp:= NewAMP("client_id:api_key")
	//giving non valid resource
	_, err := validmp.GenericReq("GET", "computers")
	if err == nil{
		t.Errorf("expected error but got nil")
	}else {
		t.Logf("success and got Error: %v", err)
	}
	//giving valid resource but not valid auth
	res, err :=  validmp.GenericReq("GET", "/computers")
	if err != nil{
		t.Errorf("expected error but got nil %v", err)
	}else {
		body := string(res)
		t.Logf("resp: %v", body)
	}

}