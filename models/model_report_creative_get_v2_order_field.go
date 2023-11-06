/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
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
	ADVANCED_CREATIVE_FORM_CLICK_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "advanced_creative_form_click"
	CLICK_CALL_DY_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_call_dy"
	ATTRIBUTION_RETENTION_6D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_6d_cost"
	STAT_UNION_LTV_7_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_7"
	COMMUTE_FIRST_PAY_COUNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "commute_first_pay_count"
	COUPON_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "coupon"
	LOAN_COMPLETION_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "loan_completion_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_2days"
	APPROVAL_COUNT_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "approval_count"
	HOME_VISITED_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "home_visited"
	PRE_LOAN_CREDIT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "pre_loan_credit"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_4d_rate"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_4d_cnt"
	INTERACT_PER_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "interact_per_cost"
	SUBMIT_CERTIFICATION_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "submit_certification_count"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCreativeGetV2OrderField                 ReportCreativeGetV2OrderField = "live_watch_one_minute_count"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_2d_cnt"
	NEXT_DAY_OPEN_COST_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "next_day_open_cost"
	CUSTOMER_EFFECTIVE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "customer_effective"
	LUBAN_LIVE_ENTER_CNT_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "luban_live_enter_cnt"
	CONVERT_COST_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "convert_cost"
	CLICK_WEBSITE_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_website"
	DEEP_CONVERT_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "deep_convert"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_4days"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_5d_cnt"
	GAME_PAY_COST_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "game_pay_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCreativeGetV2OrderField    ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	FIRST_ORDER_COUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "first_order_count"
	WECHAT_LOGIN_COUNT_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "wechat_login_count"
	IN_APP_UV_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "in_app_uv"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_wechat_login_30d_cost"
	ACTIVE_REGISTER_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "active_register_rate"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "advanced_creative_phone_click"
	WECHAT_FIRST_PAY_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "wechat_first_pay_rate"
	PLAY_DURATION_3S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_3s_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_1day"
	PLAY_DURATION_3S_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "play_duration_3s"
	DOWNLOAD_START_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "download_start_cost"
	MESSAGE_ACTION_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "message_action"
	WECHAT_FIRST_PAY_COUNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "wechat_first_pay_count"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_6days"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "luban_live_click_product_cnt"
	DOWNLOAD_START_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "download_start"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCreativeGetV2OrderField         ReportCreativeGetV2OrderField = "attribution_active_pay_7d_per_count"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_5d_rate"
	IES_MUSIC_CLICK_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "ies_music_click"
	PAY_COUNT_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "pay_count"
	REPORT_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "report"
	LUBAN_ORDER_STAT_AMOUNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "luban_order_stat_amount"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCreativeGetV2OrderField   ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	PHONE_CONNECT_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "phone_connect"
	SHARE_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "share"
	PLAY_DURATION_2S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_2s_rate"
	FORM_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "form"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_0d_ltv"
	PLAY_DURATION_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "play_duration"
	WECHAT_PAY_AMOUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "wechat_pay_amount"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCreativeGetV2OrderField          ReportCreativeGetV2OrderField = "attribution_wechat_login_30d_count"
	IN_APP_ORDER_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "in_app_order"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCreativeGetV2OrderField  ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	GAME_ADDICTION_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "game_addiction_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_4days"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_wechat_pay_30d_amount"
	ACTIVE_REGISTER_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "active_register_cost"
	LUBAN_ORDER_ROI_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "luban_order_roi"
	CONVERT_SHOW_RATE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "convert_show_rate"
	ATTRIBUTION_RETENTION_4D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_4d_cost"
	PLAY_DURATION_10S_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "play_duration_10s"
	LOAN_CREDIT_COST_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "loan_credit_cost"
	PLAY_75_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_75_feed_break"
	DOWNLOAD_FINISH_RATE_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "download_finish_rate"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "luban_live_gift_amount"
	PLAY_75_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_75_feed_break_rate"
	NEXT_DAY_OPEN_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "next_day_open"
	CLICK_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "click"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_7d_roi"
	UNION_ROI_3_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_3"
	DOWNLOAD_START_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "download_start_rate"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_game_pay_7d_count"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_3d_roi"
	LOAN_CREDIT_RATE_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "loan_credit_rate"
	GAME_PAY_COUNT_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "game_pay_count"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_7d_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_7days"
	PLAY_OVER_RATE_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "play_over_rate"
	REGISTER_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "register"
	POI_ADDRESS_CLICK_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "poi_address_click"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCreativeGetV2OrderField       ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	ACTIVE_PAY_AMOUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "active_pay_amount"
	DOWNLOAD_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "download"
	LUBAN_LIVE_COMMENT_CNT_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "luban_live_comment_cnt"
	PRE_CONVERT_COST_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "pre_convert_cost"
	BUTTON_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "button"
	PLAY_DURATION_SUM_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "play_duration_sum"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_2days"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_7days"
	LOCATION_CLICK_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "location_click"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_5days"
	ATTRIBUTION_RETENTION_7D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_7d_cost"
	DOWNLOAD_FINISH_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "download_finish"
	VALID_PLAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "valid_play"
	COMMENT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "comment"
	PLAY_DURATION_10S_RATE_ReportCreativeGetV2OrderField                      ReportCreativeGetV2OrderField = "play_duration_10s_rate"
	PRE_LOAN_CREDIT_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "pre_loan_credit_cost"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_3d_ltv"
	ATTRIBUTION_RETENTION_3D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_3d_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCreativeGetV2OrderField ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_DEEP_CONVERT_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "attribution_deep_convert"
	REDIRECT_TO_SHOP_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "redirect_to_shop"
	WIFI_PLAY_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "wifi_play"
	CONVERT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "convert"
	DISLIKE_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "dislike"
	TOTAL_PLAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "total_play"
	VALID_PLAY_COST_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "valid_play_cost"
	LOAN_COMPLETION_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "loan_completion_cost"
	LIVE_COMPONENT_CLICK_COST_ReportCreativeGetV2OrderField                   ReportCreativeGetV2OrderField = "live_component_click_cost"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_retention_7d_sum_cnt"
	INSTALL_FINISH_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "install_finish_rate"
	COST_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "cost"
	FOLLOW_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "follow"
	CLICK_INSTALL_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "click_install"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_active_pay_7d_cost"
	PLAY_25_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_25_feed_break"
	ACTIVE_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "active"
	CPC_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpc"
	CONSULT_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "consult"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCreativeGetV2OrderField      ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	LUBAN_LIVE_GIFT_CNT_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "luban_live_gift_cnt"
	NEXT_DAY_OPEN_RATE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "next_day_open_rate"
	ATTRIBUTION_CONVERT_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "attribution_convert"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "advanced_creative_form_submit"
	DEEP_CONVERT_RATE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "deep_convert_rate"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCreativeGetV2OrderField             ReportCreativeGetV2OrderField = "attribution_active_pay_7d_count"
	PAY_AMOUNT_ROI_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "pay_amount_roi"
	QQ_ReportCreativeGetV2OrderField                                          ReportCreativeGetV2OrderField = "qq"
	AVERAGE_PLAY_PROGRESS_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "average_play_progress"
	ATTRIBUTION_RETENTION_5D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_5d_cost"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "luban_live_slidecart_click_cnt"
	WECHAT_ReportCreativeGetV2OrderField                                      ReportCreativeGetV2OrderField = "wechat"
	PHONE_EFFECTIVE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "phone_effective"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_wechat_pay_30d_roi"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_next_day_open_rate"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCreativeGetV2OrderField             ReportCreativeGetV2OrderField = "advanced_creative_counsel_click"
	CLICK_SHOPWINDOW_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "click_shopwindow"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_8days"
	MAP_SEARCH_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "map_search"
	ACTIVE_COST_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "active_cost"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_game_pay_7d_cost"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_3d_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_8days"
	AVG_CLICK_COST_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "avg_click_cost"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_7d_cnt"
	VOTE_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "vote"
	FIRST_RENTAL_ORDER_COUNT_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "first_rental_order_count"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_next_day_open_cost"
	PHONE_ReportCreativeGetV2OrderField                                       ReportCreativeGetV2OrderField = "phone"
	LOTTERY_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "lottery"
	PLAY_100_FEED_BREAK_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "play_100_feed_break"
	SHOW_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "show"
	DEEP_CONVERT_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "deep_convert_cost"
	CONSULT_EFFECTIVE_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "consult_effective"
	PLAY_DURATION_2S_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "play_duration_2s"
	UNION_ROI_0_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_0"
	ACTIVE_PAY_RATE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "active_pay_rate"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_3d_cnt"
	PRE_CONVERT_COUNT_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "pre_convert_count"
	SHOPPING_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "shopping"
	PLAY_100_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "play_100_feed_break_rate"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_7d_ltv"
	WIFI_PLAY_RATE_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "wifi_play_rate"
	GAME_ADDICTION_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "game_addiction"
	LIKE_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "like"
	AVG_RANK_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "avg_rank"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCreativeGetV2OrderField         ReportCreativeGetV2OrderField = "attribution_retention_7d_total_cost"
	VALID_PLAY_RATE_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "valid_play_rate"
	PLAY_25_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_25_feed_break_rate"
	INSTALL_FINISH_COST_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "install_finish_cost"
	INSTALL_FINISH_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "install_finish"
	PHONE_CONFIRM_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "phone_confirm"
	LUBAN_ORDER_CNT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "luban_order_cnt"
	POI_COLLECT_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "poi_collect"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "attribution_active_pay_7d_rate"
	LUBAN_LIVE_SHARE_CNT_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "luban_live_share_cnt"
	PLAY_OVER_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "play_over"
	LOAN_CREDIT_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "loan_credit"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "average_play_time_per_play"
	PLAY_DURATION_5S_RATE_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "play_duration_5s_rate"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCreativeGetV2OrderField              ReportCreativeGetV2OrderField = "luban_live_pay_order_stat_cost"
	CTR_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "ctr"
	IN_APP_DETAIL_UV_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "in_app_detail_uv"
	WECHAT_FIRST_PAY_COST_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "wechat_first_pay_cost"
	STAT_UNION_LTV_3_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_3"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ACTIVE_PAY_COST_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "active_pay_cost"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_day_active_pay_count"
	LIVE_COMPONENT_CLICK_RATE_ReportCreativeGetV2OrderField                   ReportCreativeGetV2OrderField = "live_component_click_rate"
	IES_CHALLENGE_CLICK_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "ies_challenge_click"
	ATTRIBUTION_RETENTION_2D_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_2d_cost"
	CLICK_LANDING_PAGE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "click_landing_page"
	GAME_ADDICTION_RATE_ReportCreativeGetV2OrderField                         ReportCreativeGetV2OrderField = "game_addiction_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_6days"
	MESSAGE_ReportCreativeGetV2OrderField                                     ReportCreativeGetV2OrderField = "message"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "luban_live_pay_order_count"
	PLAY_50_FEED_BREAK_RATE_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "play_50_feed_break_rate"
	DOWNLOAD_FINISH_COST_ReportCreativeGetV2OrderField                        ReportCreativeGetV2OrderField = "download_finish_cost"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "advanced_creative_coupon_addition"
	CONVERT_RATE_ReportCreativeGetV2OrderField                                ReportCreativeGetV2OrderField = "convert_rate"
	UNION_ROI_7_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "union_roi_7"
	AVERAGE_VIDEO_PLAY_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "average_video_play"
	LIVE_COMPONENT_CLICK_COUNT_ReportCreativeGetV2OrderField                  ReportCreativeGetV2OrderField = "live_component_click_count"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_retention_6d_rate"
	LUBAN_LIVE_FOLLOW_CNT_ReportCreativeGetV2OrderField                       ReportCreativeGetV2OrderField = "luban_live_follow_cnt"
	AVG_SHOW_COST_ReportCreativeGetV2OrderField                               ReportCreativeGetV2OrderField = "avg_show_cost"
	PLAY_50_FEED_BREAK_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "play_50_feed_break"
	CLICK_DOWNLOAD_ReportCreativeGetV2OrderField                              ReportCreativeGetV2OrderField = "click_download"
	PRE_CONVERT_RATE_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "pre_convert_rate"
	STAT_UNION_LTV_0_ReportCreativeGetV2OrderField                            ReportCreativeGetV2OrderField = "stat_union_ltv_0"
	IN_APP_PAY_ReportCreativeGetV2OrderField                                  ReportCreativeGetV2OrderField = "in_app_pay"
	ACTIVE_RATE_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "active_rate"
	WECHAT_LOGIN_COST_ReportCreativeGetV2OrderField                           ReportCreativeGetV2OrderField = "wechat_login_cost"
	CPM_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpm"
	ATTRIBUTION_CONVERT_COST_ReportCreativeGetV2OrderField                    ReportCreativeGetV2OrderField = "attribution_convert_cost"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCreativeGetV2OrderField                ReportCreativeGetV2OrderField = "attribution_retention_6d_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_5days"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCreativeGetV2OrderField       ReportCreativeGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_next_day_open_cnt"
	IN_APP_CART_ReportCreativeGetV2OrderField                                 ReportCreativeGetV2OrderField = "in_app_cart"
	LOAN_COMPLETION_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "loan_completion"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_deep_convert_cost"
	REDIRECT_ReportCreativeGetV2OrderField                                    ReportCreativeGetV2OrderField = "redirect"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCreativeGetV2OrderField           ReportCreativeGetV2OrderField = "attribution_game_in_app_roi_3days"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_day_acitve_pay_count"
	VIEW_ReportCreativeGetV2OrderField                                        ReportCreativeGetV2OrderField = "view"
	CARD_SHOW_ReportCreativeGetV2OrderField                                   ReportCreativeGetV2OrderField = "card_show"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCreativeGetV2OrderField            ReportCreativeGetV2OrderField = "attribution_game_in_app_ltv_1day"
	STAT_PAY_AMOUNT_ReportCreativeGetV2OrderField                             ReportCreativeGetV2OrderField = "stat_pay_amount"
	COUPON_SINGLE_PAGE_ReportCreativeGetV2OrderField                          ReportCreativeGetV2OrderField = "coupon_single_page"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCreativeGetV2OrderField   ReportCreativeGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	CPA_ReportCreativeGetV2OrderField                                         ReportCreativeGetV2OrderField = "cpa"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCreativeGetV2OrderField               ReportCreativeGetV2OrderField = "attribution_micro_game_0d_roi"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCreativeGetV2OrderField                     ReportCreativeGetV2OrderField = "live_fans_club_join_cnt"
)

