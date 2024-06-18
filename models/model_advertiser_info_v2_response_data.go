/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserInfoV2ResponseData struct for AdvertiserInfoV2ResponseData
type AdvertiserInfoV2ResponseData struct {
	//
	Address *string `json:"address,omitempty"`
	//
	Brand *string `json:"brand,omitempty"`
	//
	Company *string `json:"company,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	FirstIndustryName *string `json:"first_industry_name,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Industry *string `json:"industry,omitempty"`
	//
	LicenseCity *string `json:"license_city,omitempty"`
	//
	LicenseNo *string `json:"license_no,omitempty"`
	//
	LicenseProvince *string `json:"license_province,omitempty"`
	//
	LicenseUrl *string `json:"license_url,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Note *string `json:"note,omitempty"`
	//
	PromotionArea *string `json:"promotion_area,omitempty"`
	//
	PromotionCenterCity *string `json:"promotion_center_city,omitempty"`
	//
	PromotionCenterProvince *string `json:"promotion_center_province,omitempty"`
	//
	Reason *string                   `json:"reason,omitempty"`
	Role   *AdvertiserInfoV2DataRole `json:"role,omitempty"`
	//
	SecondIndustryName *string                     `json:"second_industry_name,omitempty"`
	Status             *AdvertiserInfoV2DataStatus `json:"status,omitempty"`
}
