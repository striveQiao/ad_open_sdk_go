/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetScriptV2ResponseDataOrderScriptListInner struct for StarOrderGetScriptV2ResponseDataOrderScriptListInner
type StarOrderGetScriptV2ResponseDataOrderScriptListInner struct {
	// 订单id
	OrderId *int64 `json:"order_id,omitempty"`
	//
	ScriptList []*StarOrderGetScriptV2ResponseDataOrderScriptListInnerScriptListInner `json:"script_list,omitempty"`
}
