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

// EventManagerShareV30DataErrorListAllAccountType
type EventManagerShareV30DataErrorListAllAccountType string

// List of event_manager_share_v3.0_data_error_list_all_account_type
const (
	AD_EventManagerShareV30DataErrorListAllAccountType EventManagerShareV30DataErrorListAllAccountType = "AD"
)

// All allowed values of EventManagerShareV30DataErrorListAllAccountType enum
var AllowedEventManagerShareV30DataErrorListAllAccountTypeEnumValues = []EventManagerShareV30DataErrorListAllAccountType{
	"AD",
}

// NewEventManagerShareV30DataErrorListAllAccountTypeFromValue returns a pointer to a valid EventManagerShareV30DataErrorListAllAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareV30DataErrorListAllAccountTypeFromValue(v string) (*EventManagerShareV30DataErrorListAllAccountType, error) {
	ev := EventManagerShareV30DataErrorListAllAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareV30DataErrorListAllAccountType: valid values are %v", v, AllowedEventManagerShareV30DataErrorListAllAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareV30DataErrorListAllAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareV30DataErrorListAllAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_v3.0_data_error_list_all_account_type value
func (v EventManagerShareV30DataErrorListAllAccountType) Ptr() *EventManagerShareV30DataErrorListAllAccountType {
	return &v
}
