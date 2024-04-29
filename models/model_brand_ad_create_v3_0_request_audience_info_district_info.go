/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdCreateV30RequestAudienceInfoDistrictInfo 地域定向
type BrandAdCreateV30RequestAudienceInfoDistrictInfo struct {
	// 选择城市
	City           []int64                                                 `json:"city,omitempty"`
	CitySelectType *BrandAdCreateV30AudienceInfoDistrictInfoCitySelectType `json:"city_select_type,omitempty"`
	DistrictType   *BrandAdCreateV30AudienceInfoDistrictInfoDistrictType   `json:"district_type,omitempty"`
	LocationType   *BrandAdCreateV30AudienceInfoDistrictInfoLocationType   `json:"location_type,omitempty"`
}
