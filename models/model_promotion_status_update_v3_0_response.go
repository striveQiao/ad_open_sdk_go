/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionStatusUpdateV30Response struct for PromotionStatusUpdateV30Response
type PromotionStatusUpdateV30Response struct {
	//
	Code *int64                                `json:"code,omitempty"`
	Data *PromotionStatusUpdateV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
