/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashFullVideo 开屏全屏视频信息
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashFullVideo struct {
	// 背景图片列表
	ImageInfoBkList []*BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashFullVideoImageInfoBkListInner `json:"image_info_bk_list,omitempty"`
	// 视频列表
	VideoList []*BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashFullVideoVideoListInner `json:"video_list,omitempty"`
}
