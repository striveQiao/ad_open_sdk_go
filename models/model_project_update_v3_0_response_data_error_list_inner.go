/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30ResponseDataErrorListInner struct for ProjectUpdateV30ResponseDataErrorListInner
type ProjectUpdateV30ResponseDataErrorListInner struct {
	//
	ErrorCode int64 `json:"error_code"`
	//
	ErrorMessage string                                  `json:"error_message"`
	ObjectType   ProjectUpdateV30DataErrorListObjectType `json:"object_type"`
}
