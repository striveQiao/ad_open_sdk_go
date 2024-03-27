/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2ResponseData
type AdGetV2ResponseData struct {
	//
	ActionTrackUrl []string `json:"action_track_url,omitempty"`
	//
	AdCreateTime *string `json:"ad_create_time,omitempty"`
	//
	AdGroupId *int64 `json:"ad_group_id,omitempty"`
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdModifyTime *string `json:"ad_modify_time,omitempty"`
	//
	AdjustCpa            *int64                           `json:"adjust_cpa,omitempty"`
	AdvancedCreativeType *AdGetV2DataAdvancedCreativeType `json:"advanced_creative_type,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AdvertiserStoreIds []int64 `json:"advertiser_store_ids,omitempty"`
	//
	AppDesc *string `json:"app_desc,omitempty"`
	//
	AppIntroduction *int64 `json:"app_introduction,omitempty"`
	//
	AppThumbnails []string            `json:"app_thumbnails,omitempty"`
	AppType       *AdGetV2DataAppType `json:"app_type,omitempty"`
	//
	AssetId *int64 `json:"asset_id,omitempty"`
	//
	AssetIds       []int64                      `json:"asset_ids,omitempty"`
	Audience       *AdGetV2ResponseDataAudience `json:"audience,omitempty"`
	AudienceExtend *AdGetV2DataAudienceExtend   `json:"audience_extend,omitempty"`
	//
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	//
	AuditRejectReason *string                       `json:"audit_reject_reason,omitempty"`
	AutoInheritSwitch *AdGetV2DataAutoInheritSwitch `json:"auto_inherit_switch,omitempty"`
	AutoUpdateKeyword *AdGetV2DataAutoUpdateKeyword `json:"auto_update_keyword,omitempty"`
	//
	AwemeAccount *string `json:"aweme_account,omitempty"`
	//
	Bid *float64 `json:"bid,omitempty"`
	//
	Budget     float64                `json:"budget"`
	BudgetMode *AdGetV2DataBudgetMode `json:"budget_mode,omitempty"`
	//
	CampaignId   *int64                   `json:"campaign_id,omitempty"`
	CategoryType *AdGetV2DataCategoryType `json:"category_type,omitempty"`
	//
	ConvertId             *int64                            `json:"convert_id,omitempty"`
	ConvertedTimeDuration *AdGetV2DataConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	//
	CpaBid      *float64                `json:"cpa_bid,omitempty"`
	DeepBidType *AdGetV2DataDeepBidType `json:"deep_bid_type,omitempty"`
	//
	DeepCpabid         *float64                       `json:"deep_cpabid,omitempty"`
	DeepExternalAction *AdGetV2DataDeepExternalAction `json:"deep_external_action,omitempty"`
	DeliveryPhase      *AdGetV2DataDeliveryPhase      `json:"delivery_phase,omitempty"`
	DeliveryRange      *AdGetV2DataDeliveryRange      `json:"delivery_range,omitempty"`
	DownloadMode       *AdGetV2DataDownloadMode       `json:"download_mode,omitempty"`
	DownloadType       *AdGetV2DataDownloadType       `json:"download_type,omitempty"`
	//
	DownloadUrl *string               `json:"download_url,omitempty"`
	DpaAdtype   *AdGetV2DataDpaAdtype `json:"dpa_adtype,omitempty"`
	//
	DpaCategories []int64 `json:"dpa_categories,omitempty"`
	//
	DpaCity *int64 `json:"dpa_city,omitempty"`
	//
	DpaExternalUrlField *string `json:"dpa_external_url_field,omitempty"`
	//
	DpaExternalUrls []string `json:"dpa_external_urls,omitempty"`
	//
	DpaLbs *int64 `json:"dpa_lbs,omitempty"`
	//
	DpaOpenUrlField *string                    `json:"dpa_open_url_field,omitempty"`
	DpaOpenUrlType  *AdGetV2DataDpaOpenUrlType `json:"dpa_open_url_type,omitempty"`
	//
	DpaOpenUrls []string `json:"dpa_open_urls,omitempty"`
	//
	DpaProductTarget []*AdGetV2ResponseDataDpaProductTargetInner `json:"dpa_product_target,omitempty"`
	//
	DpaProducts []int64 `json:"dpa_products,omitempty"`
	//
	DpaProvince *int64 `json:"dpa_province,omitempty"`
	//
	EndTime        *string                    `json:"end_time,omitempty"`
	ExternalAction *AdGetV2DataExternalAction `json:"external_action,omitempty"`
	//
	ExternalActions []*AdGetV2DataExternalActions `json:"external_actions,omitempty"`
	//
	ExternalUrl *string `json:"external_url,omitempty"`
	//
	ExternalUrlParams  *string                        `json:"external_url_params,omitempty"`
	FeedDeliverySearch *AdGetV2DataFeedDeliverySearch `json:"feed_delivery_search,omitempty"`
	FlowControlMode    *AdGetV2DataFlowControlMode    `json:"flow_control_mode,omitempty"`
	//
	FormId *int64 `json:"form_id,omitempty"`
	//
	FormIndex *int64 `json:"form_index,omitempty"`
	//
	GamePackageBatchId *int64 `json:"game_package_batch_id,omitempty"`
	//
	GamePackageDesc *string `json:"game_package_desc,omitempty"`
	//
	GamePackageThumbnailIds []string                    `json:"game_package_thumbnail_ids,omitempty"`
	HideIfConverted         *AdGetV2DataHideIfConverted `json:"hide_if_converted,omitempty"`
	//
	HideIfExists *int64 `json:"hide_if_exists,omitempty"`
	//
	Id          *int64                  `json:"id,omitempty"`
	InheritType *AdGetV2DataInheritType `json:"inherit_type,omitempty"`
	//
	InheritedAdvertiserId []int64                           `json:"inherited_advertiser_id,omitempty"`
	IntelligentFlowSwitch *AdGetV2DataIntelligentFlowSwitch `json:"intelligent_flow_switch,omitempty"`
	InventoryCatalog      *AdGetV2DataInventoryCatalog      `json:"inventory_catalog,omitempty"`
	//
	InventoryType []*AdGetV2DataInventoryType `json:"inventory_type,omitempty"`
	//
	LandingPageStayTime *int64                       `json:"landing_page_stay_time,omitempty"`
	LaunchTargetType    *AdGetV2DataLaunchTargetType `json:"launch_target_type,omitempty"`
	//
	LubanRoiGoal *float64 `json:"luban_roi_goal,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	//
	OpenUrlParams *string               `json:"open_url_params,omitempty"`
	OptStatus     *AdGetV2DataOptStatus `json:"opt_status,omitempty"`
	//
	Package    *string                `json:"package,omitempty"`
	ParamsType *AdGetV2DataParamsType `json:"params_type,omitempty"`
	Pricing    *AdGetV2DataPricing    `json:"pricing,omitempty"`
	//
	ProductId *string `json:"product_id,omitempty"`
	//
	ProductPlatformId *int64                    `json:"product_platform_id,omitempty"`
	PromotionType     *AdGetV2DataPromotionType `json:"promotion_type,omitempty"`
	//
	QuickAppId *int64 `json:"quick_app_id,omitempty"`
	//
	QuickAppUrl *string `json:"quick_app_url,omitempty"`
	//
	RetentionBid *float64 `json:"retention_bid,omitempty"`
	//
	RoiGoal        *float64                   `json:"roi_goal,omitempty"`
	SceneInventory *AdGetV2DataSceneInventory `json:"scene_inventory,omitempty"`
	//
	ScheduleTime *string                  `json:"schedule_time,omitempty"`
	ScheduleType *AdGetV2DataScheduleType `json:"schedule_type,omitempty"`
	//
	SearchBidRatio *float64                     `json:"search_bid_ratio,omitempty"`
	ShopInfo       *AdGetV2ResponseDataShopInfo `json:"shop_info,omitempty"`
	SmartBidType   *AdGetV2DataSmartBidType     `json:"smart_bid_type,omitempty"`
	SmartInventory *AdGetV2DataSmartInventory   `json:"smart_inventory,omitempty"`
	//
	StartTime *string               `json:"start_time,omitempty"`
	Status    *AdGetV2DataStatus    `json:"status,omitempty"`
	StoreType *AdGetV2DataStoreType `json:"store_type,omitempty"`
	//
	StoreproPackId *int64                   `json:"storepro_pack_id,omitempty"`
	StoreproUnit   *AdGetV2DataStoreproUnit `json:"storepro_unit,omitempty"`
	//
	SubscribeUrl *string `json:"subscribe_url,omitempty"`
	//
	TargetCvr *int64 `json:"target_cvr,omitempty"`
	//
	TrackUrl []string `json:"track_url,omitempty"`
	//
	TrackUrlGroupId *int64 `json:"track_url_group_id,omitempty"`
	//
	TrackUrlGroupType *string                      `json:"track_url_group_type,omitempty"`
	TrackUrlSendType  *AdGetV2DataTrackUrlSendType `json:"track_url_send_type,omitempty"`
	UdShop            *AdGetV2ResponseDataUdShop   `json:"ud_shop,omitempty"`
	//
	Ulink          *string                    `json:"ulink,omitempty"`
	UnionVideoType *AdGetV2DataUnionVideoType `json:"union_video_type,omitempty"`
	//
	ValueOptimizedOpen *int64 `json:"value_optimized_open,omitempty"`
	//
	ValueOptimizedType *int64 `json:"value_optimized_type,omitempty"`
	//
	VideoPlayDoneTrackUrl []string `json:"video_play_done_track_url,omitempty"`
	//
	VideoPlayEffectiveTrackUrl []string `json:"video_play_effective_track_url,omitempty"`
	//
	VideoPlayTrackUrl []string `json:"video_play_track_url,omitempty"`
}
