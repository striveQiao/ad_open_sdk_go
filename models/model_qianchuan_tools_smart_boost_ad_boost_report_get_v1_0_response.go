/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostReportGetV10Response struct for QianchuanToolsSmartBoostAdBoostReportGetV10Response
type QianchuanToolsSmartBoostAdBoostReportGetV10Response struct {
	//
	Code *int64                                                   `json:"code,omitempty"`
	Data *QianchuanToolsSmartBoostAdBoostReportGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
