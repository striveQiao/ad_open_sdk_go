/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2Request struct for AudiencePackageCreateV2Request
type AudiencePackageCreateV2Request struct {
	//
	Ac []*AudiencePackageCreateV2Ac `json:"ac,omitempty"`
	//
	ActionCategories []int64                            `json:"action_categories,omitempty"`
	ActionDays       *AudiencePackageCreateV2ActionDays `json:"action_days,omitempty"`
	//
	ActionWords []int64                        `json:"action_words,omitempty"`
	AdType      *AudiencePackageCreateV2AdType `json:"ad_type,omitempty"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Age        []*AudiencePackageCreateV2Age      `json:"age,omitempty"`
	AndroidOsv *AudiencePackageCreateV2AndroidOsv `json:"android_osv,omitempty"`
	//
	AutoExtendTargets []*AudiencePackageCreateV2AutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors []*AudiencePackageCreateV2AwemeFanBehaviors `json:"aweme_fan_behaviors,omitempty"`
	//
	AwemeFanCategories []int64                                   `json:"aweme_fan_categories,omitempty"`
	AwemeFanTimeScope  *AudiencePackageCreateV2AwemeFanTimeScope `json:"aweme_fan_time_scope,omitempty"`
	//
	AwemeFansNumbers []int64 `json:"aweme_fans_numbers,omitempty"`
	//
	BusinessIds []int64 `json:"business_ids,omitempty"`
	//
	Carrier []*AudiencePackageCreateV2Carrier `json:"carrier,omitempty"`
	//
	City                  []int64                                       `json:"city,omitempty"`
	ConvertedTimeDuration *AudiencePackageCreateV2ConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	DeliveryRange         *AudiencePackageCreateV2DeliveryRange         `json:"delivery_range,omitempty"`
	//
	Description string `json:"description"`
	//
	DeviceBrand []*AudiencePackageCreateV2DeviceBrand `json:"device_brand,omitempty"`
	//
	DeviceType []*AudiencePackageCreateV2DeviceType `json:"device_type,omitempty"`
	District   *AudiencePackageCreateV2District     `json:"district,omitempty"`
	//
	ExcludeFlowPackage        []int64                                           `json:"exclude_flow_package,omitempty"`
	FilterAwemeAbnormalActive *AudiencePackageCreateV2FilterAwemeAbnormalActive `json:"filter_aweme_abnormal_active,omitempty"`
	//
	FilterAwemeFansCount *int64                                     `json:"filter_aweme_fans_count,omitempty"`
	FilterOwnAwemeFans   *AudiencePackageCreateV2FilterOwnAwemeFans `json:"filter_own_aweme_fans,omitempty"`
	//
	FlowPackage []int64                        `json:"flow_package,omitempty"`
	Gender      *AudiencePackageCreateV2Gender `json:"gender,omitempty"`
	//
	Geolocation        []*AudiencePackageCreateV2RequestGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted    *AudiencePackageCreateV2HideIfConverted           `json:"hide_if_converted,omitempty"`
	HideIfExists       *AudiencePackageCreateV2HideIfExists              `json:"hide_if_exists,omitempty"`
	InterestActionMode *AudiencePackageCreateV2InterestActionMode        `json:"interest_action_mode,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64                            `json:"interest_words,omitempty"`
	IosOsv        *AudiencePackageCreateV2IosOsv     `json:"ios_osv,omitempty"`
	LandingType   AudiencePackageCreateV2LandingType `json:"landing_type"`
	//
	LaunchPrice   []int64                               `json:"launch_price,omitempty"`
	LocationType  *AudiencePackageCreateV2LocationType  `json:"location_type,omitempty"`
	MarketingGoal *AudiencePackageCreateV2MarketingGoal `json:"marketing_goal,omitempty"`
	//
	Name string `json:"name"`
	//
	Platform []*AudiencePackageCreateV2Platform `json:"platform,omitempty"`
	//
	RegionVersion *string `json:"region_version,omitempty"`
	//
	RetargetingTags []int64 `json:"retargeting_tags,omitempty"`
	//
	RetargetingTagsExclude []int64                                        `json:"retargeting_tags_exclude,omitempty"`
	SuperiorPopularityType *AudiencePackageCreateV2SuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
