/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30RequestAnchorInfoAppEcommerceAnchorThirdSiteInfo 第三方落地页设置
type NativeAnchorCreateV30RequestAnchorInfoAppEcommerceAnchorThirdSiteInfo struct {
	//  安卓下载链接
	AndroidDownloadUrl *string `json:"android_download_url,omitempty"`
	// 第三方落地页
	ExternalUrl *string `json:"external_url,omitempty"`
	// ios下载链接
	IosDownloadUrl *string `json:"ios_download_url,omitempty"`
}
