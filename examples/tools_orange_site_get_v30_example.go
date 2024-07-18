/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
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

type ApiOpenApiV30ToolsOrangeSiteGetGetRequestExample struct {
	AdvertiserId int64                             `json:"advertiser_id"`
	Page         int32                             `json:"page"`
	PageSize     int32                             `json:"page_size"`
	OptimizeGoal ToolsOrangeSiteGetV30OptimizeGoal `json:"optimize_goal"`
	Status       ToolsOrangeSiteGetV30Status       `json:"status,omitempty"`
	Filtering    ToolsOrangeSiteGetV30Filtering    `json:"filtering,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/tools/orange_site/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ToolsOrangeSiteGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsOrangeSiteGetV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Page(request.Page).PageSize(request.PageSize).OptimizeGoal(request.OptimizeGoal).Status(request.Status).Filtering(request.Filtering).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
