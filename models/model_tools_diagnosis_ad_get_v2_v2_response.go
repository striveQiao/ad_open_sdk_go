/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2Response struct for ToolsDiagnosisAdGetV2V2Response
type ToolsDiagnosisAdGetV2V2Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *ToolsDiagnosisAdGetV2V2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
