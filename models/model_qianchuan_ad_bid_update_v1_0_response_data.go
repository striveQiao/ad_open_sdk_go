/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdBidUpdateV10ResponseData
type QianchuanAdBidUpdateV10ResponseData struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	Errors []*QianchuanAdBidUpdateV10ResponseDataErrorsInner `json:"errors,omitempty"`
}
