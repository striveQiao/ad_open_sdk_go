/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCostProtectStatusGetV30ResponseData
type ProjectCostProtectStatusGetV30ResponseData struct {
	// 项目成本保障信息列表
	CompensateStatusInfoList []*ProjectCostProtectStatusGetV30ResponseDataCompensateStatusInfoListInner `json:"compensate_status_info_list,omitempty"`
}
