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

// EventManagerShareCancelV30AllAccountType
type EventManagerShareCancelV30AllAccountType string

// List of event_manager_share_cancel_v3.0_all_account_type
const (
	AD_EventManagerShareCancelV30AllAccountType EventManagerShareCancelV30AllAccountType = "AD"
)

// All allowed values of EventManagerShareCancelV30AllAccountType enum
var AllowedEventManagerShareCancelV30AllAccountTypeEnumValues = []EventManagerShareCancelV30AllAccountType{
	"AD",
}

// NewEventManagerShareCancelV30AllAccountTypeFromValue returns a pointer to a valid EventManagerShareCancelV30AllAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareCancelV30AllAccountTypeFromValue(v string) (*EventManagerShareCancelV30AllAccountType, error) {
	ev := EventManagerShareCancelV30AllAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareCancelV30AllAccountType: valid values are %v", v, AllowedEventManagerShareCancelV30AllAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareCancelV30AllAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareCancelV30AllAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_cancel_v3.0_all_account_type value
func (v EventManagerShareCancelV30AllAccountType) Ptr() *EventManagerShareCancelV30AllAccountType {
	return &v
}
