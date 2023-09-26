/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataExternalAction
type AdGetV2DataExternalAction string

// List of ad_get_v2_data_external_action
const (
	AD_CONVERT_TYPE_MESSAGE_INTERACTION_AdGetV2DataExternalAction         AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MESSAGE_INTERACTION"
	AD_CONVERT_TYPE_WECHAT_AdGetV2DataExternalAction                      AdGetV2DataExternalAction = "AD_CONVERT_TYPE_WECHAT"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_AdGetV2DataExternalAction    AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK_AdGetV2DataExternalAction        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	AD_CONVERT_TYPE_FIRST_RENTAL_ORDER_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FIRST_RENTAL_ORDER"
	AD_CONVERT_TYPE_SALES_LEAD_AdGetV2DataExternalAction                  AdGetV2DataExternalAction = "AD_CONVERT_TYPE_SALES_LEAD"
	AD_CONVERT_TYPE_REDIRECT_TO_STORE_AdGetV2DataExternalAction           AdGetV2DataExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_STORE"
	AD_CONVERT_TYPE_HIGH_VALUE_CLUE_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_HIGH_VALUE_CLUE"
	AD_CONVERT_TYPE_MESSAGE_SERVICE_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MESSAGE_SERVICE"
	AD_CONVERT_TYPE_PRE_LOAN_CREDIT_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PRE_LOAN_CREDIT"
	AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT_AdGetV2DataExternalAction        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT"
	AD_CONVERT_TYPE_SHOPPING_CART_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_SHOPPING_CART"
	AD_CONVERT_TYPE_LIVE_APPOINTMENT_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_APPOINTMENT"
	AD_CONVERT_TYPE_MAP_SEARCH_AdGetV2DataExternalAction                  AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MAP_SEARCH"
	AD_CONVERT_TYPE_REDIRECT_AdGetV2DataExternalAction                    AdGetV2DataExternalAction = "AD_CONVERT_TYPE_REDIRECT"
	AD_CONVERT_TYPE_PAY_AdGetV2DataExternalAction                         AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_FORM_ANSWER_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FORM_ANSWER"
	AD_CONVERT_TYPE_LIKE_ACTION_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIKE_ACTION"
	AD_CONVERT_TYPE_CLICK_WEBSITE_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLICK_WEBSITE"
	AD_CONVERT_TYPE_POI_MULTIPLE_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_POI_MULTIPLE"
	AD_CONVERT_TYPE_SUBMIT_CERTIFICATION_AdGetV2DataExternalAction        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_SUBMIT_CERTIFICATION"
	AD_CONVERT_TYPE_CONSULT_EFFECTIVE_AdGetV2DataExternalAction           AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CONSULT_EFFECTIVE"
	AD_CONVERT_TYPE_NATIVE_ACTION_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_NATIVE_ACTION"
	AD_CONVERT_TYPE_FOLLOW_ACTION_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QQ_AdGetV2DataExternalAction                          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_QQ"
	AD_CONVERT_TYPE_DEEP_PURCHASE_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_DEEP_PURCHASE"
	AD_CONVERT_PAGE_VIEW_AdGetV2DataExternalAction                        AdGetV2DataExternalAction = "AD_CONVERT_PAGE_VIEW"
	AD_CONVERT_TYPE_REDIRECT_TO_SHOP_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_SHOP"
	AD_CONVERT_TYPE_EFFECTIVE_COPY_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_COPY"
	AD_CONVERT_TYPE_AUTHORIZATION_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_AUTHORIZATION"
	AD_CONVERT_TYPE_LIVE_OTO_CLICK_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_CLICK"
	AD_CONVERT_TYPE_MESSAGE_CLUE_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLUE"
	AD_CONVERT_TYPE_LOAN_AdGetV2DataExternalAction                        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LOAN"
	AD_CONVERT_TYPE_FORM_CONNECT_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FORM_CONNECT"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_AdGetV2DataExternalAction           AdGetV2DataExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK"
	AD_CONVERT_TYPE_ARPU0_AdGetV2DataExternalAction                       AdGetV2DataExternalAction = "AD_CONVERT_TYPE_ARPU0"
	AD_CONVERT_TYPE_XPATH_AdGetV2DataExternalAction                       AdGetV2DataExternalAction = "AD_CONVERT_TYPE_XPATH"
	AD_CONVERT_TYPE_INSTALL_FINISH_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_INSTALL_FINISH"
	AD_CONVERT_TYPE_LT_ROI_AdGetV2DataExternalAction                      AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LT_ROI"
	AD_CONVERT_TYPE_CLICK_CALL_DY_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLICK_CALL_DY"
	AD_CONVERT_TYPE_MESSAGE_ACTION_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MESSAGE_ACTION"
	AD_CONVERT_TYPE_POI_COLLECT_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_POI_COLLECT"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_AdGetV2DataExternalAction           AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS_AdGetV2DataExternalAction       AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS"
	AD_CONVERT_TYPE_RETENTION_7_D_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_RETENTION_7D"
	AD_CONVERT_TYPE_WECHAT_PAY_AdGetV2DataExternalAction                  AdGetV2DataExternalAction = "AD_CONVERT_TYPE_WECHAT_PAY"
	AD_CONVERT_TYPE_PURCHASE_OF_GOODS_AdGetV2DataExternalAction           AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PURCHASE_OF_GOODS"
	AD_CONVERT_TYPE_CERTIFICATION_INFORMATION_AdGetV2DataExternalAction   AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CERTIFICATION_INFORMATION"
	AD_CONVERT_TYPE_RETENTION_DAYS_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_RETENTION_DAYS"
	AD_CONVERT_TYPE_APP_DETAIL_UV_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_APP_DETAIL_UV"
	AD_CONVERT_TYPE_LIVE_HOMEPAGE_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_CREATE_GAMEROLE_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CREATE_GAMEROLE"
	AD_CONVERT_TYPE_LOCATION_ACTION_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LOCATION_ACTION"
	AD_CONVERT_TYPE_SHOPPING_AdGetV2DataExternalAction                    AdGetV2DataExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_SHOW_OFF_NUM_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	AD_CONVERT_TYPE_CLICK_NUM_AdGetV2DataExternalAction                   AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLICK_NUM"
	AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON"
	AD_CONVERT_TYPE_DIALBACK_CONFIRM_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONFIRM"
	AD_CONVERT_TYPE_STAY_TIME_AdGetV2DataExternalAction                   AdGetV2DataExternalAction = "AD_CONVERT_TYPE_STAY_TIME"
	AD_CONVERT_TYPE_DIALBACK_CONNECT_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONNECT"
	AD_CONVERT_TYPE_COMMENT_ACTION_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_COMMENT_ACTION"
	AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE"
	AD_CONVERT_TYPE_DOWNLOAD_START_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_START"
	AD_CONVERT_TYPE_CLICK_LANDING_PAGE_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"
	AD_CONVERT_TYPE_QC_MUST_BUY_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_OTHER_AdGetV2DataExternalAction                       AdGetV2DataExternalAction = "AD_CONVERT_TYPE_OTHER"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_CLICK_DOWNLOAD_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLICK_DOWNLOAD"
	AD_CONVERT_TYPE_PHONE_CONFIRM_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PHONE_CONFIRM"
	AD_CONVERT_TYPE_APP_ORDER_AdGetV2DataExternalAction                   AdGetV2DataExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_NOTIFY_FORM_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_NOTIFY_FORM"
	AD_CONVERT_TYPE_CLUE_HIGH_INTENTION_AdGetV2DataExternalAction         AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLUE_HIGH_INTENTION"
	AD_CONVERT_TYPE_PURCHASE_ROI_7_D_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_7D"
	AD_CONVERT_TYPE_LIVE_FANS_ACTION_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_FANS_ACTION"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY"
	AD_CONVERT_TYPE_LOAN_COMPLETION_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LOAN_COMPLETION"
	AD_CONVERT_TYPE_LIVE_JOIN_GROUP_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
	AD_CONVERT_TYPE_CLUE_CONFIRM_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLUE_CONFIRM"
	AD_CONVERT_TYPE_MESSAGE_CLICK_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLICK"
	AD_CONVERT_TYPE_VIEW_AdGetV2DataExternalAction                        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_VIEW"
	AD_CONVERT_TYPE_CONSULT_CLUE_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CONSULT_CLUE"
	AD_CONVERT_TYPE_CLASS_AdGetV2DataExternalAction                       AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLASS"
	AD_CONVERT_TYPE_APP_UV_AdGetV2DataExternalAction                      AdGetV2DataExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_APPLET_CLICK_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_APPLET_CLICK"
	AD_CONVERT_TYPE_LOTTERY_AdGetV2DataExternalAction                     AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LOTTERY"
	AD_CONVERT_TYPE_MULTIPLE_AdGetV2DataExternalAction                    AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MULTIPLE"
	AD_CONVERT_TYPE_CLUE_INTERFLOW_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLUE_INTERFLOW"
	AD_CONVERT_TYPE_CONSULT_AdGetV2DataExternalAction                     AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CONSULT"
	AD_CONVERT_TYPE_LINK_ACTION_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LINK_ACTION"
	AD_CONVERT_TYPE_APP_PAY_AdGetV2DataExternalAction                     AdGetV2DataExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_INVALID_CLUE_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_INVALID_CLUE"
	AD_CONVERT_TYPE_ANCHOR_CLICK_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_ANCHOR_CLICK"
	AD_CONVERT_TYPE_GAME_ADDICTION_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_GAME_ADDICTION"
	AD_CONVERT_TYPE_APP_CART_AdGetV2DataExternalAction                    AdGetV2DataExternalAction = "AD_CONVERT_TYPE_APP_CART"
	AD_CONVERT_TYPE_PHONE_EFFECTIVE_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PHONE_EFFECTIVE"
	AD_CONVERT_TYPE_INTERACTION_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_INTERACTION"
	AD_CONVERT_TYPE_UG_ROI_AdGetV2DataExternalAction                      AdGetV2DataExternalAction = "AD_CONVERT_TYPE_UG_ROI"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_AdGetV2DataExternalAction   AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_GIFT_ACTION_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP"
	AD_CONVERT_TYPE_INTENTION_CLUE_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_INTENTION_CLUE"
	AD_CONVERT_TYPE_POI_ADDRESS_CLICK_AdGetV2DataExternalAction           AdGetV2DataExternalAction = "AD_CONVERT_TYPE_POI_ADDRESS_CLICK"
	AD_CONVERT_TYPE_PURCHASE_ROI_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI"
	AD_CONVERT_TYPE_RSS_AdGetV2DataExternalAction                         AdGetV2DataExternalAction = "AD_CONVERT_TYPE_RSS"
	AD_CONVERT_TYPE_FIRST_CLASS_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FIRST_CLASS"
	AD_CONVERT_TYPE_EFFECTIVE_PLAY_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_PLAY"
	AD_CONVERT_TYPE_LIVE_NATIVE_ACITON_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_NATIVE_ACITON"
	AD_CONVERT_TYPE_WECHAT_REGISTER_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_WECHAT_REGISTER"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_PREMIUM_ROI_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PREMIUM_ROI"
	AD_CONVERT_TYPE_PREMIUM_PAYMENT_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PREMIUM_PAYMENT"
	AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION_AdGetV2DataExternalAction AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION"
	AD_CONVERT_TYPE_PURCHASE_ROI_2_D_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_2D"
	AD_CONVERT_TYPE_FORM_DEEP_AdGetV2DataExternalAction                   AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FORM_DEEP"
	AD_CONVERT_TYPE_NEXT_DAY_OPEN_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN_AdGetV2DataExternalAction         AdGetV2DataExternalAction = "AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_ACTIVE_AdGetV2DataExternalAction                      AdGetV2DataExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_COMMODITY_CLICK_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_COMMODITY_CLICK"
	AD_CONVERT_TYPE_FORM_AdGetV2DataExternalAction                        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FORM"
	AD_CONVERT_TYPE_NOTIFY_DOWNLOAD_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	AD_CONVERT_TYPE_BANKCARD_INFORMATION_AdGetV2DataExternalAction        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_BANKCARD_INFORMATION"
	AD_CONVERT_TYPE_IPU_QUALIFY_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_IPU_QUALIFY"
	AD_CONVERT_TYPE_PREMIUM_UPGEADE_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PREMIUM_UPGEADE"
	AD_CONVERT_TYPE_WECHAT_WECOM_ADD_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_WECHAT_WECOM_ADD"
	AD_CONVERT_TYPE_ENTER_HOMEPAGE_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_ENTER_HOMEPAGE"
	AD_CONVERT_TYPE_SUCCESSORDER_ACTION_AdGetV2DataExternalAction         AdGetV2DataExternalAction = "AD_CONVERT_TYPE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_CLUE_PAY_SUCCEED_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLUE_PAY_SUCCEED"
	AD_CONVERT_TYPE_IDCARD_INFORMATION_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_IDCARD_INFORMATION"
	AD_CONVERT_TYPE_BUTTON_AdGetV2DataExternalAction                      AdGetV2DataExternalAction = "AD_CONVERT_TYPE_BUTTON"
	AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER_AdGetV2DataExternalAction           AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER"
	AD_CONVERT_TYPE_UPDATE_LEVEL_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_UPDATE_LEVEL"
	AD_CONVERT_TYPE_PERSONAL_INFORMATION_AdGetV2DataExternalAction        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PERSONAL_INFORMATION"
	AD_CONVERT_TYPE_DIALBACK_AdGetV2DataExternalAction                    AdGetV2DataExternalAction = "AD_CONVERT_TYPE_DIALBACK"
	AD_CONVERT_TYPE_SHARE_ACTION_AdGetV2DataExternalAction                AdGetV2DataExternalAction = "AD_CONVERT_TYPE_SHARE_ACTION"
	AD_CONVERT_TYPE_DOWNLOAD_FINISH_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_FINISH"
	AD_CONVERT_TYPE_WECHAT_USER_FIRST_MESSAGE_AdGetV2DataExternalAction   AdGetV2DataExternalAction = "AD_CONVERT_TYPE_WECHAT_USER_FIRST_MESSAGE"
	AD_CONVERT_TYPE_PHONE_AdGetV2DataExternalAction                       AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PHONE"
	AD_CONVERT_TYPE_CLICK_SHOPWINDOW_AdGetV2DataExternalAction            AdGetV2DataExternalAction = "AD_CONVERT_TYPE_CLICK_SHOPWINDOW"
	AD_CONVERT_TYPE_COUPON_AdGetV2DataExternalAction                      AdGetV2DataExternalAction = "AD_CONVERT_TYPE_COUPON"
	AD_CONVERT_TYPE_BOOST_AdGetV2DataExternalAction                       AdGetV2DataExternalAction = "AD_CONVERT_TYPE_BOOST"
	AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH_AdGetV2DataExternalAction    AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH"
	AD_CONVERT_TYPE_PAID_CLUE_AdGetV2DataExternalAction                   AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PAID_CLUE"
	AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE_AdGetV2DataExternalAction          AdGetV2DataExternalAction = "AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_AdGetV2DataExternalAction       AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_RESERVATION_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_RESERVATION"
	AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN_AdGetV2DataExternalAction    AdGetV2DataExternalAction = "AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN"
	AD_CONVERT_TYPE_SHOPPING_ACTION_AdGetV2DataExternalAction             AdGetV2DataExternalAction = "AD_CONVERT_TYPE_SHOPPING_ACTION"
	AD_CONVERT_TYPE_PHONE_CONNECT_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_PHONE_CONNECT"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_AdGetV2DataExternalAction         AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_OTO_STAY_TIME_AdGetV2DataExternalAction               AdGetV2DataExternalAction = "AD_CONVERT_TYPE_OTO_STAY_TIME"
	AD_CONVERT_TYPE_VOTE_AdGetV2DataExternalAction                        AdGetV2DataExternalAction = "AD_CONVERT_TYPE_VOTE"
	AD_CONVERT_TYPE_OTO_PAY_AdGetV2DataExternalAction                     AdGetV2DataExternalAction = "AD_CONVERT_TYPE_OTO_PAY"
	AD_CONVERT_TYPE_LIVE_STAY_TIME_AdGetV2DataExternalAction              AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	AD_CONVERT_TYPE_LOAN_CREDIT_AdGetV2DataExternalAction                 AdGetV2DataExternalAction = "AD_CONVERT_TYPE_LOAN_CREDIT"
	AD_CONVERT_TYPE_MESSAGE_AdGetV2DataExternalAction                     AdGetV2DataExternalAction = "AD_CONVERT_TYPE_MESSAGE"
)

