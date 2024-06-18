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

type ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequestExample struct {
	StarId      int64  `json:"star_id"`
	DemandId    int64  `json:"demand_id"`
	ChannelId   string `json:"channel_id,omitempty"`
	DeveloperId int64  `json:"developer_id,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/mcn/get_contracted_challenge_url/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarMcnGetContractedChallengeUrlV2Api().
		Get(ctx).
		AccessToken(accessToken).
		StarId(request.StarId).DemandId(request.DemandId).ChannelId(request.ChannelId).DeveloperId(request.DeveloperId).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
