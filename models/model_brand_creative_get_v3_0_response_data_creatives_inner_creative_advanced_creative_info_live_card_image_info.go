/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoLiveCardImageInfo 图片
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoLiveCardImageInfo struct {
	// 图片格式:jpg/jpeg/png/gif
	Format *string `json:"format,omitempty"`
	// 图片高度
	Height *int64 `json:"height,omitempty"`
	// 图片地址
	WebUri *string `json:"web_uri,omitempty"`
	// 图片前端展示地址
	WebUriShowUrl *string `json:"web_uri_show_url,omitempty"`
	// 图片宽度
	Width *int64 `json:"width,omitempty"`
}
