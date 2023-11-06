/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectDeleteV30ResponseData
type ProjectDeleteV30ResponseData struct {
	// 错误信息
	Errors []*ProjectDeleteV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 项目id列表
	ProjectIds []int64 `json:"project_ids,omitempty"`
}
