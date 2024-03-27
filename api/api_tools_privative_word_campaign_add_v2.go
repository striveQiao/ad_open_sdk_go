/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsPrivativeWordCampaignAddV2ApiService ToolsPrivativeWordCampaignAddV2Api service
type ToolsPrivativeWordCampaignAddV2ApiService service

type ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest struct {
	ctx                                    context.Context
	ApiService                             *ToolsPrivativeWordCampaignAddV2ApiService
	toolsPrivativeWordCampaignAddV2Request *ToolsPrivativeWordCampaignAddV2Request
}

func (r *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest) ToolsPrivativeWordCampaignAddV2Request(toolsPrivativeWordCampaignAddV2Request ToolsPrivativeWordCampaignAddV2Request) *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest {
	r.toolsPrivativeWordCampaignAddV2Request = &toolsPrivativeWordCampaignAddV2Request
	return r
}

func (r *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest) Execute() (*ToolsPrivativeWordCampaignAddV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPrivativeWordCampaignAddPost Method for OpenApi2ToolsPrivativeWordCampaignAddPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest
*/
func (a *ToolsPrivativeWordCampaignAddV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest {
	return &ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPrivativeWordCampaignAddV2Response
func (a *ToolsPrivativeWordCampaignAddV2ApiService) postExecute(r *ApiOpenApi2ToolsPrivativeWordCampaignAddPostRequest) (*ToolsPrivativeWordCampaignAddV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPrivativeWordCampaignAddV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/privative_word/campaign/add/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.toolsPrivativeWordCampaignAddV2Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
