/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashDynamicImage 开屏动态图信息
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashDynamicImage struct {
	// 背景图片列表
	ImageInfoBkList []*BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashDynamicImageImageInfoBkListInner `json:"image_info_bk_list,omitempty"`
	// 图片列表
	ImageInfoList []*BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashDynamicImageImageInfoListInner `json:"image_info_list,omitempty"`
}
