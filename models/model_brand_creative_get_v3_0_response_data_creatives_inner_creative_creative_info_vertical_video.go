/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalVideo 竖版视频
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalVideo struct {
	// 标题
	Title *string                                                                              `json:"title,omitempty"`
	Video *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalVideoVideo `json:"video,omitempty"`
}
