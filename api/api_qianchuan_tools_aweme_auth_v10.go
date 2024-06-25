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

// QianchuanToolsAwemeAuthV10ApiService QianchuanToolsAwemeAuthV10Api service
type QianchuanToolsAwemeAuthV10ApiService service

type ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest struct {
	ctx                               context.Context
	ApiService                        *QianchuanToolsAwemeAuthV10ApiService
	qianchuanToolsAwemeAuthV10Request *QianchuanToolsAwemeAuthV10Request
}

func (r *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest) QianchuanToolsAwemeAuthV10Request(qianchuanToolsAwemeAuthV10Request QianchuanToolsAwemeAuthV10Request) *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest {
	r.qianchuanToolsAwemeAuthV10Request = &qianchuanToolsAwemeAuthV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest) Execute() (*QianchuanToolsAwemeAuthV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsAwemeAuthPost Method for OpenApiV10QianchuanToolsAwemeAuthPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest
*/
func (a *QianchuanToolsAwemeAuthV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest {
	return &ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsAwemeAuthV10Response
func (a *QianchuanToolsAwemeAuthV10ApiService) postExecute(r *ApiOpenApiV10QianchuanToolsAwemeAuthPostRequest) (*QianchuanToolsAwemeAuthV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanToolsAwemeAuthV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/aweme_auth/"

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
	localVarPostBody = r.qianchuanToolsAwemeAuthV10Request
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
