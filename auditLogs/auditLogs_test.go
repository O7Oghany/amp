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

func TestAuditLogs_GetLogs(t *testing.T) {
	//invalid test for type
	_, err:= NewAuditLogs("test:test").GetLogs("audit_log_type","test",false, 1)
	if err != nil {
		t.Errorf("got Error %v", err)
	}

}