// All allowed values of AdGetV2DataExternalAction enum
var AllowedAdGetV2DataExternalActionEnumValues = []AdGetV2DataExternalAction{
	"AD_CONVERT_TYPE_MESSAGE_INTERACTION",
	"AD_CONVERT_TYPE_WECHAT",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK",
	"AD_CONVERT_TYPE_FIRST_RENTAL_ORDER",
	"AD_CONVERT_TYPE_SALES_LEAD",
	"AD_CONVERT_TYPE_REDIRECT_TO_STORE",
	"AD_CONVERT_TYPE_HIGH_VALUE_CLUE",
	"AD_CONVERT_TYPE_MESSAGE_SERVICE",
	"AD_CONVERT_TYPE_PRE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT",
	"AD_CONVERT_TYPE_SHOPPING_CART",
	"AD_CONVERT_TYPE_LIVE_APPOINTMENT",
	"AD_CONVERT_TYPE_MAP_SEARCH",
	"AD_CONVERT_TYPE_REDIRECT",
	"AD_CONVERT_TYPE_PAY",
	"AD_CONVERT_TYPE_FORM_ANSWER",
	"AD_CONVERT_TYPE_LIKE_ACTION",
	"AD_CONVERT_TYPE_CLICK_WEBSITE",
	"AD_CONVERT_TYPE_POI_MULTIPLE",
	"AD_CONVERT_TYPE_SUBMIT_CERTIFICATION",
	"AD_CONVERT_TYPE_CONSULT_EFFECTIVE",
	"AD_CONVERT_TYPE_NATIVE_ACTION",
	"AD_CONVERT_TYPE_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QQ",
	"AD_CONVERT_TYPE_DEEP_PURCHASE",
	"AD_CONVERT_PAGE_VIEW",
	"AD_CONVERT_TYPE_REDIRECT_TO_SHOP",
	"AD_CONVERT_TYPE_EFFECTIVE_COPY",
	"AD_CONVERT_TYPE_AUTHORIZATION",
	"AD_CONVERT_TYPE_LIVE_OTO_CLICK",
	"AD_CONVERT_TYPE_MESSAGE_CLUE",
	"AD_CONVERT_TYPE_LOAN",
	"AD_CONVERT_TYPE_FORM_CONNECT",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK",
	"AD_CONVERT_TYPE_ARPU0",
	"AD_CONVERT_TYPE_XPATH",
	"AD_CONVERT_TYPE_INSTALL_FINISH",
	"AD_CONVERT_TYPE_LT_ROI",
	"AD_CONVERT_TYPE_CLICK_CALL_DY",
	"AD_CONVERT_TYPE_MESSAGE_ACTION",
	"AD_CONVERT_TYPE_POI_COLLECT",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS",
	"AD_CONVERT_TYPE_RETENTION_7D",
	"AD_CONVERT_TYPE_WECHAT_PAY",
	"AD_CONVERT_TYPE_PURCHASE_OF_GOODS",
	"AD_CONVERT_TYPE_CERTIFICATION_INFORMATION",
	"AD_CONVERT_TYPE_RETENTION_DAYS",
	"AD_CONVERT_TYPE_APP_DETAIL_UV",
	"AD_CONVERT_TYPE_LIVE_HOMEPAGE",
	"AD_CONVERT_TYPE_CREATE_GAMEROLE",
	"AD_CONVERT_TYPE_LOCATION_ACTION",
	"AD_CONVERT_TYPE_SHOPPING",
	"AD_CONVERT_TYPE_SHOW_OFF_NUM",
	"AD_CONVERT_TYPE_CLICK_NUM",
	"AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON",
	"AD_CONVERT_TYPE_DIALBACK_CONFIRM",
	"AD_CONVERT_TYPE_STAY_TIME",
	"AD_CONVERT_TYPE_DIALBACK_CONNECT",
	"AD_CONVERT_TYPE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE",
	"AD_CONVERT_TYPE_DOWNLOAD_START",
	"AD_CONVERT_TYPE_CLICK_LANDING_PAGE",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_OTHER",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_CLICK_DOWNLOAD",
	"AD_CONVERT_TYPE_PHONE_CONFIRM",
	"AD_CONVERT_TYPE_APP_ORDER",
	"AD_CONVERT_TYPE_NOTIFY_FORM",
	"AD_CONVERT_TYPE_CLUE_HIGH_INTENTION",
	"AD_CONVERT_TYPE_PURCHASE_ROI_7D",
	"AD_CONVERT_TYPE_LIVE_FANS_ACTION",
	"AD_CONVERT_TYPE_LIVE_OTO_PAY",
	"AD_CONVERT_TYPE_LOAN_COMPLETION",
	"AD_CONVERT_TYPE_LIVE_JOIN_GROUP",
	"AD_CONVERT_TYPE_CLUE_CONFIRM",
	"AD_CONVERT_TYPE_MESSAGE_CLICK",
	"AD_CONVERT_TYPE_VIEW",
	"AD_CONVERT_TYPE_CONSULT_CLUE",
	"AD_CONVERT_TYPE_CLASS",
	"AD_CONVERT_TYPE_APP_UV",
	"AD_CONVERT_TYPE_APPLET_CLICK",
	"AD_CONVERT_TYPE_LOTTERY",
	"AD_CONVERT_TYPE_MULTIPLE",
	"AD_CONVERT_TYPE_CLUE_INTERFLOW",
	"AD_CONVERT_TYPE_CONSULT",
	"AD_CONVERT_TYPE_LINK_ACTION",
	"AD_CONVERT_TYPE_APP_PAY",
	"AD_CONVERT_TYPE_INVALID_CLUE",
	"AD_CONVERT_TYPE_ANCHOR_CLICK",
	"AD_CONVERT_TYPE_GAME_ADDICTION",
	"AD_CONVERT_TYPE_APP_CART",
	"AD_CONVERT_TYPE_PHONE_EFFECTIVE",
	"AD_CONVERT_TYPE_INTERACTION",
	"AD_CONVERT_TYPE_UG_ROI",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_GIFT_ACTION",
	"AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP",
	"AD_CONVERT_TYPE_INTENTION_CLUE",
	"AD_CONVERT_TYPE_POI_ADDRESS_CLICK",
	"AD_CONVERT_TYPE_PURCHASE_ROI",
	"AD_CONVERT_TYPE_RSS",
	"AD_CONVERT_TYPE_FIRST_CLASS",
	"AD_CONVERT_TYPE_EFFECTIVE_PLAY",
	"AD_CONVERT_TYPE_LIVE_NATIVE_ACITON",
	"AD_CONVERT_TYPE_WECHAT_REGISTER",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_PREMIUM_ROI",
	"AD_CONVERT_TYPE_PREMIUM_PAYMENT",
	"AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION",
	"AD_CONVERT_TYPE_PURCHASE_ROI_2D",
	"AD_CONVERT_TYPE_FORM_DEEP",
	"AD_CONVERT_TYPE_NEXT_DAY_OPEN",
	"AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN",
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_COMMODITY_CLICK",
	"AD_CONVERT_TYPE_FORM",
	"AD_CONVERT_TYPE_NOTIFY_DOWNLOAD",
	"AD_CONVERT_TYPE_BANKCARD_INFORMATION",
	"AD_CONVERT_TYPE_IPU_QUALIFY",
	"AD_CONVERT_TYPE_PREMIUM_UPGEADE",
	"AD_CONVERT_TYPE_WECHAT_WECOM_ADD",
	"AD_CONVERT_TYPE_ENTER_HOMEPAGE",
	"AD_CONVERT_TYPE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_CLUE_PAY_SUCCEED",
	"AD_CONVERT_TYPE_IDCARD_INFORMATION",
	"AD_CONVERT_TYPE_BUTTON",
	"AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER",
	"AD_CONVERT_TYPE_UPDATE_LEVEL",
	"AD_CONVERT_TYPE_PERSONAL_INFORMATION",
	"AD_CONVERT_TYPE_DIALBACK",
	"AD_CONVERT_TYPE_SHARE_ACTION",
	"AD_CONVERT_TYPE_DOWNLOAD_FINISH",
	"AD_CONVERT_TYPE_WECHAT_USER_FIRST_MESSAGE",
	"AD_CONVERT_TYPE_PHONE",
	"AD_CONVERT_TYPE_CLICK_SHOPWINDOW",
	"AD_CONVERT_TYPE_COUPON",
	"AD_CONVERT_TYPE_BOOST",
	"AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH",
	"AD_CONVERT_TYPE_PAID_CLUE",
	"AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_RESERVATION",
	"AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN",
	"AD_CONVERT_TYPE_SHOPPING_ACTION",
	"AD_CONVERT_TYPE_PHONE_CONNECT",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_OTO_STAY_TIME",
	"AD_CONVERT_TYPE_VOTE",
	"AD_CONVERT_TYPE_OTO_PAY",
	"AD_CONVERT_TYPE_LIVE_STAY_TIME",
	"AD_CONVERT_TYPE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_MESSAGE",
}

// NewAdGetV2DataExternalActionFromValue returns a pointer to a valid AdGetV2DataExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataExternalActionFromValue(v string) (*AdGetV2DataExternalAction, error) {
	ev := AdGetV2DataExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataExternalAction: valid values are %v", v, AllowedAdGetV2DataExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataExternalAction) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_external_action value
func (v AdGetV2DataExternalAction) Ptr() *AdGetV2DataExternalAction {
	return &v
}
