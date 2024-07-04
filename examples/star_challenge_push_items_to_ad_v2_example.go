/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
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

type ApiOpenApi2StarChallengePushItemsToAdPostRequestExample struct {
	StarChallengePushItemsToAdV2Request StarChallengePushItemsToAdV2Request `json:"StarChallengePushItemsToAdV2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/star/challenge/push_items_to_ad/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2StarChallengePushItemsToAdPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.StarChallengePushItemsToAdV2Api().
		Post(ctx).
		AccessToken(accessToken).
		StarChallengePushItemsToAdV2Request(request.StarChallengePushItemsToAdV2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
