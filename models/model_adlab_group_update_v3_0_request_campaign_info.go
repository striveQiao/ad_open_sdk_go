/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateV30RequestCampaignInfo 广告组维度信息
type AdlabGroupUpdateV30RequestCampaignInfo struct {
	// 广告组预算，单位：元 限制范围：300 <= budget <= 9999999.99
	Budget *float64 `json:"budget,omitempty"`
	// 广告组(管家项目)名称，不允许超过100个字
	Name *string `json:"name,omitempty"`
}
