package base

import (
	"net/http"
	"testing"
)

func TestNewAMP(t *testing.T) {
	nilAmp := NewAMP("tetet")
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
	validmp := NewAMP("client_id:api_key")

	//intialize fake AMP without AUTH
	invalidAMP := &AMP{}
	request, _ := http.NewRequest("GET", "", nil)
	_, err := validmp.Do(request)
	if err == nil {
		t.Errorf("expected nil, got %v", err)
	}

	//ignoring the Error from the new request
	invalidReq, _ := http.NewRequest("test", invalidAMP.BaseURL+"/computers", nil)
	_, err = validmp.Do(invalidReq)
	if err == nil {
		t.Errorf("error: %v", err)
	} else {
		t.Logf("success and got err: %v", err)
	}
}

func TestAMP_GenericReq(t *testing.T) {
	//not checking for "client_id:api_key" because it's already checked in TestNewAMP Function
	validmp := NewAMP("client_id:api_key")
	//giving non valid resource
	_,_, err := validmp.GenericReq("GET", "computers")
	if err == nil {
		t.Errorf("expected error but got nil")
	} else {
		t.Logf("success and got Error: %v", err)
	}
	//giving valid resource but not valid auth
	_,res, err := validmp.GenericReq("GET", "/computers")
	if err != nil {
		t.Errorf("expected error but got nil %v", err)
	} else {
		body := string(res)
		t.Logf("resp: %v", body)
	}

}
