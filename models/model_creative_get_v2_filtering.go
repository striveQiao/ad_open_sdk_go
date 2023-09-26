/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeGetV2Filtering
type CreativeGetV2Filtering struct {
	// 按照ad_id过滤
	AdId *int64 `json:"ad_id,omitempty"`
	// 按照campaign_id过滤
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// 广告创意创建时间，格式yyyy-MM-dd，表示过滤出当天创建的广告创意
	CreativeCreateTime *string `json:"creative_create_time,omitempty"`
	// 按照creative_id过滤，最多传100个。创意ID需属于当前广告主，否则会报错
	CreativeIds []int64 `json:"creative_ids,omitempty"`
	// 广告创意更新时间，格式yyyy-MM-dd，表示过滤出当天更新的广告创意
	CreativeModifyTime *string `json:"creative_modify_time,omitempty"`
	// 按照creative_title过滤，支持模糊搜索。支持的最大长度为30
	CreativeTitle *string                            `json:"creative_title,omitempty"`
	ImageMode     *CreativeGetV2FilteringImageMode   `json:"image_mode,omitempty"`
	LandingType   *CreativeGetV2FilteringLandingType `json:"landing_type,omitempty"`
	Pricing       *CreativeGetV2FilteringPricing     `json:"pricing,omitempty"`
	// 按照创意状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示，[【附录-广告创意状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=1696710760171535)
	Status *string `json:"status,omitempty"`
}