// All allowed values of ReportCreativeGetV2OrderField enum
var AllowedReportCreativeGetV2OrderFieldEnumValues = []ReportCreativeGetV2OrderField{
	"advanced_creative_form_click",
	"click_call_dy",
	"attribution_retention_6d_cost",
	"stat_union_ltv_7",
	"commute_first_pay_count",
	"coupon",
	"loan_completion_rate",
	"attribution_game_in_app_ltv_2days",
	"approval_count",
	"home_visited",
	"pre_loan_credit",
	"attribution_retention_4d_rate",
	"attribution_retention_4d_cnt",
	"interact_per_cost",
	"submit_certification_count",
	"live_watch_one_minute_count",
	"attribution_retention_2d_cnt",
	"next_day_open_cost",
	"customer_effective",
	"luban_live_enter_cnt",
	"convert_cost",
	"click_website",
	"deep_convert",
	"attribution_game_in_app_ltv_4days",
	"attribution_retention_5d_cnt",
	"game_pay_cost",
	"attribution_active_pay_intra_one_day_roi",
	"first_order_count",
	"wechat_login_count",
	"in_app_uv",
	"attribution_wechat_login_30d_cost",
	"active_register_rate",
	"advanced_creative_phone_click",
	"wechat_first_pay_rate",
	"play_duration_3s_rate",
	"attribution_game_in_app_roi_1day",
	"play_duration_3s",
	"download_start_cost",
	"message_action",
	"wechat_first_pay_count",
	"attribution_game_in_app_ltv_6days",
	"luban_live_click_product_cnt",
	"download_start",
	"attribution_active_pay_7d_per_count",
	"attribution_retention_5d_rate",
	"ies_music_click",
	"pay_count",
	"report",
	"luban_order_stat_amount",
	"attribution_active_pay_intra_one_day_cost",
	"phone_connect",
	"share",
	"play_duration_2s_rate",
	"form",
	"attribution_micro_game_0d_ltv",
	"play_duration",
	"wechat_pay_amount",
	"attribution_wechat_login_30d_count",
	"in_app_order",
	"attribution_active_pay_intra_one_day_count",
	"game_addiction_cost",
	"attribution_game_in_app_roi_4days",
	"attribution_wechat_pay_30d_amount",
	"active_register_cost",
	"luban_order_roi",
	"convert_show_rate",
	"attribution_retention_4d_cost",
	"play_duration_10s",
	"loan_credit_cost",
	"play_75_feed_break",
	"download_finish_rate",
	"luban_live_gift_amount",
	"play_75_feed_break_rate",
	"next_day_open",
	"click",
	"attribution_micro_game_7d_roi",
	"union_roi_3",
	"download_start_rate",
	"attribution_game_pay_7d_count",
	"attribution_micro_game_3d_roi",
	"loan_credit_rate",
	"game_pay_count",
	"attribution_retention_7d_rate",
	"attribution_game_in_app_roi_7days",
	"play_over_rate",
	"register",
	"poi_address_click",
	"attribution_retention_2d_rate",
	"attribution_wechat_first_pay_30d_cost",
	"active_pay_amount",
	"download",
	"luban_live_comment_cnt",
	"pre_convert_cost",
	"button",
	"play_duration_sum",
	"attribution_game_in_app_roi_2days",
	"attribution_game_in_app_ltv_7days",
	"location_click",
	"attribution_game_in_app_ltv_5days",
	"attribution_retention_7d_cost",
	"download_finish",
	"valid_play",
	"comment",
	"play_duration_10s_rate",
	"pre_loan_credit_cost",
	"attribution_micro_game_3d_ltv",
	"attribution_retention_3d_cost",
	"attribution_active_pay_intra_one_day_amount",
	"attribution_deep_convert",
	"redirect_to_shop",
	"wifi_play",
	"convert",
	"dislike",
	"total_play",
	"valid_play_cost",
	"loan_completion_cost",
	"live_component_click_cost",
	"attribution_retention_7d_sum_cnt",
	"install_finish_rate",
	"cost",
	"follow",
	"click_install",
	"attribution_active_pay_7d_cost",
	"play_25_feed_break",
	"active",
	"cpc",
	"consult",
	"attribution_wechat_first_pay_30d_count",
	"luban_live_gift_cnt",
	"next_day_open_rate",
	"attribution_convert",
	"advanced_creative_form_submit",
	"deep_convert_rate",
	"attribution_active_pay_7d_count",
	"pay_amount_roi",
	"qq",
	"average_play_progress",
	"attribution_retention_5d_cost",
	"luban_live_slidecart_click_cnt",
	"wechat",
	"phone_effective",
	"attribution_wechat_pay_30d_roi",
	"attribution_next_day_open_rate",
	"advanced_creative_counsel_click",
	"click_shopwindow",
	"attribution_game_in_app_roi_8days",
	"map_search",
	"active_cost",
	"attribution_game_pay_7d_cost",
	"attribution_retention_3d_rate",
	"attribution_game_in_app_ltv_8days",
	"avg_click_cost",
	"attribution_retention_7d_cnt",
	"vote",
	"first_rental_order_count",
	"attribution_next_day_open_cost",
	"phone",
	"lottery",
	"play_100_feed_break",
	"show",
	"deep_convert_cost",
	"consult_effective",
	"play_duration_2s",
	"union_roi_0",
	"active_pay_rate",
	"attribution_retention_3d_cnt",
	"pre_convert_count",
	"shopping",
	"play_100_feed_break_rate",
	"attribution_micro_game_7d_ltv",
	"wifi_play_rate",
	"game_addiction",
	"like",
	"avg_rank",
	"attribution_retention_7d_total_cost",
	"valid_play_rate",
	"play_25_feed_break_rate",
	"install_finish_cost",
	"install_finish",
	"phone_confirm",
	"luban_order_cnt",
	"poi_collect",
	"attribution_active_pay_7d_rate",
	"luban_live_share_cnt",
	"play_over",
	"loan_credit",
	"average_play_time_per_play",
	"play_duration_5s_rate",
	"luban_live_pay_order_stat_cost",
	"ctr",
	"in_app_detail_uv",
	"wechat_first_pay_cost",
	"stat_union_ltv_3",
	"attribution_game_in_app_ltv_3days",
	"active_pay_cost",
	"attribution_day_active_pay_count",
	"live_component_click_rate",
	"ies_challenge_click",
	"attribution_retention_2d_cost",
	"click_landing_page",
	"game_addiction_rate",
	"attribution_game_in_app_roi_6days",
	"message",
	"luban_live_pay_order_count",
	"play_50_feed_break_rate",
	"download_finish_cost",
	"advanced_creative_coupon_addition",
	"convert_rate",
	"union_roi_7",
	"average_video_play",
	"live_component_click_count",
	"attribution_retention_6d_rate",
	"luban_live_follow_cnt",
	"avg_show_cost",
	"play_50_feed_break",
	"click_download",
	"pre_convert_rate",
	"stat_union_ltv_0",
	"in_app_pay",
	"active_rate",
	"wechat_login_cost",
	"cpm",
	"attribution_convert_cost",
	"attribution_retention_6d_cnt",
	"attribution_game_in_app_roi_5days",
	"attribution_wechat_first_pay_30d_rate",
	"attribution_next_day_open_cnt",
	"in_app_cart",
	"loan_completion",
	"attribution_deep_convert_cost",
	"redirect",
	"attribution_game_in_app_roi_3days",
	"attribution_day_acitve_pay_count",
	"view",
	"card_show",
	"attribution_game_in_app_ltv_1day",
	"stat_pay_amount",
	"coupon_single_page",
	"attribution_active_pay_intra_one_day_rate",
	"cpa",
	"attribution_micro_game_0d_roi",
	"live_fans_club_join_cnt",
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
