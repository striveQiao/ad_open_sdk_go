/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdStatusUpdateV10Response struct for QianchuanUniPromotionAdStatusUpdateV10Response
type QianchuanUniPromotionAdStatusUpdateV10Response struct {
	//
	Code *int64                                              `json:"code,omitempty"`
	Data *QianchuanUniPromotionAdStatusUpdateV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
