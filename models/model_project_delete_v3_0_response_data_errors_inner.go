/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectDeleteV30ResponseDataErrorsInner struct for ProjectDeleteV30ResponseDataErrorsInner
type ProjectDeleteV30ResponseDataErrorsInner struct {
	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
	// 项目ID
	ProjectId *int64 `json:"project_id,omitempty"`
}
