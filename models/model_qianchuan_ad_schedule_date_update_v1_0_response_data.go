/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdScheduleDateUpdateV10ResponseData
type QianchuanAdScheduleDateUpdateV10ResponseData struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	Errors []*QianchuanAdScheduleDateUpdateV10ResponseDataErrorsInner `json:"errors,omitempty"`
}
