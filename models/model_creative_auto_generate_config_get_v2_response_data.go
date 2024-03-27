/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigGetV2ResponseData
type CreativeAutoGenerateConfigGetV2ResponseData struct {
	// 配置ID
	ConfigId *int64 `json:"config_id,omitempty"`
	// 策略配置详情列表(仅当version=Strategy时有值)
	StrategyData []*CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInner `json:"strategy_data,omitempty"`
	// 模板列表(仅当version=Template时有值)
	Templates []*CreativeAutoGenerateConfigGetV2ResponseDataTemplatesInner `json:"templates,omitempty"`
	// 上次修改时间戳
	UpdateTime *int64                                      `json:"update_time,omitempty"`
	Version    *CreativeAutoGenerateConfigGetV2DataVersion `json:"version,omitempty"`
}
