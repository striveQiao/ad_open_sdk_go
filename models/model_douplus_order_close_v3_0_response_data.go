/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderCloseV30ResponseData
type DouplusOrderCloseV30ResponseData struct {
	//
	FailList *map[string]string `json:"fail_list,omitempty"`
	//
	SuccessList []int64 `json:"success_list"`
}
