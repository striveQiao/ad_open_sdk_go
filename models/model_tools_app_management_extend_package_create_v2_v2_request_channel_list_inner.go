/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageCreateV2V2RequestChannelListInner struct for ToolsAppManagementExtendPackageCreateV2V2RequestChannelListInner
type ToolsAppManagementExtendPackageCreateV2V2RequestChannelListInner struct {
	// 渠道号，渠道号ID支持英文，数字，下划线和连字符-，不超过50个字符，超出部分会被截断，示例：oceanengine_ads_sample-12
	ChannelId string `json:"channel_id"`
	// 备注，渠道包备注信息，至多20个字符，超出部分会被截断处理
	Remark *string `json:"remark,omitempty"`
}
