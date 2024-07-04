/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageUpdateV2V2RequestChannelInfoListInner struct for ToolsAppManagementExtendPackageUpdateV2V2RequestChannelInfoListInner
type ToolsAppManagementExtendPackageUpdateV2V2RequestChannelInfoListInner struct {
	// 渠道号，渠道号ID支持英文、数字、下划线和连字符-，不超过50个字符，超出部分会被截断。示例：oceanengine_ads_sample-12
	ChannelId string `json:"channel_id"`
	// 下载链接，（mode=Customize时需指定）渠道包（子包）的下载地址
	DownloadUrl *string `json:"download_url,omitempty"`
	// 下载链接对应包体的md5，（mode=Customize时需指定）
	Md5 *string `json:"md5,omitempty"`
	// 备注，渠道包备注信息，至多20个字符，超出部分会被截断处理
	Remark *string `json:"remark,omitempty"`
}
