/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataStatus
type AdGetV2DataStatus string

// List of ad_get_v2_data_status
const (
	AD_STATUS_SOME_DELIVERY_OK_AdGetV2DataStatus              AdGetV2DataStatus = "AD_STATUS_SOME_DELIVERY_OK"
	AD_STATUS_DONE_AdGetV2DataStatus                          AdGetV2DataStatus = "AD_STATUS_DONE"
	AD_STATUS_CANNOT_EDIT_AdGetV2DataStatus                   AdGetV2DataStatus = "AD_STATUS_CANNOT_EDIT"
	AD_STATUS_CAMPAIGN_DISABLE_AdGetV2DataStatus              AdGetV2DataStatus = "AD_STATUS_CAMPAIGN_DISABLE"
	AD_STATUS_PRE_ONLINE_AdGetV2DataStatus                    AdGetV2DataStatus = "AD_STATUS_PRE_ONLINE"
	AD_STATUS_DELETE_AdGetV2DataStatus                        AdGetV2DataStatus = "AD_STATUS_DELETE"
	AD_STATUS_AWEME_ACCOUNT_PUNISHED_AdGetV2DataStatus        AdGetV2DataStatus = "AD_STATUS_AWEME_ACCOUNT_PUNISHED"
	AD_STATUS_AUDIT_AdGetV2DataStatus                         AdGetV2DataStatus = "AD_STATUS_AUDIT"
	AD_STATUS_CAMPAIGN_PRE_OFFLINE_BUDGET_AdGetV2DataStatus   AdGetV2DataStatus = "AD_STATUS_CAMPAIGN_PRE_OFFLINE_BUDGET"
	AD_STATUS_FROZEN_AdGetV2DataStatus                        AdGetV2DataStatus = "AD_STATUS_FROZEN"
	AD_STATUS_DELIVERY_OK_AdGetV2DataStatus                   AdGetV2DataStatus = "AD_STATUS_DELIVERY_OK"
	AD_STATUS_CAMPAIGN_EXCEED_AdGetV2DataStatus               AdGetV2DataStatus = "AD_STATUS_CAMPAIGN_EXCEED"
	AD_STATUS_CREATE_AdGetV2DataStatus                        AdGetV2DataStatus = "AD_STATUS_CREATE"
	AD_STATUS_ADVERTISER_PRE_OFFLINE_BUDGET_AdGetV2DataStatus AdGetV2DataStatus = "AD_STATUS_ADVERTISER_PRE_OFFLINE_BUDGET"
	AD_STATUS_EXTERNAL_URL_DISABLE_AdGetV2DataStatus          AdGetV2DataStatus = "AD_STATUS_EXTERNAL_URL_DISABLE"
	AD_STATUS_REAUDIT_AdGetV2DataStatus                       AdGetV2DataStatus = "AD_STATUS_REAUDIT"
	AD_STATUS_AUDIT_STATUS_ERROR_AdGetV2DataStatus            AdGetV2DataStatus = "AD_STATUS_AUDIT_STATUS_ERROR"
	AD_STATUS_DATA_ERROR_AdGetV2DataStatus                    AdGetV2DataStatus = "AD_STATUS_DATA_ERROR"
	AD_STATUS_DISABLE_BY_QUOTA_AdGetV2DataStatus              AdGetV2DataStatus = "AD_STATUS_DISABLE_BY_QUOTA"
	AD_STATUS_AWEME_ACCOUNT_DISABLED_AdGetV2DataStatus        AdGetV2DataStatus = "AD_STATUS_AWEME_ACCOUNT_DISABLED"
	AD_STATUS_PRODUCT_OFFLINE_AdGetV2DataStatus               AdGetV2DataStatus = "AD_STATUS_PRODUCT_OFFLINE"
	AD_STATUS_DISABLE_AdGetV2DataStatus                       AdGetV2DataStatus = "AD_STATUS_DISABLE"
	AD_STATUS_DRAFT_AdGetV2DataStatus                         AdGetV2DataStatus = "AD_STATUS_DRAFT"
	AD_STATUS_AUDIT_DENY_AdGetV2DataStatus                    AdGetV2DataStatus = "AD_STATUS_AUDIT_DENY"
	AD_STATUS_BUDGET_EXCEED_AdGetV2DataStatus                 AdGetV2DataStatus = "AD_STATUS_BUDGET_EXCEED"
	AD_STATUS_NOT_START_AdGetV2DataStatus                     AdGetV2DataStatus = "AD_STATUS_NOT_START"
	AD_STATUS_LIVE_ROOM_OFF_AdGetV2DataStatus                 AdGetV2DataStatus = "AD_STATUS_LIVE_ROOM_OFF"
	AD_STATUS_PRE_OFFLINE_BUDGET_AdGetV2DataStatus            AdGetV2DataStatus = "AD_STATUS_PRE_OFFLINE_BUDGET"
	AD_STATUS_DSP_AD_DISABLE_AdGetV2DataStatus                AdGetV2DataStatus = "AD_STATUS_DSP_AD_DISABLE"
	AD_STATUS_BALANCE_EXCEED_AdGetV2DataStatus                AdGetV2DataStatus = "AD_STATUS_BALANCE_EXCEED"
	AD_STATUS_ADVERTISER_BUDGET_EXCEED_AdGetV2DataStatus      AdGetV2DataStatus = "AD_STATUS_ADVERTISER_BUDGET_EXCEED"
	AD_STATUS_LIVE_CANNOT_LAUNCH_AdGetV2DataStatus            AdGetV2DataStatus = "AD_STATUS_LIVE_CANNOT_LAUNCH"
	AD_STATUS_NO_SCHEDULE_AdGetV2DataStatus                   AdGetV2DataStatus = "AD_STATUS_NO_SCHEDULE"
)

// Ptr returns reference to ad_get_v2_data_status value
func (v AdGetV2DataStatus) Ptr() *AdGetV2DataStatus {
	return &v
}
