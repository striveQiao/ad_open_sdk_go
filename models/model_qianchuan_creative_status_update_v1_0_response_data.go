/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeStatusUpdateV10ResponseData
type QianchuanCreativeStatusUpdateV10ResponseData struct {
	//
	CreativeIds []int64 `json:"creative_ids,omitempty"`
	//
	Errors []*QianchuanCreativeStatusUpdateV10ResponseDataErrorsInner `json:"errors,omitempty"`
}
