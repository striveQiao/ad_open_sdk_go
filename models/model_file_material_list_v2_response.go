/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialListV2Response struct for FileMaterialListV2Response
type FileMaterialListV2Response struct {
	//
	Code *int64                          `json:"code,omitempty"`
	Data *FileMaterialListV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
