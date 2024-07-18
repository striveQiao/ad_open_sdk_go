/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInnerAppEcommerceAnchor
type NativeAnchorGetDetailV30ResponseDataListInnerAppEcommerceAnchor struct {
	//
	ActivitiyInfo *string `json:"activitiy_info,omitempty"`
	//
	AndroidAnchorTitle *string `json:"android_anchor_title,omitempty"`
	//
	AndroidPkgName *string                                                                 `json:"android_pkg_name,omitempty"`
	AppIcon        *NativeAnchorGetDetailV30ResponseDataListInnerAppEcommerceAnchorAppIcon `json:"app_icon,omitempty"`
	//
	AppName *string `json:"app_name,omitempty"`
	//
	AppOpenUrl *string `json:"app_open_url,omitempty"`
	//
	DetailInfo *string `json:"detail_info,omitempty"`
	//
	DownloadGuideText *string `json:"download_guide_text,omitempty"`
	// 1-橙子建站 2-第三方落地页
	ExternalType *int64 `json:"external_type,omitempty"`
	//
	IosAnchorTitle          *string                                                                                 `json:"ios_anchor_title,omitempty"`
	OfficialActiBannerImage *NativeAnchorGetDetailV30ResponseDataListInnerAppEcommerceAnchorOfficialActiBannerImage `json:"official_acti_banner_image,omitempty"`
	//
	OfficialActiText *string                                                                        `json:"official_acti_text,omitempty"`
	OrangeSiteInfo   *NativeAnchorGetDetailV30ResponseDataListInnerAppEcommerceAnchorOrangeSiteInfo `json:"orange_site_info,omitempty"`
	ProductImage     *NativeAnchorGetDetailV30ResponseDataListInnerAppEcommerceAnchorProductImage   `json:"product_image,omitempty"`
	//
	ProductName *string `json:"product_name,omitempty"`
	//
	ProductPrice *float64 `json:"product_price,omitempty"`
	//
	ProductTags []string `json:"product_tags,omitempty"`
	//
	ServiceInfo   *string                                                                       `json:"service_info,omitempty"`
	ThirdSiteInfo *NativeAnchorGetDetailV30ResponseDataListInnerAppEcommerceAnchorThirdSiteInfo `json:"third_site_info,omitempty"`
}
