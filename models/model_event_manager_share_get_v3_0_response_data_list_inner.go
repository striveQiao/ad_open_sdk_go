/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerShareGetV30ResponseDataListInner struct for EventManagerShareGetV30ResponseDataListInner
type EventManagerShareGetV30ResponseDataListInner struct {
	AccountInfo    *EventManagerShareGetV30ResponseDataListInnerAccountInfo `json:"account_info,omitempty"`
	AllAccountType *EventManagerShareGetV30DataListAllAccountType           `json:"all_account_type,omitempty"`
	ShareMode      *EventManagerShareGetV30DataListShareMode                `json:"share_mode,omitempty"`
}
