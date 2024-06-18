/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DiagnosisTaskAgentCreateV2ResponseData
type DiagnosisTaskAgentCreateV2ResponseData struct {
	//
	ErrCode *string `json:"err_code,omitempty"`
	//
	ErrMessage *string `json:"err_message,omitempty"`
	//
	FailVideoIds map[string]map[string]interface{} `json:"fail_video_ids,omitempty"`
	//
	TaskIds []int64 `json:"task_ids,omitempty"`
}
