/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanBrandAuthorizedGetV10Response struct for QianchuanBrandAuthorizedGetV10Response
type QianchuanBrandAuthorizedGetV10Response struct {
	//
	Code *int64                                      `json:"code,omitempty"`
	Data *QianchuanBrandAuthorizedGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
