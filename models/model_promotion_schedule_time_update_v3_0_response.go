/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionScheduleTimeUpdateV30Response struct for PromotionScheduleTimeUpdateV30Response
type PromotionScheduleTimeUpdateV30Response struct {
	//
	Code *int64                                      `json:"code,omitempty"`
	Data *PromotionScheduleTimeUpdateV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
