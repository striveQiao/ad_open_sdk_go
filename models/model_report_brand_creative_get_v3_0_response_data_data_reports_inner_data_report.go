/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandCreativeGetV30ResponseDataDataReportsInnerDataReport
type ReportBrandCreativeGetV30ResponseDataDataReportsInnerDataReport struct {
	// 活动banner图点击数
	ActivityBannerClickCount *string `json:"activity_banner_click_count,omitempty"`
	// 附加创意卡片点击数
	AttachCreativeClickCnt *string `json:"attach_creative_click_cnt,omitempty"`
	// 附加创意卡片展示数
	AttachCreativeShowCnt *string `json:"attach_creative_show_cnt,omitempty"`
	// 平均单次播放时长
	AveragePlayTimePerPlay *string `json:"average_play_time_per_play,omitempty"`
	// 支付订单平均每单交易额
	BrandAllOrderPayAverageAmount7days *string `json:"brand_all_order_pay_average_amount_7days,omitempty"`
	// 支付订单数目7天
	BrandAllOrderPayCount7days *string `json:"brand_all_order_pay_count_7days,omitempty"`
	// 支付订单gvm7天
	BrandAllOrderPayGmv7days *string `json:"brand_all_order_pay_gmv_7days,omitempty"`
	// 支付订单数7天
	BrandIndirectOrderPayCount7days *string `json:"brand_indirect_order_pay_count_7days,omitempty"`
	// 支付订单gvm7天
	BrandIndirectOrderPayGmv7days *string `json:"brand_indirect_order_pay_gmv_7days,omitempty"`
	// 3秒卡片展现数
	CardShow *string `json:"card_show,omitempty"`
	// 附加创意电话按钮点击
	ClickCallCnt *string `json:"click_call_cnt,omitempty"`
	// 主页内电话拨打点击量
	ClickCallDy *string `json:"click_call_dy,omitempty"`
	// 点击数
	ClickCnt *string `json:"click_cnt,omitempty"`
	// 附件创意在线咨询点击
	ClickCounsel *string `json:"click_counsel,omitempty"`
	// 主页下载链接点击量
	ClickDownload *string `json:"click_download,omitempty"`
	// 推广页访问量
	ClickLandingPage *string `json:"click_landing_page,omitempty"`
	// 主页商品橱窗访问量
	ClickShopwindow *string `json:"click_shopwindow,omitempty"`
	// 主页内落地页访问量
	ClickWebsite *string `json:"click_website,omitempty"`
	// 转化成本
	ConversionCost *string `json:"conversion_cost,omitempty"`
	// 转化率
	ConversionRate *string `json:"conversion_rate,omitempty"`
	// 转化数
	ConvertCnt *string `json:"convert_cnt,omitempty"`
	// 附件创意卡券领取
	CouponAddition *string `json:"coupon_addition,omitempty"`
	// 平均点击单价
	CpcPlatform *string `json:"cpc_platform,omitempty"`
	// 平均千次展现费用
	CpmPlatform *string `json:"cpm_platform,omitempty"`
	// 点击率
	Ctr *string `json:"ctr,omitempty"`
	// 深度转化数
	DeepConvertCnt *string `json:"deep_convert_cnt,omitempty"`
	// 深度转化成本
	DeepConvertCost *string `json:"deep_convert_cost,omitempty"`
	// 深度转化率
	DeepConvertRate *string `json:"deep_convert_rate,omitempty"`
	// 评论量
	DyComment *string `json:"dy_comment,omitempty"`
	// 粉丝量
	DyFollow *string `json:"dy_follow,omitempty"`
	// 主页访问量
	DyHomeVisited *string `json:"dy_home_visited,omitempty"`
	// 点赞数
	DyLike *string `json:"dy_like,omitempty"`
	// 分享量
	DyShare *string `json:"dy_share,omitempty"`
	// 粉丝视频评论数
	FansDyCommentCount *string `json:"fans_dy_comment_count,omitempty"`
	// 粉丝视频点赞数
	FansDyLikeCount *string `json:"fans_dy_like_count,omitempty"`
	// 粉丝主页浏览数
	FansHomeVisitedCount *string `json:"fans_home_visited_count,omitempty"`
	// 粉丝直播评论次数
	FansLubanLiveCommentCount *string `json:"fans_luban_live_comment_count,omitempty"`
	// 粉丝直播观看次数
	FansLubanLiveEnterCount *string `json:"fans_luban_live_enter_count,omitempty"`
	// 粉丝直播间礼物总金额
	FansLubanLiveGiftAmount *string `json:"fans_luban_live_gift_amount,omitempty"`
	// 粉丝直播间订单金额
	FansLubanLivePayOrderStatCost *string `json:"fans_luban_live_pay_order_stat_cost,omitempty"`
	// 粉丝直播间购物车点击量
	FansLubanLiveSlidecartClickCount *string `json:"fans_luban_live_slidecart_click_count,omitempty"`
	// 粉丝视频播放数
	FansVideoPlayCount *string `json:"fans_video_play_count,omitempty"`
	// 粉丝视频转发数
	FansVideoShareCount *string `json:"fans_video_share_count,omitempty"`
	// 附件创意表单按钮点击
	FormClickButton *string `json:"form_click_button,omitempty"`
	// 挑战赛查看数
	IesChallengeClick *string `json:"ies_challenge_click,omitempty"`
	// 音乐查看数
	IesMusicClick *string `json:"ies_music_click,omitempty"`
	// 直播预约人数
	LiveAppointCount *string `json:"live_appoint_count,omitempty"`
	// 组件点击成本
	LiveComponentClickCost *string `json:"live_component_click_cost,omitempty"`
	// 组件点击数
	LiveComponentClickCount *string `json:"live_component_click_count,omitempty"`
	// 组件点击率
	LiveComponentClickRate *string `json:"live_component_click_rate,omitempty"`
	// 直播间加入粉丝团
	LiveFansClubJoinCnt *string `json:"live_fans_club_join_cnt,omitempty"`
	// 直播间团购点击量
	LiveGrouponProductClickCount *string `json:"live_groupon_product_click_count,omitempty"`
	// 直播间客单价
	LivePayOrderCostPerOrder *string `json:"live_pay_order_cost_per_order,omitempty"`
	// 直播间客单价（长效影响-15天）
	LivePayOrderCostPerOrderByAuthor15days *string `json:"live_pay_order_cost_per_order_by_author_15days,omitempty"`
	// 直播间客单价（长效影响-30天）
	LivePayOrderCostPerOrderByAuthor30days *string `json:"live_pay_order_cost_per_order_by_author_30days,omitempty"`
	// 直播间客单价（长效影响-3天）
	LivePayOrderCostPerOrderByAuthor3days *string `json:"live_pay_order_cost_per_order_by_author_3days,omitempty"`
	// 直播间客单价（长效影响-7天）
	LivePayOrderCostPerOrderByAuthor7days *string `json:"live_pay_order_cost_per_order_by_author_7days,omitempty"`
	// 直播间超过一分钟观看数
	LiveWatchOneMinuteCount *string `json:"live_watch_one_minute_count,omitempty"`
	// POI点击数
	LocationClick *string `json:"location_click,omitempty"`
	// 直播间点击商品数
	LubanLiveClickProductCnt *string `json:"luban_live_click_product_cnt,omitempty"`
	// 直播间评论数
	LubanLiveCommentCnt *string `json:"luban_live_comment_cnt,omitempty"`
	// 直播间观看数
	LubanLiveEnterCnt *string `json:"luban_live_enter_cnt,omitempty"`
	// 直播间关注数
	LubanLiveFollowCnt *string `json:"luban_live_follow_cnt,omitempty"`
	// 直播间礼物总金额
	LubanLiveGiftAmount *string `json:"luban_live_gift_amount,omitempty"`
	// 直播间打赏次数
	LubanLiveGiftCnt *string `json:"luban_live_gift_cnt,omitempty"`
	// 直播间订单量
	LubanLivePayOrderCount *string `json:"luban_live_pay_order_count,omitempty"`
	// 直播间15日订单量
	LubanLivePayOrderCountByAuthor15days *string `json:"luban_live_pay_order_count_by_author_15days,omitempty"`
	// 直播间30日订单量
	LubanLivePayOrderCountByAuthor30days *string `json:"luban_live_pay_order_count_by_author_30days,omitempty"`
	// 直播间3日订单量
	LubanLivePayOrderCountByAuthor3days *string `json:"luban_live_pay_order_count_by_author_3days,omitempty"`
	// 直播间7日订单量
	LubanLivePayOrderCountByAuthor7days *string `json:"luban_live_pay_order_count_by_author_7days,omitempty"`
	// 直播间15日订单金额
	LubanLivePayOrderStatCostByAuthor15days *string `json:"luban_live_pay_order_stat_cost_by_author_15days,omitempty"`
	// 直播间30日订单金额
	LubanLivePayOrderStatCostByAuthor30days *string `json:"luban_live_pay_order_stat_cost_by_author_30days,omitempty"`
	// 直播间3日订单金额
	LubanLivePayOrderStatCostByAuthor3days *string `json:"luban_live_pay_order_stat_cost_by_author_3days,omitempty"`
	// 直播间7日订单金额
	LubanLivePayOrderStatCostByAuthor7days *string `json:"luban_live_pay_order_stat_cost_by_author_7days,omitempty"`
	// 直播间分享数
	LubanLiveShareCnt *string `json:"luban_live_share_cnt,omitempty"`
	// 直播间查看购物车数
	LubanLiveSlidecartClickCnt *string `json:"luban_live_slidecart_click_cnt,omitempty"`
	// 鲁班订单量
	LubanOrderCnt *string `json:"luban_order_cnt,omitempty"`
	// 鲁班订单ROI
	LubanOrderRoi *string `json:"luban_order_roi,omitempty"`
	// 鲁班订单金额
	LubanOrderStatAmount *string `json:"luban_order_stat_amount,omitempty"`
	// 直播间订单金额
	LubanlubanLivePayOrderStatCost *string `json:"lubanluban_live_pay_order_stat_cost,omitempty"`
	// 主素材点击数
	MainMaterialClickCount *string `json:"main_material_click_count,omitempty"`
	// 私信数
	MessageAction *string `json:"message_action,omitempty"`
	// 外跳模块点击数
	OutsideJumpAreaClickCount *string `json:"outside_jump_area_click_count,omitempty"`
	// 25%进度播放数
	Play25FeedBreak *string `json:"play_25_feed_break,omitempty"`
	// 50%进度播放数
	Play50FeedBreak *string `json:"play_50_feed_break,omitempty"`
	// 75%进度播放数
	Play75FeedBreak *string `json:"play_75_feed_break,omitempty"`
	// 99%进度播放数
	Play99FeedBreak *string `json:"play_99_feed_break,omitempty"`
	// 完播率
	PlayOverRate *string `json:"play_over_rate,omitempty"`
	// 门店浏览次数
	PoiHomepageViewCount *string `json:"poi_homepage_view_count,omitempty"`
	// 品专总点击数
	SearchBrandTotalClickCount *string `json:"search_brand_total_click_count,omitempty"`
	// 品专总点击率
	SearchBrandTotalCtr *string `json:"search_brand_total_ctr,omitempty"`
	// 店铺访问量
	ShopVisitedCount *string `json:"shop_visited_count,omitempty"`
	// 展示数
	ShowCnt *string `json:"show_cnt,omitempty"`
	// 消耗
	StatCost *string `json:"stat_cost,omitempty"`
	// 子链点击数
	SubchainClickCount *string `json:"subchain_click_count,omitempty"`
	// 超级品专话题点击数
	SuperBrandTopicClickCount *string `json:"super_brand_topic_click_count,omitempty"`
	// 超级品专视频卡点击数
	SuperBrandVideocardClickCount *string `json:"super_brand_videocard_click_count,omitempty"`
	// 播放数
	TotalPlay *string `json:"total_play,omitempty"`
	// 有效播放数
	ValidPlay *string `json:"valid_play,omitempty"`
	// 有效播放成本
	ValidPlayCost *string `json:"valid_play_cost,omitempty"`
	// 千次有效播放成本
	ValidPlayCostOfMille *string `json:"valid_play_cost_of_mille,omitempty"`
	// 千次有效播放数
	ValidPlayOfMille *string `json:"valid_play_of_mille,omitempty"`
	// 有效播放率
	ValidPlayRate *string `json:"valid_play_rate,omitempty"`
	// WIFI播放占比
	WifiPlayRate *string `json:"wifi_play_rate,omitempty"`
}
