/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus
type DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus string

// List of dmp_data_source_read_v2_data_data_list_default_audience_delivery_status
const (
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH_DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus    DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE_DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus  DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE_DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus    DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH_DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH"
)

// All allowed values of DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus enum
var AllowedDmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatusEnumValues = []DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus{
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH",
}

// NewDmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatusFromValue returns a pointer to a valid DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatusFromValue(v string) (*DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus, error) {
	ev := DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus: valid values are %v", v, AllowedDmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus) IsValid() bool {
	for _, existing := range AllowedDmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_data_source_read_v2_data_data_list_default_audience_delivery_status value
func (v DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus) Ptr() *DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus {
	return &v
}
