package model

//BaseMoodle will hold the base date for response.
//it does not contain the actual data for individual response
//
//
// needed to get the error message when errors occurred
//more info https://api-docs.amp.cisco.com/api_resources?api_host=api.eu.amp.cisco.com&api_version=v1
type BaseMoodle struct {
	Version string `json:"version"`
	Data    `json:"data"`
	Errors []Errors `json:"errors"`
}

//Data will hold the actual data for individual response
//will not be implemented in the base model
type Data struct{}

//Errors contain the ErrorCode and the details.
type Errors struct {
	ErrorCode   int      `json:"error_code"`
	Description string   `json:"description"`
	Details     []string `json:"details"`
}
