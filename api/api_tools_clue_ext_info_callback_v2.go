/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
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

// ToolsClueExtInfoCallbackV2ApiService ToolsClueExtInfoCallbackV2Api service
type ToolsClueExtInfoCallbackV2ApiService service

type ApiOpenApi2ToolsClueExtInfoCallbackPostRequest struct {
	ctx                               context.Context
	ApiService                        *ToolsClueExtInfoCallbackV2ApiService
	toolsClueExtInfoCallbackV2Request *ToolsClueExtInfoCallbackV2Request
}

func (r *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest) ToolsClueExtInfoCallbackV2Request(toolsClueExtInfoCallbackV2Request ToolsClueExtInfoCallbackV2Request) *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest {
	r.toolsClueExtInfoCallbackV2Request = &toolsClueExtInfoCallbackV2Request
	return r
}

func (r *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest) Execute() (*ToolsClueExtInfoCallbackV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueExtInfoCallbackPost Method for OpenApi2ToolsClueExtInfoCallbackPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueExtInfoCallbackPostRequest
*/
func (a *ToolsClueExtInfoCallbackV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest {
	return &ApiOpenApi2ToolsClueExtInfoCallbackPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueExtInfoCallbackV2Response
func (a *ToolsClueExtInfoCallbackV2ApiService) postExecute(r *ApiOpenApi2ToolsClueExtInfoCallbackPostRequest) (*ToolsClueExtInfoCallbackV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueExtInfoCallbackV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/ext_info/callback/"

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
	localVarPostBody = r.toolsClueExtInfoCallbackV2Request
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
