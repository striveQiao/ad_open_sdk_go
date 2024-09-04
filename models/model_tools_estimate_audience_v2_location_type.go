/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2LocationType
type ToolsEstimateAudienceV2LocationType string

// List of tools_estimate_audience_v2_location_type
const (
	CURRENT_ToolsEstimateAudienceV2LocationType ToolsEstimateAudienceV2LocationType = "CURRENT"
	TRAVEL_ToolsEstimateAudienceV2LocationType  ToolsEstimateAudienceV2LocationType = "TRAVEL"
	HOME_ToolsEstimateAudienceV2LocationType    ToolsEstimateAudienceV2LocationType = "HOME"
	ALL_ToolsEstimateAudienceV2LocationType     ToolsEstimateAudienceV2LocationType = "ALL"
)

// Ptr returns reference to tools_estimate_audience_v2_location_type value
func (v ToolsEstimateAudienceV2LocationType) Ptr() *ToolsEstimateAudienceV2LocationType {
	return &v
}
