/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigV2CreateV2Request struct for CreativeAutoGenerateConfigV2CreateV2Request
type CreativeAutoGenerateConfigV2CreateV2Request struct {
	// 计划id
	AdId int64 `json:"ad_id"`
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 策略数据(列表中strategy_id需唯一, 即同一个策略（strategy）的策略配置项列表(strategy_state)，需放到同一个对象内，不可分开传递)
	StrategyData []*CreativeAutoGenerateConfigV2CreateV2RequestStrategyDataInner `json:"strategy_data,omitempty"`
}
