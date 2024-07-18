/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetV30FilteringExternalAction
type NativeAnchorGetV30FilteringExternalAction string

// List of native_anchor_get_v3.0_filtering_external_action
const (
	AD_CONVERT_PAGE_VIEW_NativeAnchorGetV30FilteringExternalAction                        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_PAGE_VIEW"
	AD_CONVERT_TYPE_ACTIVE_NativeAnchorGetV30FilteringExternalAction                      NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_ANCHOR_CLICK_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_ANCHOR_CLICK"
	AD_CONVERT_TYPE_APPLET_CLICK_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_APPLET_CLICK"
	AD_CONVERT_TYPE_APP_CART_NativeAnchorGetV30FilteringExternalAction                    NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_APP_CART"
	AD_CONVERT_TYPE_APP_DETAIL_UV_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_APP_DETAIL_UV"
	AD_CONVERT_TYPE_APP_ORDER_NativeAnchorGetV30FilteringExternalAction                   NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_APP_PAY_NativeAnchorGetV30FilteringExternalAction                     NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_APP_UV_NativeAnchorGetV30FilteringExternalAction                      NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_ARPU0_NativeAnchorGetV30FilteringExternalAction                       NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_ARPU0"
	AD_CONVERT_TYPE_AUTHORIZATION_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_AUTHORIZATION"
	AD_CONVERT_TYPE_BANKCARD_INFORMATION_NativeAnchorGetV30FilteringExternalAction        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_BANKCARD_INFORMATION"
	AD_CONVERT_TYPE_BOOST_NativeAnchorGetV30FilteringExternalAction                       NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_BOOST"
	AD_CONVERT_TYPE_BUTTON_NativeAnchorGetV30FilteringExternalAction                      NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_BUTTON"
	AD_CONVERT_TYPE_CERTIFICATION_INFORMATION_NativeAnchorGetV30FilteringExternalAction   NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CERTIFICATION_INFORMATION"
	AD_CONVERT_TYPE_CLASS_NativeAnchorGetV30FilteringExternalAction                       NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLASS"
	AD_CONVERT_TYPE_CLICK_CALL_DY_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLICK_CALL_DY"
	AD_CONVERT_TYPE_CLICK_DOWNLOAD_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLICK_DOWNLOAD"
	AD_CONVERT_TYPE_CLICK_LANDING_PAGE_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"
	AD_CONVERT_TYPE_CLICK_NUM_NativeAnchorGetV30FilteringExternalAction                   NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLICK_NUM"
	AD_CONVERT_TYPE_CLICK_SHOPWINDOW_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLICK_SHOPWINDOW"
	AD_CONVERT_TYPE_CLICK_WEBSITE_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLICK_WEBSITE"
	AD_CONVERT_TYPE_CLUE_CONFIRM_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLUE_CONFIRM"
	AD_CONVERT_TYPE_CLUE_HIGH_INTENTION_NativeAnchorGetV30FilteringExternalAction         NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLUE_HIGH_INTENTION"
	AD_CONVERT_TYPE_CLUE_INTERFLOW_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLUE_INTERFLOW"
	AD_CONVERT_TYPE_CLUE_PAY_SUCCEED_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CLUE_PAY_SUCCEED"
	AD_CONVERT_TYPE_COMMENT_ACTION_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_COMMENT_ACTION"
	AD_CONVERT_TYPE_COMMODITY_CLICK_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_COMMODITY_CLICK"
	AD_CONVERT_TYPE_CONSULT_NativeAnchorGetV30FilteringExternalAction                     NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CONSULT"
	AD_CONVERT_TYPE_CONSULT_CLUE_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CONSULT_CLUE"
	AD_CONVERT_TYPE_CONSULT_EFFECTIVE_NativeAnchorGetV30FilteringExternalAction           NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CONSULT_EFFECTIVE"
	AD_CONVERT_TYPE_COUPON_NativeAnchorGetV30FilteringExternalAction                      NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_COUPON"
	AD_CONVERT_TYPE_CREATE_GAMEROLE_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CREATE_GAMEROLE"
	AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE"
	AD_CONVERT_TYPE_DEEP_PURCHASE_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_DEEP_PURCHASE"
	AD_CONVERT_TYPE_DIALBACK_NativeAnchorGetV30FilteringExternalAction                    NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_DIALBACK"
	AD_CONVERT_TYPE_DIALBACK_CONFIRM_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONFIRM"
	AD_CONVERT_TYPE_DIALBACK_CONNECT_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONNECT"
	AD_CONVERT_TYPE_DOWNLOAD_FINISH_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_FINISH"
	AD_CONVERT_TYPE_DOWNLOAD_START_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_START"
	AD_CONVERT_TYPE_EFFECTIVE_COPY_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_COPY"
	AD_CONVERT_TYPE_EFFECTIVE_PLAY_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_PLAY"
	AD_CONVERT_TYPE_ENTER_HOMEPAGE_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_ENTER_HOMEPAGE"
	AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_FIRST_CLASS_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FIRST_CLASS"
	AD_CONVERT_TYPE_FIRST_RENTAL_ORDER_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FIRST_RENTAL_ORDER"
	AD_CONVERT_TYPE_FOLLOW_ACTION_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FOLLOW_ACTION"
	AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT_NativeAnchorGetV30FilteringExternalAction        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT"
	AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER_NativeAnchorGetV30FilteringExternalAction           NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER"
	AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH_NativeAnchorGetV30FilteringExternalAction    NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH"
	AD_CONVERT_TYPE_FORM_NativeAnchorGetV30FilteringExternalAction                        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FORM"
	AD_CONVERT_TYPE_FORM_ANSWER_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FORM_ANSWER"
	AD_CONVERT_TYPE_FORM_CONNECT_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FORM_CONNECT"
	AD_CONVERT_TYPE_FORM_DEEP_NativeAnchorGetV30FilteringExternalAction                   NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_FORM_DEEP"
	AD_CONVERT_TYPE_GAME_ADDICTION_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_GAME_ADDICTION"
	AD_CONVERT_TYPE_HIGH_VALUE_CLUE_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_HIGH_VALUE_CLUE"
	AD_CONVERT_TYPE_IDCARD_INFORMATION_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_IDCARD_INFORMATION"
	AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN_NativeAnchorGetV30FilteringExternalAction         NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_INSTALL_FINISH_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_INSTALL_FINISH"
	AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN_NativeAnchorGetV30FilteringExternalAction    NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN"
	AD_CONVERT_TYPE_INTENTION_CLUE_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_INTENTION_CLUE"
	AD_CONVERT_TYPE_INTERACTION_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_INTERACTION"
	AD_CONVERT_TYPE_INVALID_CLUE_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_INVALID_CLUE"
	AD_CONVERT_TYPE_IPU_QUALIFY_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_IPU_QUALIFY"
	AD_CONVERT_TYPE_LIKE_ACTION_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIKE_ACTION"
	AD_CONVERT_TYPE_LINK_ACTION_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LINK_ACTION"
	AD_CONVERT_TYPE_LIVE_APPOINTMENT_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_APPOINTMENT"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_NativeAnchorGetV30FilteringExternalAction   NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_NativeAnchorGetV30FilteringExternalAction         NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK_NativeAnchorGetV30FilteringExternalAction        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_NativeAnchorGetV30FilteringExternalAction           NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_FANS_ACTION_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_FANS_ACTION"
	AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON"
	AD_CONVERT_TYPE_LIVE_GIFT_ACTION_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	AD_CONVERT_TYPE_LIVE_HOMEPAGE_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_LIVE_JOIN_GROUP_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
	AD_CONVERT_TYPE_LIVE_NATIVE_ACITON_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_NATIVE_ACITON"
	AD_CONVERT_TYPE_LIVE_OTO_CLICK_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_CLICK"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK"
	AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION_NativeAnchorGetV30FilteringExternalAction NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION"
	AD_CONVERT_TYPE_LIVE_STAY_TIME_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_NativeAnchorGetV30FilteringExternalAction    NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_NativeAnchorGetV30FilteringExternalAction       NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LOAN_NativeAnchorGetV30FilteringExternalAction                        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LOAN"
	AD_CONVERT_TYPE_LOAN_COMPLETION_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LOAN_COMPLETION"
	AD_CONVERT_TYPE_LOAN_CREDIT_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LOAN_CREDIT"
	AD_CONVERT_TYPE_LOCATION_ACTION_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LOCATION_ACTION"
	AD_CONVERT_TYPE_LOTTERY_NativeAnchorGetV30FilteringExternalAction                     NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LOTTERY"
	AD_CONVERT_TYPE_LT_ROI_NativeAnchorGetV30FilteringExternalAction                      NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_LT_ROI"
	AD_CONVERT_TYPE_MAP_SEARCH_NativeAnchorGetV30FilteringExternalAction                  NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MAP_SEARCH"
	AD_CONVERT_TYPE_MESSAGE_NativeAnchorGetV30FilteringExternalAction                     NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MESSAGE"
	AD_CONVERT_TYPE_MESSAGE_ACTION_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MESSAGE_ACTION"
	AD_CONVERT_TYPE_MESSAGE_CLICK_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLICK"
	AD_CONVERT_TYPE_MESSAGE_CLUE_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLUE"
	AD_CONVERT_TYPE_MESSAGE_INTERACTION_NativeAnchorGetV30FilteringExternalAction         NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MESSAGE_INTERACTION"
	AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP_NativeAnchorGetV30FilteringExternalAction          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP"
	AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS_NativeAnchorGetV30FilteringExternalAction       NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS"
	AD_CONVERT_TYPE_MESSAGE_SERVICE_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MESSAGE_SERVICE"
	AD_CONVERT_TYPE_MULTIPLE_NativeAnchorGetV30FilteringExternalAction                    NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_MULTIPLE"
	AD_CONVERT_TYPE_NATIVE_ACTION_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_NATIVE_ACTION"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_NativeAnchorGetV30FilteringExternalAction           NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_NEXT_DAY_OPEN_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_NOTIFY_DOWNLOAD_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	AD_CONVERT_TYPE_NOTIFY_FORM_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_NOTIFY_FORM"
	AD_CONVERT_TYPE_OTHER_NativeAnchorGetV30FilteringExternalAction                       NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_OTHER"
	AD_CONVERT_TYPE_OTO_PAY_NativeAnchorGetV30FilteringExternalAction                     NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_OTO_PAY"
	AD_CONVERT_TYPE_OTO_STAY_TIME_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_OTO_STAY_TIME"
	AD_CONVERT_TYPE_PAID_CLUE_NativeAnchorGetV30FilteringExternalAction                   NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PAID_CLUE"
	AD_CONVERT_TYPE_PAY_NativeAnchorGetV30FilteringExternalAction                         NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_PERSONAL_INFORMATION_NativeAnchorGetV30FilteringExternalAction        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PERSONAL_INFORMATION"
	AD_CONVERT_TYPE_PHONE_NativeAnchorGetV30FilteringExternalAction                       NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PHONE"
	AD_CONVERT_TYPE_PHONE_CONFIRM_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PHONE_CONFIRM"
	AD_CONVERT_TYPE_PHONE_CONNECT_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PHONE_CONNECT"
	AD_CONVERT_TYPE_PHONE_EFFECTIVE_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PHONE_EFFECTIVE"
	AD_CONVERT_TYPE_POI_ADDRESS_CLICK_NativeAnchorGetV30FilteringExternalAction           NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_POI_ADDRESS_CLICK"
	AD_CONVERT_TYPE_POI_COLLECT_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_POI_COLLECT"
	AD_CONVERT_TYPE_POI_MULTIPLE_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_POI_MULTIPLE"
	AD_CONVERT_TYPE_PREMIUM_PAYMENT_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PREMIUM_PAYMENT"
	AD_CONVERT_TYPE_PREMIUM_ROI_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PREMIUM_ROI"
	AD_CONVERT_TYPE_PREMIUM_UPGEADE_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PREMIUM_UPGEADE"
	AD_CONVERT_TYPE_PRE_LOAN_CREDIT_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PRE_LOAN_CREDIT"
	AD_CONVERT_TYPE_PURCHASE_OF_GOODS_NativeAnchorGetV30FilteringExternalAction           NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PURCHASE_OF_GOODS"
	AD_CONVERT_TYPE_PURCHASE_ROI_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI"
	AD_CONVERT_TYPE_PURCHASE_ROI_2_D_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_2D"
	AD_CONVERT_TYPE_PURCHASE_ROI_7_D_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_7D"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_QQ_NativeAnchorGetV30FilteringExternalAction                          NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_QQ"
	AD_CONVERT_TYPE_REDIRECT_NativeAnchorGetV30FilteringExternalAction                    NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_REDIRECT"
	AD_CONVERT_TYPE_REDIRECT_TO_SHOP_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_SHOP"
	AD_CONVERT_TYPE_REDIRECT_TO_STORE_NativeAnchorGetV30FilteringExternalAction           NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_STORE"
	AD_CONVERT_TYPE_RESERVATION_NativeAnchorGetV30FilteringExternalAction                 NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_RESERVATION"
	AD_CONVERT_TYPE_RETENTION_7_D_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_RETENTION_7D"
	AD_CONVERT_TYPE_RETENTION_DAYS_NativeAnchorGetV30FilteringExternalAction              NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_RETENTION_DAYS"
	AD_CONVERT_TYPE_RSS_NativeAnchorGetV30FilteringExternalAction                         NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_RSS"
	AD_CONVERT_TYPE_SALES_LEAD_NativeAnchorGetV30FilteringExternalAction                  NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_SALES_LEAD"
	AD_CONVERT_TYPE_SHARE_ACTION_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_SHARE_ACTION"
	AD_CONVERT_TYPE_SHOPPING_NativeAnchorGetV30FilteringExternalAction                    NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_SHOPPING_ACTION_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_SHOPPING_ACTION"
	AD_CONVERT_TYPE_SHOPPING_CART_NativeAnchorGetV30FilteringExternalAction               NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_SHOPPING_CART"
	AD_CONVERT_TYPE_SHOW_OFF_NUM_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	AD_CONVERT_TYPE_STAY_TIME_NativeAnchorGetV30FilteringExternalAction                   NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_STAY_TIME"
	AD_CONVERT_TYPE_SUBMIT_CERTIFICATION_NativeAnchorGetV30FilteringExternalAction        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_SUBMIT_CERTIFICATION"
	AD_CONVERT_TYPE_SUCCESSORDER_ACTION_NativeAnchorGetV30FilteringExternalAction         NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_UG_ROI_NativeAnchorGetV30FilteringExternalAction                      NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_UG_ROI"
	AD_CONVERT_TYPE_UPDATE_LEVEL_NativeAnchorGetV30FilteringExternalAction                NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_UPDATE_LEVEL"
	AD_CONVERT_TYPE_VIEW_NativeAnchorGetV30FilteringExternalAction                        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_VIEW"
	AD_CONVERT_TYPE_VOTE_NativeAnchorGetV30FilteringExternalAction                        NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_VOTE"
	AD_CONVERT_TYPE_WECHAT_NativeAnchorGetV30FilteringExternalAction                      NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_WECHAT"
	AD_CONVERT_TYPE_WECHAT_PAY_NativeAnchorGetV30FilteringExternalAction                  NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_WECHAT_PAY"
	AD_CONVERT_TYPE_WECHAT_REGISTER_NativeAnchorGetV30FilteringExternalAction             NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_WECHAT_REGISTER"
	AD_CONVERT_TYPE_WECHAT_WECOM_ADD_NativeAnchorGetV30FilteringExternalAction            NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_WECHAT_WECOM_ADD"
	AD_CONVERT_TYPE_XPATH_NativeAnchorGetV30FilteringExternalAction                       NativeAnchorGetV30FilteringExternalAction = "AD_CONVERT_TYPE_XPATH"
)

// All allowed values of NativeAnchorGetV30FilteringExternalAction enum
var AllowedNativeAnchorGetV30FilteringExternalActionEnumValues = []NativeAnchorGetV30FilteringExternalAction{
	"AD_CONVERT_PAGE_VIEW",
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
	"AD_CONVERT_TYPE_WECHAT_PAY",
	"AD_CONVERT_TYPE_WECHAT_REGISTER",
	"AD_CONVERT_TYPE_WECHAT_WECOM_ADD",
	"AD_CONVERT_TYPE_XPATH",
}

// NewNativeAnchorGetV30FilteringExternalActionFromValue returns a pointer to a valid NativeAnchorGetV30FilteringExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30FilteringExternalActionFromValue(v string) (*NativeAnchorGetV30FilteringExternalAction, error) {
	ev := NativeAnchorGetV30FilteringExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30FilteringExternalAction: valid values are %v", v, AllowedNativeAnchorGetV30FilteringExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30FilteringExternalAction) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30FilteringExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_filtering_external_action value
func (v NativeAnchorGetV30FilteringExternalAction) Ptr() *NativeAnchorGetV30FilteringExternalAction {
	return &v
}
