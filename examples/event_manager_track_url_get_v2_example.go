/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
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

type ApiOpenApi2EventManagerTrackUrlGetGetRequestExample struct {
	AdvertiserId      int64  `json:"advertiser_id"`
	AssetId           int64  `json:"asset_id"`
	DownloadUrl       string `json:"download_url,omitempty"`
	TrackUrlGroupName string `json:"track_url_group_name,omitempty"`
	TrackUrlGroupId   int64  `json:"track_url_group_id,omitempty"`
	Page              int64  `json:"page,omitempty"`
	PageSize          int64  `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/event_manager/track_url/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2EventManagerTrackUrlGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.EventManagerTrackUrlGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).AssetId(request.AssetId).DownloadUrl(request.DownloadUrl).TrackUrlGroupName(request.TrackUrlGroupName).TrackUrlGroupId(request.TrackUrlGroupId).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
