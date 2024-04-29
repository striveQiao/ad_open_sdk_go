/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseData
type QianchuanAdDetailGetV10ResponseData struct {
	//
	AdCreateTime *string `json:"ad_create_time,omitempty"`
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdKeywords []string `json:"ad_keywords,omitempty"`
	//
	AdModifyTime *string                                      `json:"ad_modify_time,omitempty"`
	Audience     *QianchuanAdDetailGetV10ResponseDataAudience `json:"audience,omitempty"`
	//
	AwemeInfo []*QianchuanAdDetailGetV10ResponseDataAwemeInfoInner `json:"aweme_info,omitempty"`
	BrandInfo *QianchuanAdDetailGetV10ResponseDataBrandInfo        `json:"brand_info,omitempty"`
	//
	CampaignId           *int64                                           `json:"campaign_id,omitempty"`
	CampaignScene        *QianchuanAdDetailGetV10DataCampaignScene        `json:"campaign_scene,omitempty"`
	CreativeAutoGenerate *QianchuanAdDetailGetV10DataCreativeAutoGenerate `json:"creative_auto_generate,omitempty"`
	//
	CreativeList         []*QianchuanAdDetailGetV10ResponseDataCreativeListInner `json:"creative_list,omitempty"`
	CreativeMaterialMode *QianchuanAdDetailGetV10DataCreativeMaterialMode        `json:"creative_material_mode,omitempty"`
	DeliverySetting      *QianchuanAdDetailGetV10ResponseDataDeliverySetting     `json:"delivery_setting,omitempty"`
	DynamicCreative      *QianchuanAdDetailGetV10DataDynamicCreative             `json:"dynamic_creative,omitempty"`
	//
	FirstIndustryId *int64                                     `json:"first_industry_id,omitempty"`
	IsHomepageHide  *QianchuanAdDetailGetV10DataIsHomepageHide `json:"is_homepage_hide,omitempty"`
	IsIntelligent   *QianchuanAdDetailGetV10DataIsIntelligent  `json:"is_intelligent,omitempty"`
	//
	Keywords       []*QianchuanAdDetailGetV10ResponseDataKeywordsInner `json:"keywords,omitempty"`
	LabAdType      *QianchuanAdDetailGetV10DataLabAdType               `json:"lab_ad_type,omitempty"`
	MarketingGoal  *QianchuanAdDetailGetV10DataMarketingGoal           `json:"marketing_goal,omitempty"`
	MarketingScene *QianchuanAdDetailGetV10DataMarketingScene          `json:"marketing_scene,omitempty"`
	//
	MultiProductCreativeList []*QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInner `json:"multi_product_creative_list,omitempty"`
	//
	Name          *string                                           `json:"name,omitempty"`
	OptStatus     *QianchuanAdDetailGetV10DataOptStatus             `json:"opt_status,omitempty"`
	PivativeWords *QianchuanAdDetailGetV10ResponseDataPivativeWords `json:"pivative_words,omitempty"`
	//
	ProductInfo              []*QianchuanAdDetailGetV10ResponseDataProductInfoInner       `json:"product_info,omitempty"`
	ProgrammaticCreativeCard *QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeCard `json:"programmatic_creative_card,omitempty"`
	//
	ProgrammaticCreativeMediaList []*QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeMediaListInner `json:"programmatic_creative_media_list,omitempty"`
	//
	ProgrammaticCreativeTitleList []*QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeTitleListInner `json:"programmatic_creative_title_list,omitempty"`
	//
	RoomInfo []*QianchuanAdDetailGetV10ResponseDataRoomInfoInner `json:"room_info,omitempty"`
	//
	SecondIndustryId *int64                                       `json:"second_industry_id,omitempty"`
	ShopInfo         *QianchuanAdDetailGetV10ResponseDataShopInfo `json:"shop_info,omitempty"`
	Status           *QianchuanAdDetailGetV10DataStatus           `json:"status,omitempty"`
	//
	ThirdIndustryId *int64                                       `json:"third_industry_id,omitempty"`
	TrackUrl        *QianchuanAdDetailGetV10ResponseDataTrackUrl `json:"track_url,omitempty"`
}
