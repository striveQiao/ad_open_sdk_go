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

// StardeliveryTaskDetailV30DataStarTaskExternalAction
type StardeliveryTaskDetailV30DataStarTaskExternalAction string

// List of stardelivery_task_detail_v3.0_data_star_task_external_action
const (
	AD_CONVERT_PAGE_VIEW_StardeliveryTaskDetailV30DataStarTaskExternalAction                        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_PAGE_VIEW"
	AD_CONVERT_PHONE_CONNECT_StardeliveryTaskDetailV30DataStarTaskExternalAction                    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_PHONE_CONNECT"
	AD_CONVERT_TYPE_ACTIVE_StardeliveryTaskDetailV30DataStarTaskExternalAction                      StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_ANCHOR_CLICK_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_ANCHOR_CLICK"
	AD_CONVERT_TYPE_APPLET_CLICK_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_APPLET_CLICK"
	AD_CONVERT_TYPE_APP_CART_StardeliveryTaskDetailV30DataStarTaskExternalAction                    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_APP_CART"
	AD_CONVERT_TYPE_APP_DETAIL_UV_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_APP_DETAIL_UV"
	AD_CONVERT_TYPE_APP_ORDER_StardeliveryTaskDetailV30DataStarTaskExternalAction                   StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_APP_PAY_StardeliveryTaskDetailV30DataStarTaskExternalAction                     StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_APP_UV_StardeliveryTaskDetailV30DataStarTaskExternalAction                      StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_ARPU0_StardeliveryTaskDetailV30DataStarTaskExternalAction                       StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_ARPU0"
	AD_CONVERT_TYPE_AUTHORIZATION_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_AUTHORIZATION"
	AD_CONVERT_TYPE_BANKCARD_INFORMATION_StardeliveryTaskDetailV30DataStarTaskExternalAction        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_BANKCARD_INFORMATION"
	AD_CONVERT_TYPE_BOOST_StardeliveryTaskDetailV30DataStarTaskExternalAction                       StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_BOOST"
	AD_CONVERT_TYPE_BUTTON_StardeliveryTaskDetailV30DataStarTaskExternalAction                      StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_BUTTON"
	AD_CONVERT_TYPE_CERTIFICATION_INFORMATION_StardeliveryTaskDetailV30DataStarTaskExternalAction   StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CERTIFICATION_INFORMATION"
	AD_CONVERT_TYPE_CLASS_StardeliveryTaskDetailV30DataStarTaskExternalAction                       StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLASS"
	AD_CONVERT_TYPE_CLICK_CALL_DY_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_CALL_DY"
	AD_CONVERT_TYPE_CLICK_DOWNLOAD_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_DOWNLOAD"
	AD_CONVERT_TYPE_CLICK_LANDING_PAGE_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"
	AD_CONVERT_TYPE_CLICK_NUM_StardeliveryTaskDetailV30DataStarTaskExternalAction                   StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_NUM"
	AD_CONVERT_TYPE_CLICK_SHOPWINDOW_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_SHOPWINDOW"
	AD_CONVERT_TYPE_CLICK_WEBSITE_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLICK_WEBSITE"
	AD_CONVERT_TYPE_CLUE_CONFIRM_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLUE_CONFIRM"
	AD_CONVERT_TYPE_CLUE_HIGH_INTENTION_StardeliveryTaskDetailV30DataStarTaskExternalAction         StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLUE_HIGH_INTENTION"
	AD_CONVERT_TYPE_CLUE_INTERFLOW_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLUE_INTERFLOW"
	AD_CONVERT_TYPE_CLUE_PAY_SUCCEED_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CLUE_PAY_SUCCEED"
	AD_CONVERT_TYPE_COMMENT_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_COMMENT_ACTION"
	AD_CONVERT_TYPE_COMMODITY_CLICK_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_COMMODITY_CLICK"
	AD_CONVERT_TYPE_CONSULT_StardeliveryTaskDetailV30DataStarTaskExternalAction                     StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CONSULT"
	AD_CONVERT_TYPE_CONSULT_CLUE_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CONSULT_CLUE"
	AD_CONVERT_TYPE_CONSULT_EFFECTIVE_StardeliveryTaskDetailV30DataStarTaskExternalAction           StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CONSULT_EFFECTIVE"
	AD_CONVERT_TYPE_COUPON_StardeliveryTaskDetailV30DataStarTaskExternalAction                      StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_COUPON"
	AD_CONVERT_TYPE_CREATE_GAMEROLE_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CREATE_GAMEROLE"
	AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE"
	AD_CONVERT_TYPE_DEEP_PURCHASE_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_DEEP_PURCHASE"
	AD_CONVERT_TYPE_DIALBACK_StardeliveryTaskDetailV30DataStarTaskExternalAction                    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_DIALBACK"
	AD_CONVERT_TYPE_DIALBACK_CONFIRM_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONFIRM"
	AD_CONVERT_TYPE_DIALBACK_CONNECT_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONNECT"
	AD_CONVERT_TYPE_DOWNLOAD_FINISH_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_FINISH"
	AD_CONVERT_TYPE_DOWNLOAD_START_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_START"
	AD_CONVERT_TYPE_EFFECTIVE_COPY_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_COPY"
	AD_CONVERT_TYPE_EFFECTIVE_PLAY_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_PLAY"
	AD_CONVERT_TYPE_ENTER_HOMEPAGE_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_ENTER_HOMEPAGE"
	AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_FIRST_CLASS_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FIRST_CLASS"
	AD_CONVERT_TYPE_FIRST_RENTAL_ORDER_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FIRST_RENTAL_ORDER"
	AD_CONVERT_TYPE_FOLLOW_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FOLLOW_ACTION"
	AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT_StardeliveryTaskDetailV30DataStarTaskExternalAction        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT"
	AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER_StardeliveryTaskDetailV30DataStarTaskExternalAction           StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER"
	AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH_StardeliveryTaskDetailV30DataStarTaskExternalAction    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH"
	AD_CONVERT_TYPE_FORM_StardeliveryTaskDetailV30DataStarTaskExternalAction                        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FORM"
	AD_CONVERT_TYPE_FORM_ANSWER_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FORM_ANSWER"
	AD_CONVERT_TYPE_FORM_CONNECT_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FORM_CONNECT"
	AD_CONVERT_TYPE_FORM_DEEP_StardeliveryTaskDetailV30DataStarTaskExternalAction                   StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FORM_DEEP"
	AD_CONVERT_TYPE_FORTUNE_SMALL_DEAL_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_FORTUNE_SMALL_DEAL"
	AD_CONVERT_TYPE_GAME_ADDICTION_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_GAME_ADDICTION"
	AD_CONVERT_TYPE_HIGH_VALUE_CLUE_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_HIGH_VALUE_CLUE"
	AD_CONVERT_TYPE_IDCARD_INFORMATION_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_IDCARD_INFORMATION"
	AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN_StardeliveryTaskDetailV30DataStarTaskExternalAction         StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_INSTALL_FINISH_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_INSTALL_FINISH"
	AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN_StardeliveryTaskDetailV30DataStarTaskExternalAction    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN"
	AD_CONVERT_TYPE_INTENTION_CLUE_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_INTENTION_CLUE"
	AD_CONVERT_TYPE_INTERACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_INTERACTION"
	AD_CONVERT_TYPE_INVALID_CLUE_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_INVALID_CLUE"
	AD_CONVERT_TYPE_IPU_QUALIFY_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_IPU_QUALIFY"
	AD_CONVERT_TYPE_LIKE_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIKE_ACTION"
	AD_CONVERT_TYPE_LINK_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LINK_ACTION"
	AD_CONVERT_TYPE_LIVE_APPOINTMENT_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_APPOINTMENT"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction   StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction         StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK_StardeliveryTaskDetailV30DataStarTaskExternalAction        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction           StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_FANS_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_FANS_ACTION"
	AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON"
	AD_CONVERT_TYPE_LIVE_GIFT_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	AD_CONVERT_TYPE_LIVE_HOMEPAGE_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_LIVE_JOIN_GROUP_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
	AD_CONVERT_TYPE_LIVE_NATIVE_ACITON_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_NATIVE_ACITON"
	AD_CONVERT_TYPE_LIVE_OTO_CLICK_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_CLICK"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK"
	AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION"
	AD_CONVERT_TYPE_LIVE_STAY_TIME_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_StardeliveryTaskDetailV30DataStarTaskExternalAction       StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LOAN_StardeliveryTaskDetailV30DataStarTaskExternalAction                        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LOAN"
	AD_CONVERT_TYPE_LOAN_COMPLETION_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LOAN_COMPLETION"
	AD_CONVERT_TYPE_LOAN_CREDIT_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LOAN_CREDIT"
	AD_CONVERT_TYPE_LOCATION_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LOCATION_ACTION"
	AD_CONVERT_TYPE_LOTTERY_StardeliveryTaskDetailV30DataStarTaskExternalAction                     StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LOTTERY"
	AD_CONVERT_TYPE_LT_ROI_StardeliveryTaskDetailV30DataStarTaskExternalAction                      StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_LT_ROI"
	AD_CONVERT_TYPE_MAP_SEARCH_StardeliveryTaskDetailV30DataStarTaskExternalAction                  StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MAP_SEARCH"
	AD_CONVERT_TYPE_MESSAGE_StardeliveryTaskDetailV30DataStarTaskExternalAction                     StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE"
	AD_CONVERT_TYPE_MESSAGE_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_ACTION"
	AD_CONVERT_TYPE_MESSAGE_CLICK_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLICK"
	AD_CONVERT_TYPE_MESSAGE_CLUE_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLUE"
	AD_CONVERT_TYPE_MESSAGE_INTERACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction         StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_INTERACTION"
	AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP"
	AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS_StardeliveryTaskDetailV30DataStarTaskExternalAction       StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS"
	AD_CONVERT_TYPE_MESSAGE_SERVICE_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MESSAGE_SERVICE"
	AD_CONVERT_TYPE_MULTIPLE_StardeliveryTaskDetailV30DataStarTaskExternalAction                    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_MULTIPLE"
	AD_CONVERT_TYPE_NATIVE_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_NATIVE_ACTION"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction           StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_NEXT_DAY_OPEN_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_NOTIFY_DOWNLOAD_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	AD_CONVERT_TYPE_NOTIFY_FORM_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_NOTIFY_FORM"
	AD_CONVERT_TYPE_OTHER_StardeliveryTaskDetailV30DataStarTaskExternalAction                       StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_OTHER"
	AD_CONVERT_TYPE_OTO_PAY_StardeliveryTaskDetailV30DataStarTaskExternalAction                     StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_OTO_PAY"
	AD_CONVERT_TYPE_OTO_STAY_TIME_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_OTO_STAY_TIME"
	AD_CONVERT_TYPE_PAID_CLUE_StardeliveryTaskDetailV30DataStarTaskExternalAction                   StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PAID_CLUE"
	AD_CONVERT_TYPE_PAY_StardeliveryTaskDetailV30DataStarTaskExternalAction                         StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_PERSONAL_INFORMATION_StardeliveryTaskDetailV30DataStarTaskExternalAction        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PERSONAL_INFORMATION"
	AD_CONVERT_TYPE_PHONE_StardeliveryTaskDetailV30DataStarTaskExternalAction                       StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PHONE"
	AD_CONVERT_TYPE_PHONE_CONFIRM_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PHONE_CONFIRM"
	AD_CONVERT_TYPE_PHONE_CONNECT_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PHONE_CONNECT"
	AD_CONVERT_TYPE_PHONE_EFFECTIVE_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PHONE_EFFECTIVE"
	AD_CONVERT_TYPE_POI_ADDRESS_CLICK_StardeliveryTaskDetailV30DataStarTaskExternalAction           StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_POI_ADDRESS_CLICK"
	AD_CONVERT_TYPE_POI_COLLECT_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_POI_COLLECT"
	AD_CONVERT_TYPE_POI_MULTIPLE_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_POI_MULTIPLE"
	AD_CONVERT_TYPE_PREMIUM_PAYMENT_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PREMIUM_PAYMENT"
	AD_CONVERT_TYPE_PREMIUM_ROI_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PREMIUM_ROI"
	AD_CONVERT_TYPE_PREMIUM_UPGEADE_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PREMIUM_UPGEADE"
	AD_CONVERT_TYPE_PRE_LOAN_CREDIT_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PRE_LOAN_CREDIT"
	AD_CONVERT_TYPE_PURCHASE_OF_GOODS_StardeliveryTaskDetailV30DataStarTaskExternalAction           StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PURCHASE_OF_GOODS"
	AD_CONVERT_TYPE_PURCHASE_ROI_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI"
	AD_CONVERT_TYPE_PURCHASE_ROI_2_D_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_2D"
	AD_CONVERT_TYPE_PURCHASE_ROI_7_D_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_7D"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_QQ_StardeliveryTaskDetailV30DataStarTaskExternalAction                          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_QQ"
	AD_CONVERT_TYPE_REDIRECT_StardeliveryTaskDetailV30DataStarTaskExternalAction                    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_REDIRECT"
	AD_CONVERT_TYPE_REDIRECT_TO_SHOP_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_SHOP"
	AD_CONVERT_TYPE_REDIRECT_TO_STORE_StardeliveryTaskDetailV30DataStarTaskExternalAction           StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_STORE"
	AD_CONVERT_TYPE_RESERVATION_StardeliveryTaskDetailV30DataStarTaskExternalAction                 StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_RESERVATION"
	AD_CONVERT_TYPE_RETENTION_7_D_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_RETENTION_7D"
	AD_CONVERT_TYPE_RETENTION_DAYS_StardeliveryTaskDetailV30DataStarTaskExternalAction              StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_RETENTION_DAYS"
	AD_CONVERT_TYPE_RSS_StardeliveryTaskDetailV30DataStarTaskExternalAction                         StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_RSS"
	AD_CONVERT_TYPE_SALES_LEAD_StardeliveryTaskDetailV30DataStarTaskExternalAction                  StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SALES_LEAD"
	AD_CONVERT_TYPE_SERVICE_WITHDRAW_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SERVICE_WITHDRAW"
	AD_CONVERT_TYPE_SHARE_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SHARE_ACTION"
	AD_CONVERT_TYPE_SHOPPING_StardeliveryTaskDetailV30DataStarTaskExternalAction                    StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_SHOPPING_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SHOPPING_ACTION"
	AD_CONVERT_TYPE_SHOPPING_CART_StardeliveryTaskDetailV30DataStarTaskExternalAction               StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SHOPPING_CART"
	AD_CONVERT_TYPE_SHOW_OFF_NUM_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	AD_CONVERT_TYPE_STAY_TIME_StardeliveryTaskDetailV30DataStarTaskExternalAction                   StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_STAY_TIME"
	AD_CONVERT_TYPE_SUBMIT_CERTIFICATION_StardeliveryTaskDetailV30DataStarTaskExternalAction        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SUBMIT_CERTIFICATION"
	AD_CONVERT_TYPE_SUCCESSORDER_ACTION_StardeliveryTaskDetailV30DataStarTaskExternalAction         StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_UG_ROI_StardeliveryTaskDetailV30DataStarTaskExternalAction                      StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_UG_ROI"
	AD_CONVERT_TYPE_UPDATE_LEVEL_StardeliveryTaskDetailV30DataStarTaskExternalAction                StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_UPDATE_LEVEL"
	AD_CONVERT_TYPE_VIEW_StardeliveryTaskDetailV30DataStarTaskExternalAction                        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_VIEW"
	AD_CONVERT_TYPE_VOTE_StardeliveryTaskDetailV30DataStarTaskExternalAction                        StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_VOTE"
	AD_CONVERT_TYPE_WECHAT_StardeliveryTaskDetailV30DataStarTaskExternalAction                      StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT"
	AD_CONVERT_TYPE_WECHAT_FIRST_MSG_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_FIRST_MSG"
	AD_CONVERT_TYPE_WECHAT_PAY_StardeliveryTaskDetailV30DataStarTaskExternalAction                  StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_PAY"
	AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW_StardeliveryTaskDetailV30DataStarTaskExternalAction          StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW"
	AD_CONVERT_TYPE_WECHAT_QRCODE_TRY_StardeliveryTaskDetailV30DataStarTaskExternalAction           StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_QRCODE_TRY"
	AD_CONVERT_TYPE_WECHAT_REGISTER_StardeliveryTaskDetailV30DataStarTaskExternalAction             StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_REGISTER"
	AD_CONVERT_TYPE_WECHAT_WECOM_ADD_StardeliveryTaskDetailV30DataStarTaskExternalAction            StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_WECHAT_WECOM_ADD"
	AD_CONVERT_TYPE_XPATH_StardeliveryTaskDetailV30DataStarTaskExternalAction                       StardeliveryTaskDetailV30DataStarTaskExternalAction = "AD_CONVERT_TYPE_XPATH"
)

