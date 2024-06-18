/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorThirdSiteInfo 第三方落地页设置
type NativeAnchorUpdateV30RequestAnchorInfoAppEcommerceAnchorThirdSiteInfo struct {
	// 安卓端落地页链接，内部需要包含应用下载链接
	AndroidExternalUrl *string `json:"android_external_url,omitempty"`
	// 第三方落地页
	ExternalUrl *string `json:"external_url,omitempty"`
	// iOS端落地页链接，内部需要包含应用下载链接
	IosExternalUrl *string `json:"ios_external_url,omitempty"`
}
