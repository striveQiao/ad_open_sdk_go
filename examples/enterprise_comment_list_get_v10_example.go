/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

type ApiOpenApiV10EnterpriseCommentListGetGetRequestExample struct {
	CcAccountId int64                                `json:"cc_account_id,omitempty"`
	EDouyinId   string                               `json:"e_douyin_id,omitempty"`
	EndTime     string                               `json:"end_time,omitempty"`
	Filter      EnterpriseCommentListGetV10Filter    `json:"filter,omitempty"`
	OrderField  string                               `json:"order_field,omitempty"`
	OrderType   EnterpriseCommentListGetV10OrderType `json:"order_type,omitempty"`
	Page        int64                                `json:"page,omitempty"`
	PageSize    int64                                `json:"page_size,omitempty"`
	StartTime   string                               `json:"start_time,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/enterprise/comment/list/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10EnterpriseCommentListGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.EnterpriseCommentListGetV10Api().
		Get(ctx).
		AccessToken(accessToken).
		CcAccountId(request.CcAccountId).EDouyinId(request.EDouyinId).EndTime(request.EndTime).Filter(request.Filter).OrderField(request.OrderField).OrderType(request.OrderType).Page(request.Page).PageSize(request.PageSize).StartTime(request.StartTime).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
