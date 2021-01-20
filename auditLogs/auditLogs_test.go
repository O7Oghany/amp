package auditLogs

import (
	"testing"
)

func TestNewAuditLogs(t *testing.T) {
	//test without providing any Auth
	invalidAuudit := NewAuditLogs("")
	if invalidAuudit !=nil{
		t.Errorf("%v", invalidAuudit)
	}
	//providing invalid format
	invalidFormat := NewAuditLogs("test")
	if invalidFormat != nil{
		t.Errorf("failed. Expected nil, got: %v", invalidFormat)
	}
	//success test  should not get nil
	validAudit := NewAuditLogs("test:test")
	if validAudit ==nil {
		t.Errorf("failed and  audit is %v", validAudit)
	}
}

//TODO test will fail always now because no access to the API
func TestAuditLogs_GetLogs(t *testing.T) {
	//invalid test for type
	res, err:= NewAuditLogs("test:test").GetAuditLogs("","test",true)
	if err == nil {
		t.Errorf("should get error, but  res %v", res)
	}

}

func TestAuditLogs_GetAuditLogsByID(t *testing.T) {
	res, err:= NewAuditLogs("test:test").GetAuditLogsByUser("",false)
	if err == nil{
		t.Errorf("expected Error for leaving the ID empty, got %v",res)
	}
	//check if the ID is GUID or not
	_, err = NewAuditLogs("test:test").GetAuditLogsByUser("test",true)
	if err == nil{
		t.Errorf("expected Error for providing invalid GUID, but got %v",res)
	}
	_, err = NewAuditLogs("test:test").GetAuditLogsByUser("e773a9eb-296c-40df-98d8-bed46322589d",true)
	if err == nil{
		t.Errorf("expected Error for leaving the limit value , got %v",res)
	}
	_, err = NewAuditLogs("test:test").GetAuditLogsByUser("e773a9eb-296c-40df-98d8-bed46322589d",true,5)
	if err == nil {
		t.Errorf("should get unauthorized but got: %v", err)
	}
	//TODO uncomment this if you have valid auth
	/*data, err := NewAuditLogs("test:test").GetAuditByID("e773a9eb-296c-40df-98d8-bed46322589d", true, 5)
	if err != nil {
		t.Errorf("should get unauthorized but got: %v", err)
	}else {
		t.Log(data.Data)
	}*/
}

func TestAuditLogs_GetAuditLogsByUser(t *testing.T) {
	res, err:= NewAuditLogs("test:test").GetAuditLogsByUser("",false)
	if err == nil{
		t.Errorf("expected Error for leaving the ID empty, got %v",res)
	}else {
		t.Logf("Error is %v", err)
	}
	_, err = NewAuditLogs("test:test").GetAuditLogsByUser("amp@cisco.com",true)
	if err == nil{
		t.Errorf("expected Error for leaving the limit value , got %v",res)
	}else {
		t.Logf("Error is %v", err)
	}
	_, err = NewAuditLogs("test:test").GetAuditLogsByUser("amp@cisco.com",true,5)
	if err == nil {
		t.Errorf("should get unauthorized but got: %v", err)
	}
	//TODO uncomment this if you have valid auth
	/*data, err := NewAuditLogs("test:test").GetAuditByID("e773a9eb-296c-40df-98d8-bed46322589d", true, 5)
	if err != nil {
		t.Errorf("should get unauthorized but got: %v", err)
	}else {
		t.Log(data.Data)
	}*/
}