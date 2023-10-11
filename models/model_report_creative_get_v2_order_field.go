/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2OrderField
type ReportCreativeGetV2OrderField string

// List of report_creative_get_v2_order_field
const (
	IN_APP_ORDER_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "in_app_order"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCreativeGetV2OrderField         ReportCreativeGetV2OrderField = "attribution_active_pay_7d_per_count"
	WECHAT_LOGIN_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "wechat_login_cost"
	DOWNLOAD_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "download"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "advanced_creative_coupon_addition"
	REDIRECT_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "redirect"
	STAT_UNION_LTV_7_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_7"
	GAME_PAY_COST_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "game_pay_cost"
	LUBAN_LIVE_GIFT_CNT_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "luban_live_gift_cnt"
	UNION_ROI_0_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_0"
	CLICK_WEBSITE_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_website"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_2days"
	IN_APP_UV_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "in_app_uv"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "luban_live_click_product_cnt"
	DOWNLOAD_START_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "download_start"
	ATTRIBUTION_RETENTION_4D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_4d_cost"
	NEXT_DAY_OPEN_RATE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "next_day_open_rate"
	LUBAN_ORDER_STAT_AMOUNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "luban_order_stat_amount"
	WIFI_PLAY_RATE_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "wifi_play_rate"
	LUBAN_LIVE_SHARE_CNT_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "luban_live_share_cnt"
	FIRST_ORDER_COUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "first_order_count"
	LOCATION_CLICK_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "location_click"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCreativeGetV2OrderField          ReportCreativeGetV2OrderField = "attribution_wechat_login_30d_count"
	LUBAN_ORDER_CNT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "luban_order_cnt"
	STAT_PAY_AMOUNT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "stat_pay_amount"
	GAME_ADDICTION_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "game_addiction"
	ACTIVE_REGISTER_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "active_register_rate"
	COMMUTE_FIRST_PAY_COUNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "commute_first_pay_count"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "luban_live_pay_order_stat_cost"
	LOAN_CREDIT_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "loan_credit"
	VALID_PLAY_COST_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "valid_play_cost"
	PAY_AMOUNT_ROI_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "pay_amount_roi"
	HOME_VISITED_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "home_visited"
	LOAN_COMPLETION_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "loan_completion"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_3d_ltv"
	PLAY_OVER_RATE_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "play_over_rate"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCreativeGetV2OrderField         ReportCreativeGetV2OrderField = "attribution_retention_7d_total_cost"
	WECHAT_FIRST_PAY_COUNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "wechat_first_pay_count"
	ATTRIBUTION_RETENTION_3D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_3d_cost"
	WECHAT_LOGIN_COUNT_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "wechat_login_count"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCreativeGetV2OrderField       ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	CPA_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpa"
	MESSAGE_ACTION_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "message_action"
	GAME_ADDICTION_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "game_addiction_rate"
	REDIRECT_TO_SHOP_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "redirect_to_shop"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_active_pay_7d_rate"
	IN_APP_DETAIL_UV_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "in_app_detail_uv"
	AVERAGE_PLAY_PROGRESS_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "average_play_progress"
	LOAN_COMPLETION_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "loan_completion_rate"
	COUPON_SINGLE_PAGE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "coupon_single_page"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_game_pay_7d_count"
	ATTRIBUTION_RETENTION_2D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_2d_cost"
	BUTTON_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "button"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_8days"
	PLAY_100_FEED_BREAK_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "play_100_feed_break"
	VALID_PLAY_RATE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "valid_play_rate"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCreativeGetV2OrderField                 ReportCreativeGetV2OrderField = "live_watch_one_minute_count"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_7d_ltv"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_4d_rate"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "luban_live_slidecart_click_cnt"
	STAT_UNION_LTV_3_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_3"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_wechat_pay_30d_amount"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_1day"
	LOAN_COMPLETION_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "loan_completion_cost"
	GAME_PAY_COUNT_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "game_pay_count"
	ATTRIBUTION_DEEP_CONVERT_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "attribution_deep_convert"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_day_acitve_pay_count"
	POI_COLLECT_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "poi_collect"
	ACTIVE_REGISTER_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "active_register_cost"
	SHOW_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "show"
	AVG_SHOW_COST_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "avg_show_cost"
	ATTRIBUTION_RETENTION_5D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_5d_cost"
	ACTIVE_PAY_RATE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "active_pay_rate"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "luban_live_pay_order_count"
	PLAY_DURATION_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "play_duration"
	ATTRIBUTION_CONVERT_COST_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "attribution_convert_cost"
	LUBAN_LIVE_FOLLOW_CNT_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "luban_live_follow_cnt"
	PLAY_75_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_75_feed_break"
	IES_MUSIC_CLICK_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "ies_music_click"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_5d_cnt"
	ATTRIBUTION_CONVERT_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "attribution_convert"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_2d_cnt"
	VIEW_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "view"
	MAP_SEARCH_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "map_search"
	DOWNLOAD_FINISH_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "download_finish_cost"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_4d_cnt"
	CONVERT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "convert"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_7days"
	LUBAN_LIVE_ENTER_CNT_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "luban_live_enter_cnt"
	INSTALL_FINISH_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "install_finish_cost"
	ADVANCED_CREATIVE_FORM_CLICK_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "advanced_creative_form_click"
	PLAY_DURATION_2S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_2s_rate"
	CONVERT_RATE_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "convert_rate"
	PRE_CONVERT_COST_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "pre_convert_cost"
	ACTIVE_PAY_COST_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "active_pay_cost"
	DOWNLOAD_FINISH_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "download_finish_rate"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCreativeGetV2OrderField             ReportCreativeGetV2OrderField = "advanced_creative_counsel_click"
	PRE_CONVERT_RATE_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "pre_convert_rate"
	DEEP_CONVERT_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "deep_convert"
	CONSULT_EFFECTIVE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "consult_effective"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCreativeGetV2OrderField ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_wechat_login_30d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_6days"
	LIVE_COMPONENT_CLICK_RATE_ReportCreativeGetV2OrderField                   ReportCreativeGetV2OrderField = "live_component_click_rate"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_0d_roi"
	STAT_UNION_LTV_0_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_0"
	CTR_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "ctr"
	FIRST_RENTAL_ORDER_COUNT_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "first_rental_order_count"
	ACTIVE_COST_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "active_cost"
	LOAN_CREDIT_COST_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "loan_credit_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCreativeGetV2OrderField       ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_5days"
	DISLIKE_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "dislike"
	AVG_RANK_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "avg_rank"
	PAY_COUNT_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "pay_count"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_2days"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_4days"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_next_day_open_rate"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCreativeGetV2OrderField      ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	LIVE_COMPONENT_CLICK_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "live_component_click_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCreativeGetV2OrderField    ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	REGISTER_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "register"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_5days"
	DOWNLOAD_START_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "download_start_cost"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_next_day_open_cnt"
	LOTTERY_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "lottery"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_5d_rate"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCreativeGetV2OrderField  ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	PLAY_DURATION_SUM_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "play_duration_sum"
	ACTIVE_RATE_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "active_rate"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "advanced_creative_form_submit"
	DEEP_CONVERT_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "deep_convert_cost"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_0d_ltv"
	WECHAT_FIRST_PAY_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "wechat_first_pay_rate"
	ACTIVE_PAY_AMOUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "active_pay_amount"
	PHONE_EFFECTIVE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "phone_effective"
	PLAY_DURATION_10S_RATE_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "play_duration_10s_rate"
	MESSAGE_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "message"
	CLICK_LANDING_PAGE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "click_landing_page"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_1day"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_3d_roi"
	FORM_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "form"
	CONSULT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "consult"
	PLAY_DURATION_3S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_3s_rate"
	IN_APP_CART_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "in_app_cart"
	NEXT_DAY_OPEN_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "next_day_open"
	CPM_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpm"
	PLAY_25_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_25_feed_break"
	VALID_PLAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "valid_play"
	PHONE_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "phone"
	UNION_ROI_7_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_7"
	LUBAN_ORDER_ROI_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "luban_order_roi"
	PLAY_OVER_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "play_over"
	AVERAGE_VIDEO_PLAY_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "average_video_play"
	INTERACT_PER_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "interact_per_cost"
	CLICK_SHOPWINDOW_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "click_shopwindow"
	WECHAT_PAY_AMOUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "wechat_pay_amount"
	CLICK_INSTALL_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_install"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_wechat_pay_30d_roi"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "average_play_time_per_play"
	REPORT_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "report"
	COST_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "cost"
	SHOPPING_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "shopping"
	PLAY_DURATION_2S_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "play_duration_2s"
	UNION_ROI_3_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_3"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_6d_cnt"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_retention_7d_sum_cnt"
	PLAY_50_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_50_feed_break_rate"
	PLAY_50_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_50_feed_break"
	CONVERT_SHOW_RATE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "convert_show_rate"
	CLICK_CALL_DY_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_call_dy"
	ATTRIBUTION_RETENTION_7D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_7d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_8days"
	DOWNLOAD_FINISH_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "download_finish"
	APPROVAL_COUNT_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "approval_count"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "luban_live_gift_amount"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCreativeGetV2OrderField             ReportCreativeGetV2OrderField = "attribution_active_pay_7d_count"
	LOAN_CREDIT_RATE_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "loan_credit_rate"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_3d_cnt"
	LUBAN_LIVE_COMMENT_CNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "luban_live_comment_cnt"
	AVG_CLICK_COST_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "avg_click_cost"
	COMMENT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "comment"
	INSTALL_FINISH_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "install_finish"
	GAME_ADDICTION_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "game_addiction_cost"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "live_fans_club_join_cnt"
	PLAY_DURATION_3S_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_6d_rate"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_next_day_open_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCreativeGetV2OrderField   ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	IN_APP_PAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "in_app_pay"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_active_pay_7d_cost"
	ATTRIBUTION_RETENTION_6D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_6d_cost"
	PLAY_75_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_75_feed_break_rate"
	QQ_ReportCreativeGetV2OrderField                                          ReportCreativeGetV2OrderField = "qq"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_day_active_pay_count"
	CLICK_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "click"
	CUSTOMER_EFFECTIVE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "customer_effective"
	INSTALL_FINISH_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "install_finish_rate"
	POI_ADDRESS_CLICK_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "poi_address_click"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_7days"
	NEXT_DAY_OPEN_COST_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "next_day_open_cost"
	FOLLOW_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "follow"
	TOTAL_PLAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "total_play"
	CLICK_DOWNLOAD_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "click_download"
	PHONE_CONFIRM_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "phone_confirm"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_deep_convert_cost"
	PLAY_25_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_25_feed_break_rate"
	PLAY_DURATION_5S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_5s_rate"
	IES_CHALLENGE_CLICK_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "ies_challenge_click"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_game_pay_7d_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_6days"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "advanced_creative_phone_click"
	SUBMIT_CERTIFICATION_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "submit_certification_count"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_4days"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_3days"
	DOWNLOAD_START_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "download_start_rate"
	PRE_LOAN_CREDIT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "pre_loan_credit"
	VOTE_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "vote"
	CPC_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpc"
	PLAY_DURATION_10S_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "play_duration_10s"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_7d_cnt"
	SHARE_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "share"
	WIFI_PLAY_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "wifi_play"
	WECHAT_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "wechat"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCreativeGetV2OrderField   ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	CARD_SHOW_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "card_show"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_7d_roi"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_3days"
	PLAY_100_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "play_100_feed_break_rate"
	COUPON_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "coupon"
	PRE_LOAN_CREDIT_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "pre_loan_credit_cost"
	LIVE_COMPONENT_CLICK_COST_ReportCreativeGetV2OrderField                   ReportCreativeGetV2OrderField = "live_component_click_cost"
	LIKE_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "like"
	PHONE_CONNECT_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "phone_connect"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_3d_rate"
	ACTIVE_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "active"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_7d_rate"
	WECHAT_FIRST_PAY_COST_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "wechat_first_pay_cost"
	CONVERT_COST_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "convert_cost"
	DEEP_CONVERT_RATE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "deep_convert_rate"
	PRE_CONVERT_COUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "pre_convert_count"
)

