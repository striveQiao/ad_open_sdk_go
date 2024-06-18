/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
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

type ApiOpenApiV30SubscribeAccountsListGetRequestExample struct {
	APPAccessToken  string                              `json:"APP-Access-Token"`
	AppId           int64                               `json:"app_id"`
	SubscribeTaskId int64                               `json:"subscribe_task_id"`
	Events          []string                            `json:"events,omitempty"`
	CoreUserId      int64                               `json:"core_user_id,omitempty"`
	AdvertiserIds   []int64                             `json:"advertiser_ids,omitempty"`
	Statuses        []*SubscribeAccountsListV30Statuses `json:"statuses,omitempty"`
	Cursor          int64                               `json:"cursor,omitempty"`
	Count           int64                               `json:"count,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/subscribe/accounts/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30SubscribeAccountsListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.SubscribeAccountsListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		APPAccessToken(request.APPAccessToken).AppId(request.AppId).SubscribeTaskId(request.SubscribeTaskId).Events(request.Events).CoreUserId(request.CoreUserId).AdvertiserIds(request.AdvertiserIds).Statuses(request.Statuses).Cursor(request.Cursor).Count(request.Count).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
