/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30ResponseDataListInnerAudience
type ProjectListV30ResponseDataListInnerAudience struct {
	//
	Ac []*ProjectListV30DataListAudienceAc `json:"ac,omitempty"`
	//
	ActionCategories []int64 `json:"action_categories,omitempty"`
	//
	ActionDays *int64 `json:"action_days,omitempty"`
	//
	ActionScene []*ProjectListV30DataListAudienceActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	ActivateType []*ProjectListV30DataListAudienceActivateType `json:"activate_type,omitempty"`
	//
	Age        []string                                  `json:"age,omitempty"`
	AndroidOsv *ProjectListV30DataListAudienceAndroidOsv `json:"android_osv,omitempty"`
	//
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	//
	AutoExtendTargets []*ProjectListV30DataListAudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors []*ProjectListV30DataListAudienceAwemeFanBehaviors `json:"aweme_fan_behaviors,omitempty"`
	//
	AwemeFanCategories []int64                                          `json:"aweme_fan_categories,omitempty"`
	AwemeFanTimeScope  *ProjectListV30DataListAudienceAwemeFanTimeScope `json:"aweme_fan_time_scope,omitempty"`
	//
	Career []*ProjectListV30DataListAudienceCareer `json:"career,omitempty"`
	// 仅通信行业且非仅选穿山甲广告位时可用
	Carrier []*ProjectListV30DataListAudienceCarrier `json:"carrier,omitempty"`
	//
	City                  []int64                                              `json:"city,omitempty"`
	ConvertedTimeDuration *ProjectListV30DataListAudienceConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	//
	DeviceBrand []*ProjectListV30DataListAudienceDeviceBrand `json:"device_brand,omitempty"`
	//
	DeviceType          []*ProjectListV30DataListAudienceDeviceType        `json:"device_type,omitempty"`
	District            *ProjectListV30DataListAudienceDistrict            `json:"district,omitempty"`
	DpaCity             *ProjectListV30DataListAudienceDpaCity             `json:"dpa_city,omitempty"`
	DpaLbs              *ProjectListV30DataListAudienceDpaLbs              `json:"dpa_lbs,omitempty"`
	DpaRtaRecommendType *ProjectListV30DataListAudienceDpaRtaRecommendType `json:"dpa_rta_recommend_type,omitempty"`
	DpaRtaSwitch        *ProjectListV30DataListAudienceDpaRtaSwitch        `json:"dpa_rta_switch,omitempty"`
	//
	ExcludeFlowPackage        []int64                                                  `json:"exclude_flow_package,omitempty"`
	FilterAwemeAbnormalActive *ProjectListV30DataListAudienceFilterAwemeAbnormalActive `json:"filter_aweme_abnormal_active,omitempty"`
	FilterAwemeFansCount      *ProjectListV30DataListAudienceFilterAwemeFansCount      `json:"filter_aweme_fans_count,omitempty"`
	FilterOwnAwemeFans        *ProjectListV30DataListAudienceFilterOwnAwemeFans        `json:"filter_own_aweme_fans,omitempty"`
	//
	FlowPackage []int64                               `json:"flow_package,omitempty"`
	Gender      *ProjectListV30DataListAudienceGender `json:"gender,omitempty"`
	//
	Geolocation        []*ProjectListV30ResponseDataListInnerAudienceGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted    *ProjectListV30DataListAudienceHideIfConverted                 `json:"hide_if_converted,omitempty"`
	HideIfExists       *ProjectListV30DataListAudienceHideIfExists                    `json:"hide_if_exists,omitempty"`
	InterestActionMode *ProjectListV30DataListAudienceInterestActionMode              `json:"interest_action_mode,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64                               `json:"interest_words,omitempty"`
	IosOsv        *ProjectListV30DataListAudienceIosOsv `json:"ios_osv,omitempty"`
	//
	LaunchPrice  []int64                                     `json:"launch_price,omitempty"`
	LocationType *ProjectListV30DataListAudienceLocationType `json:"location_type,omitempty"`
	//
	Platform []*ProjectListV30DataListAudiencePlatform `json:"platform,omitempty"`
	//
	RegionRecommend *string `json:"region_recommend,omitempty"`
	//
	RegionVersion *string `json:"region_version,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                               `json:"retargeting_tags_include,omitempty"`
	SuperiorPopularityType *ProjectListV30DataListAudienceSuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
