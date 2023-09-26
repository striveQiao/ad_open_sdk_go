/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30Request struct for PromotionCreateV30Request
type PromotionCreateV30Request struct {
	AdDownloadStatus *PromotionCreateV30AdDownloadStatus `json:"ad_download_status,omitempty"`
	//
	AdvertiserId      int64                                `json:"advertiser_id"`
	AutoExtendTraffic *PromotionCreateV30AutoExtendTraffic `json:"auto_extend_traffic,omitempty"`
	//
	Bid       *float64                            `json:"bid,omitempty"`
	BrandInfo *PromotionCreateV30RequestBrandInfo `json:"brand_info,omitempty"`
	//
	Budget     *float64                      `json:"budget,omitempty"`
	BudgetMode *PromotionCreateV30BudgetMode `json:"budget_mode,omitempty"`
	//
	ConfigId *int64 `json:"config_id,omitempty"`
	//
	CpaBid                     *float64                                      `json:"cpa_bid,omitempty"`
	CreativeAutoGenerateSwitch *PromotionCreateV30CreativeAutoGenerateSwitch `json:"creative_auto_generate_switch,omitempty"`
	//
	DeepCpabid       *float64                            `json:"deep_cpabid,omitempty"`
	IsCommentDisable *PromotionCreateV30IsCommentDisable `json:"is_comment_disable,omitempty"`
	//
	Keywords      []*PromotionCreateV30RequestKeywordsInner `json:"keywords,omitempty"`
	MaterialsType *PromotionCreateV30MaterialsType          `json:"materials_type,omitempty"`
	//
	Name          string                                  `json:"name"`
	NativeSetting *PromotionCreateV30RequestNativeSetting `json:"native_setting,omitempty"`
	Operation     *PromotionCreateV30Operation            `json:"operation,omitempty"`
	//
	ProjectId          int64                                       `json:"project_id"`
	PromotionMaterials PromotionCreateV30RequestPromotionMaterials `json:"promotion_materials"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	Source *string `json:"source,omitempty"`
	//
	UnionBidRatio *float64 `json:"union_bid_ratio,omitempty"`
}
