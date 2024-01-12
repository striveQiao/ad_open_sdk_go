/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
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

// QianchuanAdBudgetUpdateV10ApiService QianchuanAdBudgetUpdateV10Api service
type QianchuanAdBudgetUpdateV10ApiService service

type ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest struct {
	ctx                               context.Context
	ApiService                        *QianchuanAdBudgetUpdateV10ApiService
	qianchuanAdBudgetUpdateV10Request *QianchuanAdBudgetUpdateV10Request
}

func (r *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest) QianchuanAdBudgetUpdateV10Request(qianchuanAdBudgetUpdateV10Request QianchuanAdBudgetUpdateV10Request) *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest {
	r.qianchuanAdBudgetUpdateV10Request = &qianchuanAdBudgetUpdateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest) Execute() (*QianchuanAdBudgetUpdateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAdBudgetUpdatePost Method for OpenApiV10QianchuanAdBudgetUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest
*/
func (a *QianchuanAdBudgetUpdateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest {
	return &ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAdBudgetUpdateV10Response
func (a *QianchuanAdBudgetUpdateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanAdBudgetUpdatePostRequest) (*QianchuanAdBudgetUpdateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanAdBudgetUpdateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/ad/budget/update/"

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
	localVarPostBody = r.qianchuanAdBudgetUpdateV10Request
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
