/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupsDeleteV30ResponseDataData 返回数据
type AdlabGroupsDeleteV30ResponseDataData struct {
	// 删除失败的项目id信息
	Errors []*AdlabGroupsDeleteV30ResponseDataDataErrorsInner `json:"errors,omitempty"`
	// 删除成功的项目id
	Ids []int64 `json:"ids,omitempty"`
}
