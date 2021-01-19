package model

import "time"

//AuditLogs provide the query parameters
//as well as the expected response struct
//
//
//for more info about the mode:
//https://api-docs.amp.cisco.com/api_actions/details?api_action=GET+%2Fv1%2Faudit_logs&api_host=api.eu.amp.cisco.com&api_resource=AuditLog&api_version=v1
type AuditLogs struct {
	QueryParams QueryParameters
	Data ListAuditLogs
	DataLogByUser ListAuditLogsByUser
	DataLogByID ListAuditLogsByID
	AuditLogsError AuditLogsError
}

type AuditLogsError struct {
	Version  string `json:"version"`
	Data ListAuditLogsData `json:"data"`
	Errors []Errors `json:"errors"`
}
//QueryParameters to filter the request.
type QueryParameters struct {
	AuditLogID string `json:"audit_log_id"`
	AuditLogUser string `json:"audit_log_user"`
	AuditLogType string `json:"audit_log_type"`
	Limit string `json:"limit"`
	Event string `json:"event"`
	//time is ISO8601 (RFC3339)
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}

//ListAuditLogs holds the data from
///v1/audit_logs without params or with the following params:
//
//
//audit_log_type, audit_log_id, event, start_time and end_time
type ListAuditLogs struct {
	Version  string `json:"version"`
	Metadata `json:"metadata"`
	Data []ListAuditLogsData`json:"data"`
}

//ListAuditLogsByUser holds the data from
//?audit_log_user={{user}}
type ListAuditLogsByUser struct {
	Version  string `json:"version"`
	Metadata `json:"metadata"`
	Data []interface{}`json:"data"`
}

//ListAuditLogsUser holds the data from
//?audit_log_id={{ID}}
type ListAuditLogsByID struct {
	Version  string `json:"version"`
	Metadata `json:"metadata"`
	Data []ListAuditLogsDataByID `json:"data"`
}

//Metadata hold the Links and Results
type Metadata struct {
	Links `json:"links"`
	Results  `json:"results"`
}

//Links holds the info about Self and Next
type Links struct {
	Self string `json:"self"`
	Next string `json:"next"`
}

//Results holds the relevant infos but not the data itself.
type Results struct {
	Total            int `json:"total"`
	CurrentItemCount int `json:"current_item_count"`
	Index            int `json:"index"`
	ItemsPerPage     int `json:"items_per_page"`
}

//ListAuditLogsData is the actual data returned from the request.
type ListAuditLogsData struct {
	Event         string    `json:"event"`
	AuditLogType  string    `json:"audit_log_type"`
	AuditLogID    string    `json:"audit_log_id"`
	AuditLogUser  string    `json:"audit_log_user"`
	CreatedAt     time.Time `json:"created_at"`
	OldAttributes `json:"old_attributes,omitempty"`
	NewAttributes `json:"new_attributes,omitempty"`
}

//ListAuditLogsDataByID is the data returned filtering by audit_log_id.
type ListAuditLogsDataByID struct {
	Event         string    `json:"event"`
	AuditLogType  string    `json:"audit_log_type"`
	AuditLogID    string    `json:"audit_log_id"`
	AuditLogUser  string    `json:"audit_log_user"`
	CreatedAt     time.Time `json:"created_at"`
	OldAttributesSha `json:"old_attributes,omitempty"`
	NewAttributesSha `json:"new_attributes,omitempty"`
}

//OldAttributes holds previous data about the Attributes.
//if there or nil
//Attributes can be omitted in some requests.
type OldAttributes struct {
	Name     string      `json:"name"`
	Active   bool        `json:"active"`
	Default  bool        `json:"default"`
	Ancestry interface{} `json:"ancestry"`
}

//NewAttributes holds new data about the Attributes.
//if there or nil
//Attributes can be omitted in some requests.
type NewAttributes struct {
	Name     interface{} `json:"name"`
	Active   interface{} `json:"active"`
	Default  interface{} `json:"default"`
	Ancestry interface{} `json:"ancestry"`
}

//OldAttributesSha is special to
//audit_log_id request holds the old sha or nil
type OldAttributesSha struct {
	Sha string `json:"sha"`
}

//NewAttributesSha is special to
//audit_log_id request holds the new sha
type NewAttributesSha struct {
	Sha interface{} `json:"sha"`
}
