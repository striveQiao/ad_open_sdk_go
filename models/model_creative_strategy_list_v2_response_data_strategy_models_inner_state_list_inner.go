/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeStrategyListV2ResponseDataStrategyModelsInnerStateListInner struct for CreativeStrategyListV2ResponseDataStrategyModelsInnerStateListInner
type CreativeStrategyListV2ResponseDataStrategyModelsInnerStateListInner struct {
	Limit *CreativeStrategyListV2ResponseDataStrategyModelsInnerStateListInnerLimit `json:"limit,omitempty"`
	// 配置项key
	StateKey string `json:"state_key"`
	// 配置项名称
	StateName string                                                     `json:"state_name"`
	StateType CreativeStrategyListV2DataStrategyModelsStateListStateType `json:"state_type"`
}
