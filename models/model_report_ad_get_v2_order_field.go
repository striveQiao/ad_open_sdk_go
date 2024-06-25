/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2OrderField
type ReportAdGetV2OrderField string

// List of report_ad_get_v2_order_field
const (
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_form_submit"
	LUBAN_LIVE_SHARE_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_share_cnt"
	ATTRIBUTION_RETENTION_2D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_cost"
	WECHAT_FIRST_PAY_COST_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_cost"
	CLICK_SHOPWINDOW_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "click_shopwindow"
	CLICK_LANDING_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "click_landing_page"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_active_pay_count"
	VOTE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "vote"
	MAP_SEARCH_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "map_search"
	IES_CHALLENGE_CLICK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "ies_challenge_click"
	POI_ADDRESS_CLICK_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "poi_address_click"
	DEEP_CONVERT_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_rate"
	UNION_ROI_3_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_3"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	AVERAGE_VIDEO_PLAY_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "average_video_play"
	PLAY_OVER_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "play_over"
	IN_APP_DETAIL_UV_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "in_app_detail_uv"
	PLAY_100_FEED_BREAK_RATE_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "play_100_feed_break_rate"
	CONVERT_SHOW_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "convert_show_rate"
	CPA_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpa"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_4days"
	IN_APP_ORDER_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "in_app_order"
	STAT_PAY_AMOUNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "stat_pay_amount"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_rate"
	PLAY_DURATION_3S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_3s_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_2days"
	PLAY_75_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_75_feed_break_rate"
	ACTIVE_RATE_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_rate"
	ACTIVE_REGISTER_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_rate"
	ATTRIBUTION_RETENTION_5D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_5d_cnt"
	TOTAL_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "total_play"
	CLICK_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "click"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_3days"
	VALID_PLAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_cost"
	FOLLOW_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "follow"
	IN_APP_PAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "in_app_pay"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_login_30d_cost"
	INSTALL_FINISH_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "install_finish"
	PAY_AMOUNT_ROI_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "pay_amount_roi"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_3days"
	UNION_ROI_7_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_7"
	DOWNLOAD_FINISH_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_8days"
	PLAY_DURATION_10S_RATE_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "play_duration_10s_rate"
	ATTRIBUTION_RETENTION_4D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_5days"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_acitve_pay_count"
	DOWNLOAD_START_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_rate"
	PLAY_25_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_25_feed_break_rate"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_retention_7d_total_cost"
	ACTIVE_PAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_rate"
	LUBAN_LIVE_FOLLOW_CNT_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "luban_live_follow_cnt"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "attribution_active_pay_7d_count"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_ltv"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_rate"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "advanced_creative_counsel_click"
	HOME_VISITED_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "home_visited"
	NEXT_DAY_OPEN_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "next_day_open"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_6days"
	FIRST_RENTAL_ORDER_COUNT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "first_rental_order_count"
	GAME_ADDICTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_addiction"
	CONSULT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "consult"
	LIVE_COMPONENT_CLICK_RATE_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_rate"
	LIVE_FANS_CLUB_JOIN_CNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "live_fans_club_join_cnt"
	AVG_CLICK_COST_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "avg_click_cost"
	ATTRIBUTION_RETENTION_3D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_cost"
	REPORT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "report"
	ATTRIBUTION_RETENTION_6D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_rate"
	PLAY_DURATION_5S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_5s_rate"
	COMMENT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "comment"
	ATTRIBUTION_RETENTION_4D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_rate"
	PRE_LOAN_CREDIT_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "pre_loan_credit_cost"
	LUBAN_ORDER_CNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_cnt"
	PLAY_DURATION_10S_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_10s"
	CARD_SHOW_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "card_show"
	CONVERT_RATE_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_rate"
	PHONE_CONNECT_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_connect"
	LOAN_COMPLETION_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_cost"
	ATTRIBUTION_RETENTION_6D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_6d_cnt"
	ATTRIBUTION_RETENTION_3D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_rate"
	ACTIVE_PAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_cost"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_roi"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_slidecart_click_cnt"
	ATTRIBUTION_RETENTION_7D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_7d_cnt"
	LUBAN_ORDER_STAT_AMOUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "luban_order_stat_amount"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_roi"
	COST_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "cost"
	DOWNLOAD_START_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_cost"
	DOWNLOAD_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "download"
	COUPON_SINGLE_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "coupon_single_page"
	PHONE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "phone"
	WECHAT_LOGIN_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_login_cost"
	CPC_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpc"
	INTERACT_PER_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "interact_per_cost"
	QQ_ReportAdGetV2OrderField                                          ReportAdGetV2OrderField = "qq"
	ACTIVE_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "active"
	ADVANCED_CREATIVE_FORM_CLICK_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "advanced_creative_form_click"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_7days"
	VIEW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "view"
	ATTRIBUTION_RETENTION_5D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_cost"
	LOTTERY_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "lottery"
	GAME_ADDICTION_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_rate"
	LUBAN_LIVE_GIFT_AMOUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_gift_amount"
	UNION_ROI_0_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_0"
	AVG_RANK_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "avg_rank"
	ATTRIBUTION_CONVERT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "attribution_convert"
	GAME_PAY_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_pay_count"
	REDIRECT_TO_SHOP_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "redirect_to_shop"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_roi_1day"
	CLICK_WEBSITE_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_website"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "luban_live_click_product_cnt"
	PRE_CONVERT_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "pre_convert_count"
	PRE_CONVERT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	CLICK_CALL_DY_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_call_dy"
	PLAY_100_FEED_BREAK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "play_100_feed_break"
	PLAY_OVER_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "play_over_rate"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_pay_30d_amount"
	LIVE_COMPONENT_CLICK_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "live_component_click_count"
	INSTALL_FINISH_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_rate"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_retention_7d_sum_cnt"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_wechat_pay_30d_roi"
	CLICK_DOWNLOAD_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "click_download"
	MESSAGE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "message"
	CTR_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "ctr"
	DEEP_CONVERT_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_cost"
	CLICK_INSTALL_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_install"
	LOAN_CREDIT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_rate"
	PLAY_50_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_50_feed_break"
	CONVERT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "convert"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportAdGetV2OrderField  ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	ATTRIBUTION_RETENTION_2D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_2d_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_7days"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "advanced_creative_coupon_addition"
	SUBMIT_CERTIFICATION_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "submit_certification_count"
	PLAY_75_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_75_feed_break"
	VALID_PLAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_2days"
	PLAY_DURATION_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "play_duration"
	SHARE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "share"
	IES_MUSIC_CLICK_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "ies_music_click"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_pay_order_stat_cost"
	ACTIVE_REGISTER_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_cost"
	CUSTOMER_EFFECTIVE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "customer_effective"
	LOAN_CREDIT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "loan_credit"
	LUBAN_ORDER_ROI_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_roi"
	PRE_LOAN_CREDIT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "pre_loan_credit"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportAdGetV2OrderField          ReportAdGetV2OrderField = "attribution_wechat_login_30d_count"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_ltv_1day"
	WECHAT_LOGIN_COUNT_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "wechat_login_count"
	DEEP_CONVERT_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "deep_convert"
	PHONE_CONFIRM_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_confirm"
	LUBAN_LIVE_COMMENT_CNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_comment_cnt"
	PLAY_50_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_50_feed_break_rate"
	ATTRIBUTION_RETENTION_7D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_cost"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_phone_click"
	ATTRIBUTION_RETENTION_4D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_4d_cnt"
	LOAN_COMPLETION_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_rate"
	DOWNLOAD_FINISH_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "download_finish"
	LOCATION_CLICK_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "location_click"
	NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_cost"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_ltv"
	DOWNLOAD_FINISH_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_cost"
	GAME_ADDICTION_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_cost"
	DISLIKE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "dislike"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	GAME_PAY_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "game_pay_cost"
	IN_APP_UV_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "in_app_uv"
	PLAY_DURATION_2S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_2s_rate"
	AVG_SHOW_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "avg_show_cost"
	BUTTON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "button"
	PLAY_25_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_25_feed_break"
	PLAY_DURATION_3S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_RETENTION_7D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_rate"
	WIFI_PLAY_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "wifi_play_rate"
	SHOW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "show"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportAdGetV2OrderField                 ReportAdGetV2OrderField = "live_watch_one_minute_count"
	VALID_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "valid_play"
	WIFI_PLAY_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "wifi_play"
	IN_APP_CART_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "in_app_cart"
	ATTRIBUTION_RETENTION_2D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_roi"
	SHOPPING_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "shopping"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_cost"
	LIVE_COMPONENT_CLICK_COST_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_cost"
	DOWNLOAD_START_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "download_start"
	WECHAT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "wechat"
	CPM_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpm"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_game_pay_7d_count"
	ATTRIBUTION_RETENTION_3D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_3d_cnt"
	CONVERT_COST_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_6days"
	REDIRECT_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "redirect"
	PLAY_DURATION_SUM_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_sum"
	PRE_CONVERT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_rate"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportAdGetV2OrderField ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_next_day_open_cnt"
	ACTIVE_COST_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_cost"
	LOAN_CREDIT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_cost"
	WECHAT_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_pay_amount"
	LUBAN_LIVE_ENTER_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_enter_cnt"
	COMMUTE_FIRST_PAY_COUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "commute_first_pay_count"
	LUBAN_LIVE_GIFT_CNT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "luban_live_gift_cnt"
	ATTRIBUTION_CONVERT_COST_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_convert_cost"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_deep_convert_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_4days"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "luban_live_pay_order_count"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_5days"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "average_play_time_per_play"
	PLAY_DURATION_2S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_2s"
	WECHAT_FIRST_PAY_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_rate"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportAdGetV2OrderField      ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	FIRST_ORDER_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "first_order_count"
	ACTIVE_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "active_pay_amount"
	LIKE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "like"
	PAY_COUNT_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "pay_count"
	FORM_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "form"
	COUPON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "coupon"
	NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_rate"
	REGISTER_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "register"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_8days"
	WECHAT_FIRST_PAY_COUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "wechat_first_pay_count"
	MESSAGE_ACTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "message_action"
	ATTRIBUTION_RETENTION_5D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_rate"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_ltv"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportAdGetV2OrderField    ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_active_pay_7d_per_count"
	PHONE_EFFECTIVE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "phone_effective"
	STAT_UNION_LTV_0_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_0"
	CONSULT_EFFECTIVE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "consult_effective"
	STAT_UNION_LTV_7_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_7"
	AVERAGE_PLAY_PROGRESS_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "average_play_progress"
	ATTRIBUTION_RETENTION_6D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_cost"
	STAT_UNION_LTV_3_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_3"
	LOAN_COMPLETION_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "loan_completion"
	POI_COLLECT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "poi_collect"
	INSTALL_FINISH_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_cost"
	ATTRIBUTION_DEEP_CONVERT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_deep_convert"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_game_pay_7d_cost"
	APPROVAL_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "approval_count"
)

