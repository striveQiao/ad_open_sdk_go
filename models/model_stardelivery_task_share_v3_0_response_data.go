/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskShareV30ResponseData
type StardeliveryTaskShareV30ResponseData struct {
	//
	ErrorList []*StardeliveryTaskShareV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []int64 `json:"success_list"`
}
