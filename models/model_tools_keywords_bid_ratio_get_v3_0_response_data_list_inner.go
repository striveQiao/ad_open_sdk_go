/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsKeywordsBidRatioGetV30ResponseDataListInner struct for ToolsKeywordsBidRatioGetV30ResponseDataListInner
type ToolsKeywordsBidRatioGetV30ResponseDataListInner struct {
	//
	BidRatio  *float64                                      `json:"bid_ratio,omitempty"`
	Dimension *ToolsKeywordsBidRatioGetV30DataListDimension `json:"dimension,omitempty"`
	//
	ProjectNum *int64 `json:"project_num,omitempty"`
	//
	PromotionWordId *string `json:"promotion_word_id,omitempty"`
	//
	Word *string `json:"word,omitempty"`
}
