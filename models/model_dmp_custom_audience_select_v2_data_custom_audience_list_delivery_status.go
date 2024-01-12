/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus
type DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus string

// List of dmp_custom_audience_select_v2_data_custom_audience_list_delivery_status
const (
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH_DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE_DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus  DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH_DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus    DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE_DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus    DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE"
)

// All allowed values of DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus enum
var AllowedDmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatusEnumValues = []DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus{
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE",
}

// NewDmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatusFromValue returns a pointer to a valid DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatusFromValue(v string) (*DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus, error) {
	ev := DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus: valid values are %v", v, AllowedDmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus) IsValid() bool {
	for _, existing := range AllowedDmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_custom_audience_select_v2_data_custom_audience_list_delivery_status value
func (v DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus) Ptr() *DmpCustomAudienceSelectV2DataCustomAudienceListDeliveryStatus {
	return &v
}
