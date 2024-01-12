/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"io"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

type ApiOpenApiOauth2AccessTokenPostRequestExample struct {
	Oauth2AccessTokenRequest Oauth2AccessTokenRequest `json:"Oauth2AccessTokenRequest,omitempty"`
}

// url: https://api.oceanengine.com/open_api/oauth2/access_token/ Post
func main() {
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiOauth2AccessTokenPostRequestExample
	request.Oauth2AccessTokenRequest.AppId = PtrInt64(1)
	request.Oauth2AccessTokenRequest.Secret = ""
	request.Oauth2AccessTokenRequest.AuthCode = ""
	resp, httpRes, err := apiClient.Oauth2AccessTokenApi().
		Post(ctx).
		Oauth2AccessTokenRequest(request.Oauth2AccessTokenRequest).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
