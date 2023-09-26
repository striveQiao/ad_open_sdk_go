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

// EventManagerShareCancelV30DataErrorListAllAccountType
type EventManagerShareCancelV30DataErrorListAllAccountType string

// List of event_manager_share_cancel_v3.0_data_error_list_all_account_type
const (
	AD_EventManagerShareCancelV30DataErrorListAllAccountType EventManagerShareCancelV30DataErrorListAllAccountType = "AD"
)

// All allowed values of EventManagerShareCancelV30DataErrorListAllAccountType enum
var AllowedEventManagerShareCancelV30DataErrorListAllAccountTypeEnumValues = []EventManagerShareCancelV30DataErrorListAllAccountType{
	"AD",
}

// NewEventManagerShareCancelV30DataErrorListAllAccountTypeFromValue returns a pointer to a valid EventManagerShareCancelV30DataErrorListAllAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareCancelV30DataErrorListAllAccountTypeFromValue(v string) (*EventManagerShareCancelV30DataErrorListAllAccountType, error) {
	ev := EventManagerShareCancelV30DataErrorListAllAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareCancelV30DataErrorListAllAccountType: valid values are %v", v, AllowedEventManagerShareCancelV30DataErrorListAllAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareCancelV30DataErrorListAllAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareCancelV30DataErrorListAllAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_cancel_v3.0_data_error_list_all_account_type value
func (v EventManagerShareCancelV30DataErrorListAllAccountType) Ptr() *EventManagerShareCancelV30DataErrorListAllAccountType {
	return &v
}
