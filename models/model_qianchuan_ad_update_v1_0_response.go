/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10Response struct for QianchuanAdUpdateV10Response
type QianchuanAdUpdateV10Response struct {
	//
	Code *int64                            `json:"code,omitempty"`
	Data *QianchuanAdUpdateV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
