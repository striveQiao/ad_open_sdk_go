/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentQueryLinkV2ResponseDataComponentsInner struct for StarComponentQueryLinkV2ResponseDataComponentsInner
type StarComponentQueryLinkV2ResponseDataComponentsInner struct {
	//
	AuditBanReason *string `json:"audit_ban_reason,omitempty"`
	// Android落地页链接 HTTPS URL
	LinkAndroid *string `json:"link_android,omitempty"`
	// Android app 下载链接 HTTPS URL，Android下载时需要
	LinkAndroidDownload *string `json:"link_android_download,omitempty"`
	// iOS落地页链接 HTTPS URL
	LinkIos *string `json:"link_ios,omitempty"`
	// iOS app 下载链接 App store 链接，iOS下载时需要
	LinkIosDownload *string `json:"link_ios_download,omitempty"`
	// 组件名称 10字内
	LinkName *string `json:"link_name,omitempty"`
	// 组件投放目标 (1)：应用下载； (2)：线索留资； (3)：其他目标
	LinkPageTarget *int64 `json:"link_page_target,omitempty"`
	// 适用平台 (0)：不限； (1)：Android； (2)：iOS
	LinkPlatform *int64 `json:"link_platform,omitempty"`
	// 媒体渠道 仅限以下 (1)：抖音
	LinkPlatformSource *int64 `json:"link_platform_source,omitempty"`
	//
	LinkStatus *int64 `json:"link_status,omitempty"`
	// 组件标题（介绍文案） 20字内
	LinkTitle *string `json:"link_title,omitempty"`
	// 组件类型 仅限以下 (1)：落地页
	LinkType *int32 `json:"link_type,omitempty"`
}
