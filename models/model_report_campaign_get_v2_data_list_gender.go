/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2DataListGender
type ReportCampaignGetV2DataListGender string

// List of report_campaign_get_v2_data_list_gender
const (
	NONE_ReportCampaignGetV2DataListGender             ReportCampaignGetV2DataListGender = "NONE"
	GENDER_FEMALE_ReportCampaignGetV2DataListGender    ReportCampaignGetV2DataListGender = "GENDER_FEMALE"
	GENDER_UNLIMITED_ReportCampaignGetV2DataListGender ReportCampaignGetV2DataListGender = "GENDER_UNLIMITED"
	GENDER_MALE_ReportCampaignGetV2DataListGender      ReportCampaignGetV2DataListGender = "GENDER_MALE"
)

// Ptr returns reference to report_campaign_get_v2_data_list_gender value
func (v ReportCampaignGetV2DataListGender) Ptr() *ReportCampaignGetV2DataListGender {
	return &v
}
