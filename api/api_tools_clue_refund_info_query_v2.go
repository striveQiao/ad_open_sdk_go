/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
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

// ToolsClueRefundInfoQueryV2ApiService ToolsClueRefundInfoQueryV2Api service
type ToolsClueRefundInfoQueryV2ApiService service

type ApiOpenApi2ToolsClueRefundInfoQueryPostRequest struct {
	ctx                               context.Context
	ApiService                        *ToolsClueRefundInfoQueryV2ApiService
	toolsClueRefundInfoQueryV2Request *ToolsClueRefundInfoQueryV2Request
}

func (r *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest) ToolsClueRefundInfoQueryV2Request(toolsClueRefundInfoQueryV2Request ToolsClueRefundInfoQueryV2Request) *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest {
	r.toolsClueRefundInfoQueryV2Request = &toolsClueRefundInfoQueryV2Request
	return r
}

func (r *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest) Execute() (*ToolsClueRefundInfoQueryV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueRefundInfoQueryPost Method for OpenApi2ToolsClueRefundInfoQueryPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueRefundInfoQueryPostRequest
*/
func (a *ToolsClueRefundInfoQueryV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest {
	return &ApiOpenApi2ToolsClueRefundInfoQueryPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueRefundInfoQueryV2Response
func (a *ToolsClueRefundInfoQueryV2ApiService) postExecute(r *ApiOpenApi2ToolsClueRefundInfoQueryPostRequest) (*ToolsClueRefundInfoQueryV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueRefundInfoQueryV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/refund/info/query/"

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
	localVarPostBody = r.toolsClueRefundInfoQueryV2Request
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
