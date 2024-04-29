/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2OrderField
type ReportAdvertiserGetV2OrderField string

// List of report_advertiser_get_v2_order_field
const (
	QQ_ReportAdvertiserGetV2OrderField                                          ReportAdvertiserGetV2OrderField = "qq"
	CONSULT_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "consult"
	DEEP_CONVERT_COST_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "deep_convert_cost"
	COMMUTE_FIRST_PAY_COUNT_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "commute_first_pay_count"
	LOTTERY_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "lottery"
	GAME_PAY_COST_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "game_pay_cost"
	PLAY_100_FEED_BREAK_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "play_100_feed_break"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_day_acitve_pay_count"
	VOTE_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "vote"
	LOAN_CREDIT_COST_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "loan_credit_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_5days"
	PLAY_DURATION_2S_RATE_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "play_duration_2s_rate"
	SUBMIT_CERTIFICATION_COUNT_ReportAdvertiserGetV2OrderField                  ReportAdvertiserGetV2OrderField = "submit_certification_count"
	AVG_CLICK_COST_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "avg_click_cost"
	CONVERT_RATE_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "convert_rate"
	DEEP_CONVERT_RATE_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "deep_convert_rate"
	LOAN_COMPLETION_COST_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "loan_completion_cost"
	IN_APP_ORDER_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "in_app_order"
	ATTRIBUTION_RETENTION_3D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_3d_rate"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportAdvertiserGetV2OrderField       ReportAdvertiserGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	PLAY_OVER_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "play_over"
	LOAN_CREDIT_RATE_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "loan_credit_rate"
	SHOPPING_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "shopping"
	PRE_CONVERT_RATE_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "pre_convert_rate"
	ATTRIBUTION_RETENTION_7D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_7d_cnt"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_6days"
	WIFI_PLAY_RATE_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "wifi_play_rate"
	DISLIKE_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "dislike"
	ATTRIBUTION_RETENTION_3D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_3d_cnt"
	PLAY_25_FEED_BREAK_RATE_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "play_25_feed_break_rate"
	POI_ADDRESS_CLICK_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "poi_address_click"
	UNION_ROI_3_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "union_roi_3"
	NEXT_DAY_OPEN_COST_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "next_day_open_cost"
	ATTRIBUTION_RETENTION_6D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_6d_rate"
	CTR_ReportAdvertiserGetV2OrderField                                         ReportAdvertiserGetV2OrderField = "ctr"
	PLAY_50_FEED_BREAK_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "play_50_feed_break"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_next_day_open_rate"
	CONVERT_COST_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "convert_cost"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportAdvertiserGetV2OrderField                 ReportAdvertiserGetV2OrderField = "live_watch_one_minute_count"
	PLAY_DURATION_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "play_duration"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportAdvertiserGetV2OrderField         ReportAdvertiserGetV2OrderField = "attribution_active_pay_7d_per_count"
	ATTRIBUTION_RETENTION_4D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_4d_rate"
	PLAY_DURATION_2S_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "play_duration_2s"
	IN_APP_CART_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "in_app_cart"
	COST_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "cost"
	PLAY_100_FEED_BREAK_RATE_ReportAdvertiserGetV2OrderField                    ReportAdvertiserGetV2OrderField = "play_100_feed_break_rate"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportAdvertiserGetV2OrderField             ReportAdvertiserGetV2OrderField = "advanced_creative_counsel_click"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_0d_ltv"
	AVG_RANK_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "avg_rank"
	IN_APP_PAY_ReportAdvertiserGetV2OrderField                                  ReportAdvertiserGetV2OrderField = "in_app_pay"
	PHONE_CONFIRM_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "phone_confirm"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_6days"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_deep_convert_cost"
	STAT_UNION_LTV_3_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "stat_union_ltv_3"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportAdvertiserGetV2OrderField      ReportAdvertiserGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	SHARE_ReportAdvertiserGetV2OrderField                                       ReportAdvertiserGetV2OrderField = "share"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "advanced_creative_form_submit"
	INTERACT_PER_COST_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "interact_per_cost"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_3d_ltv"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_5days"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "luban_live_click_product_cnt"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportAdvertiserGetV2OrderField       ReportAdvertiserGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	GAME_PAY_COUNT_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "game_pay_count"
	AVERAGE_VIDEO_PLAY_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "average_video_play"
	STAT_UNION_LTV_0_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "stat_union_ltv_0"
	PHONE_EFFECTIVE_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "phone_effective"
	ATTRIBUTION_RETENTION_5D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_5d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_2days"
	LUBAN_LIVE_ENTER_CNT_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "luban_live_enter_cnt"
	CLICK_DOWNLOAD_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "click_download"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportAdvertiserGetV2OrderField             ReportAdvertiserGetV2OrderField = "attribution_active_pay_7d_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportAdvertiserGetV2OrderField   ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	CONVERT_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "convert"
	DOWNLOAD_START_COST_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "download_start_cost"
	IN_APP_DETAIL_UV_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "in_app_detail_uv"
	ACTIVE_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "active"
	LUBAN_ORDER_CNT_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "luban_order_cnt"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_7days"
	ATTRIBUTION_RETENTION_2D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_2d_cost"
	CARD_SHOW_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "card_show"
	GAME_ADDICTION_COST_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "game_addiction_cost"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_wechat_pay_30d_amount"
	IES_CHALLENGE_CLICK_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "ies_challenge_click"
	PLAY_OVER_RATE_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "play_over_rate"
	ATTRIBUTION_RETENTION_5D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_5d_rate"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_game_pay_7d_count"
	LUBAN_LIVE_GIFT_AMOUNT_ReportAdvertiserGetV2OrderField                      ReportAdvertiserGetV2OrderField = "luban_live_gift_amount"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_2days"
	VALID_PLAY_ReportAdvertiserGetV2OrderField                                  ReportAdvertiserGetV2OrderField = "valid_play"
	APPROVAL_COUNT_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "approval_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportAdvertiserGetV2OrderField  ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	INSTALL_FINISH_RATE_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "install_finish_rate"
	LOAN_COMPLETION_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "loan_completion"
	MESSAGE_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "message"
	WECHAT_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "wechat"
	WECHAT_FIRST_PAY_RATE_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "wechat_first_pay_rate"
	SHOW_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "show"
	PLAY_DURATION_3S_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "play_duration_3s"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "advanced_creative_phone_click"
	FOLLOW_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "follow"
	CONVERT_SHOW_RATE_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "convert_show_rate"
	ATTRIBUTION_RETENTION_6D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_6d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_8days"
	LUBAN_LIVE_GIFT_CNT_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "luban_live_gift_cnt"
	AVERAGE_PLAY_PROGRESS_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "average_play_progress"
	VALID_PLAY_RATE_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "valid_play_rate"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_7d_roi"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_wechat_login_30d_cost"
	FORM_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "form"
	WECHAT_LOGIN_COST_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "wechat_login_cost"
	CLICK_SHOPWINDOW_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "click_shopwindow"
	LOCATION_CLICK_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "location_click"
	CLICK_CALL_DY_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "click_call_dy"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_3days"
	LOAN_CREDIT_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "loan_credit"
	ACTIVE_PAY_AMOUNT_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "active_pay_amount"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_0d_roi"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_7days"
	INSTALL_FINISH_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "install_finish"
	DOWNLOAD_START_RATE_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "download_start_rate"
	PHONE_CONNECT_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "phone_connect"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_1day"
	LIVE_COMPONENT_CLICK_COUNT_ReportAdvertiserGetV2OrderField                  ReportAdvertiserGetV2OrderField = "live_component_click_count"
	PLAY_50_FEED_BREAK_RATE_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "play_50_feed_break_rate"
	LUBAN_LIVE_SHARE_CNT_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "luban_live_share_cnt"
	ATTRIBUTION_RETENTION_5D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_5d_cnt"
	DOWNLOAD_START_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "download_start"
	WECHAT_LOGIN_COUNT_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "wechat_login_count"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "luban_live_pay_order_stat_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_4days"
	PLAY_75_FEED_BREAK_RATE_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "play_75_feed_break_rate"
	VIEW_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "view"
	ADVANCED_CREATIVE_FORM_CLICK_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "advanced_creative_form_click"
	ATTRIBUTION_RETENTION_2D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_2d_cnt"
	COUPON_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "coupon"
	NEXT_DAY_OPEN_RATE_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "next_day_open_rate"
	VALID_PLAY_COST_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "valid_play_cost"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_game_pay_7d_cost"
	PLAY_75_FEED_BREAK_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "play_75_feed_break"
	MESSAGE_ACTION_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "message_action"
	GAME_ADDICTION_RATE_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "game_addiction_rate"
	DOWNLOAD_FINISH_RATE_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "download_finish_rate"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_retention_7d_sum_cnt"
	ACTIVE_RATE_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "active_rate"
	GAME_ADDICTION_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "game_addiction"
	PAY_AMOUNT_ROI_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "pay_amount_roi"
	HOME_VISITED_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "home_visited"
	ATTRIBUTION_CONVERT_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "attribution_convert"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_3days"
	REGISTER_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "register"
	FIRST_RENTAL_ORDER_COUNT_ReportAdvertiserGetV2OrderField                    ReportAdvertiserGetV2OrderField = "first_rental_order_count"
	STAT_PAY_AMOUNT_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "stat_pay_amount"
	ACTIVE_PAY_RATE_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "active_pay_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_8days"
	CLICK_INSTALL_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "click_install"
	PAY_COUNT_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "pay_count"
	CPA_ReportAdvertiserGetV2OrderField                                         ReportAdvertiserGetV2OrderField = "cpa"
	PRE_CONVERT_COST_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "pre_convert_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportAdvertiserGetV2OrderField    ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	REDIRECT_TO_SHOP_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "redirect_to_shop"
	PRE_LOAN_CREDIT_COST_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "pre_loan_credit_cost"
	WECHAT_FIRST_PAY_COST_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "wechat_first_pay_cost"
	LIVE_FANS_CLUB_JOIN_CNT_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "live_fans_club_join_cnt"
	CLICK_WEBSITE_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "click_website"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_7d_ltv"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_3d_roi"
	ATTRIBUTION_RETENTION_6D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_6d_cnt"
	PHONE_ReportAdvertiserGetV2OrderField                                       ReportAdvertiserGetV2OrderField = "phone"
	LIVE_COMPONENT_CLICK_COST_ReportAdvertiserGetV2OrderField                   ReportAdvertiserGetV2OrderField = "live_component_click_cost"
	ACTIVE_COST_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "active_cost"
	LUBAN_ORDER_STAT_AMOUNT_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "luban_order_stat_amount"
	WIFI_PLAY_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "wifi_play"
	CPM_ReportAdvertiserGetV2OrderField                                         ReportAdvertiserGetV2OrderField = "cpm"
	LIVE_COMPONENT_CLICK_RATE_ReportAdvertiserGetV2OrderField                   ReportAdvertiserGetV2OrderField = "live_component_click_rate"
	LUBAN_LIVE_FOLLOW_CNT_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "luban_live_follow_cnt"
	ATTRIBUTION_CONVERT_COST_ReportAdvertiserGetV2OrderField                    ReportAdvertiserGetV2OrderField = "attribution_convert_cost"
	DOWNLOAD_FINISH_COST_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "download_finish_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportAdvertiserGetV2OrderField   ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_wechat_pay_30d_roi"
	ACTIVE_PAY_COST_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "active_pay_cost"
	POI_COLLECT_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "poi_collect"
	WECHAT_FIRST_PAY_COUNT_ReportAdvertiserGetV2OrderField                      ReportAdvertiserGetV2OrderField = "wechat_first_pay_count"
	PRE_CONVERT_COUNT_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "pre_convert_count"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportAdvertiserGetV2OrderField         ReportAdvertiserGetV2OrderField = "attribution_retention_7d_total_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_4days"
	ATTRIBUTION_RETENTION_3D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_3d_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_1day"
	UNION_ROI_0_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "union_roi_0"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_active_pay_7d_rate"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "advanced_creative_coupon_addition"
	CLICK_LANDING_PAGE_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "click_landing_page"
	NEXT_DAY_OPEN_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "next_day_open"
	LUBAN_ORDER_ROI_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "luban_order_roi"
	COUPON_SINGLE_PAGE_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "coupon_single_page"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportAdvertiserGetV2OrderField                  ReportAdvertiserGetV2OrderField = "luban_live_pay_order_count"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportAdvertiserGetV2OrderField                  ReportAdvertiserGetV2OrderField = "average_play_time_per_play"
	CONSULT_EFFECTIVE_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "consult_effective"
	ACTIVE_REGISTER_RATE_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "active_register_rate"
	ACTIVE_REGISTER_COST_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "active_register_cost"
	IN_APP_UV_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "in_app_uv"
	CUSTOMER_EFFECTIVE_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "customer_effective"
	FIRST_ORDER_COUNT_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "first_order_count"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportAdvertiserGetV2OrderField          ReportAdvertiserGetV2OrderField = "attribution_wechat_login_30d_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportAdvertiserGetV2OrderField ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	TOTAL_PLAY_ReportAdvertiserGetV2OrderField                                  ReportAdvertiserGetV2OrderField = "total_play"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_next_day_open_cnt"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_next_day_open_cost"
	PLAY_DURATION_10S_RATE_ReportAdvertiserGetV2OrderField                      ReportAdvertiserGetV2OrderField = "play_duration_10s_rate"
	DOWNLOAD_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "download"
	INSTALL_FINISH_COST_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "install_finish_cost"
	CPC_ReportAdvertiserGetV2OrderField                                         ReportAdvertiserGetV2OrderField = "cpc"
	ATTRIBUTION_RETENTION_4D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_4d_cost"
	DOWNLOAD_FINISH_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "download_finish"
	ATTRIBUTION_RETENTION_4D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_4d_cnt"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_active_pay_7d_cost"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_day_active_pay_count"
	STAT_UNION_LTV_7_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "stat_union_ltv_7"
	PLAY_DURATION_10S_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "play_duration_10s"
	PLAY_DURATION_3S_RATE_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "play_duration_3s_rate"
	UNION_ROI_7_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "union_roi_7"
	REDIRECT_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "redirect"
	REPORT_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "report"
	MAP_SEARCH_ReportAdvertiserGetV2OrderField                                  ReportAdvertiserGetV2OrderField = "map_search"
	WECHAT_PAY_AMOUNT_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "wechat_pay_amount"
	IES_MUSIC_CLICK_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "ies_music_click"
	PLAY_DURATION_5S_RATE_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "play_duration_5s_rate"
	LUBAN_LIVE_COMMENT_CNT_ReportAdvertiserGetV2OrderField                      ReportAdvertiserGetV2OrderField = "luban_live_comment_cnt"
	PLAY_DURATION_SUM_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "play_duration_sum"
	ATTRIBUTION_RETENTION_2D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_RETENTION_7D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_7d_cost"
	ATTRIBUTION_DEEP_CONVERT_ReportAdvertiserGetV2OrderField                    ReportAdvertiserGetV2OrderField = "attribution_deep_convert"
	BUTTON_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "button"
	LOAN_COMPLETION_RATE_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "loan_completion_rate"
	PRE_LOAN_CREDIT_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "pre_loan_credit"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "luban_live_slidecart_click_cnt"
	CLICK_ReportAdvertiserGetV2OrderField                                       ReportAdvertiserGetV2OrderField = "click"
	PLAY_25_FEED_BREAK_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "play_25_feed_break"
	COMMENT_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "comment"
	DEEP_CONVERT_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "deep_convert"
	AVG_SHOW_COST_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "avg_show_cost"
	LIKE_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "like"
	ATTRIBUTION_RETENTION_7D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_7d_rate"
)