// All allowed values of ReportAdGetV2OrderField enum
var AllowedReportAdGetV2OrderFieldEnumValues = []ReportAdGetV2OrderField{
	"advanced_creative_form_submit",
	"luban_live_share_cnt",
	"attribution_retention_2d_cost",
	"wechat_first_pay_cost",
	"click_shopwindow",
	"click_landing_page",
	"attribution_day_active_pay_count",
	"vote",
	"map_search",
	"ies_challenge_click",
	"poi_address_click",
	"deep_convert_rate",
	"union_roi_3",
	"attribution_wechat_first_pay_30d_cost",
	"average_video_play",
	"play_over",
	"in_app_detail_uv",
	"play_100_feed_break_rate",
	"convert_show_rate",
	"cpa",
	"attribution_game_in_app_roi_4days",
	"in_app_order",
	"stat_pay_amount",
	"attribution_next_day_open_rate",
	"play_duration_3s_rate",
	"attribution_game_in_app_ltv_2days",
	"play_75_feed_break_rate",
	"active_rate",
	"active_register_rate",
	"attribution_retention_5d_cnt",
	"total_play",
	"click",
	"attribution_game_in_app_ltv_3days",
	"valid_play_cost",
	"follow",
	"in_app_pay",
	"attribution_wechat_login_30d_cost",
	"install_finish",
	"pay_amount_roi",
	"attribution_game_in_app_roi_3days",
	"union_roi_7",
	"download_finish_rate",
	"attribution_game_in_app_ltv_8days",
	"play_duration_10s_rate",
	"attribution_retention_4d_cost",
	"attribution_game_in_app_ltv_5days",
	"attribution_day_acitve_pay_count",
	"download_start_rate",
	"play_25_feed_break_rate",
	"attribution_retention_7d_total_cost",
	"active_pay_rate",
	"luban_live_follow_cnt",
	"attribution_active_pay_7d_count",
	"attribution_micro_game_3d_ltv",
	"attribution_active_pay_7d_rate",
	"advanced_creative_counsel_click",
	"home_visited",
	"next_day_open",
	"attribution_game_in_app_ltv_6days",
	"first_rental_order_count",
	"game_addiction",
	"consult",
	"live_component_click_rate",
	"live_fans_club_join_cnt",
	"avg_click_cost",
	"attribution_retention_3d_cost",
	"report",
	"attribution_retention_6d_rate",
	"play_duration_5s_rate",
	"comment",
	"attribution_retention_4d_rate",
	"pre_loan_credit_cost",
	"luban_order_cnt",
	"play_duration_10s",
	"card_show",
	"convert_rate",
	"phone_connect",
	"loan_completion_cost",
	"attribution_retention_6d_cnt",
	"attribution_retention_3d_rate",
	"active_pay_cost",
	"attribution_micro_game_0d_roi",
	"luban_live_slidecart_click_cnt",
	"attribution_retention_7d_cnt",
	"luban_order_stat_amount",
	"attribution_micro_game_7d_roi",
	"cost",
	"download_start_cost",
	"download",
	"coupon_single_page",
	"phone",
	"wechat_login_cost",
	"cpc",
	"interact_per_cost",
	"qq",
	"active",
	"advanced_creative_form_click",
	"attribution_game_in_app_ltv_7days",
	"view",
	"attribution_retention_5d_cost",
	"attribution_active_pay_7d_cost",
	"lottery",
	"game_addiction_rate",
	"luban_live_gift_amount",
	"union_roi_0",
	"avg_rank",
	"attribution_convert",
	"game_pay_count",
	"redirect_to_shop",
	"attribution_game_in_app_roi_1day",
	"click_website",
	"luban_live_click_product_cnt",
	"pre_convert_count",
	"pre_convert_cost",
	"attribution_active_pay_intra_one_day_rate",
	"click_call_dy",
	"play_100_feed_break",
	"play_over_rate",
	"attribution_wechat_pay_30d_amount",
	"live_component_click_count",
	"install_finish_rate",
	"attribution_retention_7d_sum_cnt",
	"attribution_wechat_pay_30d_roi",
	"click_download",
	"message",
	"ctr",
	"deep_convert_cost",
	"click_install",
	"loan_credit_rate",
	"play_50_feed_break",
	"convert",
	"attribution_active_pay_intra_one_day_count",
	"attribution_retention_2d_cnt",
	"attribution_game_in_app_roi_7days",
	"advanced_creative_coupon_addition",
	"submit_certification_count",
	"play_75_feed_break",
	"valid_play_rate",
	"attribution_game_in_app_roi_2days",
	"play_duration",
	"share",
	"ies_music_click",
	"luban_live_pay_order_stat_cost",
	"active_register_cost",
	"customer_effective",
	"loan_credit",
	"luban_order_roi",
	"pre_loan_credit",
	"attribution_wechat_login_30d_count",
	"attribution_game_in_app_ltv_1day",
	"wechat_login_count",
	"deep_convert",
	"phone_confirm",
	"luban_live_comment_cnt",
	"play_50_feed_break_rate",
	"attribution_retention_7d_cost",
	"advanced_creative_phone_click",
	"attribution_retention_4d_cnt",
	"loan_completion_rate",
	"download_finish",
	"location_click",
	"next_day_open_cost",
	"attribution_micro_game_7d_ltv",
	"download_finish_cost",
	"game_addiction_cost",
	"dislike",
	"attribution_active_pay_intra_one_day_cost",
	"game_pay_cost",
	"in_app_uv",
	"play_duration_2s_rate",
	"avg_show_cost",
	"button",
	"play_25_feed_break",
	"play_duration_3s",
	"attribution_retention_7d_rate",
	"wifi_play_rate",
	"show",
	"live_watch_one_minute_count",
	"valid_play",
	"wifi_play",
	"in_app_cart",
	"attribution_retention_2d_rate",
	"attribution_micro_game_3d_roi",
	"shopping",
	"attribution_next_day_open_cost",
	"live_component_click_cost",
	"download_start",
	"wechat",
	"cpm",
	"attribution_game_pay_7d_count",
	"attribution_retention_3d_cnt",
	"convert_cost",
	"attribution_game_in_app_roi_6days",
	"redirect",
	"play_duration_sum",
	"pre_convert_rate",
	"attribution_active_pay_intra_one_day_amount",
	"attribution_next_day_open_cnt",
	"active_cost",
	"loan_credit_cost",
	"wechat_pay_amount",
	"luban_live_enter_cnt",
	"commute_first_pay_count",
	"luban_live_gift_cnt",
	"attribution_convert_cost",
	"attribution_deep_convert_cost",
	"attribution_game_in_app_ltv_4days",
	"luban_live_pay_order_count",
	"attribution_game_in_app_roi_5days",
	"average_play_time_per_play",
	"play_duration_2s",
	"wechat_first_pay_rate",
	"attribution_wechat_first_pay_30d_count",
	"first_order_count",
	"active_pay_amount",
	"like",
	"pay_count",
	"form",
	"coupon",
	"next_day_open_rate",
	"register",
	"attribution_game_in_app_roi_8days",
	"wechat_first_pay_count",
	"message_action",
	"attribution_retention_5d_rate",
	"attribution_wechat_first_pay_30d_rate",
	"attribution_micro_game_0d_ltv",
	"attribution_active_pay_intra_one_day_roi",
	"attribution_active_pay_7d_per_count",
	"phone_effective",
	"stat_union_ltv_0",
	"consult_effective",
	"stat_union_ltv_7",
	"average_play_progress",
	"attribution_retention_6d_cost",
	"stat_union_ltv_3",
	"loan_completion",
	"poi_collect",
	"install_finish_cost",
	"attribution_deep_convert",
	"attribution_game_pay_7d_cost",
	"approval_count",
}

// NewReportAdGetV2OrderFieldFromValue returns a pointer to a valid ReportAdGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2OrderFieldFromValue(v string) (*ReportAdGetV2OrderField, error) {
	ev := ReportAdGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2OrderField: valid values are %v", v, AllowedReportAdGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_order_field value
func (v ReportAdGetV2OrderField) Ptr() *ReportAdGetV2OrderField {
	return &v
}
