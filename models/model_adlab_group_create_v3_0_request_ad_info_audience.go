/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestAdInfoAudience 定向信息
type AdlabGroupCreateV30RequestAdInfoAudience struct {
	// 年龄
	Age []*[]int64 `json:"age,omitempty"`
	// 定向包ID
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	// 是否开启智能放量
	AutoExtendEnabled *int64 `json:"auto_extend_enabled,omitempty"`
	// 可放开定向，该字段不为空时表示开启智能放量，传空时表示不开启智能放量
	AutoExtendTargets []*AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	// 地域定向省市或者区县列表，当district=CITY时必填
	City                  []int64                                                 `json:"city,omitempty"`
	ConvertedTimeDuration *AdlabGroupCreateV30AdInfoAudienceConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	District              *AdlabGroupCreateV30AdInfoAudienceDistrict              `json:"district,omitempty"`
	// 排除定向逻辑，当广告位只选择穿山甲时可选 定向逻辑和排除定向逻辑只能选其一
	ExcludeFlowPackage []int64 `json:"exclude_flow_package,omitempty"`
	// 定向逻辑，当广告位只选择穿山甲时可选 定向逻辑和排除定向逻辑只能选其一
	FlowPackage []int64                                  `json:"flow_package,omitempty"`
	Gender      *AdlabGroupCreateV30AdInfoAudienceGender `json:"gender,omitempty"`
	// 商圈定向，当district选择\"BUSINESS_DISTRICT\"时必填，最多允许选择100个
	Geolocation     []*AdlabGroupCreateV30RequestAdInfoAudienceGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted *AdlabGroupCreateV30AdInfoAudienceHideIfConverted           `json:"hide_if_converted,omitempty"`
	LocationType    *AdlabGroupCreateV30AdInfoAudienceLocationType              `json:"location_type,omitempty"`
	// 平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载链接的时候，平台可多选。
	Platform []*AdlabGroupCreateV30AdInfoAudiencePlatform `json:"platform,omitempty"`
	// 行政区域版本号
	RegionVersion *string `json:"region_version,omitempty"`
	// 排除人群包列表（自定义人群） 内容为人群包id
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	// 定向人群包列表（自定义人群） 内容为人群包id
	RetargetingTagsInclude []int64                                                  `json:"retargeting_tags_include,omitempty"`
	SuperiorPopularityType *AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