// All allowed values of ReportAdvertiserGetV2OrderField enum
var AllowedReportAdvertiserGetV2OrderFieldEnumValues = []ReportAdvertiserGetV2OrderField{
	"qq",
	"consult",
	"deep_convert_cost",
	"commute_first_pay_count",
	"lottery",
	"game_pay_cost",
	"play_100_feed_break",
	"attribution_day_acitve_pay_count",
	"vote",
	"loan_credit_cost",
	"attribution_game_in_app_roi_5days",
	"play_duration_2s_rate",
	"submit_certification_count",
	"avg_click_cost",
	"convert_rate",
	"deep_convert_rate",
	"loan_completion_cost",
	"in_app_order",
	"attribution_retention_3d_rate",
	"attribution_wechat_first_pay_30d_rate",
	"play_over",
	"loan_credit_rate",
	"shopping",
	"pre_convert_rate",
	"attribution_retention_7d_cnt",
	"attribution_game_in_app_ltv_6days",
	"wifi_play_rate",
	"dislike",
	"attribution_retention_3d_cnt",
	"play_25_feed_break_rate",
	"poi_address_click",
	"union_roi_3",
	"next_day_open_cost",
	"attribution_retention_6d_rate",
	"ctr",
	"play_50_feed_break",
	"attribution_next_day_open_rate",
	"convert_cost",
	"live_watch_one_minute_count",
	"play_duration",
	"attribution_active_pay_7d_per_count",
	"attribution_retention_4d_rate",
	"play_duration_2s",
	"in_app_cart",
	"cost",
	"play_100_feed_break_rate",
	"advanced_creative_counsel_click",
	"attribution_micro_game_0d_ltv",
	"avg_rank",
	"in_app_pay",
	"phone_confirm",
	"attribution_game_in_app_roi_6days",
	"attribution_deep_convert_cost",
	"stat_union_ltv_3",
	"attribution_wechat_first_pay_30d_count",
	"share",
	"advanced_creative_form_submit",
	"interact_per_cost",
	"attribution_micro_game_3d_ltv",
	"attribution_game_in_app_ltv_5days",
	"luban_live_click_product_cnt",
	"attribution_wechat_first_pay_30d_cost",
	"game_pay_count",
	"average_video_play",
	"stat_union_ltv_0",
	"phone_effective",
	"attribution_retention_5d_cost",
	"attribution_game_in_app_ltv_2days",
	"luban_live_enter_cnt",
	"click_download",
	"attribution_active_pay_7d_count",
	"attribution_active_pay_intra_one_day_cost",
	"convert",
	"download_start_cost",
	"in_app_detail_uv",
	"active",
	"luban_order_cnt",
	"attribution_game_in_app_ltv_7days",
	"attribution_retention_2d_cost",
	"card_show",
	"game_addiction_cost",
	"attribution_wechat_pay_30d_amount",
	"ies_challenge_click",
	"play_over_rate",
	"attribution_retention_5d_rate",
	"attribution_game_pay_7d_count",
	"luban_live_gift_amount",
	"attribution_game_in_app_roi_2days",
	"valid_play",
	"approval_count",
	"attribution_active_pay_intra_one_day_count",
	"install_finish_rate",
	"loan_completion",
	"message",
	"wechat",
	"wechat_first_pay_rate",
	"show",
	"play_duration_3s",
	"advanced_creative_phone_click",
	"follow",
	"convert_show_rate",
	"attribution_retention_6d_cost",
	"attribution_game_in_app_ltv_8days",
	"luban_live_gift_cnt",
	"average_play_progress",
	"valid_play_rate",
	"attribution_micro_game_7d_roi",
	"attribution_wechat_login_30d_cost",
	"form",
	"wechat_login_cost",
	"click_shopwindow",
	"location_click",
	"click_call_dy",
	"attribution_game_in_app_roi_3days",
	"loan_credit",
	"active_pay_amount",
	"attribution_micro_game_0d_roi",
	"attribution_game_in_app_roi_7days",
	"install_finish",
	"download_start_rate",
	"phone_connect",
	"attribution_game_in_app_ltv_1day",
	"live_component_click_count",
	"play_50_feed_break_rate",
	"luban_live_share_cnt",
	"attribution_retention_5d_cnt",
	"download_start",
	"wechat_login_count",
	"luban_live_pay_order_stat_cost",
	"attribution_game_in_app_ltv_4days",
	"play_75_feed_break_rate",
	"view",
	"advanced_creative_form_click",
	"attribution_retention_2d_cnt",
	"coupon",
	"next_day_open_rate",
	"valid_play_cost",
	"attribution_game_pay_7d_cost",
	"play_75_feed_break",
	"message_action",
	"game_addiction_rate",
	"download_finish_rate",
	"attribution_retention_7d_sum_cnt",
	"active_rate",
	"game_addiction",
	"pay_amount_roi",
	"home_visited",
	"attribution_convert",
	"attribution_game_in_app_ltv_3days",
	"register",
	"first_rental_order_count",
	"stat_pay_amount",
	"active_pay_rate",
	"attribution_game_in_app_roi_8days",
	"click_install",
	"pay_count",
	"cpa",
	"pre_convert_cost",
	"attribution_active_pay_intra_one_day_roi",
	"redirect_to_shop",
	"pre_loan_credit_cost",
	"wechat_first_pay_cost",
	"live_fans_club_join_cnt",
	"click_website",
	"attribution_micro_game_7d_ltv",
	"attribution_micro_game_3d_roi",
	"attribution_retention_6d_cnt",
	"phone",
	"live_component_click_cost",
	"active_cost",
	"luban_order_stat_amount",
	"wifi_play",
	"cpm",
	"live_component_click_rate",
	"luban_live_follow_cnt",
	"attribution_convert_cost",
	"download_finish_cost",
	"attribution_active_pay_intra_one_day_rate",
	"attribution_wechat_pay_30d_roi",
	"active_pay_cost",
	"poi_collect",
	"wechat_first_pay_count",
	"pre_convert_count",
	"attribution_retention_7d_total_cost",
	"attribution_game_in_app_roi_4days",
	"attribution_retention_3d_cost",
	"attribution_game_in_app_roi_1day",
	"union_roi_0",
	"attribution_active_pay_7d_rate",
	"advanced_creative_coupon_addition",
	"click_landing_page",
	"next_day_open",
	"luban_order_roi",
	"coupon_single_page",
	"luban_live_pay_order_count",
	"average_play_time_per_play",
	"consult_effective",
	"active_register_rate",
	"active_register_cost",
	"in_app_uv",
	"customer_effective",
	"first_order_count",
	"attribution_wechat_login_30d_count",
	"attribution_active_pay_intra_one_day_amount",
	"total_play",
	"attribution_next_day_open_cnt",
	"attribution_next_day_open_cost",
	"play_duration_10s_rate",
	"download",
	"install_finish_cost",
	"cpc",
	"attribution_retention_4d_cost",
	"download_finish",
	"attribution_retention_4d_cnt",
	"attribution_active_pay_7d_cost",
	"attribution_day_active_pay_count",
	"stat_union_ltv_7",
	"play_duration_10s",
	"play_duration_3s_rate",
	"union_roi_7",
	"redirect",
	"report",
	"map_search",
	"wechat_pay_amount",
	"ies_music_click",
	"play_duration_5s_rate",
	"luban_live_comment_cnt",
	"play_duration_sum",
	"attribution_retention_2d_rate",
	"attribution_retention_7d_cost",
	"attribution_deep_convert",
	"button",
	"loan_completion_rate",
	"pre_loan_credit",
	"luban_live_slidecart_click_cnt",
	"click",
	"play_25_feed_break",
	"comment",
	"deep_convert",
	"avg_show_cost",
	"like",
	"attribution_retention_7d_rate",
}

// NewReportAdvertiserGetV2OrderFieldFromValue returns a pointer to a valid ReportAdvertiserGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2OrderFieldFromValue(v string) (*ReportAdvertiserGetV2OrderField, error) {
	ev := ReportAdvertiserGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2OrderField: valid values are %v", v, AllowedReportAdvertiserGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_order_field value
func (v ReportAdvertiserGetV2OrderField) Ptr() *ReportAdvertiserGetV2OrderField {
	return &v
}
