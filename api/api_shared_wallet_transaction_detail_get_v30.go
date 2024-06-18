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

// SharedWalletTransactionDetailGetV30ApiService SharedWalletTransactionDetailGetV30Api service
type SharedWalletTransactionDetailGetV30ApiService service

type ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest struct {
	ctx            context.Context
	ApiService     *SharedWalletTransactionDetailGetV30ApiService
	accountId      *int64
	accountType    *SharedWalletTransactionDetailGetV30AccountType
	sharedWalletId *int64
	startDate      *string
	endDate        *string
	page           *int64
	pageSize       *int64
}

// 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) AccountId(accountId int64) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	r.accountId = &accountId
	return r
}

// 账户类型
func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) AccountType(accountType SharedWalletTransactionDetailGetV30AccountType) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	r.accountType = &accountType
	return r
}

// 共享钱包ID
func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) SharedWalletId(sharedWalletId int64) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	r.sharedWalletId = &sharedWalletId
	return r
}

// 开始时间，格式YYYY-MM-DD
func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) StartDate(startDate string) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间，格式YYYY-MM-DD
func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) EndDate(endDate string) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	r.endDate = &endDate
	return r
}

// 页码；默认为1
func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) Page(page int64) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	r.page = &page
	return r
}

// 页面大小；默认为10; 注意：page*page_size不可大于10000
func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) Execute() (*SharedWalletTransactionDetailGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) WithLog(enable bool) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SharedWalletTransactionDetailGetGet Method for OpenApiV30SharedWalletTransactionDetailGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest
*/
func (a *SharedWalletTransactionDetailGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest {
	return &ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SharedWalletTransactionDetailGetV30Response
func (a *SharedWalletTransactionDetailGetV30ApiService) getExecute(r *ApiOpenApiV30SharedWalletTransactionDetailGetGetRequest) (*SharedWalletTransactionDetailGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SharedWalletTransactionDetailGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/shared_wallet/transaction_detail/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}
	if r.sharedWalletId == nil {
		return localVarReturnValue, nil, ReportError("sharedWalletId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "shared_wallet_id", r.sharedWalletId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
