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

// AccountFundGetV30ApiService AccountFundGetV30Api service
type AccountFundGetV30ApiService service

type ApiOpenApiV30AccountFundGetGetRequest struct {
	ctx            context.Context
	ApiService     *AccountFundGetV30ApiService
	accountIds     *[]int64
	accountType    *AccountFundGetV30AccountType
	grantTypeSplit *AccountFundGetV30GrantTypeSplit
}

func (r *ApiOpenApiV30AccountFundGetGetRequest) AccountIds(accountIds []int64) *ApiOpenApiV30AccountFundGetGetRequest {
	r.accountIds = &accountIds
	return r
}

func (r *ApiOpenApiV30AccountFundGetGetRequest) AccountType(accountType AccountFundGetV30AccountType) *ApiOpenApiV30AccountFundGetGetRequest {
	r.accountType = &accountType
	return r
}

func (r *ApiOpenApiV30AccountFundGetGetRequest) GrantTypeSplit(grantTypeSplit AccountFundGetV30GrantTypeSplit) *ApiOpenApiV30AccountFundGetGetRequest {
	r.grantTypeSplit = &grantTypeSplit
	return r
}

func (r *ApiOpenApiV30AccountFundGetGetRequest) Execute() (*AccountFundGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30AccountFundGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30AccountFundGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AccountFundGetGetRequest) WithLog(enable bool) *ApiOpenApiV30AccountFundGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AccountFundGetGet Method for OpenApiV30AccountFundGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AccountFundGetGetRequest
*/
func (a *AccountFundGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30AccountFundGetGetRequest {
	return &ApiOpenApiV30AccountFundGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountFundGetV30Response
func (a *AccountFundGetV30ApiService) getExecute(r *ApiOpenApiV30AccountFundGetGetRequest) (*AccountFundGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AccountFundGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/account/fund/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountIds == nil {
		return localVarReturnValue, nil, ReportError("accountIds is required and must be specified")
	}
	if len(*r.accountIds) < 1 {
		return localVarReturnValue, nil, ReportError("accountIds must have at least 1 elements")
	}
	if len(*r.accountIds) > 20 {
		return localVarReturnValue, nil, ReportError("accountIds must have less than 20 elements")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_ids", r.accountIds)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	if r.grantTypeSplit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "grant_type_split", r.grantTypeSplit)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
