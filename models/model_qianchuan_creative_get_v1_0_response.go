/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10Response struct for QianchuanCreativeGetV10Response
type QianchuanCreativeGetV10Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *QianchuanCreativeGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