// All allowed values of ReportCreativeGetV2OrderField enum
var AllowedReportCreativeGetV2OrderFieldEnumValues = []ReportCreativeGetV2OrderField{
	"in_app_order",
	"attribution_active_pay_7d_per_count",
	"wechat_login_cost",
	"download",
	"advanced_creative_coupon_addition",
	"redirect",
	"stat_union_ltv_7",
	"game_pay_cost",
	"luban_live_gift_cnt",
	"union_roi_0",
	"click_website",
	"attribution_game_in_app_roi_2days",
	"in_app_uv",
	"luban_live_click_product_cnt",
	"download_start",
	"attribution_retention_4d_cost",
	"next_day_open_rate",
	"luban_order_stat_amount",
	"wifi_play_rate",
	"luban_live_share_cnt",
	"first_order_count",
	"location_click",
	"attribution_wechat_login_30d_count",
	"luban_order_cnt",
	"stat_pay_amount",
	"game_addiction",
	"active_register_rate",
	"commute_first_pay_count",
	"luban_live_pay_order_stat_cost",
	"loan_credit",
	"valid_play_cost",
	"pay_amount_roi",
	"home_visited",
	"loan_completion",
	"attribution_micro_game_3d_ltv",
	"play_over_rate",
	"attribution_retention_7d_total_cost",
	"wechat_first_pay_count",
	"attribution_retention_3d_cost",
	"wechat_login_count",
	"attribution_wechat_first_pay_30d_rate",
	"cpa",
	"message_action",
	"game_addiction_rate",
	"redirect_to_shop",
	"attribution_active_pay_7d_rate",
	"in_app_detail_uv",
	"average_play_progress",
	"loan_completion_rate",
	"coupon_single_page",
	"attribution_game_pay_7d_count",
	"attribution_retention_2d_cost",
	"button",
	"attribution_game_in_app_roi_8days",
	"play_100_feed_break",
	"valid_play_rate",
	"live_watch_one_minute_count",
	"attribution_micro_game_7d_ltv",
	"attribution_retention_4d_rate",
	"luban_live_slidecart_click_cnt",
	"stat_union_ltv_3",
	"attribution_wechat_pay_30d_amount",
	"attribution_game_in_app_ltv_1day",
	"loan_completion_cost",
	"game_pay_count",
	"attribution_deep_convert",
	"attribution_day_acitve_pay_count",
	"poi_collect",
	"active_register_cost",
	"show",
	"avg_show_cost",
	"attribution_retention_5d_cost",
	"active_pay_rate",
	"luban_live_pay_order_count",
	"play_duration",
	"attribution_convert_cost",
	"luban_live_follow_cnt",
	"play_75_feed_break",
	"ies_music_click",
	"attribution_retention_5d_cnt",
	"attribution_convert",
	"attribution_retention_2d_cnt",
	"view",
	"map_search",
	"download_finish_cost",
	"attribution_retention_4d_cnt",
	"convert",
	"attribution_game_in_app_roi_7days",
	"luban_live_enter_cnt",
	"install_finish_cost",
	"advanced_creative_form_click",
	"play_duration_2s_rate",
	"convert_rate",
	"pre_convert_cost",
	"active_pay_cost",
	"download_finish_rate",
	"advanced_creative_counsel_click",
	"pre_convert_rate",
	"deep_convert",
	"consult_effective",
	"attribution_active_pay_intra_one_day_amount",
	"attribution_wechat_login_30d_cost",
	"attribution_game_in_app_ltv_6days",
	"live_component_click_rate",
	"attribution_micro_game_0d_roi",
	"stat_union_ltv_0",
	"ctr",
	"first_rental_order_count",
	"active_cost",
	"loan_credit_cost",
	"attribution_wechat_first_pay_30d_cost",
	"attribution_game_in_app_ltv_5days",
	"dislike",
	"avg_rank",
	"pay_count",
	"attribution_game_in_app_ltv_2days",
	"attribution_game_in_app_ltv_4days",
	"attribution_next_day_open_rate",
	"attribution_wechat_first_pay_30d_count",
	"live_component_click_count",
	"attribution_active_pay_intra_one_day_roi",
	"register",
	"attribution_game_in_app_roi_5days",
	"download_start_cost",
	"attribution_next_day_open_cnt",
	"lottery",
	"attribution_retention_5d_rate",
	"attribution_retention_2d_rate",
	"attribution_active_pay_intra_one_day_count",
	"play_duration_sum",
	"active_rate",
	"advanced_creative_form_submit",
	"deep_convert_cost",
	"attribution_micro_game_0d_ltv",
	"wechat_first_pay_rate",
	"active_pay_amount",
	"phone_effective",
	"play_duration_10s_rate",
	"message",
	"click_landing_page",
	"attribution_game_in_app_roi_1day",
	"attribution_micro_game_3d_roi",
	"form",
	"consult",
	"play_duration_3s_rate",
	"in_app_cart",
	"next_day_open",
	"cpm",
	"play_25_feed_break",
	"valid_play",
	"phone",
	"union_roi_7",
	"luban_order_roi",
	"play_over",
	"average_video_play",
	"interact_per_cost",
	"click_shopwindow",
	"wechat_pay_amount",
	"click_install",
	"attribution_wechat_pay_30d_roi",
	"average_play_time_per_play",
	"report",
	"cost",
	"shopping",
	"play_duration_2s",
	"union_roi_3",
	"attribution_retention_6d_cnt",
	"attribution_retention_7d_sum_cnt",
	"play_50_feed_break_rate",
	"play_50_feed_break",
	"convert_show_rate",
	"click_call_dy",
	"attribution_retention_7d_cost",
	"attribution_game_in_app_ltv_8days",
	"download_finish",
	"approval_count",
	"luban_live_gift_amount",
	"attribution_active_pay_7d_count",
	"loan_credit_rate",
	"attribution_retention_3d_cnt",
	"luban_live_comment_cnt",
	"avg_click_cost",
	"comment",
	"install_finish",
	"game_addiction_cost",
	"live_fans_club_join_cnt",
	"play_duration_3s",
	"attribution_retention_6d_rate",
	"attribution_next_day_open_cost",
	"attribution_active_pay_intra_one_day_cost",
	"in_app_pay",
	"attribution_active_pay_7d_cost",
	"attribution_retention_6d_cost",
	"play_75_feed_break_rate",
	"qq",
	"attribution_day_active_pay_count",
	"click",
	"customer_effective",
	"install_finish_rate",
	"poi_address_click",
	"attribution_game_in_app_ltv_7days",
	"next_day_open_cost",
	"follow",
	"total_play",
	"click_download",
	"phone_confirm",
	"attribution_deep_convert_cost",
	"play_25_feed_break_rate",
	"play_duration_5s_rate",
	"ies_challenge_click",
	"attribution_game_pay_7d_cost",
	"attribution_game_in_app_roi_6days",
	"advanced_creative_phone_click",
	"submit_certification_count",
	"attribution_game_in_app_roi_4days",
	"attribution_game_in_app_ltv_3days",
	"download_start_rate",
	"pre_loan_credit",
	"vote",
	"cpc",
	"play_duration_10s",
	"attribution_retention_7d_cnt",
	"share",
	"wifi_play",
	"wechat",
	"attribution_active_pay_intra_one_day_rate",
	"card_show",
	"attribution_micro_game_7d_roi",
	"attribution_game_in_app_roi_3days",
	"play_100_feed_break_rate",
	"coupon",
	"pre_loan_credit_cost",
	"live_component_click_cost",
	"like",
	"phone_connect",
	"attribution_retention_3d_rate",
	"active",
	"attribution_retention_7d_rate",
	"wechat_first_pay_cost",
	"convert_cost",
	"deep_convert_rate",
	"pre_convert_count",
}

// NewReportCreativeGetV2OrderFieldFromValue returns a pointer to a valid ReportCreativeGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2OrderFieldFromValue(v string) (*ReportCreativeGetV2OrderField, error) {
	ev := ReportCreativeGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2OrderField: valid values are %v", v, AllowedReportCreativeGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_order_field value
func (v ReportCreativeGetV2OrderField) Ptr() *ReportCreativeGetV2OrderField {
	return &v
}
