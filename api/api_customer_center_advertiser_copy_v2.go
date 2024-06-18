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

// CustomerCenterAdvertiserCopyV2ApiService CustomerCenterAdvertiserCopyV2Api service
type CustomerCenterAdvertiserCopyV2ApiService service

type ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest struct {
	ctx                                   context.Context
	ApiService                            *CustomerCenterAdvertiserCopyV2ApiService
	customerCenterAdvertiserCopyV2Request *CustomerCenterAdvertiserCopyV2Request
}

func (r *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest) CustomerCenterAdvertiserCopyV2Request(customerCenterAdvertiserCopyV2Request CustomerCenterAdvertiserCopyV2Request) *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest {
	r.customerCenterAdvertiserCopyV2Request = &customerCenterAdvertiserCopyV2Request
	return r
}

func (r *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest) Execute() (*CustomerCenterAdvertiserCopyV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest) AccessToken(accessToken string) *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest) WithLog(enable bool) *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CustomerCenterAdvertiserCopyPost Method for OpenApi2CustomerCenterAdvertiserCopyPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest
*/
func (a *CustomerCenterAdvertiserCopyV2ApiService) Post(ctx context.Context) *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest {
	return &ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CustomerCenterAdvertiserCopyV2Response
func (a *CustomerCenterAdvertiserCopyV2ApiService) postExecute(r *ApiOpenApi2CustomerCenterAdvertiserCopyPostRequest) (*CustomerCenterAdvertiserCopyV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CustomerCenterAdvertiserCopyV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/customer_center/advertiser/copy/"

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
	localVarPostBody = r.customerCenterAdvertiserCopyV2Request
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