// All allowed values of StardeliveryTaskDetailV30DataStarTaskExternalAction enum
var AllowedStardeliveryTaskDetailV30DataStarTaskExternalActionEnumValues = []StardeliveryTaskDetailV30DataStarTaskExternalAction{
	"AD_CONVERT_PAGE_VIEW",
	"AD_CONVERT_PHONE_CONNECT",
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_ANCHOR_CLICK",
	"AD_CONVERT_TYPE_APPLET_CLICK",
	"AD_CONVERT_TYPE_APP_CART",
	"AD_CONVERT_TYPE_APP_DETAIL_UV",
	"AD_CONVERT_TYPE_APP_ORDER",
	"AD_CONVERT_TYPE_APP_PAY",
	"AD_CONVERT_TYPE_APP_UV",
	"AD_CONVERT_TYPE_ARPU0",
	"AD_CONVERT_TYPE_AUTHORIZATION",
	"AD_CONVERT_TYPE_BANKCARD_INFORMATION",
	"AD_CONVERT_TYPE_BOOST",
	"AD_CONVERT_TYPE_BUTTON",
	"AD_CONVERT_TYPE_CERTIFICATION_INFORMATION",
	"AD_CONVERT_TYPE_CLASS",
	"AD_CONVERT_TYPE_CLICK_CALL_DY",
	"AD_CONVERT_TYPE_CLICK_DOWNLOAD",
	"AD_CONVERT_TYPE_CLICK_LANDING_PAGE",
	"AD_CONVERT_TYPE_CLICK_NUM",
	"AD_CONVERT_TYPE_CLICK_SHOPWINDOW",
	"AD_CONVERT_TYPE_CLICK_WEBSITE",
	"AD_CONVERT_TYPE_CLUE_CONFIRM",
	"AD_CONVERT_TYPE_CLUE_HIGH_INTENTION",
	"AD_CONVERT_TYPE_CLUE_INTERFLOW",
	"AD_CONVERT_TYPE_CLUE_PAY_SUCCEED",
	"AD_CONVERT_TYPE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_COMMODITY_CLICK",
	"AD_CONVERT_TYPE_CONSULT",
	"AD_CONVERT_TYPE_CONSULT_CLUE",
	"AD_CONVERT_TYPE_CONSULT_EFFECTIVE",
	"AD_CONVERT_TYPE_COUPON",
	"AD_CONVERT_TYPE_CREATE_GAMEROLE",
	"AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE",
	"AD_CONVERT_TYPE_DEEP_PURCHASE",
	"AD_CONVERT_TYPE_DIALBACK",
	"AD_CONVERT_TYPE_DIALBACK_CONFIRM",
	"AD_CONVERT_TYPE_DIALBACK_CONNECT",
	"AD_CONVERT_TYPE_DOWNLOAD_FINISH",
	"AD_CONVERT_TYPE_DOWNLOAD_START",
	"AD_CONVERT_TYPE_EFFECTIVE_COPY",
	"AD_CONVERT_TYPE_EFFECTIVE_PLAY",
	"AD_CONVERT_TYPE_ENTER_HOMEPAGE",
	"AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE",
	"AD_CONVERT_TYPE_FIRST_CLASS",
	"AD_CONVERT_TYPE_FIRST_RENTAL_ORDER",
	"AD_CONVERT_TYPE_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT",
	"AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER",
	"AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH",
	"AD_CONVERT_TYPE_FORM",
	"AD_CONVERT_TYPE_FORM_ANSWER",
	"AD_CONVERT_TYPE_FORM_CONNECT",
	"AD_CONVERT_TYPE_FORM_DEEP",
	"AD_CONVERT_TYPE_FORTUNE_SMALL_DEAL",
	"AD_CONVERT_TYPE_GAME_ADDICTION",
	"AD_CONVERT_TYPE_HIGH_VALUE_CLUE",
	"AD_CONVERT_TYPE_IDCARD_INFORMATION",
	"AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN",
	"AD_CONVERT_TYPE_INSTALL_FINISH",
	"AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN",
	"AD_CONVERT_TYPE_INTENTION_CLUE",
	"AD_CONVERT_TYPE_INTERACTION",
	"AD_CONVERT_TYPE_INVALID_CLUE",
	"AD_CONVERT_TYPE_IPU_QUALIFY",
	"AD_CONVERT_TYPE_LIKE_ACTION",
	"AD_CONVERT_TYPE_LINK_ACTION",
	"AD_CONVERT_TYPE_LIVE_APPOINTMENT",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_FANS_ACTION",
	"AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON",
	"AD_CONVERT_TYPE_LIVE_GIFT_ACTION",
	"AD_CONVERT_TYPE_LIVE_HOMEPAGE",
	"AD_CONVERT_TYPE_LIVE_JOIN_GROUP",
	"AD_CONVERT_TYPE_LIVE_NATIVE_ACITON",
	"AD_CONVERT_TYPE_LIVE_OTO_CLICK",
	"AD_CONVERT_TYPE_LIVE_OTO_PAY",
	"AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK",
	"AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION",
	"AD_CONVERT_TYPE_LIVE_STAY_TIME",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_LOAN",
	"AD_CONVERT_TYPE_LOAN_COMPLETION",
	"AD_CONVERT_TYPE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_LOCATION_ACTION",
	"AD_CONVERT_TYPE_LOTTERY",
	"AD_CONVERT_TYPE_LT_ROI",
	"AD_CONVERT_TYPE_MAP_SEARCH",
	"AD_CONVERT_TYPE_MESSAGE",
	"AD_CONVERT_TYPE_MESSAGE_ACTION",
	"AD_CONVERT_TYPE_MESSAGE_CLICK",
	"AD_CONVERT_TYPE_MESSAGE_CLUE",
	"AD_CONVERT_TYPE_MESSAGE_INTERACTION",
	"AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP",
	"AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS",
	"AD_CONVERT_TYPE_MESSAGE_SERVICE",
	"AD_CONVERT_TYPE_MULTIPLE",
	"AD_CONVERT_TYPE_NATIVE_ACTION",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_NEXT_DAY_OPEN",
	"AD_CONVERT_TYPE_NOTIFY_DOWNLOAD",
	"AD_CONVERT_TYPE_NOTIFY_FORM",
	"AD_CONVERT_TYPE_OTHER",
	"AD_CONVERT_TYPE_OTO_PAY",
	"AD_CONVERT_TYPE_OTO_STAY_TIME",
	"AD_CONVERT_TYPE_PAID_CLUE",
	"AD_CONVERT_TYPE_PAY",
	"AD_CONVERT_TYPE_PERSONAL_INFORMATION",
	"AD_CONVERT_TYPE_PHONE",
	"AD_CONVERT_TYPE_PHONE_CONFIRM",
	"AD_CONVERT_TYPE_PHONE_CONNECT",
	"AD_CONVERT_TYPE_PHONE_EFFECTIVE",
	"AD_CONVERT_TYPE_POI_ADDRESS_CLICK",
	"AD_CONVERT_TYPE_POI_COLLECT",
	"AD_CONVERT_TYPE_POI_MULTIPLE",
	"AD_CONVERT_TYPE_PREMIUM_PAYMENT",
	"AD_CONVERT_TYPE_PREMIUM_ROI",
	"AD_CONVERT_TYPE_PREMIUM_UPGEADE",
	"AD_CONVERT_TYPE_PRE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_PURCHASE_OF_GOODS",
	"AD_CONVERT_TYPE_PURCHASE_ROI",
	"AD_CONVERT_TYPE_PURCHASE_ROI_2D",
	"AD_CONVERT_TYPE_PURCHASE_ROI_7D",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_QQ",
	"AD_CONVERT_TYPE_REDIRECT",
	"AD_CONVERT_TYPE_REDIRECT_TO_SHOP",
	"AD_CONVERT_TYPE_REDIRECT_TO_STORE",
	"AD_CONVERT_TYPE_RESERVATION",
	"AD_CONVERT_TYPE_RETENTION_7D",
	"AD_CONVERT_TYPE_RETENTION_DAYS",
	"AD_CONVERT_TYPE_RSS",
	"AD_CONVERT_TYPE_SALES_LEAD",
	"AD_CONVERT_TYPE_SERVICE_WITHDRAW",
	"AD_CONVERT_TYPE_SHARE_ACTION",
	"AD_CONVERT_TYPE_SHOPPING",
	"AD_CONVERT_TYPE_SHOPPING_ACTION",
	"AD_CONVERT_TYPE_SHOPPING_CART",
	"AD_CONVERT_TYPE_SHOW_OFF_NUM",
	"AD_CONVERT_TYPE_STAY_TIME",
	"AD_CONVERT_TYPE_SUBMIT_CERTIFICATION",
	"AD_CONVERT_TYPE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_UG_ROI",
	"AD_CONVERT_TYPE_UPDATE_LEVEL",
	"AD_CONVERT_TYPE_VIEW",
	"AD_CONVERT_TYPE_VOTE",
	"AD_CONVERT_TYPE_WECHAT",
	"AD_CONVERT_TYPE_WECHAT_FIRST_MSG",
	"AD_CONVERT_TYPE_WECHAT_PAY",
	"AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW",
	"AD_CONVERT_TYPE_WECHAT_QRCODE_TRY",
	"AD_CONVERT_TYPE_WECHAT_REGISTER",
	"AD_CONVERT_TYPE_WECHAT_WECOM_ADD",
	"AD_CONVERT_TYPE_XPATH",
}

// NewStardeliveryTaskDetailV30DataStarTaskExternalActionFromValue returns a pointer to a valid StardeliveryTaskDetailV30DataStarTaskExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskDetailV30DataStarTaskExternalActionFromValue(v string) (*StardeliveryTaskDetailV30DataStarTaskExternalAction, error) {
	ev := StardeliveryTaskDetailV30DataStarTaskExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskDetailV30DataStarTaskExternalAction: valid values are %v", v, AllowedStardeliveryTaskDetailV30DataStarTaskExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskDetailV30DataStarTaskExternalAction) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskDetailV30DataStarTaskExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_detail_v3.0_data_star_task_external_action value
func (v StardeliveryTaskDetailV30DataStarTaskExternalAction) Ptr() *StardeliveryTaskDetailV30DataStarTaskExternalAction {
	return &v
}
