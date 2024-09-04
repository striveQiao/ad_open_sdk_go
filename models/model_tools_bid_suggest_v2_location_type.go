/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2LocationType
type ToolsBidSuggestV2LocationType string

// List of tools_bid_suggest_v2_location_type
const (
	CURRENT_ToolsBidSuggestV2LocationType ToolsBidSuggestV2LocationType = "CURRENT"
	TRAVEL_ToolsBidSuggestV2LocationType  ToolsBidSuggestV2LocationType = "TRAVEL"
	HOME_ToolsBidSuggestV2LocationType    ToolsBidSuggestV2LocationType = "HOME"
	ALL_ToolsBidSuggestV2LocationType     ToolsBidSuggestV2LocationType = "ALL"
)

// Ptr returns reference to tools_bid_suggest_v2_location_type value
func (v ToolsBidSuggestV2LocationType) Ptr() *ToolsBidSuggestV2LocationType {
	return &v
}
