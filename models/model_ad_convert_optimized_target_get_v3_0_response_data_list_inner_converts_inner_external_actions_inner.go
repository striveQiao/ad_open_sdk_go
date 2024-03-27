/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdConvertOptimizedTargetGetV30ResponseDataListInnerConvertsInnerExternalActionsInner struct for AdConvertOptimizedTargetGetV30ResponseDataListInnerConvertsInnerExternalActionsInner
type AdConvertOptimizedTargetGetV30ResponseDataListInnerConvertsInnerExternalActionsInner struct {
	// 优化监测连接
	ActionTrackUrl *string `json:"action_track_url,omitempty"`
	// 线索通来源
	Belong          []*AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsBelong        `json:"belong,omitempty"`
	ConvertDataType *AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType `json:"convert_data_type,omitempty"`
	// 自定义优化ID
	ConvertId *int64 `json:"convert_id,omitempty"`
	// 深度优化目标
	DeepExternalActions []*AdConvertOptimizedTargetGetV30ResponseDataListInnerConvertsInnerExternalActionsInnerDeepExternalActionsInner `json:"deep_external_actions,omitempty"`
	// 深度优化是否禁用,true 表示已经禁用，false 表示可用
	Disabled       *bool                                                                        `json:"disabled,omitempty"`
	ExternalAction *AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsExternalAction `json:"external_action,omitempty"`
	// 优化类型名称
	ExternalActionName *string `json:"external_action_name,omitempty"`
	// 多优化目标
	ExternalActions []*AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsExternalActions `json:"external_actions,omitempty"`
	// 自定义优化名称
	ExternalName *string                                                              `json:"external_name,omitempty"`
	Source       *AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource `json:"source,omitempty"`
}
