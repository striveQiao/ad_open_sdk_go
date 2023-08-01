/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
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

// ToolsPrivativeWordProjectAddV30ApiService ToolsPrivativeWordProjectAddV30Api service
type ToolsPrivativeWordProjectAddV30ApiService service

type ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest struct {
	ctx                                    context.Context
	ApiService                             *ToolsPrivativeWordProjectAddV30ApiService
	toolsPrivativeWordProjectAddV30Request *ToolsPrivativeWordProjectAddV30Request
}

func (r *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest) ToolsPrivativeWordProjectAddV30Request(toolsPrivativeWordProjectAddV30Request ToolsPrivativeWordProjectAddV30Request) *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest {
	r.toolsPrivativeWordProjectAddV30Request = &toolsPrivativeWordProjectAddV30Request
	return r
}

func (r *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest) Execute() (*ToolsPrivativeWordProjectAddV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsPrivativeWordProjectAddPost Method for OpenApiV30ToolsPrivativeWordProjectAddPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest
*/
func (a *ToolsPrivativeWordProjectAddV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest {
	return &ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPrivativeWordProjectAddV30Response
func (a *ToolsPrivativeWordProjectAddV30ApiService) postExecute(r *ApiOpenApiV30ToolsPrivativeWordProjectAddPostRequest) (*ToolsPrivativeWordProjectAddV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsPrivativeWordProjectAddV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/privative_word/project/add/"

	localVarHeaderParams := make(map[string]string)
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
	localVarPostBody = r.toolsPrivativeWordProjectAddV30Request
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
