/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderOverviewGetV2ResponseDataComment
type StarReportOrderOverviewGetV2ResponseDataComment struct {
	//
	HighFrequencyWords []string `json:"high_frequency_words,omitempty"`
	//
	NegRate *int64 `json:"neg_rate,omitempty"`
	//
	NeuRate *int64 `json:"neu_rate,omitempty"`
	//
	PosRate *int64 `json:"pos_rate,omitempty"`
}
