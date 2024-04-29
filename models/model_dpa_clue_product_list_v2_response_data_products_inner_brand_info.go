/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductListV2ResponseDataProductsInnerBrandInfo 品牌信息
type DpaClueProductListV2ResponseDataProductsInnerBrandInfo struct {
	// 品牌ID
	BrandId *string `json:"brand_id,omitempty"`
	// 品牌logo
	BrandLogo *string `json:"brand_logo,omitempty"`
	// 品牌名称
	BrandName *string `json:"brand_name,omitempty"`
	// 品牌默认落地页
	BrandUrl *string `json:"brand_url,omitempty"`
	// 品牌安卓落地页
	BrandUrlAndroidApp *string `json:"brand_url_android_app,omitempty"`
	// 品牌iOS落地页
	BrandUrlIosApp *string `json:"brand_url_ios_app,omitempty"`
	// 品牌H5落地页
	BrandUrlMobile *string `json:"brand_url_mobile,omitempty"`
	// 品牌落地页universal link
	BrandUrlUniversalLink *string `json:"brand_url_universal_link,omitempty"`
	//
	EnBrand *string `json:"en_brand,omitempty"`
}
