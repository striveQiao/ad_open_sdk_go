/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10ResponseDataErrorListInner struct for QianchuanAdUpdateV10ResponseDataErrorListInner
type QianchuanAdUpdateV10ResponseDataErrorListInner struct {
	//
	ErrorCode *int64 `json:"error_code,omitempty"`
	//
	ErrorMessage *string `json:"error_message,omitempty"`
	//
	ObjectId *int64 `json:"object_id,omitempty"`
	//
	ObjectType *int64 `json:"object_type,omitempty"`
	//
	OptType *int64 `json:"opt_type,omitempty"`
}
