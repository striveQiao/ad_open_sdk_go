/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreative
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreative struct {
	SplashDynamicImage *BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashDynamicImage `json:"splash_dynamic_image,omitempty"`
	SplashFullVideo    *BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashFullVideo    `json:"splash_full_video,omitempty"`
	SplashStaticImage  *BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashStaticImage  `json:"splash_static_image,omitempty"`
	SplashType         *BrandCreativeGetV30DataCreativesCreativeSplashCreativeSplashType                      `json:"splash_type,omitempty"`
}
