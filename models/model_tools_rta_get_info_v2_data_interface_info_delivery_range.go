/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange
type ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange string

// List of tools_rta_get_info_v2_data_interface_info_delivery_range
const (
	LOCAL_ONLY_ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange         ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange = "LOCAL_ONLY"
	UNION_ONLY_ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange         ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange = "UNION_ONLY"
	UNIVERSAL_DELIVERY_ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange = "UNIVERSAL_DELIVERY"
)

// All allowed values of ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange enum
var AllowedToolsRtaGetInfoV2DataInterfaceInfoDeliveryRangeEnumValues = []ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange{
	"LOCAL_ONLY",
	"UNION_ONLY",
	"UNIVERSAL_DELIVERY",
}

// NewToolsRtaGetInfoV2DataInterfaceInfoDeliveryRangeFromValue returns a pointer to a valid ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRtaGetInfoV2DataInterfaceInfoDeliveryRangeFromValue(v string) (*ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange, error) {
	ev := ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange: valid values are %v", v, AllowedToolsRtaGetInfoV2DataInterfaceInfoDeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange) IsValid() bool {
	for _, existing := range AllowedToolsRtaGetInfoV2DataInterfaceInfoDeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rta_get_info_v2_data_interface_info_delivery_range value
func (v ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange) Ptr() *ToolsRtaGetInfoV2DataInterfaceInfoDeliveryRange {
	return &v
}
