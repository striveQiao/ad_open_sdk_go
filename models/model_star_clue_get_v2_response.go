/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarClueGetV2Response struct for StarClueGetV2Response
type StarClueGetV2Response struct {
	//
	Code *int64                     `json:"code,omitempty"`
	Data *StarClueGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
