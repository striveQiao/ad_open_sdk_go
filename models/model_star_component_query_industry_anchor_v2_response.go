/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentQueryIndustryAnchorV2Response struct for StarComponentQueryIndustryAnchorV2Response
type StarComponentQueryIndustryAnchorV2Response struct {
	//
	Code *int64                                          `json:"code,omitempty"`
	Data *StarComponentQueryIndustryAnchorV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
