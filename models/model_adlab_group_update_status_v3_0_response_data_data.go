/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateStatusV30ResponseDataData 返回数据
type AdlabGroupUpdateStatusV30ResponseDataData struct {
	// 更新失败的项目id信息
	Errors []*AdlabGroupUpdateStatusV30ResponseDataDataErrorsInner `json:"errors,omitempty"`
	// 更新成功的项目id
	Ids []int64 `json:"ids,omitempty"`
}
