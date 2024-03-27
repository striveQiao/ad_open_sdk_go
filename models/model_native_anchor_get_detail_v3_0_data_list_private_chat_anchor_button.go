/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetDetailV30DataListPrivateChatAnchorButton
type NativeAnchorGetDetailV30DataListPrivateChatAnchorButton string

// List of native_anchor_get_detail_v3.0_data_list_private_chat_anchor_button
const (
	PRIVATE_MESSAGE_NativeAnchorGetDetailV30DataListPrivateChatAnchorButton  NativeAnchorGetDetailV30DataListPrivateChatAnchorButton = "PRIVATE_MESSAGE"
	CONSULT_NOW_NativeAnchorGetDetailV30DataListPrivateChatAnchorButton      NativeAnchorGetDetailV30DataListPrivateChatAnchorButton = "CONSULT_NOW"
	CONSULT_ADVISOR_NativeAnchorGetDetailV30DataListPrivateChatAnchorButton  NativeAnchorGetDetailV30DataListPrivateChatAnchorButton = "CONSULT_ADVISOR"
	CONSULT_DESIGNER_NativeAnchorGetDetailV30DataListPrivateChatAnchorButton NativeAnchorGetDetailV30DataListPrivateChatAnchorButton = "CONSULT_DESIGNER"
	ASK_TEACHER_NativeAnchorGetDetailV30DataListPrivateChatAnchorButton      NativeAnchorGetDetailV30DataListPrivateChatAnchorButton = "ASK_TEACHER"
)

// All allowed values of NativeAnchorGetDetailV30DataListPrivateChatAnchorButton enum
var AllowedNativeAnchorGetDetailV30DataListPrivateChatAnchorButtonEnumValues = []NativeAnchorGetDetailV30DataListPrivateChatAnchorButton{
	"PRIVATE_MESSAGE",
	"CONSULT_NOW",
	"CONSULT_ADVISOR",
	"CONSULT_DESIGNER",
	"ASK_TEACHER",
}

// NewNativeAnchorGetDetailV30DataListPrivateChatAnchorButtonFromValue returns a pointer to a valid NativeAnchorGetDetailV30DataListPrivateChatAnchorButton
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetDetailV30DataListPrivateChatAnchorButtonFromValue(v string) (*NativeAnchorGetDetailV30DataListPrivateChatAnchorButton, error) {
	ev := NativeAnchorGetDetailV30DataListPrivateChatAnchorButton(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetDetailV30DataListPrivateChatAnchorButton: valid values are %v", v, AllowedNativeAnchorGetDetailV30DataListPrivateChatAnchorButtonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetDetailV30DataListPrivateChatAnchorButton) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetDetailV30DataListPrivateChatAnchorButtonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_detail_v3.0_data_list_private_chat_anchor_button value
func (v NativeAnchorGetDetailV30DataListPrivateChatAnchorButton) Ptr() *NativeAnchorGetDetailV30DataListPrivateChatAnchorButton {
	return &v
}
