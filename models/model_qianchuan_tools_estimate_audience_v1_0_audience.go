/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsEstimateAudienceV10Audience
type QianchuanToolsEstimateAudienceV10Audience struct {
	//
	Ac []*QianchuanToolsEstimateAudienceV10AudienceAc `json:"ac,omitempty"`
	//
	ActionCategories []int64                                              `json:"action_categories,omitempty"`
	ActionDays       *QianchuanToolsEstimateAudienceV10AudienceActionDays `json:"action_days,omitempty"`
	//
	ActionScene []*QianchuanToolsEstimateAudienceV10AudienceActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
	//
	Age               []*QianchuanToolsEstimateAudienceV10AudienceAge             `json:"age,omitempty"`
	AudienceMode      *QianchuanToolsEstimateAudienceV10AudienceAudienceMode      `json:"audience_mode,omitempty"`
	AutoExtendEnabled *QianchuanToolsEstimateAudienceV10AudienceAutoExtendEnabled `json:"auto_extend_enabled,omitempty"`
	//
	AutoExtendTargets []*QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors     []*QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors   `json:"aweme_fan_behaviors,omitempty"`
	AwemeFanBehaviorsDays *QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays `json:"aweme_fan_behaviors_days,omitempty"`
	//
	AwemeFanCategories []int64 `json:"aweme_fan_categories,omitempty"`
	//
	City     []int64                                            `json:"city,omitempty"`
	District *QianchuanToolsEstimateAudienceV10AudienceDistrict `json:"district,omitempty"`
	//
	DistrictType         *bool                                                          `json:"district_type,omitempty"`
	ExcludeLimitedRegion *QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion `json:"exclude_limited_region,omitempty"`
	Gender               *QianchuanToolsEstimateAudienceV10AudienceGender               `json:"gender,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestWords []int64 `json:"interest_words,omitempty"`
	//
	LivePlatformTags []*QianchuanToolsEstimateAudienceV10AudienceLivePlatformTags `json:"live_platform_tags,omitempty"`
	LocationType     *QianchuanToolsEstimateAudienceV10AudienceLocationType       `json:"location_type,omitempty"`
	NewCustomer      *QianchuanToolsEstimateAudienceV10AudienceNewCustomer        `json:"new_customer,omitempty"`
	//
	OrientationId *int64 `json:"orientation_id,omitempty"`
	//
	Platform []*QianchuanToolsEstimateAudienceV10AudiencePlatform `json:"platform,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                                       `json:"retargeting_tags_include,omitempty"`
	SmartInterestAction    *QianchuanToolsEstimateAudienceV10AudienceSmartInterestAction `json:"smart_interest_action,omitempty"`
}
