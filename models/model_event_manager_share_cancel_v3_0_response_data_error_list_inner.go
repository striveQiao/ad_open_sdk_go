/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerShareCancelV30ResponseDataErrorListInner struct for EventManagerShareCancelV30ResponseDataErrorListInner
type EventManagerShareCancelV30ResponseDataErrorListInner struct {
	AccountInfo    *EventManagerShareCancelV30ResponseDataErrorListInnerAccountInfo `json:"account_info,omitempty"`
	AllAccountType *EventManagerShareCancelV30DataErrorListAllAccountType           `json:"all_account_type,omitempty"`
	//
	ErrorMessage *string                                           `json:"error_message,omitempty"`
	ShareMode    *EventManagerShareCancelV30DataErrorListShareMode `json:"share_mode,omitempty"`
}
