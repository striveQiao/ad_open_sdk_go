/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataAdInfoAudience 定向信息
type AdlabGroupDetailV30ResponseDataDataAdInfoAudience struct {
	// 年龄
	Age []*[]int64 `json:"age,omitempty"`
	// 定向包ID
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	// 是否开启智能放量
	AutoExtendEnabled *int64 `json:"auto_extend_enabled,omitempty"`
	// 可放开定向
	AutoExtendTargets []*AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	// 地域定向省市或者区县列表
	City                  []int64                                                         `json:"city,omitempty"`
	ConvertedTimeDuration *AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	District              *AdlabGroupDetailV30DataDataAdInfoAudienceDistrict              `json:"district,omitempty"`
	// 排除定向逻辑
	ExcludeFlowPackage []int64 `json:"exclude_flow_package,omitempty"`
	// 定向逻辑
	FlowPackage []int64                                          `json:"flow_package,omitempty"`
	Gender      *AdlabGroupDetailV30DataDataAdInfoAudienceGender `json:"gender,omitempty"`
	// 商圈定向信息
	Geolocation     []*AdlabGroupDetailV30ResponseDataDataAdInfoAudienceGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted *AdlabGroupDetailV30DataDataAdInfoAudienceHideIfConverted            `json:"hide_if_converted,omitempty"`
	LocationType    *AdlabGroupDetailV30DataDataAdInfoAudienceLocationType               `json:"location_type,omitempty"`
	// 平台
	Platform []*AdlabGroupDetailV30DataDataAdInfoAudiencePlatform `json:"platform,omitempty"`
	// 行政区域版本号
	RegionVersion *string `json:"region_version,omitempty"`
	// 排除人群包列表（自定义人群）
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	// 定向人群包列表（自定义人群）
	RetargetingTagsInclude []int64                                                          `json:"retargeting_tags_include,omitempty"`
	SuperiorPopularityType *AdlabGroupDetailV30DataDataAdInfoAudienceSuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
