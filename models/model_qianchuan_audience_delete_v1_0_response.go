/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceDeleteV10Response struct for QianchuanAudienceDeleteV10Response
type QianchuanAudienceDeleteV10Response struct {
	//
	Code *int64                                  `json:"code,omitempty"`
	Data *QianchuanAudienceDeleteV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
