/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentReplyV10ResponseData
type EnterpriseCommentReplyV10ResponseData struct {
	//
	Errors []*EnterpriseCommentReplyV10ResponseDataErrorsInner `json:"errors,omitempty"`
	//
	Success []*EnterpriseCommentReplyV10ResponseDataSuccessInner `json:"success,omitempty"`
}
