/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentReplyV30Response struct for ToolsCommentReplyV30Response
type ToolsCommentReplyV30Response struct {
	//
	Code *int64                            `json:"code,omitempty"`
	Data *ToolsCommentReplyV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
