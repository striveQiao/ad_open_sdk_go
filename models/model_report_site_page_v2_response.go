/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportSitePageV2Response struct for ReportSitePageV2Response
type ReportSitePageV2Response struct {
	//
	Code *int64                        `json:"code,omitempty"`
	Data *ReportSitePageV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
