/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceLocationType
type AdGetV2DataAudienceLocationType string

// List of ad_get_v2_data_audience_location_type
const (
	CURRENT_AdGetV2DataAudienceLocationType AdGetV2DataAudienceLocationType = "CURRENT"
	TRAVEL_AdGetV2DataAudienceLocationType  AdGetV2DataAudienceLocationType = "TRAVEL"
	HOME_AdGetV2DataAudienceLocationType    AdGetV2DataAudienceLocationType = "HOME"
	ALL_AdGetV2DataAudienceLocationType     AdGetV2DataAudienceLocationType = "ALL"
)

// Ptr returns reference to ad_get_v2_data_audience_location_type value
func (v AdGetV2DataAudienceLocationType) Ptr() *AdGetV2DataAudienceLocationType {
	return &v
}
