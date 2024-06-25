/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandQueryStockV30AudienceInfoDistrictInfo 地域定向
type BrandQueryStockV30AudienceInfoDistrictInfo struct {
	// 选择的城市列表
	City           []int64                                                   `json:"city,omitempty"`
	CitySelectType *BrandQueryStockV30AudienceInfoDistrictInfoCitySelectType `json:"city_select_type,omitempty"`
	DistrictType   *BrandQueryStockV30AudienceInfoDistrictInfoDistrictType   `json:"district_type,omitempty"`
	LocationType   *BrandQueryStockV30AudienceInfoDistrictInfoLocationType   `json:"location_type,omitempty"`
}
