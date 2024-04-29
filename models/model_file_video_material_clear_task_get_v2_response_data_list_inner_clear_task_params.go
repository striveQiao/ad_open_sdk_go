/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoMaterialClearTaskGetV2ResponseDataListInnerClearTaskParams 任务的参数
type FileVideoMaterialClearTaskGetV2ResponseDataListInnerClearTaskParams struct {
	//
	ClearMaterialTypes []*FileVideoMaterialClearTaskGetV2DataListClearTaskParamsClearMaterialTypes `json:"clear_material_types,omitempty"`
	//
	Convert *int64 `json:"convert,omitempty"`
	//
	Cost *float64 `json:"cost,omitempty"`
	//
	CreateTimeUpper *string `json:"create_time_upper,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	// 待清理的素材列表
	MaterialIds    []int64                                                               `json:"material_ids,omitempty"`
	MaterialSource *FileVideoMaterialClearTaskGetV2DataListClearTaskParamsMaterialSource `json:"material_source,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
