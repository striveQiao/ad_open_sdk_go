/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
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

// ToolsAipThirdSiteCreateV2ApiService ToolsAipThirdSiteCreateV2Api service
type ToolsAipThirdSiteCreateV2ApiService service

type ApiOpenApi2ToolsAipThirdSiteCreatePostRequest struct {
	ctx                              context.Context
	ApiService                       *ToolsAipThirdSiteCreateV2ApiService
	toolsAipThirdSiteCreateV2Request *ToolsAipThirdSiteCreateV2Request
}

func (r *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest) ToolsAipThirdSiteCreateV2Request(toolsAipThirdSiteCreateV2Request ToolsAipThirdSiteCreateV2Request) *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest {
	r.toolsAipThirdSiteCreateV2Request = &toolsAipThirdSiteCreateV2Request
	return r
}

func (r *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest) Execute() (*ToolsAipThirdSiteCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAipThirdSiteCreatePost Method for OpenApi2ToolsAipThirdSiteCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAipThirdSiteCreatePostRequest
*/
func (a *ToolsAipThirdSiteCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest {
	return &ApiOpenApi2ToolsAipThirdSiteCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAipThirdSiteCreateV2Response
func (a *ToolsAipThirdSiteCreateV2ApiService) postExecute(r *ApiOpenApi2ToolsAipThirdSiteCreatePostRequest) (*ToolsAipThirdSiteCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAipThirdSiteCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/aip_third_site/create/"

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
	localVarPostBody = r.toolsAipThirdSiteCreateV2Request
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
