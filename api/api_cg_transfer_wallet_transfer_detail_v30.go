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

// CgTransferWalletTransferDetailV30ApiService CgTransferWalletTransferDetailV30Api service
type CgTransferWalletTransferDetailV30ApiService service

type ApiOpenApiV30CgTransferWalletTransferDetailGetRequest struct {
	ctx                  context.Context
	ApiService           *CgTransferWalletTransferDetailV30ApiService
	accountId            *int64
	accountType          *CgTransferWalletTransferDetailV30AccountType
	bizRequestNo         *string
	transferBizRequestNo *string
	transferSerial       *string
}

// 鉴权账户id
func (r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) AccountId(accountId int64) *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest {
	r.accountId = &accountId
	return r
}

// 鉴权账户类型
func (r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) AccountType(accountType CgTransferWalletTransferDetailV30AccountType) *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest {
	r.accountType = &accountType
	return r
}

// 请求id，推荐uuid，方便请求链路对齐
func (r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) BizRequestNo(bizRequestNo string) *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest {
	r.bizRequestNo = &bizRequestNo
	return r
}

// 发起转账的幂等id
func (r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) TransferBizRequestNo(transferBizRequestNo string) *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest {
	r.transferBizRequestNo = &transferBizRequestNo
	return r
}

// 转账单号，与transfer_biz_request_no两者传其一即可
func (r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) TransferSerial(transferSerial string) *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest {
	r.transferSerial = &transferSerial
	return r
}

func (r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) Execute() (*CgTransferWalletTransferDetailV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) AccessToken(accessToken string) *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) WithLog(enable bool) *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30CgTransferWalletTransferDetailGet Method for OpenApiV30CgTransferWalletTransferDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30CgTransferWalletTransferDetailGetRequest
*/
func (a *CgTransferWalletTransferDetailV30ApiService) Get(ctx context.Context) *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest {
	return &ApiOpenApiV30CgTransferWalletTransferDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CgTransferWalletTransferDetailV30Response
func (a *CgTransferWalletTransferDetailV30ApiService) getExecute(r *ApiOpenApiV30CgTransferWalletTransferDetailGetRequest) (*CgTransferWalletTransferDetailV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CgTransferWalletTransferDetailV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/cg_transfer/wallet/transfer/detail/"

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
	if r.bizRequestNo == nil {
		return localVarReturnValue, nil, ReportError("bizRequestNo is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "biz_request_no", r.bizRequestNo)
	if r.transferBizRequestNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transfer_biz_request_no", r.transferBizRequestNo)
	}
	if r.transferSerial != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transfer_serial", r.transferSerial)
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
