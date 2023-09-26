/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestAudience
type ProjectCreateV30RequestAudience struct {
	//
	Ac []*ProjectCreateV30AudienceAc `json:"ac,omitempty"`
	//
	ActionCategories []int64 `json:"action_categories,omitempty"`
	//
	ActionDays *int64 `json:"action_days,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	Age        []string                            `json:"age,omitempty"`
	AndroidOsv *ProjectCreateV30AudienceAndroidOsv `json:"android_osv,omitempty"`
	//
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	//
	AutoExtendTargets []*ProjectCreateV30AudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors []*ProjectCreateV30AudienceAwemeFanBehaviors `json:"aweme_fan_behaviors,omitempty"`
	//
	AwemeFanCategories []int64                                    `json:"aweme_fan_categories,omitempty"`
	AwemeFanTimeScope  *ProjectCreateV30AudienceAwemeFanTimeScope `json:"aweme_fan_time_scope,omitempty"`
	// 仅通信行业且非仅选穿山甲广告时可用
	Carrier []*ProjectCreateV30AudienceCarrier `json:"carrier,omitempty"`
	//
	City                  []int64                                        `json:"city,omitempty"`
	ConvertedTimeDuration *ProjectCreateV30AudienceConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	//
	DeviceBrand []*ProjectCreateV30AudienceDeviceBrand `json:"device_brand,omitempty"`
	//
	DeviceType          []*ProjectCreateV30AudienceDeviceType        `json:"device_type,omitempty"`
	District            *ProjectCreateV30AudienceDistrict            `json:"district,omitempty"`
	DpaCity             *ProjectCreateV30AudienceDpaCity             `json:"dpa_city,omitempty"`
	DpaLbs              *ProjectCreateV30AudienceDpaLbs              `json:"dpa_lbs,omitempty"`
	DpaRtaRecommendType *ProjectCreateV30AudienceDpaRtaRecommendType `json:"dpa_rta_recommend_type,omitempty"`
	DpaRtaSwitch        *ProjectCreateV30AudienceDpaRtaSwitch        `json:"dpa_rta_switch,omitempty"`
	//
	ExcludeFlowPackage        []int64                                            `json:"exclude_flow_package,omitempty"`
	FilterAwemeAbnormalActive *ProjectCreateV30AudienceFilterAwemeAbnormalActive `json:"filter_aweme_abnormal_active,omitempty"`
	FilterAwemeFansCount      *ProjectCreateV30AudienceFilterAwemeFansCount      `json:"filter_aweme_fans_count,omitempty"`
	FilterOwnAwemeFans        *ProjectCreateV30AudienceFilterOwnAwemeFans        `json:"filter_own_aweme_fans,omitempty"`
	//
	FlowPackage []int64                         `json:"flow_package,omitempty"`
	Gender      *ProjectCreateV30AudienceGender `json:"gender,omitempty"`
	//
	Geolocation        []*ProjectCreateV30RequestAudienceGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted    *ProjectCreateV30AudienceHideIfConverted           `json:"hide_if_converted,omitempty"`
	HideIfExists       *ProjectCreateV30AudienceHideIfExists              `json:"hide_if_exists,omitempty"`
	InterestActionMode *ProjectCreateV30AudienceInterestActionMode        `json:"interest_action_mode,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64                         `json:"interest_words,omitempty"`
	IosOsv        *ProjectCreateV30AudienceIosOsv `json:"ios_osv,omitempty"`
	//
	LaunchPrice  []int64                               `json:"launch_price,omitempty"`
	LocationType *ProjectCreateV30AudienceLocationType `json:"location_type,omitempty"`
	//
	Platform        []*ProjectCreateV30AudiencePlatform      `json:"platform,omitempty"`
	RegionRecommend *ProjectCreateV30AudienceRegionRecommend `json:"region_recommend,omitempty"`
	//
	RegionVersion *string `json:"region_version,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64 `json:"retargeting_tags_include,omitempty"`
	// RTA策略ID， https://open.oceanengine.com/labels/7/docs/1696710749917196
	RtaId                  *int64                                          `json:"rta_id,omitempty"`
	SuperiorPopularityType *ProjectCreateV30AudienceSuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
