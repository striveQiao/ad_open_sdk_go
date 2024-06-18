/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRecommendKeywordsGetV10Filtering
type QianchuanAdRecommendKeywordsGetV10Filtering struct {
	// 抖音id，与计划创建时设置的抖音ID一致 注：PC侧创建「直播带货」搜索广告时，默认根据aweme_id获取推荐关键
	AwemeId *int64 `json:"aweme_id,omitempty"`
	// 商品ID，此参数用于获取与商品更为相关的关键词推荐 注：PC侧创建「短视频带货」搜索广告时，默认根据product_id获取推荐关键
	ProductId *int64 `json:"product_id,omitempty"`
	// 词根
	SearchWord    *string                                                   `json:"search_word,omitempty"`
	SuggestReason *QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason `json:"suggest_reason,omitempty"`
}
